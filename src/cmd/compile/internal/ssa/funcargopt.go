// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"fmt"
)

// funcArgOpt alters the order in which function arguments
// are written to the stack to minimize spills.
func funcArgOpt(f *Func) {
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			// Ignore OpClosureCall.
			// Closures always take two arguments,
			// a function pointer and a data pointer,
			// and neither will be left over on the stack
			// from a previous function call.
			switch v.Op {
			case OpStaticCall, OpInterCall:
			default:
				continue
			}
			fmt.Println("---", v.Op, v.Aux, v.AuxInt)
			mem := v.MemoryArg()
			prevoff := int64(-1)
			for mem.Op == OpStore || mem.Op == OpMove {
				ptr := mem.Args[0]
				if ptr.Op != OpOffPtr {
					break
				}
				if ptr.Args[0].Op != OpSP {
					break
				}
				// t := mem.Aux.(*types.Type)
				// fmt.Println("\t", ptr.AuxInt, " (", t.Size(), t.Alignment(), ")")
				if prevoff != -1 && prevoff < ptr.AuxInt {
					f.Fatalf("arg stores out of order in %v: %d < %d", v.Aux, prevoff, ptr.AuxInt)
				}
				prevoff = ptr.AuxInt
				mem = mem.MemoryArg()
				if ptr.AuxInt == 0 {
					// Reached beginning of argument stores.
					// What is mem now?
					fmt.Println(mem.LongString())
				}
			}
			// fmt.Println()

			/*
				if mem.Op == OpInitMem {
					continue
				}
				if mem.Op == OpPhi || mem.Op == OpStaticCall {
					// anything we can do here?
					continue
				}
				if mem.Op == OpSelect0 || mem.Op == OpSelect1 {
					fmt.Println("opselect", f.Name, v.LongString, "...", mem.LongString())
					continue
				}
				if mem.Op == OpStore {
					fmt.Println("STORE", mem.LongString(), "PTR", mem.Args[0].PtrString())
					continue
				}
				if mem.Op == OpMove {
					fmt.Println("MOVE TODO")
					continue
				}
			*/

			// VarKill
			// AtomicStore32
			// Zero
			// ClosureCall
			// AtomicOr8

			// {name: "Store", argLength: 3, typ: "Mem", aux: "Typ"}, // Store arg1 to arg0.  arg2=memory, aux=type.  Returns memory.
			// The source and destination of Move may overlap in some cases. See e.g.
			// memmove inlining in generic.rules. When inlineablememmovesize (in ../rewrite.go)
			// returns true, we must do all loads before all stores, when lowering Move.
			// {name: "Move", argLength: 3, typ: "Mem", aux: "TypSize"}, // arg0=destptr, arg1=srcptr, arg2=mem, auxint=size, aux=type.  Returns memory.

			// fmt.Println("NEW OP", mem.Op.String())
		}
	}
}
