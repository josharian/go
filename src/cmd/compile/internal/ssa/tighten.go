// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// tighten moves Values closer to the Blocks in which they are used.
// This can reduce the amount of register spilling required,
// if it doesn't also create more live values.
// A Value can be moved to any block that
// dominates all blocks in which it is used.
func tighten(f *Func) {
	canMove := make([]bool, f.NumValues())
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			switch v.Op {
			case OpPhi, OpArg, OpSelect0, OpSelect1,
				OpAMD64LoweredGetClosurePtr, Op386LoweredGetClosurePtr,
				OpARMLoweredGetClosurePtr, OpARM64LoweredGetClosurePtr,
				OpMIPSLoweredGetClosurePtr, OpMIPS64LoweredGetClosurePtr,
				OpS390XLoweredGetClosurePtr, OpPPC64LoweredGetClosurePtr:
				// Phis need to stay in their block.
				// GetClosurePtr & Arg must stay in the entry block.
				// Tuple selectors must stay with the tuple generator.
				continue
			case OpGetG:
				// Special: getg() can sometimes be moved. It cannot fault.
				canMove[v.ID] = true
				continue
			case OpLoad:
				p := v.Args[0]
				if p.Op == OpOffPtr {
					p = p.Args[0]
				}
				switch p.Op {
				case OpSP, OpAddr:
					// Special: Loads from the stack and variables can sometimes be moved,
					// because they cannot fault.
					canMove[v.ID] = true
					continue
				}
			}
			if v.Type.IsMemory() {
				// We can't move arguments that might fault.
				continue
			}
			// Count arguments which will need a register.
			narg := 0
			for _, a := range v.Args {
				if !a.rematerializeable() {
					narg++
				}
			}
			if narg >= 2 && !v.Type.IsFlags() {
				// Don't move values with more than one input, as that may
				// increase register pressure.
				// We make an exception for flags, as we want flag generators
				// moved next to uses (because we only have 1 flag register).
				continue
			}
			canMove[v.ID] = true
		}
	}

	// Build data structure for fast least-common-ancestor queries.
	lca := makeLCArange(f)

	// For each moveable value, record the block that dominates all uses found so far.
	target := make([]*Block, f.NumValues())

	// Grab loop information.
	// We use this to make sure we don't tighten a value into a (deeper) loop.
	idom := f.Idom()
	loops := f.loopnest()
	loops.calculateDepths()

	changed := true
	for changed {
		changed = false

		// Reset target
		for i := range target {
			target[i] = nil
		}

		// Compute target locations (for moveable values only).
		// target location = the least common ancestor of all uses in the dominator tree.
		for _, b := range f.Blocks {
			for _, v := range b.Values {
				for i, a := range v.Args {
					if !canMove[a.ID] {
						continue
					}
					use := b
					if v.Op == OpPhi {
						use = b.Preds[i].b
					}
					if target[a.ID] == nil {
						target[a.ID] = use
					} else {
						target[a.ID] = lca.find(target[a.ID], use)
					}
				}
			}
			if c := b.Control; c != nil {
				if !canMove[c.ID] {
					continue
				}
				if target[c.ID] == nil {
					target[c.ID] = b
				} else {
					target[c.ID] = lca.find(target[c.ID], b)
				}
			}
		}

		// If the target location is inside a loop,
		// move the target location up to just before the loop head.
		for _, b := range f.Blocks {
			origloop := loops.b2l[b.ID]
			for _, v := range b.Values {
				t := target[v.ID]
				if t == nil {
					continue
				}
				targetloop := loops.b2l[t.ID]
				for targetloop != nil && (origloop == nil || targetloop.depth > origloop.depth) {
					t = idom[targetloop.header.ID]
					target[v.ID] = t
					targetloop = loops.b2l[t.ID]
				}
			}
		}

		// Move values to target locations.
		for _, b := range f.Blocks {
			for i := 0; i < len(b.Values); i++ {
				v := b.Values[i]
				t := target[v.ID]
				if t == nil || t == b {
					// v is not moveable, or is already in correct place.
					continue
				}
				if len(v.Args) > 0 && v.LastArg().Type.IsMemory() {
					// To move a load from the stack,
					// we must ensure that that memory arg is already
					// live in the target block.
					// For now, require that there be an existing value
					// in the target block that has the same memory arg.
					// TODO: Use a more precise analysis?
					ok := false
					for _, w := range t.Values {
						if w.Op != OpPhi && len(w.Args) > 0 && w.LastArg() == v.LastArg() {
							ok = true
							break
						}
					}
					if !ok {
						continue
					}
				}
				// Move v to the block which dominates its uses.
				t.Values = append(t.Values, v)
				v.Block = t
				last := len(b.Values) - 1
				b.Values[i] = b.Values[last]
				b.Values[last] = nil
				b.Values = b.Values[:last]
				changed = true
				i--
			}
		}
	}
}

// phiTighten moves constants closer to phi users.
// This pass avoids having lots of constants live for lots of the program.
// See issue 16407.
func phiTighten(f *Func) {
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if v.Op != OpPhi {
				continue
			}
			for i, a := range v.Args {
				if !a.rematerializeable() {
					continue // not a constant we can move around
				}
				if a.Block == b.Preds[i].b {
					continue // already in the right place
				}
				// Make a copy of a, put in predecessor block.
				v.SetArg(i, a.copyInto(b.Preds[i].b))
			}
		}
	}
}
