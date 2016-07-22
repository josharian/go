// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// from http://research.swtch.com/sparse
// in turn, from Briggs and Torczon

type sparseLiveInfoMap struct {
	dense  []liveInfo
	extra  []liveInfo
	sparse []int32
}

// newSparseLiveInfoMap returns a sparseLiveInfoMap that maps from IDs to distances.
// The caller must provide their own dense slice, which will be modified directly.
func newSparseLiveInfoMap(n int) *sparseLiveInfoMap {
	return &sparseLiveInfoMap{sparse: make([]int32, n)}
}

// get returns the dist for ID k and reports whether it succeeded.
func (s *sparseLiveInfoMap) get(k ID) (int32, bool) {
	i := int(s.sparse[k])
	if i < len(s.dense) && s.dense[i].ID == k {
		return s.dense[i].dist, true
	}
	i -= len(s.dense)
	if i >= 0 && i < len(s.extra) && s.extra[i].ID == k {
		return s.extra[i].dist, true
	}
	return 0, false
}

func (s *sparseLiveInfoMap) setDense(newdense []liveInfo) {
	s.dense = newdense
	for i, e := range s.dense {
		s.sparse[e.ID] = int32(i)
	}
	return
}

func (s *sparseLiveInfoMap) set(k ID, v int32) {
	i := int(s.sparse[k])
	if i < len(s.dense) && s.dense[i].ID == k {
		s.dense[i].dist = v
		return
	}
	i -= len(s.dense)
	if i >= 0 && i < len(s.extra) && s.extra[i].ID == k {
		s.extra[i].dist = v
		return
	}
	s.extra = append(s.extra, liveInfo{k, v})
	s.sparse[k] = int32(len(s.dense) + len(s.extra) - 1)
}

func (s *sparseLiveInfoMap) contents() []liveInfo {
	ret := s.dense
	if len(s.extra) > 0 {
		ret = append(ret, s.extra...)
	}
	s.dense = nil         // aid GC
	s.extra = s.extra[:0] // prepare for next round
	return ret
}
