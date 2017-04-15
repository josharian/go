// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"cmd/compile/internal/ssa"
	"cmd/compile/internal/types"
	"cmd/internal/dwarf"
	"cmd/internal/obj"
	"cmd/internal/src"
	"cmd/internal/sys"
	"fmt"
	"sort"
	"sync"
)

// "Portable" code generation.

var (
	ncpu         int            // the number of concurrent backend compiles, set by a compiler flag
	needscompile []*Node        // slice of functions waiting to be compiled
	compilenow   bool           // indicates whether to compile immediately or enqueue in needscompile
	compilewg    sync.WaitGroup // wait for all backend compilers to complete
	compilec     chan *Node     // channel of functions for backend compilers to drain
)

func emitptrargsmap() {
	if Curfn.Func.Nname.Sym.Name == "_" {
		return
	}
	sym := lookup(fmt.Sprintf("%s.args_stackmap", Curfn.Func.Nname.Sym.Name))
	lsym := Linksym(sym)

	nptr := int(Curfn.Type.ArgWidth() / int64(Widthptr))
	bv := bvalloc(int32(nptr) * 2)
	nbitmap := 1
	if Curfn.Type.Results().NumFields() > 0 {
		nbitmap = 2
	}
	off := duint32LSym(lsym, 0, uint32(nbitmap))
	off = duint32LSym(lsym, off, uint32(bv.n))
	var xoffset int64
	if Curfn.IsMethod() {
		xoffset = 0
		onebitwalktype1(Curfn.Type.Recvs(), &xoffset, bv)
	}

	if Curfn.Type.Params().NumFields() > 0 {
		xoffset = 0
		onebitwalktype1(Curfn.Type.Params(), &xoffset, bv)
	}

	off = dbvecLSym(lsym, off, bv)
	if Curfn.Type.Results().NumFields() > 0 {
		xoffset = 0
		onebitwalktype1(Curfn.Type.Results(), &xoffset, bv)
		off = dbvecLSym(lsym, off, bv)
	}

	ggloblLSym(lsym, int32(off), obj.RODATA|obj.LOCAL)
}

// cmpstackvarlt reports whether the stack variable a sorts before b.
//
// Sort the list of stack variables. Autos after anything else,
// within autos, unused after used, within used, things with
// pointers first, zeroed things first, and then decreasing size.
// Because autos are laid out in decreasing addresses
// on the stack, pointers first, zeroed things first and decreasing size
// really means, in memory, things with pointers needing zeroing at
// the top of the stack and increasing in size.
// Non-autos sort on offset.
func cmpstackvarlt(a, b *Node) bool {
	if (a.Class == PAUTO) != (b.Class == PAUTO) {
		return b.Class == PAUTO
	}

	if a.Class != PAUTO {
		return a.Xoffset < b.Xoffset
	}

	if a.Used() != b.Used() {
		return a.Used()
	}

	ap := types.Haspointers(a.Type)
	bp := types.Haspointers(b.Type)
	if ap != bp {
		return ap
	}

	ap = a.Name.Needzero()
	bp = b.Name.Needzero()
	if ap != bp {
		return ap
	}

	if a.Type.Width != b.Type.Width {
		return a.Type.Width > b.Type.Width
	}

	return a.Sym.Name < b.Sym.Name
}

// byStackvar implements sort.Interface for []*Node using cmpstackvarlt.
type byStackVar []*Node

func (s byStackVar) Len() int           { return len(s) }
func (s byStackVar) Less(i, j int) bool { return cmpstackvarlt(s[i], s[j]) }
func (s byStackVar) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *ssafn) AllocFrame(f *ssa.Func) {
	s.stksize = 0
	s.stkptrsize = 0
	fn := s.curfn.Func

	// Mark the PAUTO's unused.
	for _, ln := range fn.Dcl {
		if ln.Class == PAUTO {
			ln.SetUsed(false)
		}
	}

	for _, l := range f.RegAlloc {
		if ls, ok := l.(ssa.LocalSlot); ok {
			ls.N.(*Node).SetUsed(true)
		}
	}

	scratchUsed := false
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			switch a := v.Aux.(type) {
			case *ssa.ArgSymbol:
				n := a.Node.(*Node)
				// Don't modify nodfp; it is a global.
				if n != nodfp {
					n.SetUsed(true)
				}
			case *ssa.AutoSymbol:
				a.Node.(*Node).SetUsed(true)
			}

			if !scratchUsed {
				scratchUsed = v.Op.UsesScratch()
			}
		}
	}

	if f.Config.NeedsFpScratch && scratchUsed {
		s.scratchFpMem = tempAt(src.NoXPos, s.curfn, types.Types[TUINT64])
	}

	sort.Sort(byStackVar(fn.Dcl))

	// Reassign stack offsets of the locals that are used.
	for i, n := range fn.Dcl {
		if n.Op != ONAME || n.Class != PAUTO {
			continue
		}
		if !n.Used() {
			fn.Dcl = fn.Dcl[:i]
			break
		}

		dowidth(n.Type)
		w := n.Type.Width
		if w >= thearch.MAXWIDTH || w < 0 {
			Fatalf("bad width")
		}
		s.stksize += w
		s.stksize = Rnd(s.stksize, int64(n.Type.Align))
		if types.Haspointers(n.Type) {
			s.stkptrsize = s.stksize
		}
		if thearch.LinkArch.InFamily(sys.MIPS, sys.MIPS64, sys.ARM, sys.ARM64, sys.PPC64, sys.S390X) {
			s.stksize = Rnd(s.stksize, int64(Widthptr))
		}
		n.Xoffset = -s.stksize
	}

	s.stksize = Rnd(s.stksize, int64(Widthreg))
	s.stkptrsize = Rnd(s.stkptrsize, int64(Widthreg))
}

func compile(fn *Node) {
	Curfn = fn
	dowidth(fn.Type)

	if fn.Nbody.Len() == 0 {
		emitptrargsmap()
		return
	}

	saveerrors()

	order(fn)
	if nerrors != 0 {
		return
	}

	walk(fn)
	if nerrors != 0 {
		return
	}
	checkcontrolflow(fn)
	if nerrors != 0 {
		return
	}
	if instrumenting {
		instrument(fn)
	}

	// From this point, there should be no uses of Curfn. Enforce that.
	Curfn = nil

	// Set up the function's LSym early to avoid data races with the assemblers.
	fn.Func.initLSym()

	if compilenow {
		compileSSA(fn, 0)
	} else {
		needscompile = append(needscompile, fn)
	}
}

// compileSSA builds an SSA backend function,
// uses it to generate a plist,
// and flushes that plist to machine code.
func compileSSA(fn *Node, shard int) {
	cache := &ssaCaches[shard]
	ssafn := buildssa(fn, cache)
	pp := newProgs(fn, shard)
	genssa(ssafn, pp)
	fieldtrack(pp.Text.From.Sym, fn.Func.FieldTrack)
	if pp.Text.To.Offset < 1<<31 {
		pp.Flush()
	} else {
		largeStackFrames = append(largeStackFrames, fn.Pos)
	}
	pp.Free()
}

func debuginfo(fnsym *obj.LSym, curfn interface{}) []*dwarf.Var {
	fn := curfn.(*Node)
	if expect := Linksym(fn.Func.Nname.Sym); fnsym != expect {
		Fatalf("unexpected fnsym: %v != %v", fnsym, expect)
	}

	var vars []*dwarf.Var
	for _, n := range fn.Func.Dcl {
		if n.Op != ONAME { // might be OTYPE or OLITERAL
			continue
		}

		var name obj.AddrName
		var abbrev int
		offs := n.Xoffset

		switch n.Class {
		case PAUTO:
			if !n.Used() {
				Fatalf("debuginfo unused node (AllocFrame should truncate fn.Func.Dcl)")
			}
			name = obj.NAME_AUTO

			abbrev = dwarf.DW_ABRV_AUTO
			if Ctxt.FixedFrameSize() == 0 {
				offs -= int64(Widthptr)
			}
			if obj.Framepointer_enabled(obj.GOOS, obj.GOARCH) {
				offs -= int64(Widthptr)
			}

		case PPARAM, PPARAMOUT:
			name = obj.NAME_PARAM

			abbrev = dwarf.DW_ABRV_PARAM
			offs += Ctxt.FixedFrameSize()

		default:
			continue
		}

		gotype := Linksym(ngotype(n))
		fnsym.Autom = append(fnsym.Autom, &obj.Auto{
			Asym:    n.Sym.Name,
			Aoffset: int32(n.Xoffset),
			Name:    name,
			Gotype:  gotype,
		})

		if n.IsAutoTmp() {
			continue
		}

		typename := dwarf.InfoPrefix + gotype.Name[len("type."):]
		vars = append(vars, &dwarf.Var{
			Name:   n.Sym.Name,
			Abbrev: abbrev,
			Offset: int32(offs),
			Type:   Ctxt.Lookup(typename, 0),
		})
	}

	// Stable sort so that ties are broken with declaration order.
	sort.Stable(dwarf.VarsByOffset(vars))

	return vars
}

// fieldtrack adds R_USEFIELD relocations to fnsym to record any
// struct fields that it used.
func fieldtrack(fnsym *obj.LSym, tracked map[*types.Sym]struct{}) {
	if fnsym == nil {
		return
	}
	if obj.Fieldtrack_enabled == 0 || len(tracked) == 0 {
		return
	}

	trackSyms := make([]*types.Sym, 0, len(tracked))
	for sym := range tracked {
		trackSyms = append(trackSyms, sym)
	}
	sort.Sort(symByName(trackSyms))
	for _, sym := range trackSyms {
		r := obj.Addrel(fnsym)
		r.Sym = Linksym(sym)
		r.Type = obj.R_USEFIELD
	}
}

type symByName []*types.Sym

func (a symByName) Len() int           { return len(a) }
func (a symByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a symByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
