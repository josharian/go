// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// from http://research.swtch.com/sparse
// in turn, from Briggs and Torczon

type sparseSet struct {
	dense  []ID
	sparse []ID
}

// newSparseSet returns a sparseSet that can represent
// integers between 0 and n-1.
func newSparseSet(n int) *sparseSet {
	s := new(sparseSet)
	// TODO: grow dense initially?
	s.grow(n)
	return s
}

// grow increases the size of s to accommodate integers in [0, n).
func (s *sparseSet) grow(n int) {
	for len(s.sparse) < n {
		s.sparse = append(s.sparse, 0)
	}
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
	return int(s.sparse[x])
}

func (s *sparseSet) setSparseAt(x ID, v int) {
	u := ID(v)
	if int(u) != v {
		panic("setSparseAt overflow")
	}
	s.sparse[x] = u
}
