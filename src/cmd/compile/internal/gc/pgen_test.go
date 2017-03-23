// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"reflect"
	"sort"
	"testing"
)

func typeWithoutPointers() *Type {
	return &Type{Etype: TSTRUCT, Extra: &StructType{Haspointers: 1}} // haspointers -> false
}

func typeWithPointers() *Type {
	return &Type{Etype: TSTRUCT, Extra: &StructType{Haspointers: 2}} // haspointers -> true
}

// Test all code paths for cmpstackvarlt.
func TestCmpstackvar(t *testing.T) {
	testdata := []struct {
		a, b Node
		lt   bool
	}{
		{
			Node{Class: PAUTO},
			Node{Class: PFUNC},
			false,
		},
		{
			Node{Class: PFUNC},
			Node{Class: PAUTO},
			true,
		},
		{
			Node{Class: PFUNC, Xoffset: 0},
			Node{Class: PFUNC, Xoffset: 10},
			true,
		},
		{
			Node{Class: PFUNC, Xoffset: 20},
			Node{Class: PFUNC, Xoffset: 10},
			false,
		},
		{
			Node{Class: PFUNC, Xoffset: 10},
			Node{Class: PFUNC, Xoffset: 10},
			false,
		},
		{
			Node{Class: PPARAM, Xoffset: 10},
			Node{Class: PPARAMOUT, Xoffset: 20},
			true,
		},
		{
			Node{Class: PPARAMOUT, Xoffset: 10},
			Node{Class: PPARAM, Xoffset: 20},
			true,
		},
		{
			Node{Class: PAUTO, flags: nodeUsed},
			Node{Class: PAUTO},
			true,
		},
		{
			Node{Class: PAUTO},
			Node{Class: PAUTO, flags: nodeUsed},
			false,
		},
		{
			Node{Class: PAUTO, Type: typeWithoutPointers()},
			Node{Class: PAUTO, Type: typeWithPointers()},
			false,
		},
		{
			Node{Class: PAUTO, Type: typeWithPointers()},
			Node{Class: PAUTO, Type: typeWithoutPointers()},
			true,
		},
		{
			Node{Class: PAUTO, Name: &Name{flags: nameNeedzero}},
			Node{Class: PAUTO, Name: &Name{}},
			true,
		},
		{
			Node{Class: PAUTO, Name: &Name{}},
			Node{Class: PAUTO, Name: &Name{flags: nameNeedzero}},
			false,
		},
		{
			Node{Class: PAUTO, Type: Types[TINT8], Name: &Name{}},
			Node{Class: PAUTO, Type: Types[TINT16], Name: &Name{}},
			false,
		},
		{
			Node{Class: PAUTO, Type: Types[TINT16], Name: &Name{}},
			Node{Class: PAUTO, Type: Types[TINT8], Name: &Name{}},
			true,
		},
		{
			Node{Class: PAUTO, Name: &Name{}, Sym: &Sym{Name: "abc"}},
			Node{Class: PAUTO, Name: &Name{}, Sym: &Sym{Name: "xyz"}},
			true,
		},
		{
			Node{Class: PAUTO, Name: &Name{}, Sym: &Sym{Name: "abc"}},
			Node{Class: PAUTO, Name: &Name{}, Sym: &Sym{Name: "abc"}},
			false,
		},
		{
			Node{Class: PAUTO, Name: &Name{}, Sym: &Sym{Name: "xyz"}},
			Node{Class: PAUTO, Name: &Name{}, Sym: &Sym{Name: "abc"}},
			false,
		},
	}
	for _, d := range testdata {
		if (d.a.Type == nil) != (d.b.Type == nil) {
			t.Fatalf("Type==nil mismatch: %v vs %v", d.a.Type, d.b.Type)
		}
		if d.a.Type == nil {
			d.a.Type = Types[TINT]
			d.b.Type = Types[TINT]
		}
		got := cmpstackvarlt(&d.a, &d.b)
		if got != d.lt {
			t.Errorf("want %#v < %#v", d.a, d.b)
		}
		// If we expect a < b to be true, check that b < a is false.
		if d.lt && cmpstackvarlt(&d.b, &d.a) {
			t.Errorf("unexpected %#v < %#v", d.b, d.a)
		}
	}
}

func TestStackvarSort(t *testing.T) {
	inp := []*Node{
		{Class: PFUNC, Name: &Name{}, Sym: &Sym{}},
		{Class: PAUTO, Name: &Name{}, Sym: &Sym{}},
		{Class: PFUNC, Xoffset: 0, Name: &Name{}, Sym: &Sym{}},
		{Class: PFUNC, Xoffset: 10, Name: &Name{}, Sym: &Sym{}},
		{Class: PFUNC, Xoffset: 20, Name: &Name{}, Sym: &Sym{}},
		{Class: PAUTO, flags: nodeUsed, Name: &Name{}, Sym: &Sym{}},
		{Class: PAUTO, Type: typeWithoutPointers(), Name: &Name{}, Sym: &Sym{}},
		{Class: PAUTO, Name: &Name{}, Sym: &Sym{}},
		{Class: PAUTO, Name: &Name{flags: nameNeedzero}, Sym: &Sym{}},
		{Class: PAUTO, Type: Types[TINT32], Name: &Name{}, Sym: &Sym{}},
		{Class: PAUTO, Type: Types[TINT64], Name: &Name{}, Sym: &Sym{}},
		{Class: PAUTO, Name: &Name{}, Sym: &Sym{Name: "abc"}},
		{Class: PAUTO, Name: &Name{}, Sym: &Sym{Name: "xyz"}},
	}
	want := []*Node{
		{Class: PFUNC, Name: &Name{}, Sym: &Sym{}},
		{Class: PFUNC, Xoffset: 0, Name: &Name{}, Sym: &Sym{}},
		{Class: PFUNC, Xoffset: 10, Name: &Name{}, Sym: &Sym{}},
		{Class: PFUNC, Xoffset: 20, Name: &Name{}, Sym: &Sym{}},
		{Class: PAUTO, flags: nodeUsed, Name: &Name{}, Sym: &Sym{}},
		{Class: PAUTO, Name: &Name{flags: nameNeedzero}, Sym: &Sym{}},
		{Class: PAUTO, Type: Types[TINT64], Name: &Name{}, Sym: &Sym{}},
		{Class: PAUTO, Type: Types[TINT32], Name: &Name{}, Sym: &Sym{}},
		{Class: PAUTO, Name: &Name{}, Sym: &Sym{}},
		{Class: PAUTO, Name: &Name{}, Sym: &Sym{}},
		{Class: PAUTO, Name: &Name{}, Sym: &Sym{Name: "abc"}},
		{Class: PAUTO, Name: &Name{}, Sym: &Sym{Name: "xyz"}},
		{Class: PAUTO, Type: typeWithoutPointers(), Name: &Name{}, Sym: &Sym{}},
	}

	for _, n := range inp {
		if n.Type == nil {
			n.Type = Types[TINT8]
		}
		_ = n.Type.Size()
	}

	for _, n := range want {
		if n.Type == nil {
			n.Type = Types[TINT8]
		}
		_ = n.Type.Size()
	}

	// haspointers updates Type.Haspointers as a side effect, so
	// exercise this function on all inputs so that reflect.DeepEqual
	// doesn't produce false positives.
	for i := range want {
		haspointers(want[i].Type)
		haspointers(inp[i].Type)
	}

	sort.Sort(byStackVar(inp))
	if !reflect.DeepEqual(want, inp) {
		t.Error("sort failed")
		for i := range inp {
			g := inp[i]
			w := want[i]
			eq := reflect.DeepEqual(w, g)
			if !eq {
				t.Logf("%d:\nwant=%+v\ngot=%+v", i, w, g)
			}
		}
	}
}
