// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import "strconv"

// from http://research.swtch.com/sparse
// in turn, from Briggs and Torczon

type sparseSet struct {
	dense []ID
	// exactly one of the sparse slices will be non-nil
	sparse8  *[256]uint8
	sparse16 []uint16
	sparse32 []int32
}

// newSparseSet returns a sparseSet that can represent
// integers between 0 and n-1
func newSparseSet(n int) *sparseSet {
	if n < 1<<8 {
		return &sparseSet{dense: nil, sparse8: new([256]uint8)}
	}
	if n < 1<<16 {
		return &sparseSet{dense: nil, sparse16: make([]uint16, n)}
	}
	return &sparseSet{dense: nil, sparse32: make([]int32, n)}
}

func (s *sparseSet) cap() int {
	if s.sparse8 != nil {
		return 256
	}
	if s.sparse16 != nil {
		return len(s.sparse16)
	}
	return len(s.sparse32)
}

func (s *sparseSet) size() int {
	return len(s.dense)
}

func (s *sparseSet) contains(x ID) bool {
	i := s.sparseAt(x)
	return i < len(s.dense) && s.dense[i] == x
}

func (s *sparseSet) add(x ID) {
	i := s.sparseAt(x)
	if i < len(s.dense) && s.dense[i] == x {
		return
	}
	s.dense = append(s.dense, x)
	s.setSparseAt(x, len(s.dense)-1)
}

func (s *sparseSet) addAll(a []ID) {
	for _, x := range a {
		s.add(x)
	}
}

func (s *sparseSet) addAllValues(a []*Value) {
	for _, v := range a {
		s.add(v.ID)
	}
}

func (s *sparseSet) remove(x ID) {
	i := s.sparseAt(x)
	if i < len(s.dense) && s.dense[i] == x {
		y := s.dense[len(s.dense)-1]
		s.dense[i] = y
		s.setSparseAt(y, i)
		s.dense = s.dense[:len(s.dense)-1]
	}
}

// pop removes an arbitrary element from the set.
// The set must be nonempty.
func (s *sparseSet) pop() ID {
	x := s.dense[len(s.dense)-1]
	s.dense = s.dense[:len(s.dense)-1]
	return x
}

func (s *sparseSet) clear() {
	s.dense = s.dense[:0]
}

func (s *sparseSet) contents() []ID {
	return s.dense
}

func (s *sparseSet) sparseAt(x ID) int {
	if s.sparse8 != nil {
		return int(s.sparse8[x])
	}
	if s.sparse16 != nil {
		return int(s.sparse16[x])
	}
	return int(s.sparse32[x])
}

func (s *sparseSet) setSparseAt(x ID, v int) {
	if s.sparse8 != nil {
		u := uint8(v)
		if int(u) != v {
			panic("bad sparse8 index " + strconv.Itoa(v))
		}
		s.sparse8[x] = u
		return
	}
	if s.sparse16 != nil {
		u := uint16(v)
		if int(u) != v {
			panic("bad sparse16 index " + strconv.Itoa(v))
		}
		s.sparse16[x] = u
		return
	}
	u := int32(v)
	if int(u) != v {
		panic("bad sparse32 index " + strconv.Itoa(v))
	}
	s.sparse32[x] = u
}
