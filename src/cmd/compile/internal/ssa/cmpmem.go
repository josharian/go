package ssa

func flagrewrite(f *Func) {
	for _, b := range f.Blocks {
		var mem *Value
		for _, v := range b.Values {
			if v.Type.IsMemory() {
				mem = v
			}
			if v.Op == OpAMD64CMPQconst {
				if v.Args[0].Op == OpAMD64MOVQload {
					if v.Args[0].MemoryArg() == mem && v.Block == v.Args[0].Block && v.Args[0].Uses == 1 { // TODO: more options here
						var i int
						var w *Value
						for i, w = range b.Values {
							if w == v.Args[0] {
								break
							}
						}
						copy(b.Values[i:], b.Values[i+1:])
						b.Values[len(b.Values)-1] = nil
						b.Values = b.Values[:len(b.Values)-1]
						v.Op = OpAMD64CMPQconstmem
					}
					// fmt.Println("eq", v.Args[0].MemoryArg() == mem, "vnil", v.Args[0].MemoryArg() == nil, "memnil", mem == nil, "sameblock", v.Block == v.Args[0].Block)
					// fmt.Println(v.LongString())
				}
			}
		}
	}
}
