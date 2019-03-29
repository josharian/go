// TODO: copyright blah blah blah

// +build gofuzz
// +build !math_big_pure_go

package big

import (
	"encoding/binary"
	"fmt"
)

// TODO: think about our interpretation of data...
// depends on size of uint. could depend on native endianness.
// depends on architecture (for which asm we run).
// how do we account for this in our corpus and workdir?

func FuzzMulWW(data []byte) int {
	var x, y Word
	if _W == 32 {
		if len(data) < 8 {
			return 0
		}
		x = Word(binary.LittleEndian.Uint32(data))
		y = Word(binary.LittleEndian.Uint32(data[4:]))
	} else {
		if len(data) < 16 {
			return 0
		}
		x = Word(binary.LittleEndian.Uint64(data))
		y = Word(binary.LittleEndian.Uint64(data[8:]))
	}
	z1, z0 := mulWW(x, y)
	z1g, z0g := mulWW_g(x, y)
	if z1 != z1g || z0 != z0g {
		fmt.Println("x", x, "y", y, "z1", z1, "z0", z0, "z1g", z1g, "z0g", z0g)
		panic("mulWW mismatch")
	}
	return 0
}

// func divWW(x1, x0, y Word) (q, r Word)
// func addVV(z, x, y []Word) (c Word)
// func subVV(z, x, y []Word) (c Word)
// func addVW(z, x []Word, y Word) (c Word)
// func subVW(z, x []Word, y Word) (c Word)

func copyV(x []Word) []Word {
	if x == nil {
		return nil
	}
	c := make([]Word, len(x))
	copy(c, x)
	return c
}

func assertEqualV(x, y []Word) {
	if len(x) != len(y) {
		fmt.Println("x", x, "y", y)
		panic("mismatched lengths")
	}
	for i := range x {
		if x[i] != y[i] {
			fmt.Println("x", x, "y", y, "i", i, "x[i]", x[i], "y[i]", y[i])
			panic("mismatched elem")
		}
	}
}

func FuzzShlVU(data []byte) int {
	if len(data) < 3 {
		return 0
	}
	s := uint(data[0])
	cut1 := int(data[1])
	cut2 := int(data[2])
	data = data[3:]
	var v []Word
	if _W == 32 {
		v = make([]Word, len(v)/4)
		for i := range v {
			v[i] = Word(binary.LittleEndian.Uint32(data[i*4 : i*4+4]))
		}
	} else {
		v = make([]Word, len(v)/8)
		for i := range v {
			v[i] = Word(binary.LittleEndian.Uint64(data[i*8 : i*8+8]))
		}
	}

	// check nil z. this should always be a no-op.
	v2 := copyV(v)

	c := shlVU(nil, v2, s)
	if c != 0 {
		panic("want 0 after shifting nil (asm)")
	}
	assertEqualV(v, v2)

	c = shlVU_g(nil, v2, s)
	if c != 0 {
		panic("want 0 after shifting nil (Go)")
	}
	assertEqualV(v, v2)

	// check complete overlap
	c = shlVU(v2, v2, s) // v2 is still pristine
	v3 := copyV(v)
	c2 := shlVU_g(v3, v3, s)
	if c != c2 {
		fmt.Println("complete overlap c", c, "c2", c2)
	}
	assertEqualV(v2, v3)

	// check no overlap. if the slices have uneven length, the first arg must be longer.
	copy(v2, v)
	copy(v3, v)
	mid := len(v)
	if len(v)%2 == 1 {
		mid++
	}
	c = shlVU(v2[:mid], v2[mid:], s)
	c2 = shlVU_g(v3[:mid], v3[mid:], s)
	if c != c2 {
		fmt.Println("no overlap c", c, "c2", c2)
	}
	assertEqualV(v2, v3)

	// check random overlap. again, the first arg must be longer if either is.
	// we want to use subslices [:cut1] and [cut2:], but we need to ensure that
	// the second slice is not longer than the first, so we use [cut2:cut3],
	// and select cut3 to keep the second slice short enough.
	if len(v) == 0 {
		// covered above; bailing here makes math below nicer
		return 0
	}
	copy(v2, v)
	copy(v3, v)
	if cut1 > len(v)-1 {
		cut1 = len(v) - 1
	}
	if cut2 > len(v)-1 {
		cut2 = len(v) - 1
	}
	cut3 := len(v) - 1
	if cut3-cut2 > cut1 {
		cut3 = cut1 + cut2
	}
	c = shlVU(v2[:cut1], v2[cut2:cut3], s)
	c2 = shlVU_g(v3[:cut1], v3[cut2:cut3], s)
	if c != c2 {
		fmt.Println("no overlap c", c, "c2", c2)
	}
	assertEqualV(v2, v3)

	return 0
}

// func shrVU(z, x []Word, s uint) (c Word)
// func mulAddVWW(z, x []Word, y, r Word) (c Word)
// func addMulVVW(z, x []Word, y Word) (c Word)
// func divWVW(z []Word, xn Word, x []Word, y Word) (r Word)
