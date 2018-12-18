// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package ssa

import (
	"cmd/internal/src"
)

type linepair struct {
	first, last uint32
}

type xposmap struct {
	maps       map[int32]*biasedSparseMap
	cachedsize int
}

func newXposmap(x map[int]linepair) *xposmap {
	maps := make(map[int32]*biasedSparseMap)
	for i, p := range x {
		maps[int32(i)] = newBiasedSparseMap(int(p.first), int(p.last))
	}
	return &xposmap{maps: maps} // zero for the rest is okay
}

func (m *xposmap) clear() {
	for _, l := range m.maps {
		if l != nil {
			l.clear()
		}
	}
	m.cachedsize = 0
}

func (m *xposmap) size() int {
	return m.cachedsize
}

func (m *xposmap) set(p src.XPos, v int32) {
	i := p.Index()
	s := m.maps[i]
	if s == nil {
		return
	}
	o := s.size()
	s.set(p.Line(), v)
	m.cachedsize += s.size() - o
}

func (m *xposmap) add(p src.XPos) {
	m.set(p, 0)
}

func (m *xposmap) contains(p src.XPos) bool {
	i := p.Index()
	s := m.maps[i]
	if s == nil {
		return false
	}
	return s.contains(p.Line())
}

func (m *xposmap) get(p src.XPos) int32 {
	i := p.Index()
	s := m.maps[i]
	if s == nil {
		return -1
	}
	return s.get(p.Line())
}

func (m *xposmap) remove(p src.XPos) {
	i := p.Index()
	s := m.maps[i]
	if s == nil {
		return
	}
	o := s.size()
	s.remove(p.Line())
	m.cachedsize += s.size() - o
}

func (m *xposmap) foreachEntry(f func(j int32, l uint, v int32)) {
	for j, mm := range m.maps {
		s := mm.size()
		for i := 0; i < s; i++ {
			l, v := mm.getEntry(i)
			f(j, l, v)
		}
	}
}

// getEntry returns the i'th key and value stored in s,
// where 0 <= i < s.size().  "Key" is XPos index and line number.
//func (m *xposmap) getEntry(i int) (j int32, l uint, v int32) {
//	geI := m.geIndex
//	geO := m.geOffset
//	if geI < 0 { // known-good values
//		geO = 0
//		geI = 0
//	}
//	for {
//		next := geO + m.maps[geI].size()
//		if next > i {
//			break
//		}
//		geO = next
//		geI++
//	}
//	j = int32(geI)
//	l, v = m.maps[geI].getEntry(i - geO)
//	m.geIndex = geI
//	m.geOffset = geO
//	return
//}
