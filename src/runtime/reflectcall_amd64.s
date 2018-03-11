// Code generated by gen_reflect_call.go. DO NOT EDIT.

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"
#include "funcdata.h"

// reflect.call calls a function with the given argument list.
//   func call(argtype *_type, f *FuncVal, arg *byte, argsize, retoffset uint32)
// We do not have variable-sized frames, so we use a small number
// of constant-sized-frame functions to encode a few bits of size in the PC.
TEXT reflect·call(SB), NOSPLIT, $0-0
	JMP	·reflectcall(SB)

// callRet copies return values back at the end of call*. This is a
// separate function so it can allocate stack space for the arguments
// to reflectcallmove. It does not follow the Go ABI; it expects its
// arguments in registers.
TEXT callRet<>(SB), NOSPLIT, $32-0
	NO_LOCAL_POINTERS
	MOVQ	DX, 0(SP)
	MOVQ	DI, 8(SP)
	MOVQ	SI, 16(SP)
	MOVQ	CX, 24(SP)
	CALL	runtime·reflectcallmove(SB)
	RET

TEXT ·reflectcall(SB), NOSPLIT, $0-32
	MOVLQZX argsize+24(FP), CX
	CMPQ	CX, $32
	JA	3(PC)
	MOVQ	$runtime·call32(SB), AX
	JMP	AX
	CMPQ	CX, $64
	JA	3(PC)
	MOVQ	$runtime·call64(SB), AX
	JMP	AX
	CMPQ	CX, $128
	JA	3(PC)
	MOVQ	$runtime·call128(SB), AX
	JMP	AX
	CMPQ	CX, $256
	JA	3(PC)
	MOVQ	$runtime·call256(SB), AX
	JMP	AX
	CMPQ	CX, $512
	JA	3(PC)
	MOVQ	$runtime·call512(SB), AX
	JMP	AX
	CMPQ	CX, $1024
	JA	3(PC)
	MOVQ	$runtime·call1024(SB), AX
	JMP	AX
	CMPQ	CX, $2048
	JA	3(PC)
	MOVQ	$runtime·call2048(SB), AX
	JMP	AX
	CMPQ	CX, $4096
	JA	3(PC)
	MOVQ	$runtime·call4096(SB), AX
	JMP	AX
	CMPQ	CX, $8192
	JA	3(PC)
	MOVQ	$runtime·call8192(SB), AX
	JMP	AX
	CMPQ	CX, $16384
	JA	3(PC)
	MOVQ	$runtime·call16384(SB), AX
	JMP	AX
	CMPQ	CX, $32768
	JA	3(PC)
	MOVQ	$runtime·call32768(SB), AX
	JMP	AX
	CMPQ	CX, $65536
	JA	3(PC)
	MOVQ	$runtime·call65536(SB), AX
	JMP	AX
	CMPQ	CX, $131072
	JA	3(PC)
	MOVQ	$runtime·call131072(SB), AX
	JMP	AX
	CMPQ	CX, $262144
	JA	3(PC)
	MOVQ	$runtime·call262144(SB), AX
	JMP	AX
	CMPQ	CX, $524288
	JA	3(PC)
	MOVQ	$runtime·call524288(SB), AX
	JMP	AX
	CMPQ	CX, $1048576
	JA	3(PC)
	MOVQ	$runtime·call1048576(SB), AX
	JMP	AX
	CMPQ	CX, $2097152
	JA	3(PC)
	MOVQ	$runtime·call2097152(SB), AX
	JMP	AX
	CMPQ	CX, $4194304
	JA	3(PC)
	MOVQ	$runtime·call4194304(SB), AX
	JMP	AX
	CMPQ	CX, $8388608
	JA	3(PC)
	MOVQ	$runtime·call8388608(SB), AX
	JMP	AX
	CMPQ	CX, $16777216
	JA	3(PC)
	MOVQ	$runtime·call16777216(SB), AX
	JMP	AX
	CMPQ	CX, $33554432
	JA	3(PC)
	MOVQ	$runtime·call33554432(SB), AX
	JMP	AX
	CMPQ	CX, $67108864
	JA	3(PC)
	MOVQ	$runtime·call67108864(SB), AX
	JMP	AX
	CMPQ	CX, $134217728
	JA	3(PC)
	MOVQ	$runtime·call134217728(SB), AX
	JMP	AX
	CMPQ	CX, $268435456
	JA	3(PC)
	MOVQ	$runtime·call268435456(SB), AX
	JMP	AX
	CMPQ	CX, $536870912
	JA	3(PC)
	MOVQ	$runtime·call536870912(SB), AX
	JMP	AX
	CMPQ	CX, $1073741824
	JA	3(PC)
	MOVQ	$runtime·call1073741824(SB), AX
	JMP	AX

	MOVQ	$runtime·badreflectcall(SB), AX
	JMP	AX

TEXT ·call32(SB), WRAPPER, $32-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call64(SB), WRAPPER, $64-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call128(SB), WRAPPER, $128-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call256(SB), WRAPPER, $256-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call512(SB), WRAPPER, $512-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call1024(SB), WRAPPER, $1024-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call2048(SB), WRAPPER, $2048-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call4096(SB), WRAPPER, $4096-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call8192(SB), WRAPPER, $8192-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call16384(SB), WRAPPER, $16384-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call32768(SB), WRAPPER, $32768-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call65536(SB), WRAPPER, $65536-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call131072(SB), WRAPPER, $131072-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call262144(SB), WRAPPER, $262144-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call524288(SB), WRAPPER, $524288-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call1048576(SB), WRAPPER, $1048576-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call2097152(SB), WRAPPER, $2097152-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call4194304(SB), WRAPPER, $4194304-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call8388608(SB), WRAPPER, $8388608-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call16777216(SB), WRAPPER, $16777216-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call33554432(SB), WRAPPER, $33554432-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call67108864(SB), WRAPPER, $67108864-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call134217728(SB), WRAPPER, $134217728-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call268435456(SB), WRAPPER, $268435456-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call536870912(SB), WRAPPER, $536870912-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

TEXT ·call1073741824(SB), WRAPPER, $1073741824-32
	NO_LOCAL_POINTERS
	// copy arguments to stack
	MOVQ	argptr+16(FP), SI
	MOVLQZX argsize+24(FP), CX
	MOVQ	SP, DI
	REP;MOVSB
	// call function
	MOVQ	f+8(FP), DX
	PCDATA  $PCDATA_StackMapIndex, $0
	CALL	(DX)
	// copy return values back
	MOVQ	argtype+0(FP), DX
	MOVQ	argptr+16(FP), DI
	MOVLQZX	argsize+24(FP), CX
	MOVLQZX	retoffset+28(FP), BX
	MOVQ	SP, SI
	ADDQ	BX, DI
	ADDQ	BX, SI
	SUBQ	BX, CX
	CALL	callRet<>(SB)
	RET

