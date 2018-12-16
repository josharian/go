// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rewrite

import (
	"cmd/compile/internal/ssa"
	"math/bits"
)

// bfAuxInt encodes the lsb and width for arm64 bitfield ops into the expected auxInt format.
func bfAuxInt(lsb, width int64) int64 {
	if lsb < 0 || lsb > 63 {
		panic("ARM64 bit field lsb constant out of range")
	}
	if width < 1 || width > 64 {
		panic("ARM64 bit field width constant out of range")
	}
	return width | lsb<<8
}

// bfLSB returns the lsb part of the auxInt field of arm64 bitfield ops.
func bfLSB(bfc int64) int64 {
	return int64(uint64(bfc) >> 8)
}

// getBFWidth returns the width part of the auxInt field of arm64 bitfield ops.
func getBFWidth(bfc int64) int64 {
	return bfc & 0xff
}

// isBFMask reports whether mask >> rshift applied at lsb is a valid arm64 bitfield op mask.
func isBFMask(lsb, mask, rshift int64) bool {
	shiftedMask := int64(uint64(mask) >> uint64(rshift))
	return shiftedMask != 0 && ssa.IsPowerOfTwo(shiftedMask+1) && nto(shiftedMask)+lsb < 64
}

// bfWidth returns the bitfield width of mask >> rshift for arm64 bitfield ops
func bfWidth(mask, rshift int64) int64 {
	shiftedMask := int64(uint64(mask) >> uint64(rshift))
	if shiftedMask == 0 {
		panic("ARM64 BF mask is zero")
	}
	return nto(shiftedMask)
}

// negate finds the complement to an ARM64 condition code,
// for example Equal -> NotEqual or LessThan -> GreaterEqual
//
// TODO: add floating-point conditions
func negate(op ssa.Op) ssa.Op {
	switch op {
	case ssa.OpARM64LessThan:
		return ssa.OpARM64GreaterEqual
	case ssa.OpARM64LessThanU:
		return ssa.OpARM64GreaterEqualU
	case ssa.OpARM64GreaterThan:
		return ssa.OpARM64LessEqual
	case ssa.OpARM64GreaterThanU:
		return ssa.OpARM64LessEqualU
	case ssa.OpARM64LessEqual:
		return ssa.OpARM64GreaterThan
	case ssa.OpARM64LessEqualU:
		return ssa.OpARM64GreaterThanU
	case ssa.OpARM64GreaterEqual:
		return ssa.OpARM64LessThan
	case ssa.OpARM64GreaterEqualU:
		return ssa.OpARM64LessThanU
	case ssa.OpARM64Equal:
		return ssa.OpARM64NotEqual
	case ssa.OpARM64NotEqual:
		return ssa.OpARM64Equal
	default:
		panic("unreachable")
	}
}

// invert evaluates (InvertFlags op), which
// is the same as altering the condition codes such
// that the same result would be produced if the arguments
// to the flag-generating instruction were reversed, e.g.
// (InvertFlags (CMP x y)) -> (CMP y x)
//
// TODO: add floating-point conditions
func invert(op ssa.Op) ssa.Op {
	switch op {
	case ssa.OpARM64LessThan:
		return ssa.OpARM64GreaterThan
	case ssa.OpARM64LessThanU:
		return ssa.OpARM64GreaterThanU
	case ssa.OpARM64GreaterThan:
		return ssa.OpARM64LessThan
	case ssa.OpARM64GreaterThanU:
		return ssa.OpARM64LessThanU
	case ssa.OpARM64LessEqual:
		return ssa.OpARM64GreaterEqual
	case ssa.OpARM64LessEqualU:
		return ssa.OpARM64GreaterEqualU
	case ssa.OpARM64GreaterEqual:
		return ssa.OpARM64LessEqual
	case ssa.OpARM64GreaterEqualU:
		return ssa.OpARM64LessEqualU
	case ssa.OpARM64Equal, ssa.OpARM64NotEqual:
		return op
	default:
		panic("unreachable")
	}
}

// ccEval evaluates an ARM64 op against a flags value
// that is potentially constant; return 1 for true,
// -1 for false, and 0 for not constant.
func ccEval(cc interface{}, flags *ssa.Value) int {
	op := cc.(ssa.Op)
	fop := flags.Op
	switch fop {
	case ssa.OpARM64InvertFlags:
		return -ccEval(op, flags.Args[0])
	case ssa.OpARM64FlagEQ:
		switch op {
		case ssa.OpARM64Equal, ssa.OpARM64GreaterEqual, ssa.OpARM64LessEqual,
			ssa.OpARM64GreaterEqualU, ssa.OpARM64LessEqualU:
			return 1
		default:
			return -1
		}
	case ssa.OpARM64FlagLT_ULT:
		switch op {
		case ssa.OpARM64LessThan, ssa.OpARM64LessThanU,
			ssa.OpARM64LessEqual, ssa.OpARM64LessEqualU:
			return 1
		default:
			return -1
		}
	case ssa.OpARM64FlagLT_UGT:
		switch op {
		case ssa.OpARM64LessThan, ssa.OpARM64GreaterThanU,
			ssa.OpARM64LessEqual, ssa.OpARM64GreaterEqualU:
			return 1
		default:
			return -1
		}
	case ssa.OpARM64FlagGT_ULT:
		switch op {
		case ssa.OpARM64GreaterThan, ssa.OpARM64LessThanU,
			ssa.OpARM64GreaterEqual, ssa.OpARM64LessEqualU:
			return 1
		default:
			return -1
		}
	case ssa.OpARM64FlagGT_UGT:
		switch op {
		case ssa.OpARM64GreaterThan, ssa.OpARM64GreaterThanU,
			ssa.OpARM64GreaterEqual, ssa.OpARM64GreaterEqualU:
			return 1
		default:
			return -1
		}
	default:
		return 0
	}
}

func oneBit(x int64) bool {
	return bits.OnesCount64(uint64(x)) == 1
}

// nto returns the number of trailing ones.
func nto(x int64) int64 {
	return ssa.Ntz(^x)
}
