// Code generated from gen/AMD64merge.rules; DO NOT EDIT.
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

func rewriteValueAMD64merge(v *Value) bool {
	switch v.Op {
	case OpAMD64ADCQ:
		return rewriteValueAMD64merge_OpAMD64ADCQ_0(v)
	case OpAMD64BTLconst:
		return rewriteValueAMD64merge_OpAMD64BTLconst_0(v)
	case OpAMD64BTQconst:
		return rewriteValueAMD64merge_OpAMD64BTQconst_0(v)
	case OpAMD64CMPB:
		return rewriteValueAMD64merge_OpAMD64CMPB_0(v)
	case OpAMD64CMPBconst:
		return rewriteValueAMD64merge_OpAMD64CMPBconst_0(v)
	case OpAMD64CMPL:
		return rewriteValueAMD64merge_OpAMD64CMPL_0(v)
	case OpAMD64CMPLconst:
		return rewriteValueAMD64merge_OpAMD64CMPLconst_0(v)
	case OpAMD64CMPQ:
		return rewriteValueAMD64merge_OpAMD64CMPQ_0(v)
	case OpAMD64CMPQconst:
		return rewriteValueAMD64merge_OpAMD64CMPQconst_0(v)
	case OpAMD64CMPW:
		return rewriteValueAMD64merge_OpAMD64CMPW_0(v)
	case OpAMD64CMPWconst:
		return rewriteValueAMD64merge_OpAMD64CMPWconst_0(v)
	case OpAMD64SBBQ:
		return rewriteValueAMD64merge_OpAMD64SBBQ_0(v)
	}
	return false
}
func rewriteValueAMD64merge_OpAMD64ADCQ_0(v *Value) bool {
	// match: (ADCQ load:(MOVQload {sym} [off] ptr mem) x carry)
	// cond: canMergeLoadLateClobber(v, load, x) && clobber(load)
	// result: (ADCQload {sym} [off] ptr x carry mem)
	for {
		_ = v.Args[2]
		load := v.Args[0]
		if load.Op != OpAMD64MOVQload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		x := v.Args[1]
		carry := v.Args[2]
		if !(canMergeLoadLateClobber(v, load, x) && clobber(load)) {
			break
		}
		v.reset(OpAMD64ADCQload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(carry)
		v.AddArg(mem)
		return true
	}
	// match: (ADCQ x load:(MOVQload {sym} [off] ptr mem) carry)
	// cond: canMergeLoadLateClobber(v, load, x) && clobber(load)
	// result: (ADCQload {sym} [off] ptr x carry mem)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVQload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		carry := v.Args[2]
		if !(canMergeLoadLateClobber(v, load, x) && clobber(load)) {
			break
		}
		v.reset(OpAMD64ADCQload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(carry)
		v.AddArg(mem)
		return true
	}
	// match: (ADCQ load:(MOVQloadidx1 {sym} [off] ptr idx mem) x carry)
	// cond: canMergeLoadLateClobber(v, load, x) && clobber(load)
	// result: (ADCQloadidx1 {sym} [off] ptr idx x carry mem)
	for {
		_ = v.Args[2]
		load := v.Args[0]
		if load.Op != OpAMD64MOVQloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		x := v.Args[1]
		carry := v.Args[2]
		if !(canMergeLoadLateClobber(v, load, x) && clobber(load)) {
			break
		}
		v.reset(OpAMD64ADCQloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(carry)
		v.AddArg(mem)
		return true
	}
	// match: (ADCQ x load:(MOVQloadidx1 {sym} [off] ptr idx mem) carry)
	// cond: canMergeLoadLateClobber(v, load, x) && clobber(load)
	// result: (ADCQloadidx1 {sym} [off] ptr idx x carry mem)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVQloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		carry := v.Args[2]
		if !(canMergeLoadLateClobber(v, load, x) && clobber(load)) {
			break
		}
		v.reset(OpAMD64ADCQloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(carry)
		v.AddArg(mem)
		return true
	}
	// match: (ADCQ load:(MOVQloadidx8 {sym} [off] ptr idx mem) x carry)
	// cond: canMergeLoadLateClobber(v, load, x) && clobber(load)
	// result: (ADCQloadidx8 {sym} [off] ptr idx x carry mem)
	for {
		_ = v.Args[2]
		load := v.Args[0]
		if load.Op != OpAMD64MOVQloadidx8 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		x := v.Args[1]
		carry := v.Args[2]
		if !(canMergeLoadLateClobber(v, load, x) && clobber(load)) {
			break
		}
		v.reset(OpAMD64ADCQloadidx8)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(carry)
		v.AddArg(mem)
		return true
	}
	// match: (ADCQ x load:(MOVQloadidx8 {sym} [off] ptr idx mem) carry)
	// cond: canMergeLoadLateClobber(v, load, x) && clobber(load)
	// result: (ADCQloadidx8 {sym} [off] ptr idx x carry mem)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVQloadidx8 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		carry := v.Args[2]
		if !(canMergeLoadLateClobber(v, load, x) && clobber(load)) {
			break
		}
		v.reset(OpAMD64ADCQloadidx8)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(carry)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64merge_OpAMD64BTLconst_0(v *Value) bool {
	// match: (BTLconst [c] load:(MOVLload {sym} [off] ptr mem))
	// cond: validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (BTLconstload {sym} [makeValAndOff(int64(int8(c)), off)] ptr mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVLload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		if !(validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64BTLconstload)
		v.AuxInt = makeValAndOff(int64(int8(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (BTLconst [c] load:(MOVLloadidx1 {sym} [off] ptr idx mem))
	// cond: validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (BTLconstloadidx1 {sym} [makeValAndOff(int64(int8(c)), off)] ptr idx mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVLloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64BTLconstloadidx1)
		v.AuxInt = makeValAndOff(int64(int8(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (BTLconst [c] load:(MOVLloadidx4 {sym} [off] ptr idx mem))
	// cond: validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (BTLconstloadidx4 {sym} [makeValAndOff(int64(int8(c)), off)] ptr idx mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVLloadidx4 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64BTLconstloadidx4)
		v.AuxInt = makeValAndOff(int64(int8(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (BTLconst [c] load:(MOVLloadidx8 {sym} [off] ptr idx mem))
	// cond: validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (BTLconstloadidx8 {sym} [makeValAndOff(int64(int8(c)), off)] ptr idx mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVLloadidx8 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64BTLconstloadidx8)
		v.AuxInt = makeValAndOff(int64(int8(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64merge_OpAMD64BTQconst_0(v *Value) bool {
	// match: (BTQconst [c] load:(MOVQload {sym} [off] ptr mem))
	// cond: validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (BTQconstload {sym} [makeValAndOff(int64(int8(c)), off)] ptr mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVQload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		if !(validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64BTQconstload)
		v.AuxInt = makeValAndOff(int64(int8(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (BTQconst [c] load:(MOVQloadidx1 {sym} [off] ptr idx mem))
	// cond: validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (BTQconstloadidx1 {sym} [makeValAndOff(int64(int8(c)), off)] ptr idx mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVQloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64BTQconstloadidx1)
		v.AuxInt = makeValAndOff(int64(int8(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (BTQconst [c] load:(MOVQloadidx8 {sym} [off] ptr idx mem))
	// cond: validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (BTQconstloadidx8 {sym} [makeValAndOff(int64(int8(c)), off)] ptr idx mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVQloadidx8 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64BTQconstloadidx8)
		v.AuxInt = makeValAndOff(int64(int8(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64merge_OpAMD64CMPB_0(v *Value) bool {
	// match: (CMPB load:(MOVBload {sym} [off] ptr mem) x)
	// cond: canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPBload {sym} [off] ptr x mem)
	for {
		_ = v.Args[1]
		load := v.Args[0]
		if load.Op != OpAMD64MOVBload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		x := v.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPBload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPB load:(MOVBloadidx1 {sym} [off] ptr idx mem) x)
	// cond: canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPBloadidx1 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		load := v.Args[0]
		if load.Op != OpAMD64MOVBloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		x := v.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPBloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPB x load:(MOVBload {sym} [off] ptr mem))
	// cond: canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)
	// result: (CMPBload {sym} [off] ptr x mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVBload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)) {
			break
		}
		v.reset(OpAMD64CMPBload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPB x load:(MOVBloadidx1 {sym} [off] ptr idx mem))
	// cond: canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)
	// result: (CMPBloadidx1 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVBloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)) {
			break
		}
		v.reset(OpAMD64CMPBloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64merge_OpAMD64CMPBconst_0(v *Value) bool {
	// match: (CMPBconst [c] load:(MOVBload {sym} [off] ptr mem))
	// cond: validValAndOff(int64( int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPBconstload {sym} [makeValAndOff(int64( int8(c)), off)] ptr mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVBload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		if !(validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPBconstload)
		v.AuxInt = makeValAndOff(int64(int8(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (CMPBconst [c] load:(MOVBloadidx1 {sym} [off] ptr idx mem))
	// cond: validValAndOff(int64( int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPBconstloadidx1 {sym} [makeValAndOff(int64( int8(c)), off)] ptr idx mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVBloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(validValAndOff(int64(int8(c)), off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPBconstloadidx1)
		v.AuxInt = makeValAndOff(int64(int8(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64merge_OpAMD64CMPL_0(v *Value) bool {
	// match: (CMPL load:(MOVLload {sym} [off] ptr mem) x)
	// cond: canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPLload {sym} [off] ptr x mem)
	for {
		_ = v.Args[1]
		load := v.Args[0]
		if load.Op != OpAMD64MOVLload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		x := v.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPLload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPL load:(MOVLloadidx1 {sym} [off] ptr idx mem) x)
	// cond: canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPLloadidx1 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		load := v.Args[0]
		if load.Op != OpAMD64MOVLloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		x := v.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPLloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPL load:(MOVLloadidx4 {sym} [off] ptr idx mem) x)
	// cond: canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPLloadidx4 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		load := v.Args[0]
		if load.Op != OpAMD64MOVLloadidx4 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		x := v.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPLloadidx4)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPL load:(MOVLloadidx8 {sym} [off] ptr idx mem) x)
	// cond: canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPLloadidx8 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		load := v.Args[0]
		if load.Op != OpAMD64MOVLloadidx8 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		x := v.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPLloadidx8)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPL x load:(MOVLload {sym} [off] ptr mem))
	// cond: canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)
	// result: (CMPLload {sym} [off] ptr x mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVLload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)) {
			break
		}
		v.reset(OpAMD64CMPLload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPL x load:(MOVLloadidx1 {sym} [off] ptr idx mem))
	// cond: canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)
	// result: (CMPLloadidx1 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVLloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)) {
			break
		}
		v.reset(OpAMD64CMPLloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPL x load:(MOVLloadidx4 {sym} [off] ptr idx mem))
	// cond: canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)
	// result: (CMPLloadidx4 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVLloadidx4 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)) {
			break
		}
		v.reset(OpAMD64CMPLloadidx4)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPL x load:(MOVLloadidx8 {sym} [off] ptr idx mem))
	// cond: canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)
	// result: (CMPLloadidx8 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVLloadidx8 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)) {
			break
		}
		v.reset(OpAMD64CMPLloadidx8)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64merge_OpAMD64CMPLconst_0(v *Value) bool {
	// match: (CMPLconst [c] load:(MOVLload {sym} [off] ptr mem))
	// cond: validValAndOff( c , off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPLconstload {sym} [makeValAndOff( c , off)] ptr mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVLload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		if !(validValAndOff(c, off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPLconstload)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (CMPLconst [c] load:(MOVLloadidx1 {sym} [off] ptr idx mem))
	// cond: validValAndOff( c , off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPLconstloadidx1 {sym} [makeValAndOff( c , off)] ptr idx mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVLloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(validValAndOff(c, off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPLconstloadidx1)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (CMPLconst [c] load:(MOVLloadidx4 {sym} [off] ptr idx mem))
	// cond: validValAndOff( c , off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPLconstloadidx4 {sym} [makeValAndOff( c , off)] ptr idx mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVLloadidx4 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(validValAndOff(c, off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPLconstloadidx4)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (CMPLconst [c] load:(MOVLloadidx8 {sym} [off] ptr idx mem))
	// cond: validValAndOff( c , off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPLconstloadidx8 {sym} [makeValAndOff( c , off)] ptr idx mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVLloadidx8 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(validValAndOff(c, off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPLconstloadidx8)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64merge_OpAMD64CMPQ_0(v *Value) bool {
	// match: (CMPQ load:(MOVQload {sym} [off] ptr mem) x)
	// cond: canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPQload {sym} [off] ptr x mem)
	for {
		_ = v.Args[1]
		load := v.Args[0]
		if load.Op != OpAMD64MOVQload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		x := v.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPQload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPQ load:(MOVQloadidx1 {sym} [off] ptr idx mem) x)
	// cond: canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPQloadidx1 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		load := v.Args[0]
		if load.Op != OpAMD64MOVQloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		x := v.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPQloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPQ load:(MOVQloadidx8 {sym} [off] ptr idx mem) x)
	// cond: canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPQloadidx8 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		load := v.Args[0]
		if load.Op != OpAMD64MOVQloadidx8 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		x := v.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPQloadidx8)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPQ x load:(MOVQload {sym} [off] ptr mem))
	// cond: canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)
	// result: (CMPQload {sym} [off] ptr x mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVQload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)) {
			break
		}
		v.reset(OpAMD64CMPQload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPQ x load:(MOVQloadidx1 {sym} [off] ptr idx mem))
	// cond: canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)
	// result: (CMPQloadidx1 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVQloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)) {
			break
		}
		v.reset(OpAMD64CMPQloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPQ x load:(MOVQloadidx8 {sym} [off] ptr idx mem))
	// cond: canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)
	// result: (CMPQloadidx8 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVQloadidx8 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)) {
			break
		}
		v.reset(OpAMD64CMPQloadidx8)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64merge_OpAMD64CMPQconst_0(v *Value) bool {
	// match: (CMPQconst [c] load:(MOVQload {sym} [off] ptr mem))
	// cond: validValAndOff( c , off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPQconstload {sym} [makeValAndOff( c , off)] ptr mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVQload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		if !(validValAndOff(c, off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPQconstload)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (CMPQconst [c] load:(MOVQloadidx1 {sym} [off] ptr idx mem))
	// cond: validValAndOff( c , off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPQconstloadidx1 {sym} [makeValAndOff( c , off)] ptr idx mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVQloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(validValAndOff(c, off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPQconstloadidx1)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (CMPQconst [c] load:(MOVQloadidx8 {sym} [off] ptr idx mem))
	// cond: validValAndOff( c , off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPQconstloadidx8 {sym} [makeValAndOff( c , off)] ptr idx mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVQloadidx8 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(validValAndOff(c, off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPQconstloadidx8)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64merge_OpAMD64CMPW_0(v *Value) bool {
	// match: (CMPW load:(MOVWload {sym} [off] ptr mem) x)
	// cond: canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPWload {sym} [off] ptr x mem)
	for {
		_ = v.Args[1]
		load := v.Args[0]
		if load.Op != OpAMD64MOVWload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		x := v.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPWload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPW load:(MOVWloadidx1 {sym} [off] ptr idx mem) x)
	// cond: canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPWloadidx1 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		load := v.Args[0]
		if load.Op != OpAMD64MOVWloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		x := v.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPWloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPW load:(MOVWloadidx2 {sym} [off] ptr idx mem) x)
	// cond: canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPWloadidx2 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		load := v.Args[0]
		if load.Op != OpAMD64MOVWloadidx2 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		x := v.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPWloadidx2)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPW x load:(MOVWload {sym} [off] ptr mem))
	// cond: canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)
	// result: (CMPWload {sym} [off] ptr x mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVWload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		if !(canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)) {
			break
		}
		v.reset(OpAMD64CMPWload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPW x load:(MOVWloadidx1 {sym} [off] ptr idx mem))
	// cond: canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)
	// result: (CMPWloadidx1 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVWloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)) {
			break
		}
		v.reset(OpAMD64CMPWloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (CMPW x load:(MOVWloadidx2 {sym} [off] ptr idx mem))
	// cond: canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)
	// result: (CMPWloadidx2 {sym} [off] ptr idx x mem)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVWloadidx2 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(canMergeLoadLate(v, load) && clobber(load) && invertFlags(v)) {
			break
		}
		v.reset(OpAMD64CMPWloadidx2)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64merge_OpAMD64CMPWconst_0(v *Value) bool {
	// match: (CMPWconst [c] load:(MOVWload {sym} [off] ptr mem))
	// cond: validValAndOff(int64(int16(c)), off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPWconstload {sym} [makeValAndOff(int64(int16(c)), off)] ptr mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVWload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		if !(validValAndOff(int64(int16(c)), off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPWconstload)
		v.AuxInt = makeValAndOff(int64(int16(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (CMPWconst [c] load:(MOVWloadidx1 {sym} [off] ptr idx mem))
	// cond: validValAndOff(int64(int16(c)), off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPWconstloadidx1 {sym} [makeValAndOff(int64(int16(c)), off)] ptr idx mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVWloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(validValAndOff(int64(int16(c)), off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPWconstloadidx1)
		v.AuxInt = makeValAndOff(int64(int16(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (CMPWconst [c] load:(MOVWloadidx2 {sym} [off] ptr idx mem))
	// cond: validValAndOff(int64(int16(c)), off) && canMergeLoadLate(v, load) && clobber(load)
	// result: (CMPWconstloadidx2 {sym} [makeValAndOff(int64(int16(c)), off)] ptr idx mem)
	for {
		c := v.AuxInt
		load := v.Args[0]
		if load.Op != OpAMD64MOVWloadidx2 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		if !(validValAndOff(int64(int16(c)), off) && canMergeLoadLate(v, load) && clobber(load)) {
			break
		}
		v.reset(OpAMD64CMPWconstloadidx2)
		v.AuxInt = makeValAndOff(int64(int16(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueAMD64merge_OpAMD64SBBQ_0(v *Value) bool {
	// match: (SBBQ x load:(MOVQload {sym} [off] ptr mem) carry)
	// cond: canMergeLoadLateClobber(v, load, x) && clobber(load)
	// result: (SBBQload {sym} [off] ptr x carry mem)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVQload {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[1]
		ptr := load.Args[0]
		mem := load.Args[1]
		carry := v.Args[2]
		if !(canMergeLoadLateClobber(v, load, x) && clobber(load)) {
			break
		}
		v.reset(OpAMD64SBBQload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(carry)
		v.AddArg(mem)
		return true
	}
	// match: (SBBQ x load:(MOVQloadidx1 {sym} [off] ptr idx mem) carry)
	// cond: canMergeLoadLateClobber(v, load, x) && clobber(load)
	// result: (SBBQloadidx1 {sym} [off] ptr idx x carry mem)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVQloadidx1 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		carry := v.Args[2]
		if !(canMergeLoadLateClobber(v, load, x) && clobber(load)) {
			break
		}
		v.reset(OpAMD64SBBQloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(carry)
		v.AddArg(mem)
		return true
	}
	// match: (SBBQ x load:(MOVQloadidx8 {sym} [off] ptr idx mem) carry)
	// cond: canMergeLoadLateClobber(v, load, x) && clobber(load)
	// result: (SBBQloadidx8 {sym} [off] ptr idx x carry mem)
	for {
		_ = v.Args[2]
		x := v.Args[0]
		load := v.Args[1]
		if load.Op != OpAMD64MOVQloadidx8 {
			break
		}
		off := load.AuxInt
		sym := load.Aux
		_ = load.Args[2]
		ptr := load.Args[0]
		idx := load.Args[1]
		mem := load.Args[2]
		carry := v.Args[2]
		if !(canMergeLoadLateClobber(v, load, x) && clobber(load)) {
			break
		}
		v.reset(OpAMD64SBBQloadidx8)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(carry)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteBlockAMD64merge(b *Block) bool {
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
