// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rewrite

import "math/bits"

// log2uint32 returns logarithm in base 2 of uint32(n), with Log2(0) = -1.
// Rounds down.
func log2uint32(n int64) int64 {
	return int64(bits.Len32(uint32(n))) - 1
}

// isUint64PowerOfTwo reports whether uint64(n) is a power of 2.
func isUint64PowerOfTwo(in int64) bool {
	n := uint64(in)
	return n > 0 && n&(n-1) == 0
}

// isUint32PowerOfTwo reports whether uint32(n) is a power of 2.
func isUint32PowerOfTwo(in int64) bool {
	n := uint64(uint32(in))
	return n > 0 && n&(n-1) == 0
}
