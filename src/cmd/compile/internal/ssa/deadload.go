// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import "fmt"

// TODO

// For any given load, walk back through its history until:
// we find a mtaching store/zero
// we find a store that we can't prove doesn't alias our pointer

const (
	debugLoadElim    = false
	pctLoadElim      = false
	debugContainsPtr = false
)

func loadelim(f *Func) {
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if v.Op != OpLoad {
				continue
			}
			ptr := v.Args[0]
			if debugLoadElim {
				fmt.Print("loadelim ", ptr.LongString(), "? ")
			}
			_ = ptr
			mem := v.Args[1]
			zeroonly := true
			n := 0
			// TODO: OpMove? OpPhi?
			for mem.Op == OpStore || mem.Op == OpZero {
				n++
				if zeroonly && n > 0 {
					fmt.Println("ZERO MEM", n, mem)
				}

				if debugLoadElim {
					fmt.Print("mem: ", mem.LongLongString())
				}

				p := mem.Args[0]
				if isSamePtr(p, ptr) {
					// TODO: cehck equal type, equal size
					if mem.Op == OpStore {
						v.Op = OpCopy
						v.SetArgs1(mem.Args[1])
						if pctLoadElim {
							fmt.Println("DEADLOAD CONVERT SAME PTR STORE", mem.Op)
						}
					} else {
						// TODO: convert v into a const
						if pctLoadElim {
							fmt.Println("DEADLOAD CONVERT SAME PTR ZERO", mem.Op)
						}
					}
					if debugLoadElim {
						fmt.Print(" CONVERT SAME PTR", mem.Op)
					}
					break
				}

				if mem.Op == OpZero && zeroonly && containsPtr(p, ptr) {
					// TODO: convert v into a constant
					if debugLoadElim {
						fmt.Print(" CONVERT ZERO PARTIAL")
					}
					if pctLoadElim {
						fmt.Println("DEADLOAD CONVERT ZERO PARTIAL")
					}
					break
				}

				switch mem.Op {
				case OpStore:
					zeroonly = false
					mem = mem.Args[2]
				case OpZero:
					mem = mem.Args[1]
				}
				if !zeroonly && mightalias(ptr, p) {
					// can't prove they don't alias, have to stop now
					break
				}
			}
			if debugLoadElim {
				fmt.Println()
			}
		}
	}
}

// ptroff splits p into a pointer and an offset relative to that pointer.
func ptroff(p *Value) (int64, *Value) {
	off := int64(0)
	for p.Op == OpOffPtr {
		off += p.AuxInt
		p = p.Args[0]
	}
	return off, p
}

// mightalias reports whether pointers p and q might address the same memory.
func mightalias(p, q *Value) bool {
	// {name: "OffPtr", argLength: 1, aux: "Int64"},     // arg0 + auxint (arg0 and result are pointers)
	// {name: "Addr", argLength: 1, aux: "Sym"}, // Address of a variable.  Arg0=SP or SB.  Aux identifies the variable.
	// {name: "AddPtr", argLength: 2}, // For address calculations.  arg0 is a pointer and arg1 is an int.
	// if p.Op != q.Op {

	// }
	if p.Op == OpAddr && q.Op == OpAddr && (p.Args[0].Op != q.Args[0].Op || p.Aux != q.Aux) {
		return false
	}

	var poff, qoff int64
	poff, p = ptroff(p)
	qoff, q = ptroff(q)
	if isSamePtr(p, q) &&
		((poff > qoff && poff-qoff > q.Type.Size()) ||
			(qoff > poff && qoff-poff > p.Type.Size())) {
		return false
	}
	// If we don't know any better, they might alias.
	return true
}

// containsPtr reports whether we can prove that the memory pointed to by q is a subset of the memory pointed to by p.
// TODO: unicode subset for concision.
func containsPtr(p, q *Value) bool {
	if debugContainsPtr {
		fmt.Print("containsPtr? ", p.LongLongString(), " ;; ", q.LongLongString(), " ::")
	}
	if isSamePtr(p, q) {
		fmt.Println(" SAME")
		return true
	}
	// if p.Op != q.Op {
	// 	return false // TODO: strengthen
	// }
	var poff, qoff int64
	poff, p = ptroff(p)
	qoff, q = ptroff(q)
	x := isSamePtr(p, q) &&
		((poff > qoff && poff-qoff <= q.Type.Size()) ||
			(qoff > poff && qoff-poff <= p.Type.Size()))
	if debugContainsPtr {
		if x {
			fmt.Println(" ", x, "(", isSamePtr(p, q), (poff > qoff && poff-qoff <= q.Type.Size()), (qoff > poff && qoff-poff <= p.Type.Size()), ")")
		} else {
			fmt.Println()
		}
	}
	return x
}
