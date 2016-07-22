// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// from http://research.swtch.com/sparse
// in turn, from Briggs and Torczon

type sparseLiveInfoMap struct {
	dense  []liveInfo
	sparse []int32
}

// newSparseLiveInfoMap returns a sparseLiveInfoMap
// that maps from IDs to distances.
func newSparseLiveInfoMap(n int) *sparseLiveInfoMap {
	return &sparseLiveInfoMap{
		dense:  nil,
		sparse: make([]int32, n),
	}
}

func (s *sparseLiveInfoMap) size() int {
	return len(s.dense)
}

func (s *sparseLiveInfoMap) contains(k ID) bool {
	i := s.sparse[k]
	return i < int32(len(s.dense)) && s.dense[i].ID == k
}

// get returns the dist for ID k, or -1 if k does not appear in the map.
func (s *sparseLiveInfoMap) get(k ID) int32 {
	i := s.sparse[k]
	if i < int32(len(s.dense)) && s.dense[i].ID == k {
		return s.dense[i].dist
	}
	return -1
}

func (s *sparseLiveInfoMap) setto(all []liveInfo) {
	for cap(s.dense) < len(all) {
		s.dense = s.dense[:len(s.dense)]
		s.dense = append(s.dense, liveInfo{})
	}
	s.dense = s.dense[:len(all)]
	copy(s.dense, all)
	for i, e := range all {
		s.sparse[e.ID] = int32(i)
	}
	return
}

func (s *sparseLiveInfoMap) set(k ID, v int32) {
	i := s.sparse[k]
	if i < int32(len(s.dense)) && s.dense[i].ID == k {
		s.dense[i].dist = v
		return
	}
	s.dense = append(s.dense, liveInfo{k, v})
	s.sparse[k] = int32(len(s.dense)) - 1
}

func (s *sparseLiveInfoMap) contents() []liveInfo {
	return s.dense
}
