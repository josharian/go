// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"cmd/compile/internal/types"
	"fmt"
)

// AlgKind describes the kind of algorithms used for comparing and
// hashing a Type.
type AlgKind int

const (
	// These values are known by runtime.
	ANOEQ AlgKind = iota
	AMEM0
	AMEM8
	AMEM16
	AMEM32
	AMEM64
	AMEM128
	ASTRING
	AINTER
	ANILINTER
	AFLOAT32
	AFLOAT64
	ACPLX64
	ACPLX128

	// Type can be compared/hashed as regular memory.
	AMEM AlgKind = 100

	// Type needs special comparison/hashing functions.
	ASPECIAL AlgKind = -1
)

// IsComparable reports whether t is a comparable type.
func IsComparable(t *types.Type) bool {
	a, _ := algtype1(t)
	return a != ANOEQ
}

// IsRegularMemory reports whether t can be compared/hashed as regular memory.
func IsRegularMemory(t *types.Type) bool {
	a, _ := algtype1(t)
	return a == AMEM
}

// IncomparableField returns an incomparable Field of struct Type t, if any.
func IncomparableField(t *types.Type) *types.Field {
	for _, f := range t.FieldSlice() {
		if !IsComparable(f.Type) {
			return f
		}
	}
	return nil
}

// algtype is like algtype1, except it returns the fixed-width AMEMxx variants
// instead of the general AMEM kind when possible.
func algtype(t *types.Type) AlgKind {
	a, _ := algtype1(t)
	if a == AMEM {
		switch t.Width {
		case 0:
			return AMEM0
		case 1:
			return AMEM8
		case 2:
			return AMEM16
		case 4:
			return AMEM32
		case 8:
			return AMEM64
		case 16:
			return AMEM128
		}
	}

	return a
}

// algtype1 returns the AlgKind used for comparing and hashing Type t.
// If it returns ANOEQ, it also returns the component type of t that
// makes it incomparable.
func algtype1(t *types.Type) (AlgKind, *types.Type) {
	if t.Broke() {
		return AMEM, nil
	}
	if t.Noalg() {
		return ANOEQ, t
	}

	switch t.Etype {
	case TELIDED:
		// TODO: use
		return AMEM, nil
	case TANY, TFORW:
		// will be defined later.
		return ANOEQ, t

	case TINT8, TUINT8, TINT16, TUINT16,
		TINT32, TUINT32, TINT64, TUINT64,
		TINT, TUINT, TUINTPTR,
		TBOOL, TPTR32, TPTR64,
		TCHAN, TUNSAFEPTR:
		return AMEM, nil

	case TFUNC, TMAP:
		return ANOEQ, t

	case TFLOAT32:
		return AFLOAT32, nil

	case TFLOAT64:
		return AFLOAT64, nil

	case TCOMPLEX64:
		return ACPLX64, nil

	case TCOMPLEX128:
		return ACPLX128, nil

	case TSTRING:
		return ASTRING, nil

	case TINTER:
		if t.IsEmptyInterface() {
			return ANILINTER, nil
		}
		return AINTER, nil

	case TSLICE:
		return ANOEQ, t

	case TARRAY:
		a, bad := algtype1(t.Elem())
		switch a {
		case AMEM:
			return AMEM, nil
		case ANOEQ:
			return ANOEQ, bad
		}

		switch t.NumElem() {
		case 0:
			// We checked above that the element type is comparable.
			return AMEM, nil
		case 1:
			// Single-element array is same as its lone element.
			return a, nil
		}

		return ASPECIAL, nil

	case TSTRUCT:
		fields := t.FieldSlice()

		// One-field struct is same as that one field alone.
		if len(fields) == 1 && !fields[0].Sym.IsBlank() {
			return algtype1(fields[0].Type)
		}

		ret := AMEM
		for i, f := range fields {
			// All fields must be comparable.
			a, bad := algtype1(f.Type)
			if a == ANOEQ {
				return ANOEQ, bad
			}

			// Blank fields, padded fields, fields with non-memory
			// equality need special compare.
			if a != AMEM || f.Sym.IsBlank() || ispaddedfield(t, i) {
				ret = ASPECIAL
			}
		}

		return ret, nil
	}

	Fatalf("algtype1: unexpected type %v", t)
	return 0, nil
}

// Generate a helper function to compute the hash of a value of type t.
func genhash(sym *types.Sym, t *types.Type) {
	if Debug['r'] != 0 {
		fmt.Printf("genhash %v %v\n", sym, t)
	}

	lineno = autogeneratedPos // less confusing than end of input
	dclcontext = PEXTERN
	types.Markdcl()

	// func sym(p *T, h uintptr) uintptr
	tfn := nod(OTFUNC, nil, nil)
	n := namedfield("p", types.NewPtr(t))
	tfn.List.Append(n)
	np := n.Left
	n = namedfield("h", types.Types[TUINTPTR])
	tfn.List.Append(n)
	nh := n.Left
	n = anonfield(types.Types[TUINTPTR]) // return value
	tfn.Rlist.Append(n)

	fn := dclfunc(sym, tfn)

	// genhash is only called for types that have equality but
	// cannot be handled by the standard algorithms,
	// so t must be either an array or a struct.
	switch t.Etype {
	default:
		Fatalf("genhash %v", t)

	case types.TARRAY:
		// An array of pure memory would be handled by the
		// standard algorithm, so the element type must not be
		// pure memory.
		hashel := hashfor(t.Elem())

		n := nod(ORANGE, nil, nod(OIND, np, nil))
		ni := newname(lookup("i"))
		ni.Type = types.Types[TINT]
		n.List.Set1(ni)
		n.SetColas(true)
		colasdefn(n.List.Slice(), n)
		ni = n.List.First()

		// h = hashel(&p[i], h)
		call := nod(OCALL, hashel, nil)

		nx := nod(OINDEX, np, ni)
		nx.SetBounded(true)
		na := nod(OADDR, nx, nil)
		na.Etype = 1 // no escape to heap
		call.List.Append(na)
		call.List.Append(nh)
		n.Nbody.Append(nod(OAS, nh, call))

		fn.Nbody.Append(n)

	case types.TSTRUCT:
		// Walk the struct using memhash for runs of AMEM
		// and calling specific hash functions for the others.
		for i, fields := 0, t.FieldSlice(); i < len(fields); {
			f := fields[i]

			// Skip blank fields.
			if f.Sym.IsBlank() {
				i++
				continue
			}

			// Hash non-memory fields with appropriate hash function.
			if !IsRegularMemory(f.Type) {
				hashel := hashfor(f.Type)
				call := nod(OCALL, hashel, nil)
				nx := nodSym(OXDOT, np, f.Sym) // TODO: fields from other packages?
				na := nod(OADDR, nx, nil)
				na.Etype = 1 // no escape to heap
				call.List.Append(na)
				call.List.Append(nh)
				fn.Nbody.Append(nod(OAS, nh, call))
				i++
				continue
			}

			// Otherwise, hash a maximal length run of raw memory.
			size, next := memrun(t, i)

			// h = hashel(&p.first, size, h)
			hashel := hashmem(f.Type)
			call := nod(OCALL, hashel, nil)
			nx := nodSym(OXDOT, np, f.Sym) // TODO: fields from other packages?
			na := nod(OADDR, nx, nil)
			na.Etype = 1 // no escape to heap
			call.List.Append(na)
			call.List.Append(nh)
			call.List.Append(nodintconst(size))
			fn.Nbody.Append(nod(OAS, nh, call))

			i = next
		}
	}

	r := nod(ORETURN, nil, nil)
	r.List.Append(nh)
	fn.Nbody.Append(r)

	if Debug['r'] != 0 {
		dumplist("genhash body", fn.Nbody)
	}

	funcbody(fn)
	Curfn = fn
	fn.Func.SetDupok(true)
	fn = typecheck(fn, Etop)
	typecheckslice(fn.Nbody.Slice(), Etop)
	Curfn = nil
	types.Popdcl()
	if debug_dclstack != 0 {
		testdclstack()
	}

	// Disable safemode while compiling this code: the code we
	// generate internally can refer to unsafe.Pointer.
	// In this case it can happen if we need to generate an ==
	// for a struct containing a reflect.Value, which itself has
	// an unexported field of type unsafe.Pointer.
	old_safemode := safemode
	safemode = false

	disable_checknil++
	funccompile(fn)
	disable_checknil--

	safemode = old_safemode
}

func hashfor(t *types.Type) *Node {
	var sym *types.Sym

	switch a, _ := algtype1(t); a {
	case AMEM:
		Fatalf("hashfor with AMEM type")
	case AINTER:
		sym = Runtimepkg.Lookup("interhash")
	case ANILINTER:
		sym = Runtimepkg.Lookup("nilinterhash")
	case ASTRING:
		sym = Runtimepkg.Lookup("strhash")
	case AFLOAT32:
		sym = Runtimepkg.Lookup("f32hash")
	case AFLOAT64:
		sym = Runtimepkg.Lookup("f64hash")
	case ACPLX64:
		sym = Runtimepkg.Lookup("c64hash")
	case ACPLX128:
		sym = Runtimepkg.Lookup("c128hash")
	default:
		sym = typesymprefix(".hash", t)
	}

	n := newname(sym)
	n.Class = PFUNC
	tfn := nod(OTFUNC, nil, nil)
	tfn.List.Append(anonfield(types.NewPtr(t)))
	tfn.List.Append(anonfield(types.Types[TUINTPTR]))
	tfn.Rlist.Append(anonfield(types.Types[TUINTPTR]))
	tfn = typecheck(tfn, Etype)
	n.Type = tfn.Type
	return n
}

// geneq generates a helper function to
// check equality of two values of type t.
func geneq(sym *types.Sym, t *types.Type) {
	if Debug['r'] != 0 {
		fmt.Printf("geneq %v %v\n", sym, t)
	}

	lineno = autogeneratedPos // less confusing than end of input
	dclcontext = PEXTERN
	types.Markdcl()

	// func sym(p, q *T) bool
	tfn := nod(OTFUNC, nil, nil)
	n := namedfield("p", types.NewPtr(t))
	tfn.List.Append(n)
	np := n.Left
	n = namedfield("q", types.NewPtr(t))
	tfn.List.Append(n)
	nq := n.Left
	n = anonfield(types.Types[TBOOL])
	tfn.Rlist.Append(n)

	fn := dclfunc(sym, tfn)

	// geneq is only called for types that have equality but
	// cannot be handled by the standard algorithms,
	// so t must be either an array or a struct.
	switch t.Etype {
	default:
		Fatalf("geneq %v", t)

	case TARRAY:
		// An array of pure memory would be handled by the
		// standard memequal, so the element type must not be
		// pure memory. Even if we unrolled the range loop,
		// each iteration would be a function call, so don't bother
		// unrolling.
		nrange := nod(ORANGE, nil, nod(OIND, np, nil))

		ni := newname(lookup("i"))
		ni.Type = types.Types[TINT]
		nrange.List.Set1(ni)
		nrange.SetColas(true)
		colasdefn(nrange.List.Slice(), nrange)
		ni = nrange.List.First()

		// if p[i] != q[i] { return false }
		nx := nod(OINDEX, np, ni)

		nx.SetBounded(true)
		ny := nod(OINDEX, nq, ni)
		ny.SetBounded(true)

		nif := nod(OIF, nil, nil)
		nif.Left = nod(ONE, nx, ny)
		r := nod(ORETURN, nil, nil)
		r.List.Append(nodbool(false))
		nif.Nbody.Append(r)
		nrange.Nbody.Append(nif)
		fn.Nbody.Append(nrange)

		// return true
		ret := nod(ORETURN, nil, nil)
		ret.List.Append(nodbool(true))
		fn.Nbody.Append(ret)

	case TSTRUCT:
		var cond *Node
		and := func(n *Node) {
			if cond == nil {
				cond = n
				return
			}
			cond = nod(OANDAND, cond, n)
		}

		// Walk the struct using memequal for runs of AMEM
		// and calling specific equality tests for the others.
		for i, fields := 0, t.FieldSlice(); i < len(fields); {
			f := fields[i]

			// Skip blank-named fields.
			if f.Sym.IsBlank() {
				i++
				continue
			}

			// Compare non-memory fields with field equality.
			if !IsRegularMemory(f.Type) {
				and(eqfield(np, nq, f.Sym))
				i++
				continue
			}

			// Find maximal length run of memory-only fields.
			size, next := memrun(t, i)

			// TODO(rsc): All the calls to newname are wrong for
			// cross-package unexported fields.
			if s := fields[i:next]; len(s) <= 2 {
				// Two or fewer fields: use plain field equality.
				for _, f := range s {
					and(eqfield(np, nq, f.Sym))
				}
			} else {
				// More than two fields: use memequal.
				and(eqmem(np, nq, f.Sym, size))
			}
			i = next
		}

		if cond == nil {
			cond = nodbool(true)
		}

		ret := nod(ORETURN, nil, nil)
		ret.List.Append(cond)
		fn.Nbody.Append(ret)
	}

	if Debug['r'] != 0 {
		dumplist("geneq body", fn.Nbody)
	}

	funcbody(fn)
	Curfn = fn
	fn.Func.SetDupok(true)
	fn = typecheck(fn, Etop)
	typecheckslice(fn.Nbody.Slice(), Etop)
	Curfn = nil
	types.Popdcl()
	if debug_dclstack != 0 {
		testdclstack()
	}

	// Disable safemode while compiling this code: the code we
	// generate internally can refer to unsafe.Pointer.
	// In this case it can happen if we need to generate an ==
	// for a struct containing a reflect.Value, which itself has
	// an unexported field of type unsafe.Pointer.
	old_safemode := safemode
	safemode = false

	// Disable checknils while compiling this code.
	// We are comparing a struct or an array,
	// neither of which can be nil, and our comparisons
	// are shallow.
	disable_checknil++

	funccompile(fn)

	safemode = old_safemode
	disable_checknil--
}

// eqfield returns the node
// 	p.field == q.field
func eqfield(p *Node, q *Node, field *types.Sym) *Node {
	nx := nodSym(OXDOT, p, field)
	ny := nodSym(OXDOT, q, field)
	ne := nod(OEQ, nx, ny)
	return ne
}

// eqmem returns the node
// 	memequal(&p.field, &q.field [, size])
func eqmem(p *Node, q *Node, field *types.Sym, size int64) *Node {
	nx := nod(OADDR, nodSym(OXDOT, p, field), nil)
	nx.Etype = 1 // does not escape
	ny := nod(OADDR, nodSym(OXDOT, q, field), nil)
	ny.Etype = 1 // does not escape
	nx = typecheck(nx, Erv)
	ny = typecheck(ny, Erv)

	fn, needsize := eqmemfunc(size, nx.Type.Elem())
	call := nod(OCALL, fn, nil)
	call.List.Append(nx)
	call.List.Append(ny)
	if needsize {
		call.List.Append(nodintconst(size))
	}

	return call
}

func eqmemfunc(size int64, t *types.Type) (fn *Node, needsize bool) {
	switch size {
	default:
		fn = syslook("memequal")
		needsize = true
	case 1, 2, 4, 8, 16:
		buf := fmt.Sprintf("memequal%d", int(size)*8)
		fn = syslook(buf)
	}

	fn = substArgTypes(fn, t, t)
	return fn, needsize
}

// memrun finds runs of struct fields for which memory-only algs are appropriate.
// t is the parent struct type, and start is the field index at which to start the run.
// size is the length in bytes of the memory included in the run.
// next is the index just after the end of the memory run.
func memrun(t *types.Type, start int) (size int64, next int) {
	next = start
	for {
		next++
		if next == t.NumFields() {
			break
		}
		// Stop run after a padded field.
		if ispaddedfield(t, next-1) {
			break
		}
		// Also, stop before a blank or non-memory field.
		if f := t.Field(next); f.Sym.IsBlank() || !IsRegularMemory(f.Type) {
			break
		}
	}
	return t.Field(next-1).End() - t.Field(start).Offset, next
}

// ispaddedfield reports whether the i'th field of struct type t is followed
// by padding.
func ispaddedfield(t *types.Type, i int) bool {
	if !t.IsStruct() {
		Fatalf("ispaddedfield called non-struct %v", t)
	}
	end := t.Width
	if i+1 < t.NumFields() {
		end = t.Field(i + 1).Offset
	}
	return t.Field(i).End() != end
}
