// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rewrite

import (
	"cmd/compile/internal/ssa"
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

// nto returns the number of trailing ones.
func nto(x int64) int64 {
	return ssa.Ntz(^x)
}
