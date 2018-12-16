// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"cmd/compile/internal/types"
	"cmd/internal/obj"
	"cmd/internal/objabi"
	"cmd/internal/src"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"math/bits"
	"os"
	"path/filepath"
)

func applyRewrite(f *Func, rb BlockRewriter, rv ValueRewriter) {
	// repeat rewrites until we find no more rewrites
	pendingLines := f.cachedLineStarts // Holds statement boundaries that need to be moved to a new value/block
	pendingLines.clear()
	for {
		change := false
		for _, b := range f.Blocks {
			if b.Control != nil && b.Control.Op == OpCopy {
				for b.Control.Op == OpCopy {
					b.SetControl(b.Control.Args[0])
				}
			}
			if rb(b) {
				change = true
			}
			for j, v := range b.Values {
				change = phielimValue(v) || change

				// Eliminate copy inputs.
				// If any copy input becomes unused, mark it
				// as invalid and discard its argument. Repeat
				// recursively on the discarded argument.
				// This phase helps remove phantom "dead copy" uses
				// of a value so that a x.Uses==1 rule condition
				// fires reliably.
				for i, a := range v.Args {
					if a.Op != OpCopy {
						continue
					}
					aa := copySource(a)
					v.SetArg(i, aa)
					// If a, a copy, has a line boundary indicator, attempt to find a new value
					// to hold it.  The first candidate is the value that will replace a (aa),
					// if it shares the same block and line and is eligible.
					// The second option is v, which has a as an input.  Because aa is earlier in
					// the data flow, it is the better choice.
					if a.Pos.IsStmt() == src.PosIsStmt {
						if aa.Block == a.Block && aa.Pos.Line() == a.Pos.Line() && aa.Pos.IsStmt() != src.PosNotStmt {
							aa.Pos = aa.Pos.WithIsStmt()
						} else if v.Block == a.Block && v.Pos.Line() == a.Pos.Line() && v.Pos.IsStmt() != src.PosNotStmt {
							v.Pos = v.Pos.WithIsStmt()
						} else {
							// Record the lost line and look for a new home after all rewrites are complete.
							// TODO: it's possible (in FOR loops, in particular) for statement boundaries for the same
							// line to appear in more than one block, but only one block is stored, so if both end
							// up here, then one will be lost.
							pendingLines.set(a.Pos.Line(), int32(a.Block.ID))
						}
						a.Pos = a.Pos.WithNotStmt()
					}
					change = true
					for a.Uses == 0 {
						b := a.Args[0]
						a.Reset(OpInvalid)
						a = b
					}
				}

				// apply rewrite function
				if rv(v) {
					change = true
					// If value changed to a poor choice for a statement boundary, move the boundary
					if v.Pos.IsStmt() == src.PosIsStmt {
						if k := nextGoodStatementIndex(v, j, b); k != j {
							v.Pos = v.Pos.WithNotStmt()
							b.Values[k].Pos = b.Values[k].Pos.WithIsStmt()
						}
					}
				}
			}
		}
		if !change {
			break
		}
	}
	// remove clobbered values
	for _, b := range f.Blocks {
		j := 0
		for i, v := range b.Values {
			vl := v.Pos.Line()
			if v.Op == OpInvalid {
				if v.Pos.IsStmt() == src.PosIsStmt {
					pendingLines.set(vl, int32(b.ID))
				}
				f.freeValue(v)
				continue
			}
			if v.Pos.IsStmt() != src.PosNotStmt && pendingLines.get(vl) == int32(b.ID) {
				pendingLines.remove(vl)
				v.Pos = v.Pos.WithIsStmt()
			}
			if i != j {
				b.Values[j] = v
			}
			j++
		}
		if pendingLines.get(b.Pos.Line()) == int32(b.ID) {
			b.Pos = b.Pos.WithIsStmt()
			pendingLines.remove(b.Pos.Line())
		}
		if j != len(b.Values) {
			tail := b.Values[j:]
			for j := range tail {
				tail[j] = nil
			}
			b.Values = b.Values[:j]
		}
	}
}

// Common functions called from rewriting rules

func Is64BitFloat(t *types.Type) bool {
	return t.Size() == 8 && t.IsFloat()
}

func Is32BitFloat(t *types.Type) bool {
	return t.Size() == 4 && t.IsFloat()
}

func Is64BitInt(t *types.Type) bool {
	return t.Size() == 8 && t.IsInteger()
}

func Is32BitInt(t *types.Type) bool {
	return t.Size() == 4 && t.IsInteger()
}

func Is16BitInt(t *types.Type) bool {
	return t.Size() == 2 && t.IsInteger()
}

func Is8BitInt(t *types.Type) bool {
	return t.Size() == 1 && t.IsInteger()
}

func IsPtr(t *types.Type) bool {
	return t.IsPtrShaped()
}

func IsSigned(t *types.Type) bool {
	return t.IsSigned()
}

// MergeSym merges two symbolic offsets. There is no real merging of
// offsets, we just pick the non-nil one.
func MergeSym(x, y interface{}) interface{} {
	if x == nil {
		return y
	}
	if y == nil {
		return x
	}
	panic(fmt.Sprintf("MergeSym with two non-nil syms %s %s", x, y))
}
func CanMergeSym(x, y interface{}) bool {
	return x == nil || y == nil
}

// CanMergeLoadClobber reports whether the load can be merged into target without
// invalidating the schedule.
// It also checks that the other non-load argument x is something we
// are ok with Clobbering.
func CanMergeLoadClobber(target, load, x *Value) bool {
	// The register containing x is going to get Clobbered.
	// Don't merge if we still need the value of x.
	// We don't have liveness information here, but we can
	// approximate x dying with:
	//  1) target is x's only use.
	//  2) target is not in a deeper loop than x.
	if x.Uses != 1 {
		return false
	}
	loopnest := x.Block.Func.loopnest()
	loopnest.calculateDepths()
	if loopnest.depth(target.Block.ID) > loopnest.depth(x.Block.ID) {
		return false
	}
	return CanMergeLoad(target, load)
}

// CanMergeLoad reports whether the load can be merged into target without
// invalidating the schedule.
func CanMergeLoad(target, load *Value) bool {
	if target.Block.ID != load.Block.ID {
		// If the load is in a different block do not merge it.
		return false
	}

	// We can't merge the load into the target if the load
	// has more than one use.
	if load.Uses != 1 {
		return false
	}

	mem := load.MemoryArg()

	// We need the load's memory arg to still be alive at target. That
	// can't be the case if one of target's args depends on a memory
	// state that is a successor of load's memory arg.
	//
	// For example, it would be invalid to merge load into target in
	// the following situation because newmem has killed oldmem
	// before target is reached:
	//     load = read ... oldmem
	//   newmem = write ... oldmem
	//     arg0 = read ... newmem
	//   target = add arg0 load
	//
	// If the argument comes from a different block then we can exclude
	// it immediately because it must dominate load (which is in the
	// same block as target).
	var args []*Value
	for _, a := range target.Args {
		if a != load && a.Block.ID == target.Block.ID {
			args = append(args, a)
		}
	}

	// memPreds contains memory states known to be predecessors of load's
	// memory state. It is lazily initialized.
	var memPreds map[*Value]bool
	for i := 0; len(args) > 0; i++ {
		const limit = 100
		if i >= limit {
			// Give up if we have done a lot of iterations.
			return false
		}
		v := args[len(args)-1]
		args = args[:len(args)-1]
		if target.Block.ID != v.Block.ID {
			// Since target and load are in the same block
			// we can stop searching when we leave the block.
			continue
		}
		if v.Op == OpPhi {
			// A Phi implies we have reached the top of the block.
			// The memory phi, if it exists, is always
			// the first logical store in the block.
			continue
		}
		if v.Type.IsTuple() && v.Type.FieldType(1).IsMemory() {
			// We could handle this situation however it is likely
			// to be very rare.
			return false
		}
		if v.Op.SymEffect()&SymAddr != 0 {
			// This case prevents an operation that calculates the
			// address of a local variable from being forced to schedule
			// before its corresponding VarDef.
			// See issue 28445.
			//   v1 = LOAD ...
			//   v2 = VARDEF
			//   v3 = LEAQ
			//   v4 = CMPQ v1 v3
			// We don't want to combine the CMPQ with the load, because
			// that would force the CMPQ to schedule before the VARDEF, which
			// in turn requires the LEAQ to schedule before the VARDEF.
			return false
		}
		if v.Type.IsMemory() {
			if memPreds == nil {
				// Initialise a map containing memory states
				// known to be predecessors of load's memory
				// state.
				memPreds = make(map[*Value]bool)
				m := mem
				const limit = 50
				for i := 0; i < limit; i++ {
					if m.Op == OpPhi {
						// The memory phi, if it exists, is always
						// the first logical store in the block.
						break
					}
					if m.Block.ID != target.Block.ID {
						break
					}
					if !m.Type.IsMemory() {
						break
					}
					memPreds[m] = true
					if len(m.Args) == 0 {
						break
					}
					m = m.MemoryArg()
				}
			}

			// We can merge if v is a predecessor of mem.
			//
			// For example, we can merge load into target in the
			// following scenario:
			//      x = read ... v
			//    mem = write ... v
			//   load = read ... mem
			// target = add x load
			if memPreds[v] {
				continue
			}
			return false
		}
		if len(v.Args) > 0 && v.Args[len(v.Args)-1] == mem {
			// If v takes mem as an input then we know mem
			// is valid at this point.
			continue
		}
		for _, a := range v.Args {
			if target.Block.ID == a.Block.ID {
				args = append(args, a)
			}
		}
	}

	return true
}

// isSameSym reports whether sym is the same as the given named symbol
func isSameSym(sym interface{}, name string) bool {
	s, ok := sym.(fmt.Stringer)
	return ok && s.String() == name
}

// nlz returns the number of leading zeros.
func nlz(x int64) int64 {
	return int64(bits.LeadingZeros64(uint64(x)))
}

// Ntz returns the number of trailing zeros.
func Ntz(x int64) int64 {
	return int64(bits.TrailingZeros64(uint64(x)))
}

// nlo returns the number of leading ones.
func nlo(x int64) int64 {
	return nlz(^x)
}

// nto returns the number of trailing ones.
func nto(x int64) int64 {
	return Ntz(^x)
}

// Log2 returns logarithm in base 2 of uint64(n), with Log2(0) = -1.
// Rounds down.
func Log2(n int64) int64 {
	return int64(bits.Len64(uint64(n))) - 1
}

// IsPowerOfTwo reports whether n is a power of 2.
func IsPowerOfTwo(n int64) bool {
	return n > 0 && n&(n-1) == 0
}

// Is32Bit reports whether n can be represented as a signed 32 bit integer.
func Is32Bit(n int64) bool {
	return n == int64(int32(n))
}

// Is16Bit reports whether n can be represented as a signed 16 bit integer.
func Is16Bit(n int64) bool {
	return n == int64(int16(n))
}

// IsU12Bit reports whether n can be represented as an unsigned 12 bit integer.
func IsU12Bit(n int64) bool {
	return 0 <= n && n < (1<<12)
}

// IsU16Bit reports whether n can be represented as an unsigned 16 bit integer.
func IsU16Bit(n int64) bool {
	return n == int64(uint16(n))
}

// IsU32Bit reports whether n can be represented as an unsigned 32 bit integer.
func IsU32Bit(n int64) bool {
	return n == int64(uint32(n))
}

// Is20Bit reports whether n can be represented as a signed 20 bit integer.
func Is20Bit(n int64) bool {
	return -(1<<19) <= n && n < (1<<19)
}

// b2i translates a boolean value to 0 or 1 for assigning to auxInt.
func b2i(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

// ShiftIsBounded reports whether (left/right) shift Value v is known to be bounded.
// A shift is bounded if it is shifting by less than the width of the shifted value.
func ShiftIsBounded(v *Value) bool {
	return v.AuxInt != 0
}

// truncate64Fto32F converts a float64 value to a float32 preserving the bit pattern
// of the mantissa. It will panic if the truncation results in lost information.
func truncate64Fto32F(f float64) float32 {
	if !isExactFloat32(f) {
		panic("truncate64Fto32F: truncation is not exact")
	}
	if !math.IsNaN(f) {
		return float32(f)
	}
	// NaN bit patterns aren't necessarily preserved across conversion
	// instructions so we need to do the conversion manually.
	b := math.Float64bits(f)
	m := b & ((1 << 52) - 1) // mantissa (a.k.a. significand)
	//          | sign                  | exponent   | mantissa       |
	r := uint32(((b >> 32) & (1 << 31)) | 0x7f800000 | (m >> (52 - 23)))
	return math.Float32frombits(r)
}

// extend32Fto64F converts a float32 value to a float64 value preserving the bit
// pattern of the mantissa.
func extend32Fto64F(f float32) float64 {
	if !math.IsNaN(float64(f)) {
		return float64(f)
	}
	// NaN bit patterns aren't necessarily preserved across conversion
	// instructions so we need to do the conversion manually.
	b := uint64(math.Float32bits(f))
	//   | sign                  | exponent      | mantissa                    |
	r := ((b << 32) & (1 << 63)) | (0x7ff << 52) | ((b & 0x7fffff) << (52 - 23))
	return math.Float64frombits(r)
}

// NeedsFixUp reports whether the division needs fix-up code.
func NeedsFixUp(v *Value) bool {
	return v.AuxInt == 0
}

// i2f is used in rules for converting from an AuxInt to a float.
func i2f(i int64) float64 {
	return math.Float64frombits(uint64(i))
}

// AuxFrom64F encodes a float64 value so it can be stored in an AuxInt.
func AuxFrom64F(f float64) int64 {
	return int64(math.Float64bits(f))
}

// AuxFrom32F encodes a float32 value so it can be stored in an AuxInt.
func AuxFrom32F(f float32) int64 {
	return int64(math.Float64bits(extend32Fto64F(f)))
}

// AuxTo32F decodes a float32 from the AuxInt value provided.
func AuxTo32F(i int64) float32 {
	return truncate64Fto32F(math.Float64frombits(uint64(i)))
}

// AuxTo64F decodes a float64 from the AuxInt value provided.
func AuxTo64F(i int64) float64 {
	return math.Float64frombits(uint64(i))
}

// uaddOvf reports whether unsigned a+b would overflow.
func uaddOvf(a, b int64) bool {
	return uint64(a)+uint64(b) < uint64(a)
}

// de-virtualize an InterCall
// 'sym' is the symbol for the itab
func devirt(v *Value, sym interface{}, offset int64) *obj.LSym {
	f := v.Block.Func
	n, ok := sym.(*obj.LSym)
	if !ok {
		return nil
	}
	lsym := f.Fe.DerefItab(n, offset)
	if f.pass.debug > 0 {
		if lsym != nil {
			f.Warnl(v.Pos, "de-virtualizing call")
		} else {
			f.Warnl(v.Pos, "couldn't de-virtualize call")
		}
	}
	return lsym
}

// IsSamePtr reports whether p1 and p2 point to the same address.
func IsSamePtr(p1, p2 *Value) bool {
	if p1 == p2 {
		return true
	}
	if p1.Op != p2.Op {
		return false
	}
	switch p1.Op {
	case OpOffPtr:
		return p1.AuxInt == p2.AuxInt && IsSamePtr(p1.Args[0], p2.Args[0])
	case OpAddr, OpLocalAddr:
		// OpAddr's 0th arg is either OpSP or OpSB, which means that it is uniquely identified by its Op.
		// Checking for value equality only works after [z]cse has run.
		return p1.Aux == p2.Aux && p1.Args[0].Op == p2.Args[0].Op
	case OpAddPtr:
		return p1.Args[1] == p2.Args[1] && IsSamePtr(p1.Args[0], p2.Args[0])
	}
	return false
}

func isStackPtr(v *Value) bool {
	for v.Op == OpOffPtr || v.Op == OpAddPtr {
		v = v.Args[0]
	}
	return v.Op == OpSP || v.Op == OpLocalAddr
}

// disjoint reports whether the memory region specified by [p1:p1+n1)
// does not overlap with [p2:p2+n2).
// A return value of false does not imply the regions overlap.
func disjoint(p1 *Value, n1 int64, p2 *Value, n2 int64) bool {
	if n1 == 0 || n2 == 0 {
		return true
	}
	if p1 == p2 {
		return false
	}
	baseAndOffset := func(ptr *Value) (base *Value, offset int64) {
		base, offset = ptr, 0
		for base.Op == OpOffPtr {
			offset += base.AuxInt
			base = base.Args[0]
		}
		return base, offset
	}
	p1, off1 := baseAndOffset(p1)
	p2, off2 := baseAndOffset(p2)
	if IsSamePtr(p1, p2) {
		return !overlap(off1, n1, off2, n2)
	}
	// p1 and p2 are not the same, so if they are both OpAddrs then
	// they point to different variables.
	// If one pointer is on the stack and the other is an argument
	// then they can't overlap.
	switch p1.Op {
	case OpAddr, OpLocalAddr:
		if p2.Op == OpAddr || p2.Op == OpLocalAddr || p2.Op == OpSP {
			return true
		}
		return p2.Op == OpArg && p1.Args[0].Op == OpSP
	case OpArg:
		if p2.Op == OpSP || p2.Op == OpLocalAddr {
			return true
		}
	case OpSP:
		return p2.Op == OpAddr || p2.Op == OpLocalAddr || p2.Op == OpArg || p2.Op == OpSP
	}
	return false
}

// moveSize returns the number of bytes an aligned MOV instruction moves
func MoveSize(align int64, c *Config) int64 {
	switch {
	case align%8 == 0 && c.PtrSize == 8:
		return 8
	case align%4 == 0:
		return 4
	case align%2 == 0:
		return 2
	}
	return 1
}

// MergePoint finds a block among a's blocks which dominates b and is itself
// dominated by all of a's blocks. Returns nil if it can't find one.
// Might return nil even if one does exist.
func MergePoint(b *Block, a ...*Value) *Block {
	// Walk backward from b looking for one of the a's blocks.

	// Max distance
	d := 100

	for d > 0 {
		for _, x := range a {
			if b == x.Block {
				goto found
			}
		}
		if len(b.Preds) > 1 {
			// Don't know which way to go back. Abort.
			return nil
		}
		b = b.Preds[0].b
		d--
	}
	return nil // too far away
found:
	// At this point, r is the first value in a that we find by walking backwards.
	// if we return anything, r will be it.
	r := b

	// Keep going, counting the other a's that we find. They must all dominate r.
	na := 0
	for d > 0 {
		for _, x := range a {
			if b == x.Block {
				na++
			}
		}
		if na == len(a) {
			// Found all of a in a backwards walk. We can return r.
			return r
		}
		if len(b.Preds) > 1 {
			return nil
		}
		b = b.Preds[0].b
		d--

	}
	return nil // too far away
}

// Clobber invalidates v.  Returns true.
// Clobber is used by rewrite rules to:
//   A) make sure v is really dead and never used again.
//   B) decrement use counts of v's args.
func Clobber(v *Value) bool {
	v.Reset(OpInvalid)
	// Note: leave v.Block intact.  The Block field is used after Clobber.
	return true
}

// ClobberIfDead resets v when use count is 1. Returns true.
// ClobberIfDead is used by rewrite rules to decrement
// use counts of v's args when v is dead and never used.
func ClobberIfDead(v *Value) bool {
	if v.Uses == 1 {
		v.Reset(OpInvalid)
	}
	// Note: leave v.Block intact.  The Block field is used after ClobberIfDead.
	return true
}

// noteRule is an easy way to track if a rule is matched when writing
// new ones.  Make the rule of interest also conditional on
//     noteRule("note to self: rule of interest matched")
// and that message will print when the rule matches.
func noteRule(s string) bool {
	fmt.Println(s)
	return true
}

// warnRule generates compiler debug output with string s when
// v is not in autogenerated code, cond is true and the rule has fired.
func warnRule(cond bool, v *Value, s string) bool {
	if pos := v.Pos; pos.Line() > 1 && cond {
		v.Block.Func.Warnl(pos, s)
	}
	return true
}

// for a pseudo-op like (LessThan x), extract x
func FlagArg(v *Value) *Value {
	if len(v.Args) != 1 || !v.Args[0].Type.IsFlags() {
		return nil
	}
	return v.Args[0]
}

// logRule logs the use of the rule s. This will only be enabled if
// rewrite rules were generated with the -log option, see gen/rulegen.go.
func logRule(s string) {
	if ruleFile == nil {
		// Open a log file to write log to. We open in append
		// mode because all.bash runs the compiler lots of times,
		// and we want the concatenation of all of those logs.
		// This means, of course, that users need to rm the old log
		// to get fresh data.
		// TODO: all.bash runs compilers in parallel. Need to synchronize logging somehow?
		w, err := os.OpenFile(filepath.Join(os.Getenv("GOROOT"), "src", "rulelog"),
			os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		ruleFile = w
	}
	_, err := fmt.Fprintf(ruleFile, "rewrite %s\n", s)
	if err != nil {
		panic(err)
	}
}

var ruleFile io.Writer

func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func isConstZero(v *Value) bool {
	switch v.Op {
	case OpConstNil:
		return true
	case OpConst64, OpConst32, OpConst16, OpConst8, OpConstBool, OpConst32F, OpConst64F:
		return v.AuxInt == 0
	}
	return false
}

// reciprocalExact64 reports whether 1/c is exactly representable.
func reciprocalExact64(c float64) bool {
	b := math.Float64bits(c)
	man := b & (1<<52 - 1)
	if man != 0 {
		return false // not a power of 2, denormal, or NaN
	}
	exp := b >> 52 & (1<<11 - 1)
	// exponent bias is 0x3ff.  So taking the reciprocal of a number
	// changes the exponent to 0x7fe-exp.
	switch exp {
	case 0:
		return false // ±0
	case 0x7ff:
		return false // ±inf
	case 0x7fe:
		return false // exponent is not representable
	default:
		return true
	}
}

// reciprocalExact32 reports whether 1/c is exactly representable.
func reciprocalExact32(c float32) bool {
	b := math.Float32bits(c)
	man := b & (1<<23 - 1)
	if man != 0 {
		return false // not a power of 2, denormal, or NaN
	}
	exp := b >> 23 & (1<<8 - 1)
	// exponent bias is 0x7f.  So taking the reciprocal of a number
	// changes the exponent to 0xfe-exp.
	switch exp {
	case 0:
		return false // ±0
	case 0xff:
		return false // ±inf
	case 0xfe:
		return false // exponent is not representable
	default:
		return true
	}
}

// check if an immediate can be directly encoded into an ARM's instruction
func IsARMImmRot(v uint32) bool {
	for i := 0; i < 16; i++ {
		if v&^0xff == 0 {
			return true
		}
		v = v<<2 | v>>30
	}

	return false
}

// overlap reports whether the ranges given by the given offset and
// size pairs overlap.
func overlap(offset1, size1, offset2, size2 int64) bool {
	if offset1 >= offset2 && offset2+size2 > offset1 {
		return true
	}
	if offset2 >= offset1 && offset1+size1 > offset2 {
		return true
	}
	return false
}

func AreAdjacentOffsets(off1, off2, size int64) bool {
	return off1+size == off2 || off1 == off2+size
}

// check if value zeroes out upper 32-bit of 64-bit register.
// depth limits recursion depth. In AMD64.rules 3 is used as limit,
// because it catches same amount of cases as 4.
func ZeroUpper32Bits(x *Value, depth int) bool {
	switch x.Op {
	case OpAMD64MOVLconst, OpAMD64MOVLload, OpAMD64MOVLQZX, OpAMD64MOVLloadidx1,
		OpAMD64MOVWload, OpAMD64MOVWloadidx1, OpAMD64MOVBload, OpAMD64MOVBloadidx1,
		OpAMD64MOVLloadidx4, OpAMD64ADDLload, OpAMD64SUBLload, OpAMD64ANDLload,
		OpAMD64ORLload, OpAMD64XORLload, OpAMD64CVTTSD2SL,
		OpAMD64ADDL, OpAMD64ADDLconst, OpAMD64SUBL, OpAMD64SUBLconst,
		OpAMD64ANDL, OpAMD64ANDLconst, OpAMD64ORL, OpAMD64ORLconst,
		OpAMD64XORL, OpAMD64XORLconst, OpAMD64NEGL, OpAMD64NOTL:
		return true
	case OpArg:
		return x.Type.Width == 4
	case OpPhi, OpSelect0, OpSelect1:
		// Phis can use each-other as an arguments, instead of tracking visited values,
		// just limit recursion depth.
		if depth <= 0 {
			return false
		}
		for i := range x.Args {
			if !ZeroUpper32Bits(x.Args[i], depth-1) {
				return false
			}
		}
		return true

	}
	return false
}

// ZeroUpper48Bits is similar to ZeroUpper32Bits, but for upper 48 bits
func ZeroUpper48Bits(x *Value, depth int) bool {
	switch x.Op {
	case OpAMD64MOVWQZX, OpAMD64MOVWload, OpAMD64MOVWloadidx1, OpAMD64MOVWloadidx2:
		return true
	case OpArg:
		return x.Type.Width == 2
	case OpPhi, OpSelect0, OpSelect1:
		// Phis can use each-other as an arguments, instead of tracking visited values,
		// just limit recursion depth.
		if depth <= 0 {
			return false
		}
		for i := range x.Args {
			if !ZeroUpper48Bits(x.Args[i], depth-1) {
				return false
			}
		}
		return true

	}
	return false
}

// ZeroUpper56Bits is similar to ZeroUpper32Bits, but for upper 56 bits
func ZeroUpper56Bits(x *Value, depth int) bool {
	switch x.Op {
	case OpAMD64MOVBQZX, OpAMD64MOVBload, OpAMD64MOVBloadidx1:
		return true
	case OpArg:
		return x.Type.Width == 1
	case OpPhi, OpSelect0, OpSelect1:
		// Phis can use each-other as an arguments, instead of tracking visited values,
		// just limit recursion depth.
		if depth <= 0 {
			return false
		}
		for i := range x.Args {
			if !ZeroUpper56Bits(x.Args[i], depth-1) {
				return false
			}
		}
		return true

	}
	return false
}

// isInlinableMemmove reports whether the given arch performs a Move of the given size
// faster than memmove. It will only return true if replacing the memmove with a Move is
// safe, either because Move is small or because the arguments are disjoint.
// This is used as a check for replacing memmove with Move ops.
func isInlinableMemmove(dst, src *Value, sz int64, c *Config) bool {
	// It is always safe to convert memmove into Move when its arguments are disjoint.
	// Move ops may or may not be faster for large sizes depending on how the platform
	// lowers them, so we only perform this optimization on platforms that we know to
	// have fast Move ops.
	switch c.arch {
	case "amd64", "amd64p32":
		return sz <= 16 || (sz < 1024 && disjoint(dst, sz, src, sz))
	case "386", "ppc64", "ppc64le", "arm64":
		return sz <= 8
	case "s390x":
		return sz <= 8 || disjoint(dst, sz, src, sz)
	case "arm", "mips", "mips64", "mipsle", "mips64le":
		return sz <= 4
	}
	return false
}

// SizeOf returns the size of t in bytes.
// It will panic if t is not a *types.Type.
func SizeOf(t interface{}) int64 {
	return t.(*types.Type).Size()
}

// alignof returns the alignment of t in bytes.
// It will panic if t is not a *types.Type.
func alignof(t interface{}) int64 {
	return t.(*types.Type).Alignment()
}

// registerizable reports whether t is a primitive type that fits in
// a register. It assumes float64 values will always fit into registers
// even if that isn't strictly true.
// It will panic if t is not a *types.Type.
func registerizable(b *Block, t interface{}) bool {
	typ := t.(*types.Type)
	if typ.IsPtrShaped() || typ.IsFloat() {
		return true
	}
	if typ.IsInteger() {
		return typ.Size() <= b.Func.Config.RegSize
	}
	return false
}

// needRaceCleanup reports whether this call to racefuncenter/exit isn't needed.
func needRaceCleanup(sym interface{}, v *Value) bool {
	f := v.Block.Func
	if !f.Config.Race {
		return false
	}
	if !isSameSym(sym, "runtime.racefuncenter") && !isSameSym(sym, "runtime.racefuncexit") {
		return false
	}
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if v.Op == OpStaticCall {
				switch v.Aux.(fmt.Stringer).String() {
				case "runtime.racefuncenter", "runtime.racefuncexit", "runtime.panicindex",
					"runtime.panicslice", "runtime.panicdivide", "runtime.panicwrap":
				// Check for racefuncenter will encounter racefuncexit and vice versa.
				// Allow calls to panic*
				default:
					// If we encountered any call, we need to keep racefunc*,
					// for accurate stacktraces.
					return false
				}
			}
		}
	}
	return true
}

// SymIsRO reports whether sym is a read-only global.
func SymIsRO(sym interface{}) bool {
	lsym := sym.(*obj.LSym)
	return lsym.Type == objabi.SRODATA && len(lsym.R) == 0
}

// Read8 reads one byte from the read-only global sym at offset off.
func Read8(sym interface{}, off int64) uint8 {
	lsym := sym.(*obj.LSym)
	return lsym.P[off]
}

// Read16 reads two bytes from the read-only global sym at offset off.
func Read16(sym interface{}, off int64, bigEndian bool) uint16 {
	lsym := sym.(*obj.LSym)
	if bigEndian {
		return binary.BigEndian.Uint16(lsym.P[off:])
	} else {
		return binary.LittleEndian.Uint16(lsym.P[off:])
	}
}

// Read32 reads four bytes from the read-only global sym at offset off.
func Read32(sym interface{}, off int64, bigEndian bool) uint32 {
	lsym := sym.(*obj.LSym)
	if bigEndian {
		return binary.BigEndian.Uint32(lsym.P[off:])
	} else {
		return binary.LittleEndian.Uint32(lsym.P[off:])
	}
}

// Read64 reads eight bytes from the read-only global sym at offset off.
func Read64(sym interface{}, off int64, bigEndian bool) uint64 {
	lsym := sym.(*obj.LSym)
	if bigEndian {
		return binary.BigEndian.Uint64(lsym.P[off:])
	} else {
		return binary.LittleEndian.Uint64(lsym.P[off:])
	}
}
