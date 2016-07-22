// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// from http://research.swtch.com/sparse
// in turn, from Briggs and Torczon

type sparseLiveInfoMap struct {
	dense  []liveInfo
	sparse []int32
	// copy-on-write support
	densecopy []liveInfo
	usecopy   bool
}

// newSparseLiveInfoMap returns a sparseLiveInfoMap
// that maps from IDs to distances.
func newSparseLiveInfoMap(n int) *sparseLiveInfoMap {
	return &sparseLiveInfoMap{sparse: make([]int32, n)}
}

func (s *sparseLiveInfoMap) size() int {
	return len(s.contents())
}

func (s *sparseLiveInfoMap) contains(k ID) bool {
	i := s.sparse[k]
	dense := s.contents()
	return i < int32(len(dense)) && dense[i].ID == k
}

// get returns the dist for ID k, or -1 if k does not appear in the map.
func (s *sparseLiveInfoMap) get(k ID) int32 {
	i := s.sparse[k]
	dense := s.contents()
	if i < int32(len(dense)) && dense[i].ID == k {
		return dense[i].dist
	}
	return -1
}

func (s *sparseLiveInfoMap) setto(all []liveInfo) {
	s.usecopy = true
	s.densecopy = all
	for i, e := range all {
		s.sparse[e.ID] = int32(i)
	}
}

func (s *sparseLiveInfoMap) copyForWrite() {
	if !s.usecopy {
		// already copied
		return
	}
	for cap(s.dense) < len(s.densecopy) {
		s.dense = s.dense[:len(s.dense)]
		s.dense = append(s.dense, liveInfo{})
	}
	s.dense = s.dense[:len(s.densecopy)]
	copy(s.dense, s.densecopy)
	s.usecopy = false
}

func (s *sparseLiveInfoMap) set(k ID, v int32) {
	s.copyForWrite()
	i := s.sparse[k]
	if i < int32(len(s.dense)) && s.dense[i].ID == k {
		s.dense[i].dist = v
		return
	}
	s.dense = append(s.dense, liveInfo{k, v})
	s.sparse[k] = int32(len(s.dense)) - 1
}

func (s *sparseLiveInfoMap) contents() []liveInfo {
	if s.usecopy {
		return s.densecopy
	}
	return s.dense
}
