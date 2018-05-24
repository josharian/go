// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"bytes"
	"fmt"
)

// funcArgOpt alters the order in which function arguments
// are written to the stack to minimize spills.
func funcArgOpt(f *Func) {
	return

	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if v.Op == OpMove {
				srcmem := v.MemoryArg()
				srcptr := v.Args[0]
				if srcmem.Op == OpMove {
					for srcptr.Op == OpOffPtr {
						srcptr = srcptr.Args[0]
					}
					fmt.Println("FROM", srcmem.LongString(), "TO", v.LongString(), "SRCPTR", srcptr.LongString())
				}
			}

			// {name: "Move", argLength: 3, typ: "Mem", aux: "TypSize"}, // arg0=destptr, arg1=srcptr, arg2=mem, auxint=size, aux=type.  Returns memory.

			continue

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

			// Chase the memory chain of Stores/Moves that set up the arguments.
			var origmem *Value
			var firstmem *Value
			mem := v.MemoryArg()
			prevoff := int64(-1)
			for mem.Op == OpStore || mem.Op == OpMove {
				ptr := mem.Args[0]
				if ptr.Op != OpOffPtr || ptr.Args[0].Op != OpSP {
					break
				}
				// TODO: can this be strengthed? Is it always right?
				if prevoff != -1 && prevoff <= ptr.AuxInt {
					f.Fatalf("arg stores out of order in %v: %d < %d", v.Aux, prevoff, ptr.AuxInt)
				}
				prevoff = ptr.AuxInt
				firstmem = mem
				mem = mem.MemoryArg()
				if ptr.AuxInt == 0 {
					// Reached beginning of argument stores.
					// What is mem now?
					origmem = mem
					break
				}
			}

			if origmem == nil { //|| (origmem.Op != OpStaticCall && origmem != OpInterCall && origmem.Op != OpClosureCall) {
				// origmem == nil happens when making a function call that takes no arguments.
				// And if we're not potentially loading the results
				continue
			}

			buf := new(bytes.Buffer)
			fmt.Fprintln(buf, "---", v.Op, v.Aux, v.AuxInt, "origmem=", origmem)

			var target *Value
			mem = v.MemoryArg()
			for mem.Op == OpStore || mem.Op == OpMove {
				if mem == origmem {
					break
				}
				arg := mem.Args[1]

				// if mem.Op == OpStore {
				// 	// Write down the value being stored.
				// 	fmt.Fprintln(buf, "\tSTORE", mem.Args[1].LongString(), "TO", ptr.AuxInt, "(SP)")
				// } else {
				// 	// Write down the source of the value being moved.
				// 	// OpMove
				// 	fmt.Fprintln(buf, "\tMOVE", mem.Args[1].LongString(), "TO", ptr.AuxInt, "(SP)")
				// }

				if arg.Op == OpLoad && arg.MemoryArg() == origmem && mem.Op == OpStore {
					// TODO: OpLoad is for OpStore; what about OpMove??
					fmt.Fprintln(buf, "\tusing origmem", origmem.Op.String(), "LOAD FROM PTR", arg.Args[0].LongString(), "TO STORE TO PTR", mem.Args[0].LongString())
					target = mem
					// break
				}

				// fmt.Fprintln(buf, "\t", mem.LongString(), "@@", ptr.AuxInt, " (", t.Size(), t.Alignment(), ")")
				mem = mem.MemoryArg()
			}
			if target != nil {
				// rewrite store order
				// TODO don't reorder if firstmem == target (or does it matter?)
				fmt.Fprintln(buf, "target=", target.LongString(), "first=", firstmem.LongString())
				target.Args[0], firstmem.Args[0] = firstmem.Args[0], target.Args[0]
				target.Args[1], firstmem.Args[1] = firstmem.Args[1], target.Args[1]
				target.Type, firstmem.Type = firstmem.Type, target.Type
				fmt.Println(buf.String())
			}

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
