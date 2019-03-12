// Code generated from gen/AMD64splitload.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "fmt"
import "math"
import "cmd/internal/obj"
import "cmd/internal/objabi"
import "cmd/compile/internal/types"

var _ = fmt.Println   // in case not otherwise used
var _ = math.MinInt8  // in case not otherwise used
var _ = obj.ANOP      // in case not otherwise used
var _ = objabi.GOROOT // in case not otherwise used
var _ = types.TypeMem // in case not otherwise used

func rewriteValueAMD64splitload(v *Value) bool {
	switch v.Op {
	case OpAMD64BTLconstload:
		return rewriteValueAMD64splitload_OpAMD64BTLconstload_0(v)
	case OpAMD64BTLconstloadidx1:
		return rewriteValueAMD64splitload_OpAMD64BTLconstloadidx1_0(v)
	case OpAMD64BTLconstloadidx4:
		return rewriteValueAMD64splitload_OpAMD64BTLconstloadidx4_0(v)
	case OpAMD64BTLconstloadidx8:
		return rewriteValueAMD64splitload_OpAMD64BTLconstloadidx8_0(v)
	case OpAMD64BTQconstload:
		return rewriteValueAMD64splitload_OpAMD64BTQconstload_0(v)
	case OpAMD64BTQconstloadidx1:
		return rewriteValueAMD64splitload_OpAMD64BTQconstloadidx1_0(v)
	case OpAMD64BTQconstloadidx8:
		return rewriteValueAMD64splitload_OpAMD64BTQconstloadidx8_0(v)
	case OpAMD64CMPBconstload:
		return rewriteValueAMD64splitload_OpAMD64CMPBconstload_0(v)
	case OpAMD64CMPBconstloadidx1:
		return rewriteValueAMD64splitload_OpAMD64CMPBconstloadidx1_0(v)
	case OpAMD64CMPBload:
		return rewriteValueAMD64splitload_OpAMD64CMPBload_0(v)
	case OpAMD64CMPBloadidx1:
		return rewriteValueAMD64splitload_OpAMD64CMPBloadidx1_0(v)
	case OpAMD64CMPLconstload:
		return rewriteValueAMD64splitload_OpAMD64CMPLconstload_0(v)
	case OpAMD64CMPLconstloadidx1:
		return rewriteValueAMD64splitload_OpAMD64CMPLconstloadidx1_0(v)
	case OpAMD64CMPLconstloadidx4:
		return rewriteValueAMD64splitload_OpAMD64CMPLconstloadidx4_0(v)
	case OpAMD64CMPLconstloadidx8:
		return rewriteValueAMD64splitload_OpAMD64CMPLconstloadidx8_0(v)
	case OpAMD64CMPLload:
		return rewriteValueAMD64splitload_OpAMD64CMPLload_0(v)
	case OpAMD64CMPLloadidx1:
		return rewriteValueAMD64splitload_OpAMD64CMPLloadidx1_0(v)
	case OpAMD64CMPLloadidx4:
		return rewriteValueAMD64splitload_OpAMD64CMPLloadidx4_0(v)
	case OpAMD64CMPLloadidx8:
		return rewriteValueAMD64splitload_OpAMD64CMPLloadidx8_0(v)
	case OpAMD64CMPQconstload:
		return rewriteValueAMD64splitload_OpAMD64CMPQconstload_0(v)
	case OpAMD64CMPQconstloadidx1:
		return rewriteValueAMD64splitload_OpAMD64CMPQconstloadidx1_0(v)
	case OpAMD64CMPQconstloadidx8:
		return rewriteValueAMD64splitload_OpAMD64CMPQconstloadidx8_0(v)
	case OpAMD64CMPQload:
		return rewriteValueAMD64splitload_OpAMD64CMPQload_0(v)
	case OpAMD64CMPQloadidx1:
		return rewriteValueAMD64splitload_OpAMD64CMPQloadidx1_0(v)
	case OpAMD64CMPQloadidx8:
		return rewriteValueAMD64splitload_OpAMD64CMPQloadidx8_0(v)
	case OpAMD64CMPWconstload:
		return rewriteValueAMD64splitload_OpAMD64CMPWconstload_0(v)
	case OpAMD64CMPWconstloadidx1:
		return rewriteValueAMD64splitload_OpAMD64CMPWconstloadidx1_0(v)
	case OpAMD64CMPWconstloadidx2:
		return rewriteValueAMD64splitload_OpAMD64CMPWconstloadidx2_0(v)
	case OpAMD64CMPWload:
		return rewriteValueAMD64splitload_OpAMD64CMPWload_0(v)
	case OpAMD64CMPWloadidx1:
		return rewriteValueAMD64splitload_OpAMD64CMPWloadidx1_0(v)
	case OpAMD64CMPWloadidx2:
		return rewriteValueAMD64splitload_OpAMD64CMPWloadidx2_0(v)
	}
	return false
}
func rewriteValueAMD64splitload_OpAMD64BTLconstload_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (BTLconstload {sym} [vo] ptr mem)
	// cond:
	// result: (BTLconst [valOnly(vo)] (MOVLload {sym} [offOnly(vo)] ptr mem))
	for {
		vo := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpAMD64BTLconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLload, typ.UInt32)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64BTLconstloadidx1_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (BTLconstloadidx1 {sym} [vo] ptr idx mem)
	// cond:
	// result: (BTLconst [valOnly(vo)] (MOVLloadidx1 {sym} [offOnly(vo)] ptr idx mem))
	for {
		vo := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(OpAMD64BTLconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx1, typ.UInt32)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64BTLconstloadidx4_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (BTLconstloadidx4 {sym} [vo] ptr idx mem)
	// cond:
	// result: (BTLconst [valOnly(vo)] (MOVLloadidx4 {sym} [offOnly(vo)] ptr idx mem))
	for {
		vo := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(OpAMD64BTLconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx4, typ.UInt32)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64BTLconstloadidx8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (BTLconstloadidx8 {sym} [vo] ptr idx mem)
	// cond:
	// result: (BTLconst [valOnly(vo)] (MOVLloadidx8 {sym} [offOnly(vo)] ptr idx mem))
	for {
		vo := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(OpAMD64BTLconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx8, typ.UInt32)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64BTQconstload_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (BTQconstload {sym} [vo] ptr mem)
	// cond:
	// result: (BTQconst [valOnly(vo)] (MOVQload {sym} [offOnly(vo)] ptr mem))
	for {
		vo := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpAMD64BTQconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQload, typ.UInt64)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64BTQconstloadidx1_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (BTQconstloadidx1 {sym} [vo] ptr idx mem)
	// cond:
	// result: (BTQconst [valOnly(vo)] (MOVQloadidx1 {sym} [offOnly(vo)] ptr idx mem))
	for {
		vo := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(OpAMD64BTQconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQloadidx1, typ.UInt64)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64BTQconstloadidx8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (BTQconstloadidx8 {sym} [vo] ptr idx mem)
	// cond:
	// result: (BTQconst [valOnly(vo)] (MOVQloadidx8 {sym} [offOnly(vo)] ptr idx mem))
	for {
		vo := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(OpAMD64BTQconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQloadidx8, typ.UInt64)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPBconstload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPBconstload {sym} [vo] ptr mem)
	// cond:
	// result: (CMPBconst (MOVBload {sym} [offOnly(vo)] ptr mem) [valOnly(vo)])
	for {
		vo := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpAMD64CMPBconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVBload, typ.UInt8)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPBconstloadidx1_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPBconstloadidx1 {sym} [vo] ptr idx mem)
	// cond:
	// result: (CMPBconst (MOVBloadidx1 {sym} [offOnly(vo)] ptr idx mem) [valOnly(vo)])
	for {
		vo := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v.reset(OpAMD64CMPBconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVBloadidx1, typ.UInt8)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPBload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPBload {sym} [off] ptr x mem)
	// cond:
	// result: (CMPB (MOVBload {sym} [off] ptr mem) x)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		x := v.Args[1]
		v.reset(OpAMD64CMPB)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVBload, typ.UInt8)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPBloadidx1_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPBloadidx1 {sym} [off] ptr idx x mem)
	// cond:
	// result: (CMPB (MOVBloadidx1 {sym} [off] ptr idx mem) x)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		x := v.Args[2]
		v.reset(OpAMD64CMPB)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVBloadidx1, typ.UInt8)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPLconstload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPLconstload {sym} [vo] ptr mem)
	// cond:
	// result: (CMPLconst (MOVLload {sym} [offOnly(vo)] ptr mem) [valOnly(vo)])
	for {
		vo := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpAMD64CMPLconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLload, typ.UInt32)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPLconstloadidx1_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPLconstloadidx1 {sym} [vo] ptr idx mem)
	// cond:
	// result: (CMPLconst (MOVLloadidx1 {sym} [offOnly(vo)] ptr idx mem) [valOnly(vo)])
	for {
		vo := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v.reset(OpAMD64CMPLconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx1, typ.UInt32)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPLconstloadidx4_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPLconstloadidx4 {sym} [vo] ptr idx mem)
	// cond:
	// result: (CMPLconst (MOVLloadidx4 {sym} [offOnly(vo)] ptr idx mem) [valOnly(vo)])
	for {
		vo := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v.reset(OpAMD64CMPLconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx4, typ.UInt32)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPLconstloadidx8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPLconstloadidx8 {sym} [vo] ptr idx mem)
	// cond:
	// result: (CMPLconst (MOVLloadidx8 {sym} [offOnly(vo)] ptr idx mem) [valOnly(vo)])
	for {
		vo := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v.reset(OpAMD64CMPLconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx8, typ.UInt32)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPLload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPLload {sym} [off] ptr x mem)
	// cond:
	// result: (CMPL (MOVLload {sym} [off] ptr mem) x)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		x := v.Args[1]
		v.reset(OpAMD64CMPL)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLload, typ.UInt32)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPLloadidx1_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPLloadidx1 {sym} [off] ptr idx x mem)
	// cond:
	// result: (CMPL (MOVLloadidx1 {sym} [off] ptr idx mem) x)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		x := v.Args[2]
		v.reset(OpAMD64CMPL)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx1, typ.UInt32)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPLloadidx4_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPLloadidx4 {sym} [off] ptr idx x mem)
	// cond:
	// result: (CMPL (MOVLloadidx4 {sym} [off] ptr idx mem) x)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		x := v.Args[2]
		v.reset(OpAMD64CMPL)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx4, typ.UInt32)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPLloadidx8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPLloadidx8 {sym} [off] ptr idx x mem)
	// cond:
	// result: (CMPL (MOVLloadidx8 {sym} [off] ptr idx mem) x)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		x := v.Args[2]
		v.reset(OpAMD64CMPL)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVLloadidx8, typ.UInt32)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPQconstload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPQconstload {sym} [vo] ptr mem)
	// cond:
	// result: (CMPQconst (MOVQload {sym} [offOnly(vo)] ptr mem) [valOnly(vo)])
	for {
		vo := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpAMD64CMPQconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQload, typ.UInt64)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPQconstloadidx1_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPQconstloadidx1 {sym} [vo] ptr idx mem)
	// cond:
	// result: (CMPQconst (MOVQloadidx1 {sym} [offOnly(vo)] ptr idx mem) [valOnly(vo)])
	for {
		vo := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v.reset(OpAMD64CMPQconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQloadidx1, typ.UInt64)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPQconstloadidx8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPQconstloadidx8 {sym} [vo] ptr idx mem)
	// cond:
	// result: (CMPQconst (MOVQloadidx8 {sym} [offOnly(vo)] ptr idx mem) [valOnly(vo)])
	for {
		vo := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v.reset(OpAMD64CMPQconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQloadidx8, typ.UInt64)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPQload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPQload {sym} [off] ptr x mem)
	// cond:
	// result: (CMPQ (MOVQload {sym} [off] ptr mem) x)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		x := v.Args[1]
		v.reset(OpAMD64CMPQ)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQload, typ.UInt64)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPQloadidx1_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPQloadidx1 {sym} [off] ptr idx x mem)
	// cond:
	// result: (CMPQ (MOVQloadidx1 {sym} [off] ptr idx mem) x)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		x := v.Args[2]
		v.reset(OpAMD64CMPQ)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQloadidx1, typ.UInt64)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPQloadidx8_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPQloadidx8 {sym} [off] ptr idx x mem)
	// cond:
	// result: (CMPQ (MOVQloadidx8 {sym} [off] ptr idx mem) x)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		x := v.Args[2]
		v.reset(OpAMD64CMPQ)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVQloadidx8, typ.UInt64)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPWconstload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPWconstload {sym} [vo] ptr mem)
	// cond:
	// result: (CMPWconst (MOVWload {sym} [offOnly(vo)] ptr mem) [valOnly(vo)])
	for {
		vo := v.AuxInt
		sym := v.Aux
		mem := v.Args[1]
		ptr := v.Args[0]
		v.reset(OpAMD64CMPWconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVWload, typ.UInt16)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPWconstloadidx1_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPWconstloadidx1 {sym} [vo] ptr idx mem)
	// cond:
	// result: (CMPWconst (MOVWloadidx1 {sym} [offOnly(vo)] ptr idx mem) [valOnly(vo)])
	for {
		vo := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v.reset(OpAMD64CMPWconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVWloadidx1, typ.UInt16)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPWconstloadidx2_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPWconstloadidx2 {sym} [vo] ptr idx mem)
	// cond:
	// result: (CMPWconst (MOVWloadidx2 {sym} [offOnly(vo)] ptr idx mem) [valOnly(vo)])
	for {
		vo := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v.reset(OpAMD64CMPWconst)
		v.AuxInt = valOnly(vo)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVWloadidx2, typ.UInt16)
		v0.AuxInt = offOnly(vo)
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPWload_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPWload {sym} [off] ptr x mem)
	// cond:
	// result: (CMPW (MOVWload {sym} [off] ptr mem) x)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[2]
		ptr := v.Args[0]
		x := v.Args[1]
		v.reset(OpAMD64CMPW)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVWload, typ.UInt16)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPWloadidx1_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPWloadidx1 {sym} [off] ptr idx x mem)
	// cond:
	// result: (CMPW (MOVWloadidx1 {sym} [off] ptr idx mem) x)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		x := v.Args[2]
		v.reset(OpAMD64CMPW)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVWloadidx1, typ.UInt16)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueAMD64splitload_OpAMD64CMPWloadidx2_0(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (CMPWloadidx2 {sym} [off] ptr idx x mem)
	// cond:
	// result: (CMPW (MOVWloadidx2 {sym} [off] ptr idx mem) x)
	for {
		off := v.AuxInt
		sym := v.Aux
		mem := v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		x := v.Args[2]
		v.reset(OpAMD64CMPW)
		v0 := b.NewValue0(v.Pos, OpAMD64MOVWloadidx2, typ.UInt16)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteBlockAMD64splitload(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	typ := &config.Types
	_ = typ
	switch b.Kind {
	}
	return false
}
