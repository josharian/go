// Code generated from gen/*Ops.go; DO NOT EDIT.

package ssa

import (
	"cmd/internal/obj"
	"cmd/internal/obj/arm"
	"cmd/internal/obj/arm64"
	"cmd/internal/obj/mips"
	"cmd/internal/obj/ppc64"
	"cmd/internal/obj/s390x"
	"cmd/internal/obj/wasm"
	"cmd/internal/obj/x86"
)

const (
	BlockInvalid BlockKind = iota

	Block386EQ
	Block386NE
	Block386LT
	Block386LE
	Block386GT
	Block386GE
	Block386OS
	Block386OC
	Block386ULT
	Block386ULE
	Block386UGT
	Block386UGE
	Block386EQF
	Block386NEF
	Block386ORD
	Block386NAN

	BlockAMD64EQ
	BlockAMD64NE
	BlockAMD64LT
	BlockAMD64LE
	BlockAMD64GT
	BlockAMD64GE
	BlockAMD64OS
	BlockAMD64OC
	BlockAMD64ULT
	BlockAMD64ULE
	BlockAMD64UGT
	BlockAMD64UGE
	BlockAMD64EQF
	BlockAMD64NEF
	BlockAMD64ORD
	BlockAMD64NAN

	BlockARMEQ
	BlockARMNE
	BlockARMLT
	BlockARMLE
	BlockARMGT
	BlockARMGE
	BlockARMULT
	BlockARMULE
	BlockARMUGT
	BlockARMUGE

	BlockARM64EQ
	BlockARM64NE
	BlockARM64LT
	BlockARM64LE
	BlockARM64GT
	BlockARM64GE
	BlockARM64ULT
	BlockARM64ULE
	BlockARM64UGT
	BlockARM64UGE
	BlockARM64Z
	BlockARM64NZ
	BlockARM64ZW
	BlockARM64NZW
	BlockARM64TBZ
	BlockARM64TBNZ
	BlockARM64FLT
	BlockARM64FLE
	BlockARM64FGT
	BlockARM64FGE

	BlockMIPSEQ
	BlockMIPSNE
	BlockMIPSLTZ
	BlockMIPSLEZ
	BlockMIPSGTZ
	BlockMIPSGEZ
	BlockMIPSFPT
	BlockMIPSFPF

	BlockMIPS64EQ
	BlockMIPS64NE
	BlockMIPS64LTZ
	BlockMIPS64LEZ
	BlockMIPS64GTZ
	BlockMIPS64GEZ
	BlockMIPS64FPT
	BlockMIPS64FPF

	BlockPPC64EQ
	BlockPPC64NE
	BlockPPC64LT
	BlockPPC64LE
	BlockPPC64GT
	BlockPPC64GE
	BlockPPC64FLT
	BlockPPC64FLE
	BlockPPC64FGT
	BlockPPC64FGE

	BlockS390XEQ
	BlockS390XNE
	BlockS390XLT
	BlockS390XLE
	BlockS390XGT
	BlockS390XGE
	BlockS390XGTF
	BlockS390XGEF

	BlockPlain
	BlockIf
	BlockDefer
	BlockRet
	BlockRetJmp
	BlockExit
	BlockFirst
)

var blockString = [...]string{
	BlockInvalid: "BlockInvalid",

	Block386EQ:  "EQ",
	Block386NE:  "NE",
	Block386LT:  "LT",
	Block386LE:  "LE",
	Block386GT:  "GT",
	Block386GE:  "GE",
	Block386OS:  "OS",
	Block386OC:  "OC",
	Block386ULT: "ULT",
	Block386ULE: "ULE",
	Block386UGT: "UGT",
	Block386UGE: "UGE",
	Block386EQF: "EQF",
	Block386NEF: "NEF",
	Block386ORD: "ORD",
	Block386NAN: "NAN",

	BlockAMD64EQ:  "EQ",
	BlockAMD64NE:  "NE",
	BlockAMD64LT:  "LT",
	BlockAMD64LE:  "LE",
	BlockAMD64GT:  "GT",
	BlockAMD64GE:  "GE",
	BlockAMD64OS:  "OS",
	BlockAMD64OC:  "OC",
	BlockAMD64ULT: "ULT",
	BlockAMD64ULE: "ULE",
	BlockAMD64UGT: "UGT",
	BlockAMD64UGE: "UGE",
	BlockAMD64EQF: "EQF",
	BlockAMD64NEF: "NEF",
	BlockAMD64ORD: "ORD",
	BlockAMD64NAN: "NAN",

	BlockARMEQ:  "EQ",
	BlockARMNE:  "NE",
	BlockARMLT:  "LT",
	BlockARMLE:  "LE",
	BlockARMGT:  "GT",
	BlockARMGE:  "GE",
	BlockARMULT: "ULT",
	BlockARMULE: "ULE",
	BlockARMUGT: "UGT",
	BlockARMUGE: "UGE",

	BlockARM64EQ:   "EQ",
	BlockARM64NE:   "NE",
	BlockARM64LT:   "LT",
	BlockARM64LE:   "LE",
	BlockARM64GT:   "GT",
	BlockARM64GE:   "GE",
	BlockARM64ULT:  "ULT",
	BlockARM64ULE:  "ULE",
	BlockARM64UGT:  "UGT",
	BlockARM64UGE:  "UGE",
	BlockARM64Z:    "Z",
	BlockARM64NZ:   "NZ",
	BlockARM64ZW:   "ZW",
	BlockARM64NZW:  "NZW",
	BlockARM64TBZ:  "TBZ",
	BlockARM64TBNZ: "TBNZ",
	BlockARM64FLT:  "FLT",
	BlockARM64FLE:  "FLE",
	BlockARM64FGT:  "FGT",
	BlockARM64FGE:  "FGE",

	BlockMIPSEQ:  "EQ",
	BlockMIPSNE:  "NE",
	BlockMIPSLTZ: "LTZ",
	BlockMIPSLEZ: "LEZ",
	BlockMIPSGTZ: "GTZ",
	BlockMIPSGEZ: "GEZ",
	BlockMIPSFPT: "FPT",
	BlockMIPSFPF: "FPF",

	BlockMIPS64EQ:  "EQ",
	BlockMIPS64NE:  "NE",
	BlockMIPS64LTZ: "LTZ",
	BlockMIPS64LEZ: "LEZ",
	BlockMIPS64GTZ: "GTZ",
	BlockMIPS64GEZ: "GEZ",
	BlockMIPS64FPT: "FPT",
	BlockMIPS64FPF: "FPF",

	BlockPPC64EQ:  "EQ",
	BlockPPC64NE:  "NE",
	BlockPPC64LT:  "LT",
	BlockPPC64LE:  "LE",
	BlockPPC64GT:  "GT",
	BlockPPC64GE:  "GE",
	BlockPPC64FLT: "FLT",
	BlockPPC64FLE: "FLE",
	BlockPPC64FGT: "FGT",
	BlockPPC64FGE: "FGE",

	BlockS390XEQ:  "EQ",
	BlockS390XNE:  "NE",
	BlockS390XLT:  "LT",
	BlockS390XLE:  "LE",
	BlockS390XGT:  "GT",
	BlockS390XGE:  "GE",
	BlockS390XGTF: "GTF",
	BlockS390XGEF: "GEF",

	BlockPlain:  "Plain",
	BlockIf:     "If",
	BlockDefer:  "Defer",
	BlockRet:    "Ret",
	BlockRetJmp: "RetJmp",
	BlockExit:   "Exit",
	BlockFirst:  "First",
}

func (k BlockKind) String() string { return blockString[k] }

const (
	OpInvalid Op = iota

	Op386ADDSS                Op = 1
	Op386ADDSD                Op = 2
	Op386SUBSS                Op = 3
	Op386SUBSD                Op = 4
	Op386MULSS                Op = 5
	Op386MULSD                Op = 6
	Op386DIVSS                Op = 7
	Op386DIVSD                Op = 8
	Op386MOVSSload            Op = 9
	Op386MOVSDload            Op = 10
	Op386MOVSSconst           Op = 11
	Op386MOVSDconst           Op = 12
	Op386MOVSSloadidx1        Op = 13
	Op386MOVSSloadidx4        Op = 14
	Op386MOVSDloadidx1        Op = 15
	Op386MOVSDloadidx8        Op = 16
	Op386MOVSSstore           Op = 17
	Op386MOVSDstore           Op = 18
	Op386MOVSSstoreidx1       Op = 19
	Op386MOVSSstoreidx4       Op = 20
	Op386MOVSDstoreidx1       Op = 21
	Op386MOVSDstoreidx8       Op = 22
	Op386ADDSSload            Op = 23
	Op386ADDSDload            Op = 24
	Op386SUBSSload            Op = 25
	Op386SUBSDload            Op = 26
	Op386MULSSload            Op = 27
	Op386MULSDload            Op = 28
	Op386DIVSSload            Op = 29
	Op386DIVSDload            Op = 30
	Op386ADDL                 Op = 31
	Op386ADDLconst            Op = 32
	Op386ADDLcarry            Op = 33
	Op386ADDLconstcarry       Op = 34
	Op386ADCL                 Op = 35
	Op386ADCLconst            Op = 36
	Op386SUBL                 Op = 37
	Op386SUBLconst            Op = 38
	Op386SUBLcarry            Op = 39
	Op386SUBLconstcarry       Op = 40
	Op386SBBL                 Op = 41
	Op386SBBLconst            Op = 42
	Op386MULL                 Op = 43
	Op386MULLconst            Op = 44
	Op386MULLU                Op = 45
	Op386HMULL                Op = 46
	Op386HMULLU               Op = 47
	Op386MULLQU               Op = 48
	Op386AVGLU                Op = 49
	Op386DIVL                 Op = 50
	Op386DIVW                 Op = 51
	Op386DIVLU                Op = 52
	Op386DIVWU                Op = 53
	Op386MODL                 Op = 54
	Op386MODW                 Op = 55
	Op386MODLU                Op = 56
	Op386MODWU                Op = 57
	Op386ANDL                 Op = 58
	Op386ANDLconst            Op = 59
	Op386ORL                  Op = 60
	Op386ORLconst             Op = 61
	Op386XORL                 Op = 62
	Op386XORLconst            Op = 63
	Op386CMPL                 Op = 64
	Op386CMPW                 Op = 65
	Op386CMPB                 Op = 66
	Op386CMPLconst            Op = 67
	Op386CMPWconst            Op = 68
	Op386CMPBconst            Op = 69
	Op386CMPLload             Op = 70
	Op386CMPWload             Op = 71
	Op386CMPBload             Op = 72
	Op386CMPLconstload        Op = 73
	Op386CMPWconstload        Op = 74
	Op386CMPBconstload        Op = 75
	Op386UCOMISS              Op = 76
	Op386UCOMISD              Op = 77
	Op386TESTL                Op = 78
	Op386TESTW                Op = 79
	Op386TESTB                Op = 80
	Op386TESTLconst           Op = 81
	Op386TESTWconst           Op = 82
	Op386TESTBconst           Op = 83
	Op386SHLL                 Op = 84
	Op386SHLLconst            Op = 85
	Op386SHRL                 Op = 86
	Op386SHRW                 Op = 87
	Op386SHRB                 Op = 88
	Op386SHRLconst            Op = 89
	Op386SHRWconst            Op = 90
	Op386SHRBconst            Op = 91
	Op386SARL                 Op = 92
	Op386SARW                 Op = 93
	Op386SARB                 Op = 94
	Op386SARLconst            Op = 95
	Op386SARWconst            Op = 96
	Op386SARBconst            Op = 97
	Op386ROLLconst            Op = 98
	Op386ROLWconst            Op = 99
	Op386ROLBconst            Op = 100
	Op386ADDLload             Op = 101
	Op386SUBLload             Op = 102
	Op386MULLload             Op = 103
	Op386ANDLload             Op = 104
	Op386ORLload              Op = 105
	Op386XORLload             Op = 106
	Op386ADDLloadidx4         Op = 107
	Op386SUBLloadidx4         Op = 108
	Op386MULLloadidx4         Op = 109
	Op386ANDLloadidx4         Op = 110
	Op386ORLloadidx4          Op = 111
	Op386XORLloadidx4         Op = 112
	Op386NEGL                 Op = 113
	Op386NOTL                 Op = 114
	Op386BSFL                 Op = 115
	Op386BSFW                 Op = 116
	Op386BSRL                 Op = 117
	Op386BSRW                 Op = 118
	Op386BSWAPL               Op = 119
	Op386SQRTSD               Op = 120
	Op386SBBLcarrymask        Op = 121
	Op386SETEQ                Op = 122
	Op386SETNE                Op = 123
	Op386SETL                 Op = 124
	Op386SETLE                Op = 125
	Op386SETG                 Op = 126
	Op386SETGE                Op = 127
	Op386SETB                 Op = 128
	Op386SETBE                Op = 129
	Op386SETA                 Op = 130
	Op386SETAE                Op = 131
	Op386SETO                 Op = 132
	Op386SETEQF               Op = 133
	Op386SETNEF               Op = 134
	Op386SETORD               Op = 135
	Op386SETNAN               Op = 136
	Op386SETGF                Op = 137
	Op386SETGEF               Op = 138
	Op386MOVBLSX              Op = 139
	Op386MOVBLZX              Op = 140
	Op386MOVWLSX              Op = 141
	Op386MOVWLZX              Op = 142
	Op386MOVLconst            Op = 143
	Op386CVTTSD2SL            Op = 144
	Op386CVTTSS2SL            Op = 145
	Op386CVTSL2SS             Op = 146
	Op386CVTSL2SD             Op = 147
	Op386CVTSD2SS             Op = 148
	Op386CVTSS2SD             Op = 149
	Op386PXOR                 Op = 150
	Op386LEAL                 Op = 151
	Op386LEAL1                Op = 152
	Op386LEAL2                Op = 153
	Op386LEAL4                Op = 154
	Op386LEAL8                Op = 155
	Op386MOVBload             Op = 156
	Op386MOVBLSXload          Op = 157
	Op386MOVWload             Op = 158
	Op386MOVWLSXload          Op = 159
	Op386MOVLload             Op = 160
	Op386MOVBstore            Op = 161
	Op386MOVWstore            Op = 162
	Op386MOVLstore            Op = 163
	Op386ADDLmodify           Op = 164
	Op386SUBLmodify           Op = 165
	Op386ANDLmodify           Op = 166
	Op386ORLmodify            Op = 167
	Op386XORLmodify           Op = 168
	Op386ADDLmodifyidx4       Op = 169
	Op386SUBLmodifyidx4       Op = 170
	Op386ANDLmodifyidx4       Op = 171
	Op386ORLmodifyidx4        Op = 172
	Op386XORLmodifyidx4       Op = 173
	Op386ADDLconstmodify      Op = 174
	Op386ANDLconstmodify      Op = 175
	Op386ORLconstmodify       Op = 176
	Op386XORLconstmodify      Op = 177
	Op386ADDLconstmodifyidx4  Op = 178
	Op386ANDLconstmodifyidx4  Op = 179
	Op386ORLconstmodifyidx4   Op = 180
	Op386XORLconstmodifyidx4  Op = 181
	Op386MOVBloadidx1         Op = 182
	Op386MOVWloadidx1         Op = 183
	Op386MOVWloadidx2         Op = 184
	Op386MOVLloadidx1         Op = 185
	Op386MOVLloadidx4         Op = 186
	Op386MOVBstoreidx1        Op = 187
	Op386MOVWstoreidx1        Op = 188
	Op386MOVWstoreidx2        Op = 189
	Op386MOVLstoreidx1        Op = 190
	Op386MOVLstoreidx4        Op = 191
	Op386MOVBstoreconst       Op = 192
	Op386MOVWstoreconst       Op = 193
	Op386MOVLstoreconst       Op = 194
	Op386MOVBstoreconstidx1   Op = 195
	Op386MOVWstoreconstidx1   Op = 196
	Op386MOVWstoreconstidx2   Op = 197
	Op386MOVLstoreconstidx1   Op = 198
	Op386MOVLstoreconstidx4   Op = 199
	Op386DUFFZERO             Op = 200
	Op386REPSTOSL             Op = 201
	Op386CALLstatic           Op = 202
	Op386CALLclosure          Op = 203
	Op386CALLinter            Op = 204
	Op386DUFFCOPY             Op = 205
	Op386REPMOVSL             Op = 206
	Op386InvertFlags          Op = 207
	Op386LoweredGetG          Op = 208
	Op386LoweredGetClosurePtr Op = 209
	Op386LoweredGetCallerPC   Op = 210
	Op386LoweredGetCallerSP   Op = 211
	Op386LoweredNilCheck      Op = 212
	Op386LoweredWB            Op = 213
	Op386LoweredPanicBoundsA  Op = 214
	Op386LoweredPanicBoundsB  Op = 215
	Op386LoweredPanicBoundsC  Op = 216
	Op386LoweredPanicExtendA  Op = 217
	Op386LoweredPanicExtendB  Op = 218
	Op386LoweredPanicExtendC  Op = 219
	Op386FlagEQ               Op = 220
	Op386FlagLT_ULT           Op = 221
	Op386FlagLT_UGT           Op = 222
	Op386FlagGT_UGT           Op = 223
	Op386FlagGT_ULT           Op = 224
	Op386FCHS                 Op = 225
	Op386MOVSSconst1          Op = 226
	Op386MOVSDconst1          Op = 227
	Op386MOVSSconst2          Op = 228
	Op386MOVSDconst2          Op = 229

	OpAMD64ADDSS                Op = 230
	OpAMD64ADDSD                Op = 231
	OpAMD64SUBSS                Op = 232
	OpAMD64SUBSD                Op = 233
	OpAMD64MULSS                Op = 234
	OpAMD64MULSD                Op = 235
	OpAMD64DIVSS                Op = 236
	OpAMD64DIVSD                Op = 237
	OpAMD64MOVSSload            Op = 238
	OpAMD64MOVSDload            Op = 239
	OpAMD64MOVSSconst           Op = 240
	OpAMD64MOVSDconst           Op = 241
	OpAMD64MOVSSloadidx1        Op = 242
	OpAMD64MOVSSloadidx4        Op = 243
	OpAMD64MOVSDloadidx1        Op = 244
	OpAMD64MOVSDloadidx8        Op = 245
	OpAMD64MOVSSstore           Op = 246
	OpAMD64MOVSDstore           Op = 247
	OpAMD64MOVSSstoreidx1       Op = 248
	OpAMD64MOVSSstoreidx4       Op = 249
	OpAMD64MOVSDstoreidx1       Op = 250
	OpAMD64MOVSDstoreidx8       Op = 251
	OpAMD64ADDSSload            Op = 252
	OpAMD64ADDSDload            Op = 253
	OpAMD64SUBSSload            Op = 254
	OpAMD64SUBSDload            Op = 255
	OpAMD64MULSSload            Op = 256
	OpAMD64MULSDload            Op = 257
	OpAMD64DIVSSload            Op = 258
	OpAMD64DIVSDload            Op = 259
	OpAMD64ADDQ                 Op = 260
	OpAMD64ADDL                 Op = 261
	OpAMD64ADDQconst            Op = 262
	OpAMD64ADDLconst            Op = 263
	OpAMD64ADDQconstmodify      Op = 264
	OpAMD64ADDLconstmodify      Op = 265
	OpAMD64SUBQ                 Op = 266
	OpAMD64SUBL                 Op = 267
	OpAMD64SUBQconst            Op = 268
	OpAMD64SUBLconst            Op = 269
	OpAMD64MULQ                 Op = 270
	OpAMD64MULL                 Op = 271
	OpAMD64MULQconst            Op = 272
	OpAMD64MULLconst            Op = 273
	OpAMD64MULLU                Op = 274
	OpAMD64MULQU                Op = 275
	OpAMD64HMULQ                Op = 276
	OpAMD64HMULL                Op = 277
	OpAMD64HMULQU               Op = 278
	OpAMD64HMULLU               Op = 279
	OpAMD64AVGQU                Op = 280
	OpAMD64DIVQ                 Op = 281
	OpAMD64DIVL                 Op = 282
	OpAMD64DIVW                 Op = 283
	OpAMD64DIVQU                Op = 284
	OpAMD64DIVLU                Op = 285
	OpAMD64DIVWU                Op = 286
	OpAMD64NEGLflags            Op = 287
	OpAMD64ADDQcarry            Op = 288
	OpAMD64ADCQ                 Op = 289
	OpAMD64ADDQconstcarry       Op = 290
	OpAMD64ADCQconst            Op = 291
	OpAMD64SUBQborrow           Op = 292
	OpAMD64SBBQ                 Op = 293
	OpAMD64SUBQconstborrow      Op = 294
	OpAMD64SBBQconst            Op = 295
	OpAMD64MULQU2               Op = 296
	OpAMD64DIVQU2               Op = 297
	OpAMD64ANDQ                 Op = 298
	OpAMD64ANDL                 Op = 299
	OpAMD64ANDQconst            Op = 300
	OpAMD64ANDLconst            Op = 301
	OpAMD64ANDQconstmodify      Op = 302
	OpAMD64ANDLconstmodify      Op = 303
	OpAMD64ORQ                  Op = 304
	OpAMD64ORL                  Op = 305
	OpAMD64ORQconst             Op = 306
	OpAMD64ORLconst             Op = 307
	OpAMD64ORQconstmodify       Op = 308
	OpAMD64ORLconstmodify       Op = 309
	OpAMD64XORQ                 Op = 310
	OpAMD64XORL                 Op = 311
	OpAMD64XORQconst            Op = 312
	OpAMD64XORLconst            Op = 313
	OpAMD64XORQconstmodify      Op = 314
	OpAMD64XORLconstmodify      Op = 315
	OpAMD64CMPQ                 Op = 316
	OpAMD64CMPL                 Op = 317
	OpAMD64CMPW                 Op = 318
	OpAMD64CMPB                 Op = 319
	OpAMD64CMPQconst            Op = 320
	OpAMD64CMPLconst            Op = 321
	OpAMD64CMPWconst            Op = 322
	OpAMD64CMPBconst            Op = 323
	OpAMD64CMPQload             Op = 324
	OpAMD64CMPLload             Op = 325
	OpAMD64CMPWload             Op = 326
	OpAMD64CMPBload             Op = 327
	OpAMD64CMPQconstload        Op = 328
	OpAMD64CMPLconstload        Op = 329
	OpAMD64CMPWconstload        Op = 330
	OpAMD64CMPBconstload        Op = 331
	OpAMD64UCOMISS              Op = 332
	OpAMD64UCOMISD              Op = 333
	OpAMD64BTL                  Op = 334
	OpAMD64BTQ                  Op = 335
	OpAMD64BTCL                 Op = 336
	OpAMD64BTCQ                 Op = 337
	OpAMD64BTRL                 Op = 338
	OpAMD64BTRQ                 Op = 339
	OpAMD64BTSL                 Op = 340
	OpAMD64BTSQ                 Op = 341
	OpAMD64BTLconst             Op = 342
	OpAMD64BTQconst             Op = 343
	OpAMD64BTCLconst            Op = 344
	OpAMD64BTCQconst            Op = 345
	OpAMD64BTRLconst            Op = 346
	OpAMD64BTRQconst            Op = 347
	OpAMD64BTSLconst            Op = 348
	OpAMD64BTSQconst            Op = 349
	OpAMD64BTCQmodify           Op = 350
	OpAMD64BTCLmodify           Op = 351
	OpAMD64BTSQmodify           Op = 352
	OpAMD64BTSLmodify           Op = 353
	OpAMD64BTRQmodify           Op = 354
	OpAMD64BTRLmodify           Op = 355
	OpAMD64BTCQconstmodify      Op = 356
	OpAMD64BTCLconstmodify      Op = 357
	OpAMD64BTSQconstmodify      Op = 358
	OpAMD64BTSLconstmodify      Op = 359
	OpAMD64BTRQconstmodify      Op = 360
	OpAMD64BTRLconstmodify      Op = 361
	OpAMD64TESTQ                Op = 362
	OpAMD64TESTL                Op = 363
	OpAMD64TESTW                Op = 364
	OpAMD64TESTB                Op = 365
	OpAMD64TESTQconst           Op = 366
	OpAMD64TESTLconst           Op = 367
	OpAMD64TESTWconst           Op = 368
	OpAMD64TESTBconst           Op = 369
	OpAMD64SHLQ                 Op = 370
	OpAMD64SHLL                 Op = 371
	OpAMD64SHLQconst            Op = 372
	OpAMD64SHLLconst            Op = 373
	OpAMD64SHRQ                 Op = 374
	OpAMD64SHRL                 Op = 375
	OpAMD64SHRW                 Op = 376
	OpAMD64SHRB                 Op = 377
	OpAMD64SHRQconst            Op = 378
	OpAMD64SHRLconst            Op = 379
	OpAMD64SHRWconst            Op = 380
	OpAMD64SHRBconst            Op = 381
	OpAMD64SARQ                 Op = 382
	OpAMD64SARL                 Op = 383
	OpAMD64SARW                 Op = 384
	OpAMD64SARB                 Op = 385
	OpAMD64SARQconst            Op = 386
	OpAMD64SARLconst            Op = 387
	OpAMD64SARWconst            Op = 388
	OpAMD64SARBconst            Op = 389
	OpAMD64ROLQ                 Op = 390
	OpAMD64ROLL                 Op = 391
	OpAMD64ROLW                 Op = 392
	OpAMD64ROLB                 Op = 393
	OpAMD64RORQ                 Op = 394
	OpAMD64RORL                 Op = 395
	OpAMD64RORW                 Op = 396
	OpAMD64RORB                 Op = 397
	OpAMD64ROLQconst            Op = 398
	OpAMD64ROLLconst            Op = 399
	OpAMD64ROLWconst            Op = 400
	OpAMD64ROLBconst            Op = 401
	OpAMD64ADDLload             Op = 402
	OpAMD64ADDQload             Op = 403
	OpAMD64SUBQload             Op = 404
	OpAMD64SUBLload             Op = 405
	OpAMD64ANDLload             Op = 406
	OpAMD64ANDQload             Op = 407
	OpAMD64ORQload              Op = 408
	OpAMD64ORLload              Op = 409
	OpAMD64XORQload             Op = 410
	OpAMD64XORLload             Op = 411
	OpAMD64ADDQmodify           Op = 412
	OpAMD64SUBQmodify           Op = 413
	OpAMD64ANDQmodify           Op = 414
	OpAMD64ORQmodify            Op = 415
	OpAMD64XORQmodify           Op = 416
	OpAMD64ADDLmodify           Op = 417
	OpAMD64SUBLmodify           Op = 418
	OpAMD64ANDLmodify           Op = 419
	OpAMD64ORLmodify            Op = 420
	OpAMD64XORLmodify           Op = 421
	OpAMD64NEGQ                 Op = 422
	OpAMD64NEGL                 Op = 423
	OpAMD64NOTQ                 Op = 424
	OpAMD64NOTL                 Op = 425
	OpAMD64BSFQ                 Op = 426
	OpAMD64BSFL                 Op = 427
	OpAMD64BSRQ                 Op = 428
	OpAMD64BSRL                 Op = 429
	OpAMD64CMOVQEQ              Op = 430
	OpAMD64CMOVQNE              Op = 431
	OpAMD64CMOVQLT              Op = 432
	OpAMD64CMOVQGT              Op = 433
	OpAMD64CMOVQLE              Op = 434
	OpAMD64CMOVQGE              Op = 435
	OpAMD64CMOVQLS              Op = 436
	OpAMD64CMOVQHI              Op = 437
	OpAMD64CMOVQCC              Op = 438
	OpAMD64CMOVQCS              Op = 439
	OpAMD64CMOVLEQ              Op = 440
	OpAMD64CMOVLNE              Op = 441
	OpAMD64CMOVLLT              Op = 442
	OpAMD64CMOVLGT              Op = 443
	OpAMD64CMOVLLE              Op = 444
	OpAMD64CMOVLGE              Op = 445
	OpAMD64CMOVLLS              Op = 446
	OpAMD64CMOVLHI              Op = 447
	OpAMD64CMOVLCC              Op = 448
	OpAMD64CMOVLCS              Op = 449
	OpAMD64CMOVWEQ              Op = 450
	OpAMD64CMOVWNE              Op = 451
	OpAMD64CMOVWLT              Op = 452
	OpAMD64CMOVWGT              Op = 453
	OpAMD64CMOVWLE              Op = 454
	OpAMD64CMOVWGE              Op = 455
	OpAMD64CMOVWLS              Op = 456
	OpAMD64CMOVWHI              Op = 457
	OpAMD64CMOVWCC              Op = 458
	OpAMD64CMOVWCS              Op = 459
	OpAMD64CMOVQEQF             Op = 460
	OpAMD64CMOVQNEF             Op = 461
	OpAMD64CMOVQGTF             Op = 462
	OpAMD64CMOVQGEF             Op = 463
	OpAMD64CMOVLEQF             Op = 464
	OpAMD64CMOVLNEF             Op = 465
	OpAMD64CMOVLGTF             Op = 466
	OpAMD64CMOVLGEF             Op = 467
	OpAMD64CMOVWEQF             Op = 468
	OpAMD64CMOVWNEF             Op = 469
	OpAMD64CMOVWGTF             Op = 470
	OpAMD64CMOVWGEF             Op = 471
	OpAMD64BSWAPQ               Op = 472
	OpAMD64BSWAPL               Op = 473
	OpAMD64POPCNTQ              Op = 474
	OpAMD64POPCNTL              Op = 475
	OpAMD64SQRTSD               Op = 476
	OpAMD64ROUNDSD              Op = 477
	OpAMD64SBBQcarrymask        Op = 478
	OpAMD64SBBLcarrymask        Op = 479
	OpAMD64SETEQ                Op = 480
	OpAMD64SETNE                Op = 481
	OpAMD64SETL                 Op = 482
	OpAMD64SETLE                Op = 483
	OpAMD64SETG                 Op = 484
	OpAMD64SETGE                Op = 485
	OpAMD64SETB                 Op = 486
	OpAMD64SETBE                Op = 487
	OpAMD64SETA                 Op = 488
	OpAMD64SETAE                Op = 489
	OpAMD64SETO                 Op = 490
	OpAMD64SETEQstore           Op = 491
	OpAMD64SETNEstore           Op = 492
	OpAMD64SETLstore            Op = 493
	OpAMD64SETLEstore           Op = 494
	OpAMD64SETGstore            Op = 495
	OpAMD64SETGEstore           Op = 496
	OpAMD64SETBstore            Op = 497
	OpAMD64SETBEstore           Op = 498
	OpAMD64SETAstore            Op = 499
	OpAMD64SETAEstore           Op = 500
	OpAMD64SETEQF               Op = 501
	OpAMD64SETNEF               Op = 502
	OpAMD64SETORD               Op = 503
	OpAMD64SETNAN               Op = 504
	OpAMD64SETGF                Op = 505
	OpAMD64SETGEF               Op = 506
	OpAMD64MOVBQSX              Op = 507
	OpAMD64MOVBQZX              Op = 508
	OpAMD64MOVWQSX              Op = 509
	OpAMD64MOVWQZX              Op = 510
	OpAMD64MOVLQSX              Op = 511
	OpAMD64MOVLQZX              Op = 512
	OpAMD64MOVLconst            Op = 513
	OpAMD64MOVQconst            Op = 514
	OpAMD64CVTTSD2SL            Op = 515
	OpAMD64CVTTSD2SQ            Op = 516
	OpAMD64CVTTSS2SL            Op = 517
	OpAMD64CVTTSS2SQ            Op = 518
	OpAMD64CVTSL2SS             Op = 519
	OpAMD64CVTSL2SD             Op = 520
	OpAMD64CVTSQ2SS             Op = 521
	OpAMD64CVTSQ2SD             Op = 522
	OpAMD64CVTSD2SS             Op = 523
	OpAMD64CVTSS2SD             Op = 524
	OpAMD64MOVQi2f              Op = 525
	OpAMD64MOVQf2i              Op = 526
	OpAMD64MOVLi2f              Op = 527
	OpAMD64MOVLf2i              Op = 528
	OpAMD64PXOR                 Op = 529
	OpAMD64LEAQ                 Op = 530
	OpAMD64LEAL                 Op = 531
	OpAMD64LEAW                 Op = 532
	OpAMD64LEAQ1                Op = 533
	OpAMD64LEAL1                Op = 534
	OpAMD64LEAW1                Op = 535
	OpAMD64LEAQ2                Op = 536
	OpAMD64LEAL2                Op = 537
	OpAMD64LEAW2                Op = 538
	OpAMD64LEAQ4                Op = 539
	OpAMD64LEAL4                Op = 540
	OpAMD64LEAW4                Op = 541
	OpAMD64LEAQ8                Op = 542
	OpAMD64LEAL8                Op = 543
	OpAMD64LEAW8                Op = 544
	OpAMD64MOVBload             Op = 545
	OpAMD64MOVBQSXload          Op = 546
	OpAMD64MOVWload             Op = 547
	OpAMD64MOVWQSXload          Op = 548
	OpAMD64MOVLload             Op = 549
	OpAMD64MOVLQSXload          Op = 550
	OpAMD64MOVQload             Op = 551
	OpAMD64MOVBstore            Op = 552
	OpAMD64MOVWstore            Op = 553
	OpAMD64MOVLstore            Op = 554
	OpAMD64MOVQstore            Op = 555
	OpAMD64MOVOload             Op = 556
	OpAMD64MOVOstore            Op = 557
	OpAMD64MOVBloadidx1         Op = 558
	OpAMD64MOVWloadidx1         Op = 559
	OpAMD64MOVWloadidx2         Op = 560
	OpAMD64MOVLloadidx1         Op = 561
	OpAMD64MOVLloadidx4         Op = 562
	OpAMD64MOVLloadidx8         Op = 563
	OpAMD64MOVQloadidx1         Op = 564
	OpAMD64MOVQloadidx8         Op = 565
	OpAMD64MOVBstoreidx1        Op = 566
	OpAMD64MOVWstoreidx1        Op = 567
	OpAMD64MOVWstoreidx2        Op = 568
	OpAMD64MOVLstoreidx1        Op = 569
	OpAMD64MOVLstoreidx4        Op = 570
	OpAMD64MOVLstoreidx8        Op = 571
	OpAMD64MOVQstoreidx1        Op = 572
	OpAMD64MOVQstoreidx8        Op = 573
	OpAMD64MOVBstoreconst       Op = 574
	OpAMD64MOVWstoreconst       Op = 575
	OpAMD64MOVLstoreconst       Op = 576
	OpAMD64MOVQstoreconst       Op = 577
	OpAMD64MOVBstoreconstidx1   Op = 578
	OpAMD64MOVWstoreconstidx1   Op = 579
	OpAMD64MOVWstoreconstidx2   Op = 580
	OpAMD64MOVLstoreconstidx1   Op = 581
	OpAMD64MOVLstoreconstidx4   Op = 582
	OpAMD64MOVQstoreconstidx1   Op = 583
	OpAMD64MOVQstoreconstidx8   Op = 584
	OpAMD64DUFFZERO             Op = 585
	OpAMD64MOVOconst            Op = 586
	OpAMD64REPSTOSQ             Op = 587
	OpAMD64CALLstatic           Op = 588
	OpAMD64CALLclosure          Op = 589
	OpAMD64CALLinter            Op = 590
	OpAMD64DUFFCOPY             Op = 591
	OpAMD64REPMOVSQ             Op = 592
	OpAMD64InvertFlags          Op = 593
	OpAMD64LoweredGetG          Op = 594
	OpAMD64LoweredGetClosurePtr Op = 595
	OpAMD64LoweredGetCallerPC   Op = 596
	OpAMD64LoweredGetCallerSP   Op = 597
	OpAMD64LoweredNilCheck      Op = 598
	OpAMD64LoweredWB            Op = 599
	OpAMD64LoweredPanicBoundsA  Op = 600
	OpAMD64LoweredPanicBoundsB  Op = 601
	OpAMD64LoweredPanicBoundsC  Op = 602
	OpAMD64LoweredPanicExtendA  Op = 603
	OpAMD64LoweredPanicExtendB  Op = 604
	OpAMD64LoweredPanicExtendC  Op = 605
	OpAMD64FlagEQ               Op = 606
	OpAMD64FlagLT_ULT           Op = 607
	OpAMD64FlagLT_UGT           Op = 608
	OpAMD64FlagGT_UGT           Op = 609
	OpAMD64FlagGT_ULT           Op = 610
	OpAMD64MOVBatomicload       Op = 611
	OpAMD64MOVLatomicload       Op = 612
	OpAMD64MOVQatomicload       Op = 613
	OpAMD64XCHGL                Op = 614
	OpAMD64XCHGQ                Op = 615
	OpAMD64XADDLlock            Op = 616
	OpAMD64XADDQlock            Op = 617
	OpAMD64AddTupleFirst32      Op = 618
	OpAMD64AddTupleFirst64      Op = 619
	OpAMD64CMPXCHGLlock         Op = 620
	OpAMD64CMPXCHGQlock         Op = 621
	OpAMD64ANDBlock             Op = 622
	OpAMD64ORBlock              Op = 623

	OpARMADD                  Op = 624
	OpARMADDconst             Op = 625
	OpARMSUB                  Op = 626
	OpARMSUBconst             Op = 627
	OpARMRSB                  Op = 628
	OpARMRSBconst             Op = 629
	OpARMMUL                  Op = 630
	OpARMHMUL                 Op = 631
	OpARMHMULU                Op = 632
	OpARMCALLudiv             Op = 633
	OpARMADDS                 Op = 634
	OpARMADDSconst            Op = 635
	OpARMADC                  Op = 636
	OpARMADCconst             Op = 637
	OpARMSUBS                 Op = 638
	OpARMSUBSconst            Op = 639
	OpARMRSBSconst            Op = 640
	OpARMSBC                  Op = 641
	OpARMSBCconst             Op = 642
	OpARMRSCconst             Op = 643
	OpARMMULLU                Op = 644
	OpARMMULA                 Op = 645
	OpARMMULS                 Op = 646
	OpARMADDF                 Op = 647
	OpARMADDD                 Op = 648
	OpARMSUBF                 Op = 649
	OpARMSUBD                 Op = 650
	OpARMMULF                 Op = 651
	OpARMMULD                 Op = 652
	OpARMNMULF                Op = 653
	OpARMNMULD                Op = 654
	OpARMDIVF                 Op = 655
	OpARMDIVD                 Op = 656
	OpARMMULAF                Op = 657
	OpARMMULAD                Op = 658
	OpARMMULSF                Op = 659
	OpARMMULSD                Op = 660
	OpARMAND                  Op = 661
	OpARMANDconst             Op = 662
	OpARMOR                   Op = 663
	OpARMORconst              Op = 664
	OpARMXOR                  Op = 665
	OpARMXORconst             Op = 666
	OpARMBIC                  Op = 667
	OpARMBICconst             Op = 668
	OpARMBFX                  Op = 669
	OpARMBFXU                 Op = 670
	OpARMMVN                  Op = 671
	OpARMNEGF                 Op = 672
	OpARMNEGD                 Op = 673
	OpARMSQRTD                Op = 674
	OpARMCLZ                  Op = 675
	OpARMREV                  Op = 676
	OpARMREV16                Op = 677
	OpARMRBIT                 Op = 678
	OpARMSLL                  Op = 679
	OpARMSLLconst             Op = 680
	OpARMSRL                  Op = 681
	OpARMSRLconst             Op = 682
	OpARMSRA                  Op = 683
	OpARMSRAconst             Op = 684
	OpARMSRRconst             Op = 685
	OpARMADDshiftLL           Op = 686
	OpARMADDshiftRL           Op = 687
	OpARMADDshiftRA           Op = 688
	OpARMSUBshiftLL           Op = 689
	OpARMSUBshiftRL           Op = 690
	OpARMSUBshiftRA           Op = 691
	OpARMRSBshiftLL           Op = 692
	OpARMRSBshiftRL           Op = 693
	OpARMRSBshiftRA           Op = 694
	OpARMANDshiftLL           Op = 695
	OpARMANDshiftRL           Op = 696
	OpARMANDshiftRA           Op = 697
	OpARMORshiftLL            Op = 698
	OpARMORshiftRL            Op = 699
	OpARMORshiftRA            Op = 700
	OpARMXORshiftLL           Op = 701
	OpARMXORshiftRL           Op = 702
	OpARMXORshiftRA           Op = 703
	OpARMXORshiftRR           Op = 704
	OpARMBICshiftLL           Op = 705
	OpARMBICshiftRL           Op = 706
	OpARMBICshiftRA           Op = 707
	OpARMMVNshiftLL           Op = 708
	OpARMMVNshiftRL           Op = 709
	OpARMMVNshiftRA           Op = 710
	OpARMADCshiftLL           Op = 711
	OpARMADCshiftRL           Op = 712
	OpARMADCshiftRA           Op = 713
	OpARMSBCshiftLL           Op = 714
	OpARMSBCshiftRL           Op = 715
	OpARMSBCshiftRA           Op = 716
	OpARMRSCshiftLL           Op = 717
	OpARMRSCshiftRL           Op = 718
	OpARMRSCshiftRA           Op = 719
	OpARMADDSshiftLL          Op = 720
	OpARMADDSshiftRL          Op = 721
	OpARMADDSshiftRA          Op = 722
	OpARMSUBSshiftLL          Op = 723
	OpARMSUBSshiftRL          Op = 724
	OpARMSUBSshiftRA          Op = 725
	OpARMRSBSshiftLL          Op = 726
	OpARMRSBSshiftRL          Op = 727
	OpARMRSBSshiftRA          Op = 728
	OpARMADDshiftLLreg        Op = 729
	OpARMADDshiftRLreg        Op = 730
	OpARMADDshiftRAreg        Op = 731
	OpARMSUBshiftLLreg        Op = 732
	OpARMSUBshiftRLreg        Op = 733
	OpARMSUBshiftRAreg        Op = 734
	OpARMRSBshiftLLreg        Op = 735
	OpARMRSBshiftRLreg        Op = 736
	OpARMRSBshiftRAreg        Op = 737
	OpARMANDshiftLLreg        Op = 738
	OpARMANDshiftRLreg        Op = 739
	OpARMANDshiftRAreg        Op = 740
	OpARMORshiftLLreg         Op = 741
	OpARMORshiftRLreg         Op = 742
	OpARMORshiftRAreg         Op = 743
	OpARMXORshiftLLreg        Op = 744
	OpARMXORshiftRLreg        Op = 745
	OpARMXORshiftRAreg        Op = 746
	OpARMBICshiftLLreg        Op = 747
	OpARMBICshiftRLreg        Op = 748
	OpARMBICshiftRAreg        Op = 749
	OpARMMVNshiftLLreg        Op = 750
	OpARMMVNshiftRLreg        Op = 751
	OpARMMVNshiftRAreg        Op = 752
	OpARMADCshiftLLreg        Op = 753
	OpARMADCshiftRLreg        Op = 754
	OpARMADCshiftRAreg        Op = 755
	OpARMSBCshiftLLreg        Op = 756
	OpARMSBCshiftRLreg        Op = 757
	OpARMSBCshiftRAreg        Op = 758
	OpARMRSCshiftLLreg        Op = 759
	OpARMRSCshiftRLreg        Op = 760
	OpARMRSCshiftRAreg        Op = 761
	OpARMADDSshiftLLreg       Op = 762
	OpARMADDSshiftRLreg       Op = 763
	OpARMADDSshiftRAreg       Op = 764
	OpARMSUBSshiftLLreg       Op = 765
	OpARMSUBSshiftRLreg       Op = 766
	OpARMSUBSshiftRAreg       Op = 767
	OpARMRSBSshiftLLreg       Op = 768
	OpARMRSBSshiftRLreg       Op = 769
	OpARMRSBSshiftRAreg       Op = 770
	OpARMCMP                  Op = 771
	OpARMCMPconst             Op = 772
	OpARMCMN                  Op = 773
	OpARMCMNconst             Op = 774
	OpARMTST                  Op = 775
	OpARMTSTconst             Op = 776
	OpARMTEQ                  Op = 777
	OpARMTEQconst             Op = 778
	OpARMCMPF                 Op = 779
	OpARMCMPD                 Op = 780
	OpARMCMPshiftLL           Op = 781
	OpARMCMPshiftRL           Op = 782
	OpARMCMPshiftRA           Op = 783
	OpARMCMNshiftLL           Op = 784
	OpARMCMNshiftRL           Op = 785
	OpARMCMNshiftRA           Op = 786
	OpARMTSTshiftLL           Op = 787
	OpARMTSTshiftRL           Op = 788
	OpARMTSTshiftRA           Op = 789
	OpARMTEQshiftLL           Op = 790
	OpARMTEQshiftRL           Op = 791
	OpARMTEQshiftRA           Op = 792
	OpARMCMPshiftLLreg        Op = 793
	OpARMCMPshiftRLreg        Op = 794
	OpARMCMPshiftRAreg        Op = 795
	OpARMCMNshiftLLreg        Op = 796
	OpARMCMNshiftRLreg        Op = 797
	OpARMCMNshiftRAreg        Op = 798
	OpARMTSTshiftLLreg        Op = 799
	OpARMTSTshiftRLreg        Op = 800
	OpARMTSTshiftRAreg        Op = 801
	OpARMTEQshiftLLreg        Op = 802
	OpARMTEQshiftRLreg        Op = 803
	OpARMTEQshiftRAreg        Op = 804
	OpARMCMPF0                Op = 805
	OpARMCMPD0                Op = 806
	OpARMMOVWconst            Op = 807
	OpARMMOVFconst            Op = 808
	OpARMMOVDconst            Op = 809
	OpARMMOVWaddr             Op = 810
	OpARMMOVBload             Op = 811
	OpARMMOVBUload            Op = 812
	OpARMMOVHload             Op = 813
	OpARMMOVHUload            Op = 814
	OpARMMOVWload             Op = 815
	OpARMMOVFload             Op = 816
	OpARMMOVDload             Op = 817
	OpARMMOVBstore            Op = 818
	OpARMMOVHstore            Op = 819
	OpARMMOVWstore            Op = 820
	OpARMMOVFstore            Op = 821
	OpARMMOVDstore            Op = 822
	OpARMMOVWloadidx          Op = 823
	OpARMMOVWloadshiftLL      Op = 824
	OpARMMOVWloadshiftRL      Op = 825
	OpARMMOVWloadshiftRA      Op = 826
	OpARMMOVBUloadidx         Op = 827
	OpARMMOVBloadidx          Op = 828
	OpARMMOVHUloadidx         Op = 829
	OpARMMOVHloadidx          Op = 830
	OpARMMOVWstoreidx         Op = 831
	OpARMMOVWstoreshiftLL     Op = 832
	OpARMMOVWstoreshiftRL     Op = 833
	OpARMMOVWstoreshiftRA     Op = 834
	OpARMMOVBstoreidx         Op = 835
	OpARMMOVHstoreidx         Op = 836
	OpARMMOVBreg              Op = 837
	OpARMMOVBUreg             Op = 838
	OpARMMOVHreg              Op = 839
	OpARMMOVHUreg             Op = 840
	OpARMMOVWreg              Op = 841
	OpARMMOVWnop              Op = 842
	OpARMMOVWF                Op = 843
	OpARMMOVWD                Op = 844
	OpARMMOVWUF               Op = 845
	OpARMMOVWUD               Op = 846
	OpARMMOVFW                Op = 847
	OpARMMOVDW                Op = 848
	OpARMMOVFWU               Op = 849
	OpARMMOVDWU               Op = 850
	OpARMMOVFD                Op = 851
	OpARMMOVDF                Op = 852
	OpARMCMOVWHSconst         Op = 853
	OpARMCMOVWLSconst         Op = 854
	OpARMSRAcond              Op = 855
	OpARMCALLstatic           Op = 856
	OpARMCALLclosure          Op = 857
	OpARMCALLinter            Op = 858
	OpARMLoweredNilCheck      Op = 859
	OpARMEqual                Op = 860
	OpARMNotEqual             Op = 861
	OpARMLessThan             Op = 862
	OpARMLessEqual            Op = 863
	OpARMGreaterThan          Op = 864
	OpARMGreaterEqual         Op = 865
	OpARMLessThanU            Op = 866
	OpARMLessEqualU           Op = 867
	OpARMGreaterThanU         Op = 868
	OpARMGreaterEqualU        Op = 869
	OpARMDUFFZERO             Op = 870
	OpARMDUFFCOPY             Op = 871
	OpARMLoweredZero          Op = 872
	OpARMLoweredMove          Op = 873
	OpARMLoweredGetClosurePtr Op = 874
	OpARMLoweredGetCallerSP   Op = 875
	OpARMLoweredGetCallerPC   Op = 876
	OpARMLoweredPanicBoundsA  Op = 877
	OpARMLoweredPanicBoundsB  Op = 878
	OpARMLoweredPanicBoundsC  Op = 879
	OpARMLoweredPanicExtendA  Op = 880
	OpARMLoweredPanicExtendB  Op = 881
	OpARMLoweredPanicExtendC  Op = 882
	OpARMFlagEQ               Op = 883
	OpARMFlagLT_ULT           Op = 884
	OpARMFlagLT_UGT           Op = 885
	OpARMFlagGT_UGT           Op = 886
	OpARMFlagGT_ULT           Op = 887
	OpARMInvertFlags          Op = 888
	OpARMLoweredWB            Op = 889

	OpARM64ADCSflags                 Op = 890
	OpARM64ADCzerocarry              Op = 891
	OpARM64ADD                       Op = 892
	OpARM64ADDconst                  Op = 893
	OpARM64ADDSconstflags            Op = 894
	OpARM64ADDSflags                 Op = 895
	OpARM64SUB                       Op = 896
	OpARM64SUBconst                  Op = 897
	OpARM64SBCSflags                 Op = 898
	OpARM64SUBSflags                 Op = 899
	OpARM64MUL                       Op = 900
	OpARM64MULW                      Op = 901
	OpARM64MNEG                      Op = 902
	OpARM64MNEGW                     Op = 903
	OpARM64MULH                      Op = 904
	OpARM64UMULH                     Op = 905
	OpARM64MULL                      Op = 906
	OpARM64UMULL                     Op = 907
	OpARM64DIV                       Op = 908
	OpARM64UDIV                      Op = 909
	OpARM64DIVW                      Op = 910
	OpARM64UDIVW                     Op = 911
	OpARM64MOD                       Op = 912
	OpARM64UMOD                      Op = 913
	OpARM64MODW                      Op = 914
	OpARM64UMODW                     Op = 915
	OpARM64FADDS                     Op = 916
	OpARM64FADDD                     Op = 917
	OpARM64FSUBS                     Op = 918
	OpARM64FSUBD                     Op = 919
	OpARM64FMULS                     Op = 920
	OpARM64FMULD                     Op = 921
	OpARM64FNMULS                    Op = 922
	OpARM64FNMULD                    Op = 923
	OpARM64FDIVS                     Op = 924
	OpARM64FDIVD                     Op = 925
	OpARM64AND                       Op = 926
	OpARM64ANDconst                  Op = 927
	OpARM64OR                        Op = 928
	OpARM64ORconst                   Op = 929
	OpARM64XOR                       Op = 930
	OpARM64XORconst                  Op = 931
	OpARM64BIC                       Op = 932
	OpARM64EON                       Op = 933
	OpARM64ORN                       Op = 934
	OpARM64LoweredMuluhilo           Op = 935
	OpARM64MVN                       Op = 936
	OpARM64NEG                       Op = 937
	OpARM64NEGSflags                 Op = 938
	OpARM64NGCzerocarry              Op = 939
	OpARM64FABSD                     Op = 940
	OpARM64FNEGS                     Op = 941
	OpARM64FNEGD                     Op = 942
	OpARM64FSQRTD                    Op = 943
	OpARM64REV                       Op = 944
	OpARM64REVW                      Op = 945
	OpARM64REV16W                    Op = 946
	OpARM64RBIT                      Op = 947
	OpARM64RBITW                     Op = 948
	OpARM64CLZ                       Op = 949
	OpARM64CLZW                      Op = 950
	OpARM64VCNT                      Op = 951
	OpARM64VUADDLV                   Op = 952
	OpARM64LoweredRound32F           Op = 953
	OpARM64LoweredRound64F           Op = 954
	OpARM64FMADDS                    Op = 955
	OpARM64FMADDD                    Op = 956
	OpARM64FNMADDS                   Op = 957
	OpARM64FNMADDD                   Op = 958
	OpARM64FMSUBS                    Op = 959
	OpARM64FMSUBD                    Op = 960
	OpARM64FNMSUBS                   Op = 961
	OpARM64FNMSUBD                   Op = 962
	OpARM64MADD                      Op = 963
	OpARM64MADDW                     Op = 964
	OpARM64MSUB                      Op = 965
	OpARM64MSUBW                     Op = 966
	OpARM64SLL                       Op = 967
	OpARM64SLLconst                  Op = 968
	OpARM64SRL                       Op = 969
	OpARM64SRLconst                  Op = 970
	OpARM64SRA                       Op = 971
	OpARM64SRAconst                  Op = 972
	OpARM64ROR                       Op = 973
	OpARM64RORW                      Op = 974
	OpARM64RORconst                  Op = 975
	OpARM64RORWconst                 Op = 976
	OpARM64EXTRconst                 Op = 977
	OpARM64EXTRWconst                Op = 978
	OpARM64CMP                       Op = 979
	OpARM64CMPconst                  Op = 980
	OpARM64CMPW                      Op = 981
	OpARM64CMPWconst                 Op = 982
	OpARM64CMN                       Op = 983
	OpARM64CMNconst                  Op = 984
	OpARM64CMNW                      Op = 985
	OpARM64CMNWconst                 Op = 986
	OpARM64TST                       Op = 987
	OpARM64TSTconst                  Op = 988
	OpARM64TSTW                      Op = 989
	OpARM64TSTWconst                 Op = 990
	OpARM64FCMPS                     Op = 991
	OpARM64FCMPD                     Op = 992
	OpARM64FCMPS0                    Op = 993
	OpARM64FCMPD0                    Op = 994
	OpARM64MVNshiftLL                Op = 995
	OpARM64MVNshiftRL                Op = 996
	OpARM64MVNshiftRA                Op = 997
	OpARM64NEGshiftLL                Op = 998
	OpARM64NEGshiftRL                Op = 999
	OpARM64NEGshiftRA                Op = 1000
	OpARM64ADDshiftLL                Op = 1001
	OpARM64ADDshiftRL                Op = 1002
	OpARM64ADDshiftRA                Op = 1003
	OpARM64SUBshiftLL                Op = 1004
	OpARM64SUBshiftRL                Op = 1005
	OpARM64SUBshiftRA                Op = 1006
	OpARM64ANDshiftLL                Op = 1007
	OpARM64ANDshiftRL                Op = 1008
	OpARM64ANDshiftRA                Op = 1009
	OpARM64ORshiftLL                 Op = 1010
	OpARM64ORshiftRL                 Op = 1011
	OpARM64ORshiftRA                 Op = 1012
	OpARM64XORshiftLL                Op = 1013
	OpARM64XORshiftRL                Op = 1014
	OpARM64XORshiftRA                Op = 1015
	OpARM64BICshiftLL                Op = 1016
	OpARM64BICshiftRL                Op = 1017
	OpARM64BICshiftRA                Op = 1018
	OpARM64EONshiftLL                Op = 1019
	OpARM64EONshiftRL                Op = 1020
	OpARM64EONshiftRA                Op = 1021
	OpARM64ORNshiftLL                Op = 1022
	OpARM64ORNshiftRL                Op = 1023
	OpARM64ORNshiftRA                Op = 1024
	OpARM64CMPshiftLL                Op = 1025
	OpARM64CMPshiftRL                Op = 1026
	OpARM64CMPshiftRA                Op = 1027
	OpARM64CMNshiftLL                Op = 1028
	OpARM64CMNshiftRL                Op = 1029
	OpARM64CMNshiftRA                Op = 1030
	OpARM64TSTshiftLL                Op = 1031
	OpARM64TSTshiftRL                Op = 1032
	OpARM64TSTshiftRA                Op = 1033
	OpARM64BFI                       Op = 1034
	OpARM64BFXIL                     Op = 1035
	OpARM64SBFIZ                     Op = 1036
	OpARM64SBFX                      Op = 1037
	OpARM64UBFIZ                     Op = 1038
	OpARM64UBFX                      Op = 1039
	OpARM64MOVDconst                 Op = 1040
	OpARM64FMOVSconst                Op = 1041
	OpARM64FMOVDconst                Op = 1042
	OpARM64MOVDaddr                  Op = 1043
	OpARM64MOVBload                  Op = 1044
	OpARM64MOVBUload                 Op = 1045
	OpARM64MOVHload                  Op = 1046
	OpARM64MOVHUload                 Op = 1047
	OpARM64MOVWload                  Op = 1048
	OpARM64MOVWUload                 Op = 1049
	OpARM64MOVDload                  Op = 1050
	OpARM64FMOVSload                 Op = 1051
	OpARM64FMOVDload                 Op = 1052
	OpARM64MOVDloadidx               Op = 1053
	OpARM64MOVWloadidx               Op = 1054
	OpARM64MOVWUloadidx              Op = 1055
	OpARM64MOVHloadidx               Op = 1056
	OpARM64MOVHUloadidx              Op = 1057
	OpARM64MOVBloadidx               Op = 1058
	OpARM64MOVBUloadidx              Op = 1059
	OpARM64FMOVSloadidx              Op = 1060
	OpARM64FMOVDloadidx              Op = 1061
	OpARM64MOVHloadidx2              Op = 1062
	OpARM64MOVHUloadidx2             Op = 1063
	OpARM64MOVWloadidx4              Op = 1064
	OpARM64MOVWUloadidx4             Op = 1065
	OpARM64MOVDloadidx8              Op = 1066
	OpARM64MOVBstore                 Op = 1067
	OpARM64MOVHstore                 Op = 1068
	OpARM64MOVWstore                 Op = 1069
	OpARM64MOVDstore                 Op = 1070
	OpARM64STP                       Op = 1071
	OpARM64FMOVSstore                Op = 1072
	OpARM64FMOVDstore                Op = 1073
	OpARM64MOVBstoreidx              Op = 1074
	OpARM64MOVHstoreidx              Op = 1075
	OpARM64MOVWstoreidx              Op = 1076
	OpARM64MOVDstoreidx              Op = 1077
	OpARM64FMOVSstoreidx             Op = 1078
	OpARM64FMOVDstoreidx             Op = 1079
	OpARM64MOVHstoreidx2             Op = 1080
	OpARM64MOVWstoreidx4             Op = 1081
	OpARM64MOVDstoreidx8             Op = 1082
	OpARM64MOVBstorezero             Op = 1083
	OpARM64MOVHstorezero             Op = 1084
	OpARM64MOVWstorezero             Op = 1085
	OpARM64MOVDstorezero             Op = 1086
	OpARM64MOVQstorezero             Op = 1087
	OpARM64MOVBstorezeroidx          Op = 1088
	OpARM64MOVHstorezeroidx          Op = 1089
	OpARM64MOVWstorezeroidx          Op = 1090
	OpARM64MOVDstorezeroidx          Op = 1091
	OpARM64MOVHstorezeroidx2         Op = 1092
	OpARM64MOVWstorezeroidx4         Op = 1093
	OpARM64MOVDstorezeroidx8         Op = 1094
	OpARM64FMOVDgpfp                 Op = 1095
	OpARM64FMOVDfpgp                 Op = 1096
	OpARM64FMOVSgpfp                 Op = 1097
	OpARM64FMOVSfpgp                 Op = 1098
	OpARM64MOVBreg                   Op = 1099
	OpARM64MOVBUreg                  Op = 1100
	OpARM64MOVHreg                   Op = 1101
	OpARM64MOVHUreg                  Op = 1102
	OpARM64MOVWreg                   Op = 1103
	OpARM64MOVWUreg                  Op = 1104
	OpARM64MOVDreg                   Op = 1105
	OpARM64MOVDnop                   Op = 1106
	OpARM64SCVTFWS                   Op = 1107
	OpARM64SCVTFWD                   Op = 1108
	OpARM64UCVTFWS                   Op = 1109
	OpARM64UCVTFWD                   Op = 1110
	OpARM64SCVTFS                    Op = 1111
	OpARM64SCVTFD                    Op = 1112
	OpARM64UCVTFS                    Op = 1113
	OpARM64UCVTFD                    Op = 1114
	OpARM64FCVTZSSW                  Op = 1115
	OpARM64FCVTZSDW                  Op = 1116
	OpARM64FCVTZUSW                  Op = 1117
	OpARM64FCVTZUDW                  Op = 1118
	OpARM64FCVTZSS                   Op = 1119
	OpARM64FCVTZSD                   Op = 1120
	OpARM64FCVTZUS                   Op = 1121
	OpARM64FCVTZUD                   Op = 1122
	OpARM64FCVTSD                    Op = 1123
	OpARM64FCVTDS                    Op = 1124
	OpARM64FRINTAD                   Op = 1125
	OpARM64FRINTMD                   Op = 1126
	OpARM64FRINTND                   Op = 1127
	OpARM64FRINTPD                   Op = 1128
	OpARM64FRINTZD                   Op = 1129
	OpARM64CSEL                      Op = 1130
	OpARM64CSEL0                     Op = 1131
	OpARM64CALLstatic                Op = 1132
	OpARM64CALLclosure               Op = 1133
	OpARM64CALLinter                 Op = 1134
	OpARM64LoweredNilCheck           Op = 1135
	OpARM64Equal                     Op = 1136
	OpARM64NotEqual                  Op = 1137
	OpARM64LessThan                  Op = 1138
	OpARM64LessEqual                 Op = 1139
	OpARM64GreaterThan               Op = 1140
	OpARM64GreaterEqual              Op = 1141
	OpARM64LessThanU                 Op = 1142
	OpARM64LessEqualU                Op = 1143
	OpARM64GreaterThanU              Op = 1144
	OpARM64GreaterEqualU             Op = 1145
	OpARM64LessThanF                 Op = 1146
	OpARM64LessEqualF                Op = 1147
	OpARM64GreaterThanF              Op = 1148
	OpARM64GreaterEqualF             Op = 1149
	OpARM64DUFFZERO                  Op = 1150
	OpARM64LoweredZero               Op = 1151
	OpARM64DUFFCOPY                  Op = 1152
	OpARM64LoweredMove               Op = 1153
	OpARM64LoweredGetClosurePtr      Op = 1154
	OpARM64LoweredGetCallerSP        Op = 1155
	OpARM64LoweredGetCallerPC        Op = 1156
	OpARM64FlagEQ                    Op = 1157
	OpARM64FlagLT_ULT                Op = 1158
	OpARM64FlagLT_UGT                Op = 1159
	OpARM64FlagGT_UGT                Op = 1160
	OpARM64FlagGT_ULT                Op = 1161
	OpARM64InvertFlags               Op = 1162
	OpARM64LDAR                      Op = 1163
	OpARM64LDARB                     Op = 1164
	OpARM64LDARW                     Op = 1165
	OpARM64STLR                      Op = 1166
	OpARM64STLRW                     Op = 1167
	OpARM64LoweredAtomicExchange64   Op = 1168
	OpARM64LoweredAtomicExchange32   Op = 1169
	OpARM64LoweredAtomicAdd64        Op = 1170
	OpARM64LoweredAtomicAdd32        Op = 1171
	OpARM64LoweredAtomicAdd64Variant Op = 1172
	OpARM64LoweredAtomicAdd32Variant Op = 1173
	OpARM64LoweredAtomicCas64        Op = 1174
	OpARM64LoweredAtomicCas32        Op = 1175
	OpARM64LoweredAtomicAnd8         Op = 1176
	OpARM64LoweredAtomicOr8          Op = 1177
	OpARM64LoweredWB                 Op = 1178
	OpARM64LoweredPanicBoundsA       Op = 1179
	OpARM64LoweredPanicBoundsB       Op = 1180
	OpARM64LoweredPanicBoundsC       Op = 1181

	OpMIPSADD                    Op = 1182
	OpMIPSADDconst               Op = 1183
	OpMIPSSUB                    Op = 1184
	OpMIPSSUBconst               Op = 1185
	OpMIPSMUL                    Op = 1186
	OpMIPSMULT                   Op = 1187
	OpMIPSMULTU                  Op = 1188
	OpMIPSDIV                    Op = 1189
	OpMIPSDIVU                   Op = 1190
	OpMIPSADDF                   Op = 1191
	OpMIPSADDD                   Op = 1192
	OpMIPSSUBF                   Op = 1193
	OpMIPSSUBD                   Op = 1194
	OpMIPSMULF                   Op = 1195
	OpMIPSMULD                   Op = 1196
	OpMIPSDIVF                   Op = 1197
	OpMIPSDIVD                   Op = 1198
	OpMIPSAND                    Op = 1199
	OpMIPSANDconst               Op = 1200
	OpMIPSOR                     Op = 1201
	OpMIPSORconst                Op = 1202
	OpMIPSXOR                    Op = 1203
	OpMIPSXORconst               Op = 1204
	OpMIPSNOR                    Op = 1205
	OpMIPSNORconst               Op = 1206
	OpMIPSNEG                    Op = 1207
	OpMIPSNEGF                   Op = 1208
	OpMIPSNEGD                   Op = 1209
	OpMIPSSQRTD                  Op = 1210
	OpMIPSSLL                    Op = 1211
	OpMIPSSLLconst               Op = 1212
	OpMIPSSRL                    Op = 1213
	OpMIPSSRLconst               Op = 1214
	OpMIPSSRA                    Op = 1215
	OpMIPSSRAconst               Op = 1216
	OpMIPSCLZ                    Op = 1217
	OpMIPSSGT                    Op = 1218
	OpMIPSSGTconst               Op = 1219
	OpMIPSSGTzero                Op = 1220
	OpMIPSSGTU                   Op = 1221
	OpMIPSSGTUconst              Op = 1222
	OpMIPSSGTUzero               Op = 1223
	OpMIPSCMPEQF                 Op = 1224
	OpMIPSCMPEQD                 Op = 1225
	OpMIPSCMPGEF                 Op = 1226
	OpMIPSCMPGED                 Op = 1227
	OpMIPSCMPGTF                 Op = 1228
	OpMIPSCMPGTD                 Op = 1229
	OpMIPSMOVWconst              Op = 1230
	OpMIPSMOVFconst              Op = 1231
	OpMIPSMOVDconst              Op = 1232
	OpMIPSMOVWaddr               Op = 1233
	OpMIPSMOVBload               Op = 1234
	OpMIPSMOVBUload              Op = 1235
	OpMIPSMOVHload               Op = 1236
	OpMIPSMOVHUload              Op = 1237
	OpMIPSMOVWload               Op = 1238
	OpMIPSMOVFload               Op = 1239
	OpMIPSMOVDload               Op = 1240
	OpMIPSMOVBstore              Op = 1241
	OpMIPSMOVHstore              Op = 1242
	OpMIPSMOVWstore              Op = 1243
	OpMIPSMOVFstore              Op = 1244
	OpMIPSMOVDstore              Op = 1245
	OpMIPSMOVBstorezero          Op = 1246
	OpMIPSMOVHstorezero          Op = 1247
	OpMIPSMOVWstorezero          Op = 1248
	OpMIPSMOVBreg                Op = 1249
	OpMIPSMOVBUreg               Op = 1250
	OpMIPSMOVHreg                Op = 1251
	OpMIPSMOVHUreg               Op = 1252
	OpMIPSMOVWreg                Op = 1253
	OpMIPSMOVWnop                Op = 1254
	OpMIPSCMOVZ                  Op = 1255
	OpMIPSCMOVZzero              Op = 1256
	OpMIPSMOVWF                  Op = 1257
	OpMIPSMOVWD                  Op = 1258
	OpMIPSTRUNCFW                Op = 1259
	OpMIPSTRUNCDW                Op = 1260
	OpMIPSMOVFD                  Op = 1261
	OpMIPSMOVDF                  Op = 1262
	OpMIPSCALLstatic             Op = 1263
	OpMIPSCALLclosure            Op = 1264
	OpMIPSCALLinter              Op = 1265
	OpMIPSLoweredAtomicLoad      Op = 1266
	OpMIPSLoweredAtomicStore     Op = 1267
	OpMIPSLoweredAtomicStorezero Op = 1268
	OpMIPSLoweredAtomicExchange  Op = 1269
	OpMIPSLoweredAtomicAdd       Op = 1270
	OpMIPSLoweredAtomicAddconst  Op = 1271
	OpMIPSLoweredAtomicCas       Op = 1272
	OpMIPSLoweredAtomicAnd       Op = 1273
	OpMIPSLoweredAtomicOr        Op = 1274
	OpMIPSLoweredZero            Op = 1275
	OpMIPSLoweredMove            Op = 1276
	OpMIPSLoweredNilCheck        Op = 1277
	OpMIPSFPFlagTrue             Op = 1278
	OpMIPSFPFlagFalse            Op = 1279
	OpMIPSLoweredGetClosurePtr   Op = 1280
	OpMIPSLoweredGetCallerSP     Op = 1281
	OpMIPSLoweredGetCallerPC     Op = 1282
	OpMIPSLoweredWB              Op = 1283
	OpMIPSLoweredPanicBoundsA    Op = 1284
	OpMIPSLoweredPanicBoundsB    Op = 1285
	OpMIPSLoweredPanicBoundsC    Op = 1286
	OpMIPSLoweredPanicExtendA    Op = 1287
	OpMIPSLoweredPanicExtendB    Op = 1288
	OpMIPSLoweredPanicExtendC    Op = 1289

	OpMIPS64ADDV                     Op = 1290
	OpMIPS64ADDVconst                Op = 1291
	OpMIPS64SUBV                     Op = 1292
	OpMIPS64SUBVconst                Op = 1293
	OpMIPS64MULV                     Op = 1294
	OpMIPS64MULVU                    Op = 1295
	OpMIPS64DIVV                     Op = 1296
	OpMIPS64DIVVU                    Op = 1297
	OpMIPS64ADDF                     Op = 1298
	OpMIPS64ADDD                     Op = 1299
	OpMIPS64SUBF                     Op = 1300
	OpMIPS64SUBD                     Op = 1301
	OpMIPS64MULF                     Op = 1302
	OpMIPS64MULD                     Op = 1303
	OpMIPS64DIVF                     Op = 1304
	OpMIPS64DIVD                     Op = 1305
	OpMIPS64AND                      Op = 1306
	OpMIPS64ANDconst                 Op = 1307
	OpMIPS64OR                       Op = 1308
	OpMIPS64ORconst                  Op = 1309
	OpMIPS64XOR                      Op = 1310
	OpMIPS64XORconst                 Op = 1311
	OpMIPS64NOR                      Op = 1312
	OpMIPS64NORconst                 Op = 1313
	OpMIPS64NEGV                     Op = 1314
	OpMIPS64NEGF                     Op = 1315
	OpMIPS64NEGD                     Op = 1316
	OpMIPS64SQRTD                    Op = 1317
	OpMIPS64SLLV                     Op = 1318
	OpMIPS64SLLVconst                Op = 1319
	OpMIPS64SRLV                     Op = 1320
	OpMIPS64SRLVconst                Op = 1321
	OpMIPS64SRAV                     Op = 1322
	OpMIPS64SRAVconst                Op = 1323
	OpMIPS64SGT                      Op = 1324
	OpMIPS64SGTconst                 Op = 1325
	OpMIPS64SGTU                     Op = 1326
	OpMIPS64SGTUconst                Op = 1327
	OpMIPS64CMPEQF                   Op = 1328
	OpMIPS64CMPEQD                   Op = 1329
	OpMIPS64CMPGEF                   Op = 1330
	OpMIPS64CMPGED                   Op = 1331
	OpMIPS64CMPGTF                   Op = 1332
	OpMIPS64CMPGTD                   Op = 1333
	OpMIPS64MOVVconst                Op = 1334
	OpMIPS64MOVFconst                Op = 1335
	OpMIPS64MOVDconst                Op = 1336
	OpMIPS64MOVVaddr                 Op = 1337
	OpMIPS64MOVBload                 Op = 1338
	OpMIPS64MOVBUload                Op = 1339
	OpMIPS64MOVHload                 Op = 1340
	OpMIPS64MOVHUload                Op = 1341
	OpMIPS64MOVWload                 Op = 1342
	OpMIPS64MOVWUload                Op = 1343
	OpMIPS64MOVVload                 Op = 1344
	OpMIPS64MOVFload                 Op = 1345
	OpMIPS64MOVDload                 Op = 1346
	OpMIPS64MOVBstore                Op = 1347
	OpMIPS64MOVHstore                Op = 1348
	OpMIPS64MOVWstore                Op = 1349
	OpMIPS64MOVVstore                Op = 1350
	OpMIPS64MOVFstore                Op = 1351
	OpMIPS64MOVDstore                Op = 1352
	OpMIPS64MOVBstorezero            Op = 1353
	OpMIPS64MOVHstorezero            Op = 1354
	OpMIPS64MOVWstorezero            Op = 1355
	OpMIPS64MOVVstorezero            Op = 1356
	OpMIPS64MOVBreg                  Op = 1357
	OpMIPS64MOVBUreg                 Op = 1358
	OpMIPS64MOVHreg                  Op = 1359
	OpMIPS64MOVHUreg                 Op = 1360
	OpMIPS64MOVWreg                  Op = 1361
	OpMIPS64MOVWUreg                 Op = 1362
	OpMIPS64MOVVreg                  Op = 1363
	OpMIPS64MOVVnop                  Op = 1364
	OpMIPS64MOVWF                    Op = 1365
	OpMIPS64MOVWD                    Op = 1366
	OpMIPS64MOVVF                    Op = 1367
	OpMIPS64MOVVD                    Op = 1368
	OpMIPS64TRUNCFW                  Op = 1369
	OpMIPS64TRUNCDW                  Op = 1370
	OpMIPS64TRUNCFV                  Op = 1371
	OpMIPS64TRUNCDV                  Op = 1372
	OpMIPS64MOVFD                    Op = 1373
	OpMIPS64MOVDF                    Op = 1374
	OpMIPS64CALLstatic               Op = 1375
	OpMIPS64CALLclosure              Op = 1376
	OpMIPS64CALLinter                Op = 1377
	OpMIPS64DUFFZERO                 Op = 1378
	OpMIPS64LoweredZero              Op = 1379
	OpMIPS64LoweredMove              Op = 1380
	OpMIPS64LoweredAtomicLoad8       Op = 1381
	OpMIPS64LoweredAtomicLoad32      Op = 1382
	OpMIPS64LoweredAtomicLoad64      Op = 1383
	OpMIPS64LoweredAtomicStore32     Op = 1384
	OpMIPS64LoweredAtomicStore64     Op = 1385
	OpMIPS64LoweredAtomicStorezero32 Op = 1386
	OpMIPS64LoweredAtomicStorezero64 Op = 1387
	OpMIPS64LoweredAtomicExchange32  Op = 1388
	OpMIPS64LoweredAtomicExchange64  Op = 1389
	OpMIPS64LoweredAtomicAdd32       Op = 1390
	OpMIPS64LoweredAtomicAdd64       Op = 1391
	OpMIPS64LoweredAtomicAddconst32  Op = 1392
	OpMIPS64LoweredAtomicAddconst64  Op = 1393
	OpMIPS64LoweredAtomicCas32       Op = 1394
	OpMIPS64LoweredAtomicCas64       Op = 1395
	OpMIPS64LoweredNilCheck          Op = 1396
	OpMIPS64FPFlagTrue               Op = 1397
	OpMIPS64FPFlagFalse              Op = 1398
	OpMIPS64LoweredGetClosurePtr     Op = 1399
	OpMIPS64LoweredGetCallerSP       Op = 1400
	OpMIPS64LoweredGetCallerPC       Op = 1401
	OpMIPS64LoweredWB                Op = 1402
	OpMIPS64LoweredPanicBoundsA      Op = 1403
	OpMIPS64LoweredPanicBoundsB      Op = 1404
	OpMIPS64LoweredPanicBoundsC      Op = 1405

	OpPPC64ADD                     Op = 1406
	OpPPC64ADDconst                Op = 1407
	OpPPC64FADD                    Op = 1408
	OpPPC64FADDS                   Op = 1409
	OpPPC64SUB                     Op = 1410
	OpPPC64FSUB                    Op = 1411
	OpPPC64FSUBS                   Op = 1412
	OpPPC64MULLD                   Op = 1413
	OpPPC64MULLW                   Op = 1414
	OpPPC64MULHD                   Op = 1415
	OpPPC64MULHW                   Op = 1416
	OpPPC64MULHDU                  Op = 1417
	OpPPC64MULHWU                  Op = 1418
	OpPPC64LoweredMuluhilo         Op = 1419
	OpPPC64FMUL                    Op = 1420
	OpPPC64FMULS                   Op = 1421
	OpPPC64FMADD                   Op = 1422
	OpPPC64FMADDS                  Op = 1423
	OpPPC64FMSUB                   Op = 1424
	OpPPC64FMSUBS                  Op = 1425
	OpPPC64SRAD                    Op = 1426
	OpPPC64SRAW                    Op = 1427
	OpPPC64SRD                     Op = 1428
	OpPPC64SRW                     Op = 1429
	OpPPC64SLD                     Op = 1430
	OpPPC64SLW                     Op = 1431
	OpPPC64ROTL                    Op = 1432
	OpPPC64ROTLW                   Op = 1433
	OpPPC64LoweredAdd64Carry       Op = 1434
	OpPPC64ADDconstForCarry        Op = 1435
	OpPPC64MaskIfNotCarry          Op = 1436
	OpPPC64SRADconst               Op = 1437
	OpPPC64SRAWconst               Op = 1438
	OpPPC64SRDconst                Op = 1439
	OpPPC64SRWconst                Op = 1440
	OpPPC64SLDconst                Op = 1441
	OpPPC64SLWconst                Op = 1442
	OpPPC64ROTLconst               Op = 1443
	OpPPC64ROTLWconst              Op = 1444
	OpPPC64CNTLZD                  Op = 1445
	OpPPC64CNTLZW                  Op = 1446
	OpPPC64CNTTZD                  Op = 1447
	OpPPC64CNTTZW                  Op = 1448
	OpPPC64POPCNTD                 Op = 1449
	OpPPC64POPCNTW                 Op = 1450
	OpPPC64POPCNTB                 Op = 1451
	OpPPC64FDIV                    Op = 1452
	OpPPC64FDIVS                   Op = 1453
	OpPPC64DIVD                    Op = 1454
	OpPPC64DIVW                    Op = 1455
	OpPPC64DIVDU                   Op = 1456
	OpPPC64DIVWU                   Op = 1457
	OpPPC64FCTIDZ                  Op = 1458
	OpPPC64FCTIWZ                  Op = 1459
	OpPPC64FCFID                   Op = 1460
	OpPPC64FCFIDS                  Op = 1461
	OpPPC64FRSP                    Op = 1462
	OpPPC64MFVSRD                  Op = 1463
	OpPPC64MTVSRD                  Op = 1464
	OpPPC64AND                     Op = 1465
	OpPPC64ANDN                    Op = 1466
	OpPPC64ANDCC                   Op = 1467
	OpPPC64OR                      Op = 1468
	OpPPC64ORN                     Op = 1469
	OpPPC64ORCC                    Op = 1470
	OpPPC64NOR                     Op = 1471
	OpPPC64XOR                     Op = 1472
	OpPPC64XORCC                   Op = 1473
	OpPPC64EQV                     Op = 1474
	OpPPC64NEG                     Op = 1475
	OpPPC64FNEG                    Op = 1476
	OpPPC64FSQRT                   Op = 1477
	OpPPC64FSQRTS                  Op = 1478
	OpPPC64FFLOOR                  Op = 1479
	OpPPC64FCEIL                   Op = 1480
	OpPPC64FTRUNC                  Op = 1481
	OpPPC64FROUND                  Op = 1482
	OpPPC64FABS                    Op = 1483
	OpPPC64FNABS                   Op = 1484
	OpPPC64FCPSGN                  Op = 1485
	OpPPC64ORconst                 Op = 1486
	OpPPC64XORconst                Op = 1487
	OpPPC64ANDconst                Op = 1488
	OpPPC64ANDCCconst              Op = 1489
	OpPPC64MOVBreg                 Op = 1490
	OpPPC64MOVBZreg                Op = 1491
	OpPPC64MOVHreg                 Op = 1492
	OpPPC64MOVHZreg                Op = 1493
	OpPPC64MOVWreg                 Op = 1494
	OpPPC64MOVWZreg                Op = 1495
	OpPPC64MOVBZload               Op = 1496
	OpPPC64MOVHload                Op = 1497
	OpPPC64MOVHZload               Op = 1498
	OpPPC64MOVWload                Op = 1499
	OpPPC64MOVWZload               Op = 1500
	OpPPC64MOVDload                Op = 1501
	OpPPC64MOVDBRload              Op = 1502
	OpPPC64MOVWBRload              Op = 1503
	OpPPC64MOVHBRload              Op = 1504
	OpPPC64MOVBZloadidx            Op = 1505
	OpPPC64MOVHloadidx             Op = 1506
	OpPPC64MOVHZloadidx            Op = 1507
	OpPPC64MOVWloadidx             Op = 1508
	OpPPC64MOVWZloadidx            Op = 1509
	OpPPC64MOVDloadidx             Op = 1510
	OpPPC64MOVHBRloadidx           Op = 1511
	OpPPC64MOVWBRloadidx           Op = 1512
	OpPPC64MOVDBRloadidx           Op = 1513
	OpPPC64FMOVDloadidx            Op = 1514
	OpPPC64FMOVSloadidx            Op = 1515
	OpPPC64MOVDBRstore             Op = 1516
	OpPPC64MOVWBRstore             Op = 1517
	OpPPC64MOVHBRstore             Op = 1518
	OpPPC64FMOVDload               Op = 1519
	OpPPC64FMOVSload               Op = 1520
	OpPPC64MOVBstore               Op = 1521
	OpPPC64MOVHstore               Op = 1522
	OpPPC64MOVWstore               Op = 1523
	OpPPC64MOVDstore               Op = 1524
	OpPPC64FMOVDstore              Op = 1525
	OpPPC64FMOVSstore              Op = 1526
	OpPPC64MOVBstoreidx            Op = 1527
	OpPPC64MOVHstoreidx            Op = 1528
	OpPPC64MOVWstoreidx            Op = 1529
	OpPPC64MOVDstoreidx            Op = 1530
	OpPPC64FMOVDstoreidx           Op = 1531
	OpPPC64FMOVSstoreidx           Op = 1532
	OpPPC64MOVHBRstoreidx          Op = 1533
	OpPPC64MOVWBRstoreidx          Op = 1534
	OpPPC64MOVDBRstoreidx          Op = 1535
	OpPPC64MOVBstorezero           Op = 1536
	OpPPC64MOVHstorezero           Op = 1537
	OpPPC64MOVWstorezero           Op = 1538
	OpPPC64MOVDstorezero           Op = 1539
	OpPPC64MOVDaddr                Op = 1540
	OpPPC64MOVDconst               Op = 1541
	OpPPC64FMOVDconst              Op = 1542
	OpPPC64FMOVSconst              Op = 1543
	OpPPC64FCMPU                   Op = 1544
	OpPPC64CMP                     Op = 1545
	OpPPC64CMPU                    Op = 1546
	OpPPC64CMPW                    Op = 1547
	OpPPC64CMPWU                   Op = 1548
	OpPPC64CMPconst                Op = 1549
	OpPPC64CMPUconst               Op = 1550
	OpPPC64CMPWconst               Op = 1551
	OpPPC64CMPWUconst              Op = 1552
	OpPPC64Equal                   Op = 1553
	OpPPC64NotEqual                Op = 1554
	OpPPC64LessThan                Op = 1555
	OpPPC64FLessThan               Op = 1556
	OpPPC64LessEqual               Op = 1557
	OpPPC64FLessEqual              Op = 1558
	OpPPC64GreaterThan             Op = 1559
	OpPPC64FGreaterThan            Op = 1560
	OpPPC64GreaterEqual            Op = 1561
	OpPPC64FGreaterEqual           Op = 1562
	OpPPC64LoweredGetClosurePtr    Op = 1563
	OpPPC64LoweredGetCallerSP      Op = 1564
	OpPPC64LoweredGetCallerPC      Op = 1565
	OpPPC64LoweredNilCheck         Op = 1566
	OpPPC64LoweredRound32F         Op = 1567
	OpPPC64LoweredRound64F         Op = 1568
	OpPPC64CALLstatic              Op = 1569
	OpPPC64CALLclosure             Op = 1570
	OpPPC64CALLinter               Op = 1571
	OpPPC64LoweredZero             Op = 1572
	OpPPC64LoweredMove             Op = 1573
	OpPPC64LoweredAtomicStore32    Op = 1574
	OpPPC64LoweredAtomicStore64    Op = 1575
	OpPPC64LoweredAtomicLoad8      Op = 1576
	OpPPC64LoweredAtomicLoad32     Op = 1577
	OpPPC64LoweredAtomicLoad64     Op = 1578
	OpPPC64LoweredAtomicLoadPtr    Op = 1579
	OpPPC64LoweredAtomicAdd32      Op = 1580
	OpPPC64LoweredAtomicAdd64      Op = 1581
	OpPPC64LoweredAtomicExchange32 Op = 1582
	OpPPC64LoweredAtomicExchange64 Op = 1583
	OpPPC64LoweredAtomicCas64      Op = 1584
	OpPPC64LoweredAtomicCas32      Op = 1585
	OpPPC64LoweredAtomicAnd8       Op = 1586
	OpPPC64LoweredAtomicOr8        Op = 1587
	OpPPC64LoweredWB               Op = 1588
	OpPPC64LoweredPanicBoundsA     Op = 1589
	OpPPC64LoweredPanicBoundsB     Op = 1590
	OpPPC64LoweredPanicBoundsC     Op = 1591
	OpPPC64InvertFlags             Op = 1592
	OpPPC64FlagEQ                  Op = 1593
	OpPPC64FlagLT                  Op = 1594
	OpPPC64FlagGT                  Op = 1595

	OpS390XFADDS                   Op = 1596
	OpS390XFADD                    Op = 1597
	OpS390XFSUBS                   Op = 1598
	OpS390XFSUB                    Op = 1599
	OpS390XFMULS                   Op = 1600
	OpS390XFMUL                    Op = 1601
	OpS390XFDIVS                   Op = 1602
	OpS390XFDIV                    Op = 1603
	OpS390XFNEGS                   Op = 1604
	OpS390XFNEG                    Op = 1605
	OpS390XFMADDS                  Op = 1606
	OpS390XFMADD                   Op = 1607
	OpS390XFMSUBS                  Op = 1608
	OpS390XFMSUB                   Op = 1609
	OpS390XLPDFR                   Op = 1610
	OpS390XLNDFR                   Op = 1611
	OpS390XCPSDR                   Op = 1612
	OpS390XFIDBR                   Op = 1613
	OpS390XFMOVSload               Op = 1614
	OpS390XFMOVDload               Op = 1615
	OpS390XFMOVSconst              Op = 1616
	OpS390XFMOVDconst              Op = 1617
	OpS390XFMOVSloadidx            Op = 1618
	OpS390XFMOVDloadidx            Op = 1619
	OpS390XFMOVSstore              Op = 1620
	OpS390XFMOVDstore              Op = 1621
	OpS390XFMOVSstoreidx           Op = 1622
	OpS390XFMOVDstoreidx           Op = 1623
	OpS390XADD                     Op = 1624
	OpS390XADDW                    Op = 1625
	OpS390XADDconst                Op = 1626
	OpS390XADDWconst               Op = 1627
	OpS390XADDload                 Op = 1628
	OpS390XADDWload                Op = 1629
	OpS390XSUB                     Op = 1630
	OpS390XSUBW                    Op = 1631
	OpS390XSUBconst                Op = 1632
	OpS390XSUBWconst               Op = 1633
	OpS390XSUBload                 Op = 1634
	OpS390XSUBWload                Op = 1635
	OpS390XMULLD                   Op = 1636
	OpS390XMULLW                   Op = 1637
	OpS390XMULLDconst              Op = 1638
	OpS390XMULLWconst              Op = 1639
	OpS390XMULLDload               Op = 1640
	OpS390XMULLWload               Op = 1641
	OpS390XMULHD                   Op = 1642
	OpS390XMULHDU                  Op = 1643
	OpS390XDIVD                    Op = 1644
	OpS390XDIVW                    Op = 1645
	OpS390XDIVDU                   Op = 1646
	OpS390XDIVWU                   Op = 1647
	OpS390XMODD                    Op = 1648
	OpS390XMODW                    Op = 1649
	OpS390XMODDU                   Op = 1650
	OpS390XMODWU                   Op = 1651
	OpS390XAND                     Op = 1652
	OpS390XANDW                    Op = 1653
	OpS390XANDconst                Op = 1654
	OpS390XANDWconst               Op = 1655
	OpS390XANDload                 Op = 1656
	OpS390XANDWload                Op = 1657
	OpS390XOR                      Op = 1658
	OpS390XORW                     Op = 1659
	OpS390XORconst                 Op = 1660
	OpS390XORWconst                Op = 1661
	OpS390XORload                  Op = 1662
	OpS390XORWload                 Op = 1663
	OpS390XXOR                     Op = 1664
	OpS390XXORW                    Op = 1665
	OpS390XXORconst                Op = 1666
	OpS390XXORWconst               Op = 1667
	OpS390XXORload                 Op = 1668
	OpS390XXORWload                Op = 1669
	OpS390XADDC                    Op = 1670
	OpS390XADDCconst               Op = 1671
	OpS390XADDE                    Op = 1672
	OpS390XSUBC                    Op = 1673
	OpS390XSUBE                    Op = 1674
	OpS390XCMP                     Op = 1675
	OpS390XCMPW                    Op = 1676
	OpS390XCMPU                    Op = 1677
	OpS390XCMPWU                   Op = 1678
	OpS390XCMPconst                Op = 1679
	OpS390XCMPWconst               Op = 1680
	OpS390XCMPUconst               Op = 1681
	OpS390XCMPWUconst              Op = 1682
	OpS390XFCMPS                   Op = 1683
	OpS390XFCMP                    Op = 1684
	OpS390XSLD                     Op = 1685
	OpS390XSLW                     Op = 1686
	OpS390XSLDconst                Op = 1687
	OpS390XSLWconst                Op = 1688
	OpS390XSRD                     Op = 1689
	OpS390XSRW                     Op = 1690
	OpS390XSRDconst                Op = 1691
	OpS390XSRWconst                Op = 1692
	OpS390XSRAD                    Op = 1693
	OpS390XSRAW                    Op = 1694
	OpS390XSRADconst               Op = 1695
	OpS390XSRAWconst               Op = 1696
	OpS390XRLLG                    Op = 1697
	OpS390XRLL                     Op = 1698
	OpS390XRLLGconst               Op = 1699
	OpS390XRLLconst                Op = 1700
	OpS390XNEG                     Op = 1701
	OpS390XNEGW                    Op = 1702
	OpS390XNOT                     Op = 1703
	OpS390XNOTW                    Op = 1704
	OpS390XFSQRT                   Op = 1705
	OpS390XMOVDEQ                  Op = 1706
	OpS390XMOVDNE                  Op = 1707
	OpS390XMOVDLT                  Op = 1708
	OpS390XMOVDLE                  Op = 1709
	OpS390XMOVDGT                  Op = 1710
	OpS390XMOVDGE                  Op = 1711
	OpS390XMOVDGTnoinv             Op = 1712
	OpS390XMOVDGEnoinv             Op = 1713
	OpS390XMOVBreg                 Op = 1714
	OpS390XMOVBZreg                Op = 1715
	OpS390XMOVHreg                 Op = 1716
	OpS390XMOVHZreg                Op = 1717
	OpS390XMOVWreg                 Op = 1718
	OpS390XMOVWZreg                Op = 1719
	OpS390XMOVDreg                 Op = 1720
	OpS390XMOVDnop                 Op = 1721
	OpS390XMOVDconst               Op = 1722
	OpS390XLDGR                    Op = 1723
	OpS390XLGDR                    Op = 1724
	OpS390XCFDBRA                  Op = 1725
	OpS390XCGDBRA                  Op = 1726
	OpS390XCFEBRA                  Op = 1727
	OpS390XCGEBRA                  Op = 1728
	OpS390XCEFBRA                  Op = 1729
	OpS390XCDFBRA                  Op = 1730
	OpS390XCEGBRA                  Op = 1731
	OpS390XCDGBRA                  Op = 1732
	OpS390XLEDBR                   Op = 1733
	OpS390XLDEBR                   Op = 1734
	OpS390XMOVDaddr                Op = 1735
	OpS390XMOVDaddridx             Op = 1736
	OpS390XMOVBZload               Op = 1737
	OpS390XMOVBload                Op = 1738
	OpS390XMOVHZload               Op = 1739
	OpS390XMOVHload                Op = 1740
	OpS390XMOVWZload               Op = 1741
	OpS390XMOVWload                Op = 1742
	OpS390XMOVDload                Op = 1743
	OpS390XMOVWBR                  Op = 1744
	OpS390XMOVDBR                  Op = 1745
	OpS390XMOVHBRload              Op = 1746
	OpS390XMOVWBRload              Op = 1747
	OpS390XMOVDBRload              Op = 1748
	OpS390XMOVBstore               Op = 1749
	OpS390XMOVHstore               Op = 1750
	OpS390XMOVWstore               Op = 1751
	OpS390XMOVDstore               Op = 1752
	OpS390XMOVHBRstore             Op = 1753
	OpS390XMOVWBRstore             Op = 1754
	OpS390XMOVDBRstore             Op = 1755
	OpS390XMVC                     Op = 1756
	OpS390XMOVBZloadidx            Op = 1757
	OpS390XMOVBloadidx             Op = 1758
	OpS390XMOVHZloadidx            Op = 1759
	OpS390XMOVHloadidx             Op = 1760
	OpS390XMOVWZloadidx            Op = 1761
	OpS390XMOVWloadidx             Op = 1762
	OpS390XMOVDloadidx             Op = 1763
	OpS390XMOVHBRloadidx           Op = 1764
	OpS390XMOVWBRloadidx           Op = 1765
	OpS390XMOVDBRloadidx           Op = 1766
	OpS390XMOVBstoreidx            Op = 1767
	OpS390XMOVHstoreidx            Op = 1768
	OpS390XMOVWstoreidx            Op = 1769
	OpS390XMOVDstoreidx            Op = 1770
	OpS390XMOVHBRstoreidx          Op = 1771
	OpS390XMOVWBRstoreidx          Op = 1772
	OpS390XMOVDBRstoreidx          Op = 1773
	OpS390XMOVBstoreconst          Op = 1774
	OpS390XMOVHstoreconst          Op = 1775
	OpS390XMOVWstoreconst          Op = 1776
	OpS390XMOVDstoreconst          Op = 1777
	OpS390XCLEAR                   Op = 1778
	OpS390XCALLstatic              Op = 1779
	OpS390XCALLclosure             Op = 1780
	OpS390XCALLinter               Op = 1781
	OpS390XInvertFlags             Op = 1782
	OpS390XLoweredGetG             Op = 1783
	OpS390XLoweredGetClosurePtr    Op = 1784
	OpS390XLoweredGetCallerSP      Op = 1785
	OpS390XLoweredGetCallerPC      Op = 1786
	OpS390XLoweredNilCheck         Op = 1787
	OpS390XLoweredRound32F         Op = 1788
	OpS390XLoweredRound64F         Op = 1789
	OpS390XLoweredWB               Op = 1790
	OpS390XLoweredPanicBoundsA     Op = 1791
	OpS390XLoweredPanicBoundsB     Op = 1792
	OpS390XLoweredPanicBoundsC     Op = 1793
	OpS390XFlagEQ                  Op = 1794
	OpS390XFlagLT                  Op = 1795
	OpS390XFlagGT                  Op = 1796
	OpS390XFlagOV                  Op = 1797
	OpS390XMOVBZatomicload         Op = 1798
	OpS390XMOVWZatomicload         Op = 1799
	OpS390XMOVDatomicload          Op = 1800
	OpS390XMOVWatomicstore         Op = 1801
	OpS390XMOVDatomicstore         Op = 1802
	OpS390XLAA                     Op = 1803
	OpS390XLAAG                    Op = 1804
	OpS390XAddTupleFirst32         Op = 1805
	OpS390XAddTupleFirst64         Op = 1806
	OpS390XLoweredAtomicCas32      Op = 1807
	OpS390XLoweredAtomicCas64      Op = 1808
	OpS390XLoweredAtomicExchange32 Op = 1809
	OpS390XLoweredAtomicExchange64 Op = 1810
	OpS390XFLOGR                   Op = 1811
	OpS390XPOPCNT                  Op = 1812
	OpS390XSumBytes2               Op = 1813
	OpS390XSumBytes4               Op = 1814
	OpS390XSumBytes8               Op = 1815
	OpS390XSTMG2                   Op = 1816
	OpS390XSTMG3                   Op = 1817
	OpS390XSTMG4                   Op = 1818
	OpS390XSTM2                    Op = 1819
	OpS390XSTM3                    Op = 1820
	OpS390XSTM4                    Op = 1821
	OpS390XLoweredMove             Op = 1822
	OpS390XLoweredZero             Op = 1823

	OpWasmLoweredStaticCall    Op = 1824
	OpWasmLoweredClosureCall   Op = 1825
	OpWasmLoweredInterCall     Op = 1826
	OpWasmLoweredAddr          Op = 1827
	OpWasmLoweredMove          Op = 1828
	OpWasmLoweredZero          Op = 1829
	OpWasmLoweredGetClosurePtr Op = 1830
	OpWasmLoweredGetCallerPC   Op = 1831
	OpWasmLoweredGetCallerSP   Op = 1832
	OpWasmLoweredNilCheck      Op = 1833
	OpWasmLoweredWB            Op = 1834
	OpWasmLoweredRound32F      Op = 1835
	OpWasmLoweredConvert       Op = 1836
	OpWasmSelect               Op = 1837
	OpWasmI64Load8U            Op = 1838
	OpWasmI64Load8S            Op = 1839
	OpWasmI64Load16U           Op = 1840
	OpWasmI64Load16S           Op = 1841
	OpWasmI64Load32U           Op = 1842
	OpWasmI64Load32S           Op = 1843
	OpWasmI64Load              Op = 1844
	OpWasmI64Store8            Op = 1845
	OpWasmI64Store16           Op = 1846
	OpWasmI64Store32           Op = 1847
	OpWasmI64Store             Op = 1848
	OpWasmF32Load              Op = 1849
	OpWasmF64Load              Op = 1850
	OpWasmF32Store             Op = 1851
	OpWasmF64Store             Op = 1852
	OpWasmI64Const             Op = 1853
	OpWasmF64Const             Op = 1854
	OpWasmI64Eqz               Op = 1855
	OpWasmI64Eq                Op = 1856
	OpWasmI64Ne                Op = 1857
	OpWasmI64LtS               Op = 1858
	OpWasmI64LtU               Op = 1859
	OpWasmI64GtS               Op = 1860
	OpWasmI64GtU               Op = 1861
	OpWasmI64LeS               Op = 1862
	OpWasmI64LeU               Op = 1863
	OpWasmI64GeS               Op = 1864
	OpWasmI64GeU               Op = 1865
	OpWasmF64Eq                Op = 1866
	OpWasmF64Ne                Op = 1867
	OpWasmF64Lt                Op = 1868
	OpWasmF64Gt                Op = 1869
	OpWasmF64Le                Op = 1870
	OpWasmF64Ge                Op = 1871
	OpWasmI64Add               Op = 1872
	OpWasmI64AddConst          Op = 1873
	OpWasmI64Sub               Op = 1874
	OpWasmI64Mul               Op = 1875
	OpWasmI64DivS              Op = 1876
	OpWasmI64DivU              Op = 1877
	OpWasmI64RemS              Op = 1878
	OpWasmI64RemU              Op = 1879
	OpWasmI64And               Op = 1880
	OpWasmI64Or                Op = 1881
	OpWasmI64Xor               Op = 1882
	OpWasmI64Shl               Op = 1883
	OpWasmI64ShrS              Op = 1884
	OpWasmI64ShrU              Op = 1885
	OpWasmF64Neg               Op = 1886
	OpWasmF64Add               Op = 1887
	OpWasmF64Sub               Op = 1888
	OpWasmF64Mul               Op = 1889
	OpWasmF64Div               Op = 1890
	OpWasmI64TruncSatF64S      Op = 1891
	OpWasmI64TruncSatF64U      Op = 1892
	OpWasmF64ConvertI64S       Op = 1893
	OpWasmF64ConvertI64U       Op = 1894
	OpWasmI64Extend8S          Op = 1895
	OpWasmI64Extend16S         Op = 1896
	OpWasmI64Extend32S         Op = 1897
	OpWasmF64Sqrt              Op = 1898
	OpWasmF64Trunc             Op = 1899
	OpWasmF64Ceil              Op = 1900
	OpWasmF64Floor             Op = 1901
	OpWasmF64Nearest           Op = 1902
	OpWasmF64Abs               Op = 1903
	OpWasmF64Copysign          Op = 1904
	OpWasmI64Ctz               Op = 1905
	OpWasmI64Clz               Op = 1906
	OpWasmI64Rotl              Op = 1907
	OpWasmI64Popcnt            Op = 1908

	OpAdd8                      Op = 1909
	OpAdd16                     Op = 1910
	OpAdd32                     Op = 1911
	OpAdd64                     Op = 1912
	OpAddPtr                    Op = 1913
	OpAdd32F                    Op = 1914
	OpAdd64F                    Op = 1915
	OpSub8                      Op = 1916
	OpSub16                     Op = 1917
	OpSub32                     Op = 1918
	OpSub64                     Op = 1919
	OpSubPtr                    Op = 1920
	OpSub32F                    Op = 1921
	OpSub64F                    Op = 1922
	OpMul8                      Op = 1923
	OpMul16                     Op = 1924
	OpMul32                     Op = 1925
	OpMul64                     Op = 1926
	OpMul32F                    Op = 1927
	OpMul64F                    Op = 1928
	OpDiv32F                    Op = 1929
	OpDiv64F                    Op = 1930
	OpHmul32                    Op = 1931
	OpHmul32u                   Op = 1932
	OpHmul64                    Op = 1933
	OpHmul64u                   Op = 1934
	OpMul32uhilo                Op = 1935
	OpMul64uhilo                Op = 1936
	OpMul32uover                Op = 1937
	OpMul64uover                Op = 1938
	OpAvg32u                    Op = 1939
	OpAvg64u                    Op = 1940
	OpDiv8                      Op = 1941
	OpDiv8u                     Op = 1942
	OpDiv16                     Op = 1943
	OpDiv16u                    Op = 1944
	OpDiv32                     Op = 1945
	OpDiv32u                    Op = 1946
	OpDiv64                     Op = 1947
	OpDiv64u                    Op = 1948
	OpDiv128u                   Op = 1949
	OpMod8                      Op = 1950
	OpMod8u                     Op = 1951
	OpMod16                     Op = 1952
	OpMod16u                    Op = 1953
	OpMod32                     Op = 1954
	OpMod32u                    Op = 1955
	OpMod64                     Op = 1956
	OpMod64u                    Op = 1957
	OpAnd8                      Op = 1958
	OpAnd16                     Op = 1959
	OpAnd32                     Op = 1960
	OpAnd64                     Op = 1961
	OpOr8                       Op = 1962
	OpOr16                      Op = 1963
	OpOr32                      Op = 1964
	OpOr64                      Op = 1965
	OpXor8                      Op = 1966
	OpXor16                     Op = 1967
	OpXor32                     Op = 1968
	OpXor64                     Op = 1969
	OpLsh8x8                    Op = 1970
	OpLsh8x16                   Op = 1971
	OpLsh8x32                   Op = 1972
	OpLsh8x64                   Op = 1973
	OpLsh16x8                   Op = 1974
	OpLsh16x16                  Op = 1975
	OpLsh16x32                  Op = 1976
	OpLsh16x64                  Op = 1977
	OpLsh32x8                   Op = 1978
	OpLsh32x16                  Op = 1979
	OpLsh32x32                  Op = 1980
	OpLsh32x64                  Op = 1981
	OpLsh64x8                   Op = 1982
	OpLsh64x16                  Op = 1983
	OpLsh64x32                  Op = 1984
	OpLsh64x64                  Op = 1985
	OpRsh8x8                    Op = 1986
	OpRsh8x16                   Op = 1987
	OpRsh8x32                   Op = 1988
	OpRsh8x64                   Op = 1989
	OpRsh16x8                   Op = 1990
	OpRsh16x16                  Op = 1991
	OpRsh16x32                  Op = 1992
	OpRsh16x64                  Op = 1993
	OpRsh32x8                   Op = 1994
	OpRsh32x16                  Op = 1995
	OpRsh32x32                  Op = 1996
	OpRsh32x64                  Op = 1997
	OpRsh64x8                   Op = 1998
	OpRsh64x16                  Op = 1999
	OpRsh64x32                  Op = 2000
	OpRsh64x64                  Op = 2001
	OpRsh8Ux8                   Op = 2002
	OpRsh8Ux16                  Op = 2003
	OpRsh8Ux32                  Op = 2004
	OpRsh8Ux64                  Op = 2005
	OpRsh16Ux8                  Op = 2006
	OpRsh16Ux16                 Op = 2007
	OpRsh16Ux32                 Op = 2008
	OpRsh16Ux64                 Op = 2009
	OpRsh32Ux8                  Op = 2010
	OpRsh32Ux16                 Op = 2011
	OpRsh32Ux32                 Op = 2012
	OpRsh32Ux64                 Op = 2013
	OpRsh64Ux8                  Op = 2014
	OpRsh64Ux16                 Op = 2015
	OpRsh64Ux32                 Op = 2016
	OpRsh64Ux64                 Op = 2017
	OpEq8                       Op = 2018
	OpEq16                      Op = 2019
	OpEq32                      Op = 2020
	OpEq64                      Op = 2021
	OpEqPtr                     Op = 2022
	OpEqInter                   Op = 2023
	OpEqSlice                   Op = 2024
	OpEq32F                     Op = 2025
	OpEq64F                     Op = 2026
	OpNeq8                      Op = 2027
	OpNeq16                     Op = 2028
	OpNeq32                     Op = 2029
	OpNeq64                     Op = 2030
	OpNeqPtr                    Op = 2031
	OpNeqInter                  Op = 2032
	OpNeqSlice                  Op = 2033
	OpNeq32F                    Op = 2034
	OpNeq64F                    Op = 2035
	OpLess8                     Op = 2036
	OpLess8U                    Op = 2037
	OpLess16                    Op = 2038
	OpLess16U                   Op = 2039
	OpLess32                    Op = 2040
	OpLess32U                   Op = 2041
	OpLess64                    Op = 2042
	OpLess64U                   Op = 2043
	OpLess32F                   Op = 2044
	OpLess64F                   Op = 2045
	OpLeq8                      Op = 2046
	OpLeq8U                     Op = 2047
	OpLeq16                     Op = 2048
	OpLeq16U                    Op = 2049
	OpLeq32                     Op = 2050
	OpLeq32U                    Op = 2051
	OpLeq64                     Op = 2052
	OpLeq64U                    Op = 2053
	OpLeq32F                    Op = 2054
	OpLeq64F                    Op = 2055
	OpGreater8                  Op = 2056
	OpGreater8U                 Op = 2057
	OpGreater16                 Op = 2058
	OpGreater16U                Op = 2059
	OpGreater32                 Op = 2060
	OpGreater32U                Op = 2061
	OpGreater64                 Op = 2062
	OpGreater64U                Op = 2063
	OpGreater32F                Op = 2064
	OpGreater64F                Op = 2065
	OpGeq8                      Op = 2066
	OpGeq8U                     Op = 2067
	OpGeq16                     Op = 2068
	OpGeq16U                    Op = 2069
	OpGeq32                     Op = 2070
	OpGeq32U                    Op = 2071
	OpGeq64                     Op = 2072
	OpGeq64U                    Op = 2073
	OpGeq32F                    Op = 2074
	OpGeq64F                    Op = 2075
	OpCondSelect                Op = 2076
	OpAndB                      Op = 2077
	OpOrB                       Op = 2078
	OpEqB                       Op = 2079
	OpNeqB                      Op = 2080
	OpNot                       Op = 2081
	OpNeg8                      Op = 2082
	OpNeg16                     Op = 2083
	OpNeg32                     Op = 2084
	OpNeg64                     Op = 2085
	OpNeg32F                    Op = 2086
	OpNeg64F                    Op = 2087
	OpCom8                      Op = 2088
	OpCom16                     Op = 2089
	OpCom32                     Op = 2090
	OpCom64                     Op = 2091
	OpCtz8                      Op = 2092
	OpCtz16                     Op = 2093
	OpCtz32                     Op = 2094
	OpCtz64                     Op = 2095
	OpCtz8NonZero               Op = 2096
	OpCtz16NonZero              Op = 2097
	OpCtz32NonZero              Op = 2098
	OpCtz64NonZero              Op = 2099
	OpBitLen8                   Op = 2100
	OpBitLen16                  Op = 2101
	OpBitLen32                  Op = 2102
	OpBitLen64                  Op = 2103
	OpBswap32                   Op = 2104
	OpBswap64                   Op = 2105
	OpBitRev8                   Op = 2106
	OpBitRev16                  Op = 2107
	OpBitRev32                  Op = 2108
	OpBitRev64                  Op = 2109
	OpPopCount8                 Op = 2110
	OpPopCount16                Op = 2111
	OpPopCount32                Op = 2112
	OpPopCount64                Op = 2113
	OpRotateLeft8               Op = 2114
	OpRotateLeft16              Op = 2115
	OpRotateLeft32              Op = 2116
	OpRotateLeft64              Op = 2117
	OpSqrt                      Op = 2118
	OpFloor                     Op = 2119
	OpCeil                      Op = 2120
	OpTrunc                     Op = 2121
	OpRound                     Op = 2122
	OpRoundToEven               Op = 2123
	OpAbs                       Op = 2124
	OpCopysign                  Op = 2125
	OpPhi                       Op = 2126
	OpCopy                      Op = 2127
	OpConvert                   Op = 2128
	OpConstBool                 Op = 2129
	OpConstString               Op = 2130
	OpConstNil                  Op = 2131
	OpConst8                    Op = 2132
	OpConst16                   Op = 2133
	OpConst32                   Op = 2134
	OpConst64                   Op = 2135
	OpConst32F                  Op = 2136
	OpConst64F                  Op = 2137
	OpConstInterface            Op = 2138
	OpConstSlice                Op = 2139
	OpInitMem                   Op = 2140
	OpArg                       Op = 2141
	OpAddr                      Op = 2142
	OpLocalAddr                 Op = 2143
	OpSP                        Op = 2144
	OpSB                        Op = 2145
	OpLoad                      Op = 2147
	OpStore                     Op = 2148
	OpMove                      Op = 2149
	OpZero                      Op = 2150
	OpStoreWB                   Op = 2151
	OpMoveWB                    Op = 2152
	OpZeroWB                    Op = 2153
	OpWB                        Op = 2154
	OpPanicBounds               Op = 2155
	OpPanicExtend               Op = 2156
	OpClosureCall               Op = 2157
	OpStaticCall                Op = 2158
	OpInterCall                 Op = 2159
	OpSignExt8to16              Op = 2160
	OpSignExt8to32              Op = 2161
	OpSignExt8to64              Op = 2162
	OpSignExt16to32             Op = 2163
	OpSignExt16to64             Op = 2164
	OpSignExt32to64             Op = 2165
	OpZeroExt8to16              Op = 2166
	OpZeroExt8to32              Op = 2167
	OpZeroExt8to64              Op = 2168
	OpZeroExt16to32             Op = 2169
	OpZeroExt16to64             Op = 2170
	OpZeroExt32to64             Op = 2171
	OpTrunc16to8                Op = 2172
	OpTrunc32to8                Op = 2173
	OpTrunc32to16               Op = 2174
	OpTrunc64to8                Op = 2175
	OpTrunc64to16               Op = 2176
	OpTrunc64to32               Op = 2177
	OpCvt32to32F                Op = 2178
	OpCvt32to64F                Op = 2179
	OpCvt64to32F                Op = 2180
	OpCvt64to64F                Op = 2181
	OpCvt32Fto32                Op = 2182
	OpCvt32Fto64                Op = 2183
	OpCvt64Fto32                Op = 2184
	OpCvt64Fto64                Op = 2185
	OpCvt32Fto64F               Op = 2186
	OpCvt64Fto32F               Op = 2187
	OpRound32F                  Op = 2188
	OpRound64F                  Op = 2189
	OpIsNonNil                  Op = 2190
	OpIsInBounds                Op = 2191
	OpIsSliceInBounds           Op = 2192
	OpNilCheck                  Op = 2193
	OpGetG                      Op = 2194
	OpGetClosurePtr             Op = 2195
	OpGetCallerPC               Op = 2196
	OpGetCallerSP               Op = 2197
	OpPtrIndex                  Op = 2198
	OpOffPtr                    Op = 2199
	OpSliceMake                 Op = 2200
	OpSlicePtr                  Op = 2201
	OpSliceLen                  Op = 2202
	OpSliceCap                  Op = 2203
	OpComplexMake               Op = 2204
	OpComplexReal               Op = 2205
	OpComplexImag               Op = 2206
	OpStringMake                Op = 2207
	OpStringPtr                 Op = 2208
	OpStringLen                 Op = 2209
	OpIMake                     Op = 2210
	OpITab                      Op = 2211
	OpIData                     Op = 2212
	OpStructMake0               Op = 2213
	OpStructMake1               Op = 2214
	OpStructMake2               Op = 2215
	OpStructMake3               Op = 2216
	OpStructMake4               Op = 2217
	OpStructSelect              Op = 2218
	OpArrayMake0                Op = 2219
	OpArrayMake1                Op = 2220
	OpArraySelect               Op = 2221
	OpStoreReg                  Op = 2222
	OpLoadReg                   Op = 2223
	OpFwdRef                    Op = 2224
	OpUnknown                   Op = 2225
	OpVarDef                    Op = 2226
	OpVarKill                   Op = 2227
	OpVarLive                   Op = 2228
	OpKeepAlive                 Op = 2229
	OpInlMark                   Op = 2230
	OpInt64Make                 Op = 2231
	OpInt64Hi                   Op = 2232
	OpInt64Lo                   Op = 2233
	OpAdd32carry                Op = 2234
	OpAdd32withcarry            Op = 2235
	OpSub32carry                Op = 2236
	OpSub32withcarry            Op = 2237
	OpAdd64carry                Op = 2238
	OpSub64borrow               Op = 2239
	OpSignmask                  Op = 2240
	OpZeromask                  Op = 2241
	OpSlicemask                 Op = 2242
	OpCvt32Uto32F               Op = 2243
	OpCvt32Uto64F               Op = 2244
	OpCvt32Fto32U               Op = 2245
	OpCvt64Fto32U               Op = 2246
	OpCvt64Uto32F               Op = 2247
	OpCvt64Uto64F               Op = 2248
	OpCvt32Fto64U               Op = 2249
	OpCvt64Fto64U               Op = 2250
	OpSelect0                   Op = 2251
	OpSelect1                   Op = 2252
	OpAtomicLoad8               Op = 2253
	OpAtomicLoad32              Op = 2254
	OpAtomicLoad64              Op = 2255
	OpAtomicLoadPtr             Op = 2256
	OpAtomicLoadAcq32           Op = 2257
	OpAtomicStore32             Op = 2258
	OpAtomicStore64             Op = 2259
	OpAtomicStorePtrNoWB        Op = 2260
	OpAtomicStoreRel32          Op = 2261
	OpAtomicExchange32          Op = 2262
	OpAtomicExchange64          Op = 2263
	OpAtomicAdd32               Op = 2264
	OpAtomicAdd64               Op = 2265
	OpAtomicCompareAndSwap32    Op = 2266
	OpAtomicCompareAndSwap64    Op = 2267
	OpAtomicCompareAndSwapRel32 Op = 2268
	OpAtomicAnd8                Op = 2269
	OpAtomicOr8                 Op = 2270
	OpAtomicAdd32Variant        Op = 2271
	OpAtomicAdd64Variant        Op = 2272
	OpClobber                   Op = 2273
)

var opcodeTable = [...]opInfo{
	{name: "OpInvalid"},

	{
		name:         "ADDSS",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		usesScratch:  true,
		asm:          x86.AADDSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:         "ADDSD",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.AADDSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:         "SUBSS",
		argLen:       2,
		resultInArg0: true,
		usesScratch:  true,
		asm:          x86.ASUBSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:         "SUBSD",
		argLen:       2,
		resultInArg0: true,
		asm:          x86.ASUBSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:         "MULSS",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		usesScratch:  true,
		asm:          x86.AMULSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:         "MULSD",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.AMULSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:         "DIVSS",
		argLen:       2,
		resultInArg0: true,
		usesScratch:  true,
		asm:          x86.ADIVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:         "DIVSD",
		argLen:       2,
		resultInArg0: true,
		asm:          x86.ADIVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:           "MOVSSload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:           "MOVSDload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:              "MOVSSconst",
		auxType:           auxFloat32,
		argLen:            0,
		rematerializeable: true,
		asm:               x86.AMOVSS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:              "MOVSDconst",
		auxType:           auxFloat64,
		argLen:            0,
		rematerializeable: true,
		asm:               x86.AMOVSD,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:      "MOVSSloadidx1",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:      "MOVSSloadidx4",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:      "MOVSDloadidx1",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:      "MOVSDloadidx8",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:           "MOVSSstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "MOVSDstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:      "MOVSSstoreidx1",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{2, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:      "MOVSSstoreidx4",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{2, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:      "MOVSDstoreidx1",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{2, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:      "MOVSDstoreidx8",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{2, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "ADDSSload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AADDSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:           "ADDSDload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AADDSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:           "SUBSSload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ASUBSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:           "SUBSDload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ASUBSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:           "MULSSload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AMULSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:           "MULSDload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AMULSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:           "DIVSSload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ADIVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:           "DIVSDload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ADIVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:         "ADDL",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 239}, // AX CX DX BX BP SI DI
				{0, 255}, // AX CX DX BX SP BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "ADDLconst",
		auxType:      auxInt32,
		argLen:       1,
		clobberFlags: true,
		asm:          x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 255}, // AX CX DX BX SP BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "ADDLcarry",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
				{1, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "ADDLconstcarry",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		asm:          x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "ADCL",
		argLen:       3,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AADCL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
				{1, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "ADCLconst",
		auxType:      auxInt32,
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AADCL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SUBL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
				{1, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SUBLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SUBLcarry",
		argLen:       2,
		resultInArg0: true,
		asm:          x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
				{1, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SUBLconstcarry",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		asm:          x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SBBL",
		argLen:       3,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASBBL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
				{1, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SBBLconst",
		auxType:      auxInt32,
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASBBL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "MULL",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AIMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
				{1, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "MULLconst",
		auxType:      auxInt32,
		argLen:       1,
		clobberFlags: true,
		asm:          x86.AIMUL3L,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "MULLU",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},   // AX
				{1, 255}, // AX CX DX BX SP BP SI DI
			},
			clobbers: 4, // DX
			outputs: []outputInfo{
				{1, 0},
				{0, 1}, // AX
			},
		},
	},
	{
		name:         "HMULL",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AIMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},   // AX
				{1, 255}, // AX CX DX BX SP BP SI DI
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:         "HMULLU",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},   // AX
				{1, 255}, // AX CX DX BX SP BP SI DI
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:         "MULLQU",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},   // AX
				{1, 255}, // AX CX DX BX SP BP SI DI
			},
			outputs: []outputInfo{
				{0, 4}, // DX
				{1, 1}, // AX
			},
		},
	},
	{
		name:         "AVGLU",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
				{1, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "DIVL",
		auxType:      auxBool,
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AIDIVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},   // AX
				{1, 251}, // AX CX BX SP BP SI DI
			},
			clobbers: 4, // DX
			outputs: []outputInfo{
				{0, 1}, // AX
			},
		},
	},
	{
		name:         "DIVW",
		auxType:      auxBool,
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AIDIVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},   // AX
				{1, 251}, // AX CX BX SP BP SI DI
			},
			clobbers: 4, // DX
			outputs: []outputInfo{
				{0, 1}, // AX
			},
		},
	},
	{
		name:         "DIVLU",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.ADIVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},   // AX
				{1, 251}, // AX CX BX SP BP SI DI
			},
			clobbers: 4, // DX
			outputs: []outputInfo{
				{0, 1}, // AX
			},
		},
	},
	{
		name:         "DIVWU",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.ADIVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},   // AX
				{1, 251}, // AX CX BX SP BP SI DI
			},
			clobbers: 4, // DX
			outputs: []outputInfo{
				{0, 1}, // AX
			},
		},
	},
	{
		name:         "MODL",
		auxType:      auxBool,
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AIDIVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},   // AX
				{1, 251}, // AX CX BX SP BP SI DI
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:         "MODW",
		auxType:      auxBool,
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AIDIVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},   // AX
				{1, 251}, // AX CX BX SP BP SI DI
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:         "MODLU",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.ADIVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},   // AX
				{1, 251}, // AX CX BX SP BP SI DI
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:         "MODWU",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.ADIVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},   // AX
				{1, 251}, // AX CX BX SP BP SI DI
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:         "ANDL",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
				{1, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "ANDLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "ORL",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
				{1, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "ORLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "XORL",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
				{1, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "XORLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "CMPL",
		argLen: 2,
		asm:    x86.ACMPL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 255}, // AX CX DX BX SP BP SI DI
				{1, 255}, // AX CX DX BX SP BP SI DI
			},
		},
	},
	{
		name:   "CMPW",
		argLen: 2,
		asm:    x86.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 255}, // AX CX DX BX SP BP SI DI
				{1, 255}, // AX CX DX BX SP BP SI DI
			},
		},
	},
	{
		name:   "CMPB",
		argLen: 2,
		asm:    x86.ACMPB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 255}, // AX CX DX BX SP BP SI DI
				{1, 255}, // AX CX DX BX SP BP SI DI
			},
		},
	},
	{
		name:    "CMPLconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     x86.ACMPL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 255}, // AX CX DX BX SP BP SI DI
			},
		},
	},
	{
		name:    "CMPWconst",
		auxType: auxInt16,
		argLen:  1,
		asm:     x86.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 255}, // AX CX DX BX SP BP SI DI
			},
		},
	},
	{
		name:    "CMPBconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     x86.ACMPB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 255}, // AX CX DX BX SP BP SI DI
			},
		},
	},
	{
		name:           "CMPLload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "CMPWload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "CMPBload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "CMPLconstload",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "CMPWconstload",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "CMPBconstload",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:        "UCOMISS",
		argLen:      2,
		usesScratch: true,
		asm:         x86.AUCOMISS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:        "UCOMISD",
		argLen:      2,
		usesScratch: true,
		asm:         x86.AUCOMISD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:        "TESTL",
		argLen:      2,
		commutative: true,
		asm:         x86.ATESTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 255}, // AX CX DX BX SP BP SI DI
				{1, 255}, // AX CX DX BX SP BP SI DI
			},
		},
	},
	{
		name:        "TESTW",
		argLen:      2,
		commutative: true,
		asm:         x86.ATESTW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 255}, // AX CX DX BX SP BP SI DI
				{1, 255}, // AX CX DX BX SP BP SI DI
			},
		},
	},
	{
		name:        "TESTB",
		argLen:      2,
		commutative: true,
		asm:         x86.ATESTB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 255}, // AX CX DX BX SP BP SI DI
				{1, 255}, // AX CX DX BX SP BP SI DI
			},
		},
	},
	{
		name:    "TESTLconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     x86.ATESTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 255}, // AX CX DX BX SP BP SI DI
			},
		},
	},
	{
		name:    "TESTWconst",
		auxType: auxInt16,
		argLen:  1,
		asm:     x86.ATESTW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 255}, // AX CX DX BX SP BP SI DI
			},
		},
	},
	{
		name:    "TESTBconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     x86.ATESTB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 255}, // AX CX DX BX SP BP SI DI
			},
		},
	},
	{
		name:         "SHLL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHLL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},   // CX
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SHLLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHLL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SHRL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},   // CX
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SHRW",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},   // CX
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SHRB",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},   // CX
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SHRLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SHRWconst",
		auxType:      auxInt16,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SHRBconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SARL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},   // CX
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SARW",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},   // CX
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SARB",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},   // CX
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SARLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SARWconst",
		auxType:      auxInt16,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SARBconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "ROLLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "ROLWconst",
		auxType:      auxInt16,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "ROLBconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "ADDLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239},   // AX CX DX BX BP SI DI
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "SUBLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239},   // AX CX DX BX BP SI DI
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "MULLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AIMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239},   // AX CX DX BX BP SI DI
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "ANDLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239},   // AX CX DX BX BP SI DI
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "ORLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239},   // AX CX DX BX BP SI DI
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "XORLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239},   // AX CX DX BX BP SI DI
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "ADDLloadidx4",
		auxType:        auxSymOff,
		argLen:         4,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239},   // AX CX DX BX BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "SUBLloadidx4",
		auxType:        auxSymOff,
		argLen:         4,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239},   // AX CX DX BX BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "MULLloadidx4",
		auxType:        auxSymOff,
		argLen:         4,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AIMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239},   // AX CX DX BX BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "ANDLloadidx4",
		auxType:        auxSymOff,
		argLen:         4,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239},   // AX CX DX BX BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "ORLloadidx4",
		auxType:        auxSymOff,
		argLen:         4,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239},   // AX CX DX BX BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "XORLloadidx4",
		auxType:        auxSymOff,
		argLen:         4,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239},   // AX CX DX BX BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{1, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "NEGL",
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ANEGL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "NOTL",
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ANOTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "BSFL",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ABSFL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "BSFW",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ABSFW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "BSRL",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ABSRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "BSRW",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ABSRW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "BSWAPL",
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABSWAPL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SQRTSD",
		argLen: 1,
		asm:    x86.ASQRTSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:   "SBBLcarrymask",
		argLen: 1,
		asm:    x86.ASBBL,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SETEQ",
		argLen: 1,
		asm:    x86.ASETEQ,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SETNE",
		argLen: 1,
		asm:    x86.ASETNE,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SETL",
		argLen: 1,
		asm:    x86.ASETLT,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SETLE",
		argLen: 1,
		asm:    x86.ASETLE,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SETG",
		argLen: 1,
		asm:    x86.ASETGT,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SETGE",
		argLen: 1,
		asm:    x86.ASETGE,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SETB",
		argLen: 1,
		asm:    x86.ASETCS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SETBE",
		argLen: 1,
		asm:    x86.ASETLS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SETA",
		argLen: 1,
		asm:    x86.ASETHI,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SETAE",
		argLen: 1,
		asm:    x86.ASETCC,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SETO",
		argLen: 1,
		asm:    x86.ASETOS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SETEQF",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ASETEQ,
		reg: regInfo{
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 238}, // CX DX BX BP SI DI
			},
		},
	},
	{
		name:         "SETNEF",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ASETNE,
		reg: regInfo{
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 238}, // CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SETORD",
		argLen: 1,
		asm:    x86.ASETPC,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SETNAN",
		argLen: 1,
		asm:    x86.ASETPS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SETGF",
		argLen: 1,
		asm:    x86.ASETHI,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "SETGEF",
		argLen: 1,
		asm:    x86.ASETCC,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "MOVBLSX",
		argLen: 1,
		asm:    x86.AMOVBLSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "MOVBLZX",
		argLen: 1,
		asm:    x86.AMOVBLZX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "MOVWLSX",
		argLen: 1,
		asm:    x86.AMOVWLSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "MOVWLZX",
		argLen: 1,
		asm:    x86.AMOVWLZX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:              "MOVLconst",
		auxType:           auxInt32,
		argLen:            0,
		rematerializeable: true,
		asm:               x86.AMOVL,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:        "CVTTSD2SL",
		argLen:      1,
		usesScratch: true,
		asm:         x86.ACVTTSD2SL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:        "CVTTSS2SL",
		argLen:      1,
		usesScratch: true,
		asm:         x86.ACVTTSS2SL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:        "CVTSL2SS",
		argLen:      1,
		usesScratch: true,
		asm:         x86.ACVTSL2SS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:        "CVTSL2SD",
		argLen:      1,
		usesScratch: true,
		asm:         x86.ACVTSL2SD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:        "CVTSD2SS",
		argLen:      1,
		usesScratch: true,
		asm:         x86.ACVTSD2SS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:   "CVTSS2SD",
		argLen: 1,
		asm:    x86.ACVTSS2SD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:         "PXOR",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.APXOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:              "LEAL",
		auxType:           auxSymOff,
		argLen:            1,
		rematerializeable: true,
		symEffect:         SymAddr,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:        "LEAL1",
		auxType:     auxSymOff,
		argLen:      2,
		commutative: true,
		symEffect:   SymAddr,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:      "LEAL2",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:      "LEAL4",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:      "LEAL8",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "MOVBload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVBLZX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "MOVBLSXload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVBLSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "MOVWload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVWLZX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "MOVWLSXload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVWLSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "MOVLload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "MOVBstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "MOVWstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "MOVLstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "ADDLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "SUBLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "ANDLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "ORLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "XORLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "ADDLmodifyidx4",
		auxType:        auxSymOff,
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "SUBLmodifyidx4",
		auxType:        auxSymOff,
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "ANDLmodifyidx4",
		auxType:        auxSymOff,
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "ORLmodifyidx4",
		auxType:        auxSymOff,
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "XORLmodifyidx4",
		auxType:        auxSymOff,
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "ADDLconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "ANDLconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "ORLconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "XORLconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "ADDLconstmodifyidx4",
		auxType:        auxSymValAndOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "ANDLconstmodifyidx4",
		auxType:        auxSymValAndOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "ORLconstmodifyidx4",
		auxType:        auxSymValAndOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "XORLconstmodifyidx4",
		auxType:        auxSymValAndOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:        "MOVBloadidx1",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.AMOVBLZX,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:        "MOVWloadidx1",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.AMOVWLZX,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:      "MOVWloadidx2",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVWLZX,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:        "MOVLloadidx1",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:      "MOVLloadidx4",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:        "MOVBstoreidx1",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         x86.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:        "MOVWstoreidx1",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         x86.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:      "MOVWstoreidx2",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:        "MOVLstoreidx1",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:      "MOVLstoreidx4",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{2, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "MOVBstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "MOVWstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "MOVLstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:      "MOVBstoreconstidx1",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymWrite,
		asm:       x86.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:      "MOVWstoreconstidx1",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymWrite,
		asm:       x86.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:      "MOVWstoreconstidx2",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymWrite,
		asm:       x86.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:      "MOVLstoreconstidx1",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymWrite,
		asm:       x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:      "MOVLstoreconstidx4",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymWrite,
		asm:       x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 255},   // AX CX DX BX SP BP SI DI
				{0, 65791}, // AX CX DX BX SP BP SI DI SB
			},
		},
	},
	{
		name:           "DUFFZERO",
		auxType:        auxInt64,
		argLen:         3,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 128}, // DI
				{1, 1},   // AX
			},
			clobbers: 130, // CX DI
		},
	},
	{
		name:           "REPSTOSL",
		argLen:         4,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 128}, // DI
				{1, 2},   // CX
				{2, 1},   // AX
			},
			clobbers: 130, // CX DI
		},
	},
	{
		name:         "CALLstatic",
		auxType:      auxSymOff,
		argLen:       1,
		clobberFlags: true,
		call:         true,
		symEffect:    SymNone,
		reg: regInfo{
			clobbers: 65519, // AX CX DX BX BP SI DI X0 X1 X2 X3 X4 X5 X6 X7
		},
	},
	{
		name:         "CALLclosure",
		auxType:      auxInt64,
		argLen:       3,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 4},   // DX
				{0, 255}, // AX CX DX BX SP BP SI DI
			},
			clobbers: 65519, // AX CX DX BX BP SI DI X0 X1 X2 X3 X4 X5 X6 X7
		},
	},
	{
		name:         "CALLinter",
		auxType:      auxInt64,
		argLen:       2,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			clobbers: 65519, // AX CX DX BX BP SI DI X0 X1 X2 X3 X4 X5 X6 X7
		},
	},
	{
		name:           "DUFFCOPY",
		auxType:        auxInt64,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		faultOnNilArg1: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 128}, // DI
				{1, 64},  // SI
			},
			clobbers: 194, // CX SI DI
		},
	},
	{
		name:           "REPMOVSL",
		argLen:         4,
		faultOnNilArg0: true,
		faultOnNilArg1: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 128}, // DI
				{1, 64},  // SI
				{2, 2},   // CX
			},
			clobbers: 194, // CX SI DI
		},
	},
	{
		name:   "InvertFlags",
		argLen: 1,
		reg:    regInfo{},
	},
	{
		name:   "LoweredGetG",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:      "LoweredGetClosurePtr",
		argLen:    0,
		zeroWidth: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:              "LoweredGetCallerPC",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:              "LoweredGetCallerSP",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:           "LoweredNilCheck",
		argLen:         2,
		clobberFlags:   true,
		nilCheck:       true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 255}, // AX CX DX BX SP BP SI DI
			},
		},
	},
	{
		name:         "LoweredWB",
		auxType:      auxSym,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymNone,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 128}, // DI
				{1, 1},   // AX
			},
			clobbers: 65280, // X0 X1 X2 X3 X4 X5 X6 X7
		},
	},
	{
		name:    "LoweredPanicBoundsA",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4}, // DX
				{1, 8}, // BX
			},
		},
	},
	{
		name:    "LoweredPanicBoundsB",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2}, // CX
				{1, 4}, // DX
			},
		},
	},
	{
		name:    "LoweredPanicBoundsC",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1}, // AX
				{1, 2}, // CX
			},
		},
	},
	{
		name:    "LoweredPanicExtendA",
		auxType: auxInt64,
		argLen:  4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 64}, // SI
				{1, 4},  // DX
				{2, 8},  // BX
			},
		},
	},
	{
		name:    "LoweredPanicExtendB",
		auxType: auxInt64,
		argLen:  4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 64}, // SI
				{1, 2},  // CX
				{2, 4},  // DX
			},
		},
	},
	{
		name:    "LoweredPanicExtendC",
		auxType: auxInt64,
		argLen:  4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 64}, // SI
				{1, 1},  // AX
				{2, 2},  // CX
			},
		},
	},
	{
		name:   "FlagEQ",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagLT_ULT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagLT_UGT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagGT_UGT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagGT_ULT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FCHS",
		argLen: 1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:    "MOVSSconst1",
		auxType: auxFloat32,
		argLen:  0,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:    "MOVSDconst1",
		auxType: auxFloat64,
		argLen:  0,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
		},
	},
	{
		name:   "MOVSSconst2",
		argLen: 1,
		asm:    x86.AMOVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},
	{
		name:   "MOVSDconst2",
		argLen: 1,
		asm:    x86.AMOVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 239}, // AX CX DX BX BP SI DI
			},
			outputs: []outputInfo{
				{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
	},

	{
		name:         "ADDSS",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.AADDSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:         "ADDSD",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.AADDSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:         "SUBSS",
		argLen:       2,
		resultInArg0: true,
		asm:          x86.ASUBSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:         "SUBSD",
		argLen:       2,
		resultInArg0: true,
		asm:          x86.ASUBSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:         "MULSS",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.AMULSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:         "MULSD",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.AMULSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:         "DIVSS",
		argLen:       2,
		resultInArg0: true,
		asm:          x86.ADIVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:         "DIVSD",
		argLen:       2,
		resultInArg0: true,
		asm:          x86.ADIVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:           "MOVSSload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:           "MOVSDload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:              "MOVSSconst",
		auxType:           auxFloat32,
		argLen:            0,
		rematerializeable: true,
		asm:               x86.AMOVSS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:              "MOVSDconst",
		auxType:           auxFloat64,
		argLen:            0,
		rematerializeable: true,
		asm:               x86.AMOVSD,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:      "MOVSSloadidx1",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVSS,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:      "MOVSSloadidx4",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVSS,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:      "MOVSDloadidx1",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVSD,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:      "MOVSDloadidx8",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVSD,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:           "MOVSSstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "MOVSDstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVSSstoreidx1",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVSS,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{2, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVSSstoreidx4",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVSS,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{2, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVSDstoreidx1",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVSD,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{2, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVSDstoreidx8",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVSD,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{2, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "ADDSSload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AADDSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:           "ADDSDload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AADDSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:           "SUBSSload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ASUBSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:           "SUBSDload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ASUBSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:           "MULSSload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AMULSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:           "MULSDload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AMULSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:           "DIVSSload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ADIVSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:           "DIVSDload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ADIVSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:         "ADDQ",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ADDL",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ADDQconst",
		auxType:      auxInt32,
		argLen:       1,
		clobberFlags: true,
		asm:          x86.AADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ADDLconst",
		auxType:      auxInt32,
		argLen:       1,
		clobberFlags: true,
		asm:          x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "ADDQconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "ADDLconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:         "SUBQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASUBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SUBL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SUBQconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASUBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SUBLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "MULQ",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AIMULQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "MULL",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AIMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "MULQconst",
		auxType:      auxInt32,
		argLen:       1,
		clobberFlags: true,
		asm:          x86.AIMUL3Q,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "MULLconst",
		auxType:      auxInt32,
		argLen:       1,
		clobberFlags: true,
		asm:          x86.AIMUL3L,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "MULLU",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			clobbers: 4, // DX
			outputs: []outputInfo{
				{1, 0},
				{0, 1}, // AX
			},
		},
	},
	{
		name:         "MULQU",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AMULQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			clobbers: 4, // DX
			outputs: []outputInfo{
				{1, 0},
				{0, 1}, // AX
			},
		},
	},
	{
		name:         "HMULQ",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AIMULQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:         "HMULL",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AIMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:         "HMULQU",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AMULQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:         "HMULLU",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:         "AVGQU",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "DIVQ",
		auxType:      auxBool,
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AIDIVQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 65531}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 1}, // AX
				{1, 4}, // DX
			},
		},
	},
	{
		name:         "DIVL",
		auxType:      auxBool,
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AIDIVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 65531}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 1}, // AX
				{1, 4}, // DX
			},
		},
	},
	{
		name:         "DIVW",
		auxType:      auxBool,
		argLen:       2,
		clobberFlags: true,
		asm:          x86.AIDIVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 65531}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 1}, // AX
				{1, 4}, // DX
			},
		},
	},
	{
		name:         "DIVQU",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.ADIVQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 65531}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 1}, // AX
				{1, 4}, // DX
			},
		},
	},
	{
		name:         "DIVLU",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.ADIVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 65531}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 1}, // AX
				{1, 4}, // DX
			},
		},
	},
	{
		name:         "DIVWU",
		argLen:       2,
		clobberFlags: true,
		asm:          x86.ADIVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 65531}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 1}, // AX
				{1, 4}, // DX
			},
		},
	},
	{
		name:         "NEGLflags",
		argLen:       1,
		resultInArg0: true,
		asm:          x86.ANEGL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ADDQcarry",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.AADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ADCQ",
		argLen:       3,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.AADCQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ADDQconstcarry",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		asm:          x86.AADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ADCQconst",
		auxType:      auxInt32,
		argLen:       2,
		resultInArg0: true,
		asm:          x86.AADCQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SUBQborrow",
		argLen:       2,
		resultInArg0: true,
		asm:          x86.ASUBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SBBQ",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ASBBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SUBQconstborrow",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		asm:          x86.ASUBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SBBQconst",
		auxType:      auxInt32,
		argLen:       2,
		resultInArg0: true,
		asm:          x86.ASBBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "MULQU2",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          x86.AMULQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1},     // AX
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 4}, // DX
				{1, 1}, // AX
			},
		},
	},
	{
		name:         "DIVQU2",
		argLen:       3,
		clobberFlags: true,
		asm:          x86.ADIVQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4},     // DX
				{1, 1},     // AX
				{2, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 1}, // AX
				{1, 4}, // DX
			},
		},
	},
	{
		name:         "ANDQ",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AANDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ANDL",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ANDQconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AANDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ANDLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "ANDQconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AANDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "ANDLconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:         "ORQ",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ORL",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ORQconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ORLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "ORQconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "ORLconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:         "XORQ",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AXORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "XORL",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "XORQconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AXORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "XORLconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "XORQconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AXORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "XORLconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:   "CMPQ",
		argLen: 2,
		asm:    x86.ACMPQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "CMPL",
		argLen: 2,
		asm:    x86.ACMPL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "CMPW",
		argLen: 2,
		asm:    x86.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "CMPB",
		argLen: 2,
		asm:    x86.ACMPB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "CMPQconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     x86.ACMPQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "CMPLconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     x86.ACMPL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "CMPWconst",
		auxType: auxInt16,
		argLen:  1,
		asm:     x86.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "CMPBconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     x86.ACMPB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "CMPQload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "CMPLload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "CMPWload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "CMPBload",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "CMPQconstload",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "CMPLconstload",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "CMPWconstload",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "CMPBconstload",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.ACMPB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:   "UCOMISS",
		argLen: 2,
		asm:    x86.AUCOMISS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:   "UCOMISD",
		argLen: 2,
		asm:    x86.AUCOMISD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:   "BTL",
		argLen: 2,
		asm:    x86.ABTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "BTQ",
		argLen: 2,
		asm:    x86.ABTQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BTCL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTCL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BTCQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTCQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BTRL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BTRQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTRQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BTSL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTSL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BTSQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTSQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "BTLconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     x86.ABTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "BTQconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     x86.ABTQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BTCLconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTCL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BTCQconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTCQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BTRLconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BTRQconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTRQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BTSLconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTSL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BTSQconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABTSQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "BTCQmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ABTCQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "BTCLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ABTCL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "BTSQmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ABTSQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "BTSLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ABTSL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "BTRQmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ABTRQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "BTRLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ABTRL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "BTCQconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ABTCQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "BTCLconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ABTCL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "BTSQconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ABTSQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "BTSLconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ABTSL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "BTRQconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ABTRQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "BTRLconstmodify",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ABTRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:        "TESTQ",
		argLen:      2,
		commutative: true,
		asm:         x86.ATESTQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:        "TESTL",
		argLen:      2,
		commutative: true,
		asm:         x86.ATESTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:        "TESTW",
		argLen:      2,
		commutative: true,
		asm:         x86.ATESTW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:        "TESTB",
		argLen:      2,
		commutative: true,
		asm:         x86.ATESTB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "TESTQconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     x86.ATESTQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "TESTLconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     x86.ATESTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "TESTWconst",
		auxType: auxInt16,
		argLen:  1,
		asm:     x86.ATESTW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "TESTBconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     x86.ATESTB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SHLQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHLQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SHLL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHLL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SHLQconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHLQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SHLLconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHLL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SHRQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SHRL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SHRW",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SHRB",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SHRQconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SHRLconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SHRWconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SHRBconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASHRB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SARQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SARL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SARW",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SARB",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SARQconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SARLconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SARWconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SARBconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ASARB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ROLQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ROLL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ROLW",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ROLB",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "RORQ",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ARORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "RORL",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ARORL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "RORW",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ARORW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "RORB",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ARORB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // CX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ROLQconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ROLLconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ROLWconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "ROLBconst",
		auxType:      auxInt8,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.AROLB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "ADDLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "ADDQload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "SUBQload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ASUBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "SUBLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "ANDLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "ANDQload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AANDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "ORQload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "ORLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "XORQload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AXORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "XORLload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "ADDQmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "SUBQmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ASUBQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "ANDQmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AANDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "ORQmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "XORQmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AXORQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "ADDLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "SUBLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.ASUBL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "ANDLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AANDL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "ORLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AORL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "XORLmodify",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymRead | SymWrite,
		asm:            x86.AXORL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:         "NEGQ",
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ANEGQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "NEGL",
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ANEGL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "NOTQ",
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ANOTQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "NOTL",
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ANOTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "BSFQ",
		argLen: 1,
		asm:    x86.ABSFQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BSFL",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ABSFL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "BSRQ",
		argLen: 1,
		asm:    x86.ABSRQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BSRL",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ABSRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVQEQ",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVQNE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVQLT",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQLT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVQGT",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQGT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVQLE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQLE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVQGE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQGE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVQLS",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQLS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVQHI",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQHI,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVQCC",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVQCS",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQCS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVLEQ",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVLNE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVLLT",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLLT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVLGT",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLGT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVLLE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLLE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVLGE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLGE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVLLS",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLLS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVLHI",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLHI,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVLCC",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVLCS",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLCS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVWEQ",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVWNE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVWLT",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWLT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVWGT",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWGT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVWLE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWLE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVWGE",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWGE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVWLS",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWLS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVWHI",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWHI,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVWCC",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVWCS",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWCS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVQEQF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65518}, // CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 65518}, // CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVQNEF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVQGTF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQHI,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVQGEF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVQCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVLEQF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65518}, // CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 65518}, // CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVLNEF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVLGTF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLHI,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVLGEF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVLCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVWEQF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65518}, // CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 65518}, // CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVWNEF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVWGTF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWHI,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "CMOVWGEF",
		argLen:       3,
		resultInArg0: true,
		asm:          x86.ACMOVWCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BSWAPQ",
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABSWAPQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "BSWAPL",
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          x86.ABSWAPL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "POPCNTQ",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.APOPCNTQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "POPCNTL",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.APOPCNTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SQRTSD",
		argLen: 1,
		asm:    x86.ASQRTSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:    "ROUNDSD",
		auxType: auxInt8,
		argLen:  1,
		asm:     x86.AROUNDSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:   "SBBQcarrymask",
		argLen: 1,
		asm:    x86.ASBBQ,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SBBLcarrymask",
		argLen: 1,
		asm:    x86.ASBBL,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SETEQ",
		argLen: 1,
		asm:    x86.ASETEQ,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SETNE",
		argLen: 1,
		asm:    x86.ASETNE,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SETL",
		argLen: 1,
		asm:    x86.ASETLT,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SETLE",
		argLen: 1,
		asm:    x86.ASETLE,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SETG",
		argLen: 1,
		asm:    x86.ASETGT,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SETGE",
		argLen: 1,
		asm:    x86.ASETGE,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SETB",
		argLen: 1,
		asm:    x86.ASETCS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SETBE",
		argLen: 1,
		asm:    x86.ASETLS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SETA",
		argLen: 1,
		asm:    x86.ASETHI,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SETAE",
		argLen: 1,
		asm:    x86.ASETCC,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SETO",
		argLen: 1,
		asm:    x86.ASETOS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "SETEQstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "SETNEstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "SETLstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETLT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "SETLEstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETLE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "SETGstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETGT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "SETGEstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETGE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "SETBstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETCS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "SETBEstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETLS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "SETAstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETHI,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "SETAEstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.ASETCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:         "SETEQF",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ASETEQ,
		reg: regInfo{
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 65518}, // CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "SETNEF",
		argLen:       1,
		clobberFlags: true,
		asm:          x86.ASETNE,
		reg: regInfo{
			clobbers: 1, // AX
			outputs: []outputInfo{
				{0, 65518}, // CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SETORD",
		argLen: 1,
		asm:    x86.ASETPC,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SETNAN",
		argLen: 1,
		asm:    x86.ASETPS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SETGF",
		argLen: 1,
		asm:    x86.ASETHI,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "SETGEF",
		argLen: 1,
		asm:    x86.ASETCC,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "MOVBQSX",
		argLen: 1,
		asm:    x86.AMOVBQSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "MOVBQZX",
		argLen: 1,
		asm:    x86.AMOVBLZX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "MOVWQSX",
		argLen: 1,
		asm:    x86.AMOVWQSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "MOVWQZX",
		argLen: 1,
		asm:    x86.AMOVWLZX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "MOVLQSX",
		argLen: 1,
		asm:    x86.AMOVLQSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "MOVLQZX",
		argLen: 1,
		asm:    x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:              "MOVLconst",
		auxType:           auxInt32,
		argLen:            0,
		rematerializeable: true,
		asm:               x86.AMOVL,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:              "MOVQconst",
		auxType:           auxInt64,
		argLen:            0,
		rematerializeable: true,
		asm:               x86.AMOVQ,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "CVTTSD2SL",
		argLen: 1,
		asm:    x86.ACVTTSD2SL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "CVTTSD2SQ",
		argLen: 1,
		asm:    x86.ACVTTSD2SQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "CVTTSS2SL",
		argLen: 1,
		asm:    x86.ACVTTSS2SL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "CVTTSS2SQ",
		argLen: 1,
		asm:    x86.ACVTTSS2SQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "CVTSL2SS",
		argLen: 1,
		asm:    x86.ACVTSL2SS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:   "CVTSL2SD",
		argLen: 1,
		asm:    x86.ACVTSL2SD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:   "CVTSQ2SS",
		argLen: 1,
		asm:    x86.ACVTSQ2SS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:   "CVTSQ2SD",
		argLen: 1,
		asm:    x86.ACVTSQ2SD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:   "CVTSD2SS",
		argLen: 1,
		asm:    x86.ACVTSD2SS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:   "CVTSS2SD",
		argLen: 1,
		asm:    x86.ACVTSS2SD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:   "MOVQi2f",
		argLen: 1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:   "MOVQf2i",
		argLen: 1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "MOVLi2f",
		argLen: 1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:   "MOVLf2i",
		argLen: 1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "PXOR",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          x86.APXOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:              "LEAQ",
		auxType:           auxSymOff,
		argLen:            1,
		rematerializeable: true,
		symEffect:         SymAddr,
		asm:               x86.ALEAQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:              "LEAL",
		auxType:           auxSymOff,
		argLen:            1,
		rematerializeable: true,
		symEffect:         SymAddr,
		asm:               x86.ALEAL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:              "LEAW",
		auxType:           auxSymOff,
		argLen:            1,
		rematerializeable: true,
		symEffect:         SymAddr,
		asm:               x86.ALEAW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:        "LEAQ1",
		auxType:     auxSymOff,
		argLen:      2,
		commutative: true,
		symEffect:   SymAddr,
		asm:         x86.ALEAQ,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:        "LEAL1",
		auxType:     auxSymOff,
		argLen:      2,
		commutative: true,
		symEffect:   SymAddr,
		asm:         x86.ALEAL,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:        "LEAW1",
		auxType:     auxSymOff,
		argLen:      2,
		commutative: true,
		symEffect:   SymAddr,
		asm:         x86.ALEAW,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "LEAQ2",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAQ,
		scale:     2,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "LEAL2",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAL,
		scale:     2,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "LEAW2",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAW,
		scale:     2,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "LEAQ4",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAQ,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "LEAL4",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAL,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "LEAW4",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAW,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "LEAQ8",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAQ,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "LEAL8",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAL,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "LEAW8",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymAddr,
		asm:       x86.ALEAW,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "MOVBload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVBLZX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "MOVBQSXload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVBQSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "MOVWload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVWLZX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "MOVWQSXload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVWQSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "MOVLload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "MOVLQSXload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVLQSX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "MOVQload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "MOVBstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "MOVWstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "MOVLstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "MOVQstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "MOVOload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVUPS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:           "MOVOstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVUPS,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:        "MOVBloadidx1",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.AMOVBLZX,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:        "MOVWloadidx1",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.AMOVWLZX,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "MOVWloadidx2",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVWLZX,
		scale:     2,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:        "MOVLloadidx1",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.AMOVL,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "MOVLloadidx4",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVL,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "MOVLloadidx8",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVL,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:        "MOVQloadidx1",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         x86.AMOVQ,
		scale:       1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "MOVQloadidx8",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       x86.AMOVQ,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "MOVBstoreidx1",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVB,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVWstoreidx1",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVW,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVWstoreidx2",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVW,
		scale:     2,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVLstoreidx1",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVL,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVLstoreidx4",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVL,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVLstoreidx8",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVL,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVQstoreidx1",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVQ,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVQstoreidx8",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       x86.AMOVQ,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "MOVBstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "MOVWstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "MOVLstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "MOVQstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            x86.AMOVQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVBstoreconstidx1",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymWrite,
		asm:       x86.AMOVB,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVWstoreconstidx1",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymWrite,
		asm:       x86.AMOVW,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVWstoreconstidx2",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymWrite,
		asm:       x86.AMOVW,
		scale:     2,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVLstoreconstidx1",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymWrite,
		asm:       x86.AMOVL,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVLstoreconstidx4",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymWrite,
		asm:       x86.AMOVL,
		scale:     4,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVQstoreconstidx1",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymWrite,
		asm:       x86.AMOVQ,
		scale:     1,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:      "MOVQstoreconstidx8",
		auxType:   auxSymValAndOff,
		argLen:    3,
		symEffect: SymWrite,
		asm:       x86.AMOVQ,
		scale:     8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "DUFFZERO",
		auxType:        auxInt64,
		argLen:         3,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 128},   // DI
				{1, 65536}, // X0
			},
			clobbers: 128, // DI
		},
	},
	{
		name:              "MOVOconst",
		auxType:           auxInt128,
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
	},
	{
		name:           "REPSTOSQ",
		argLen:         4,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 128}, // DI
				{1, 2},   // CX
				{2, 1},   // AX
			},
			clobbers: 130, // CX DI
		},
	},
	{
		name:         "CALLstatic",
		auxType:      auxSymOff,
		argLen:       1,
		clobberFlags: true,
		call:         true,
		symEffect:    SymNone,
		reg: regInfo{
			clobbers: 4294967279, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
		},
	},
	{
		name:         "CALLclosure",
		auxType:      auxInt64,
		argLen:       3,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 4},     // DX
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			clobbers: 4294967279, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
		},
	},
	{
		name:         "CALLinter",
		auxType:      auxInt64,
		argLen:       2,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			clobbers: 4294967279, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
		},
	},
	{
		name:           "DUFFCOPY",
		auxType:        auxInt64,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		faultOnNilArg1: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 128}, // DI
				{1, 64},  // SI
			},
			clobbers: 65728, // SI DI X0
		},
	},
	{
		name:           "REPMOVSQ",
		argLen:         4,
		faultOnNilArg0: true,
		faultOnNilArg1: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 128}, // DI
				{1, 64},  // SI
				{2, 2},   // CX
			},
			clobbers: 194, // CX SI DI
		},
	},
	{
		name:   "InvertFlags",
		argLen: 1,
		reg:    regInfo{},
	},
	{
		name:   "LoweredGetG",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "LoweredGetClosurePtr",
		argLen:    0,
		zeroWidth: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 4}, // DX
			},
		},
	},
	{
		name:              "LoweredGetCallerPC",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:              "LoweredGetCallerSP",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "LoweredNilCheck",
		argLen:         2,
		clobberFlags:   true,
		nilCheck:       true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:         "LoweredWB",
		auxType:      auxSym,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymNone,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 128}, // DI
				{1, 1},   // AX
			},
			clobbers: 4294901760, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
		},
	},
	{
		name:    "LoweredPanicBoundsA",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4}, // DX
				{1, 8}, // BX
			},
		},
	},
	{
		name:    "LoweredPanicBoundsB",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2}, // CX
				{1, 4}, // DX
			},
		},
	},
	{
		name:    "LoweredPanicBoundsC",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1}, // AX
				{1, 2}, // CX
			},
		},
	},
	{
		name:    "LoweredPanicExtendA",
		auxType: auxInt64,
		argLen:  4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 64}, // SI
				{1, 4},  // DX
				{2, 8},  // BX
			},
		},
	},
	{
		name:    "LoweredPanicExtendB",
		auxType: auxInt64,
		argLen:  4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 64}, // SI
				{1, 2},  // CX
				{2, 4},  // DX
			},
		},
	},
	{
		name:    "LoweredPanicExtendC",
		auxType: auxInt64,
		argLen:  4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 64}, // SI
				{1, 1},  // AX
				{2, 2},  // CX
			},
		},
	},
	{
		name:   "FlagEQ",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagLT_ULT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagLT_UGT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagGT_UGT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagGT_ULT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:           "MOVBatomicload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "MOVLatomicload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "MOVQatomicload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            x86.AMOVQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "XCHGL",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.AXCHGL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "XCHGQ",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		faultOnNilArg1: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.AXCHGQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "XADDLlock",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.AXADDL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "XADDQlock",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.AXADDQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
			outputs: []outputInfo{
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "AddTupleFirst32",
		argLen: 2,
		reg:    regInfo{},
	},
	{
		name:   "AddTupleFirst64",
		argLen: 2,
		reg:    regInfo{},
	},
	{
		name:           "CMPXCHGLlock",
		auxType:        auxSymOff,
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.ACMPXCHGL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1},     // AX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{2, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{1, 0},
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "CMPXCHGQlock",
		auxType:        auxSymOff,
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.ACMPXCHGQ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1},     // AX
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{2, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
			clobbers: 1, // AX
			outputs: []outputInfo{
				{1, 0},
				{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "ANDBlock",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.AANDB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},
	{
		name:           "ORBlock",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            x86.AORB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
			},
		},
	},

	{
		name:        "ADD",
		argLen:      2,
		commutative: true,
		asm:         arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ADDconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 30719}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SUB",
		argLen: 2,
		asm:    arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SUBconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "RSB",
		argLen: 2,
		asm:    arm.ARSB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "RSBconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.ARSB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:        "MUL",
		argLen:      2,
		commutative: true,
		asm:         arm.AMUL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:        "HMUL",
		argLen:      2,
		commutative: true,
		asm:         arm.AMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:        "HMULU",
		argLen:      2,
		commutative: true,
		asm:         arm.AMULLU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:         "CALLudiv",
		argLen:       2,
		clobberFlags: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2}, // R1
				{1, 1}, // R0
			},
			clobbers: 16396, // R2 R3 R14
			outputs: []outputInfo{
				{0, 1}, // R0
				{1, 2}, // R1
			},
		},
	},
	{
		name:        "ADDS",
		argLen:      2,
		commutative: true,
		asm:         arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ADDSconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:        "ADC",
		argLen:      3,
		commutative: true,
		asm:         arm.AADC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ADCconst",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AADC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SUBS",
		argLen: 2,
		asm:    arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SUBSconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "RSBSconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.ARSB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SBC",
		argLen: 3,
		asm:    arm.ASBC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SBCconst",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ASBC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "RSCconst",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ARSC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:        "MULLU",
		argLen:      2,
		commutative: true,
		asm:         arm.AMULLU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MULA",
		argLen: 3,
		asm:    arm.AMULA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MULS",
		argLen: 3,
		asm:    arm.AMULS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:        "ADDF",
		argLen:      2,
		commutative: true,
		asm:         arm.AADDF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:        "ADDD",
		argLen:      2,
		commutative: true,
		asm:         arm.AADDD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "SUBF",
		argLen: 2,
		asm:    arm.ASUBF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "SUBD",
		argLen: 2,
		asm:    arm.ASUBD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:        "MULF",
		argLen:      2,
		commutative: true,
		asm:         arm.AMULF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:        "MULD",
		argLen:      2,
		commutative: true,
		asm:         arm.AMULD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:        "NMULF",
		argLen:      2,
		commutative: true,
		asm:         arm.ANMULF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:        "NMULD",
		argLen:      2,
		commutative: true,
		asm:         arm.ANMULD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "DIVF",
		argLen: 2,
		asm:    arm.ADIVF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "DIVD",
		argLen: 2,
		asm:    arm.ADIVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "MULAF",
		argLen:       3,
		resultInArg0: true,
		asm:          arm.AMULAF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "MULAD",
		argLen:       3,
		resultInArg0: true,
		asm:          arm.AMULAD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "MULSF",
		argLen:       3,
		resultInArg0: true,
		asm:          arm.AMULSF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "MULSD",
		argLen:       3,
		resultInArg0: true,
		asm:          arm.AMULSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:        "AND",
		argLen:      2,
		commutative: true,
		asm:         arm.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ANDconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:        "OR",
		argLen:      2,
		commutative: true,
		asm:         arm.AORR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ORconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.AORR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:        "XOR",
		argLen:      2,
		commutative: true,
		asm:         arm.AEOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "XORconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.AEOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "BIC",
		argLen: 2,
		asm:    arm.ABIC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "BICconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.ABIC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "BFX",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.ABFX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "BFXU",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.ABFXU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MVN",
		argLen: 1,
		asm:    arm.AMVN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "NEGF",
		argLen: 1,
		asm:    arm.ANEGF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "NEGD",
		argLen: 1,
		asm:    arm.ANEGD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "SQRTD",
		argLen: 1,
		asm:    arm.ASQRTD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "CLZ",
		argLen: 1,
		asm:    arm.ACLZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "REV",
		argLen: 1,
		asm:    arm.AREV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "REV16",
		argLen: 1,
		asm:    arm.AREV16,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "RBIT",
		argLen: 1,
		asm:    arm.ARBIT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SLL",
		argLen: 2,
		asm:    arm.ASLL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SLLconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.ASLL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SRL",
		argLen: 2,
		asm:    arm.ASRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SRLconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.ASRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SRA",
		argLen: 2,
		asm:    arm.ASRA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SRAconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.ASRA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SRRconst",
		auxType: auxInt32,
		argLen:  1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ADDshiftLL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ADDshiftRL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ADDshiftRA",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SUBshiftLL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SUBshiftRL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SUBshiftRA",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "RSBshiftLL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ARSB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "RSBshiftRL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ARSB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "RSBshiftRA",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ARSB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ANDshiftLL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ANDshiftRL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ANDshiftRA",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ORshiftLL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AORR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ORshiftRL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AORR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ORshiftRA",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AORR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "XORshiftLL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AEOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "XORshiftRL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AEOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "XORshiftRA",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AEOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "XORshiftRR",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AEOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "BICshiftLL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ABIC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "BICshiftRL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ABIC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "BICshiftRA",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ABIC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "MVNshiftLL",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.AMVN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "MVNshiftRL",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.AMVN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "MVNshiftRA",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.AMVN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ADCshiftLL",
		auxType: auxInt32,
		argLen:  3,
		asm:     arm.AADC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ADCshiftRL",
		auxType: auxInt32,
		argLen:  3,
		asm:     arm.AADC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ADCshiftRA",
		auxType: auxInt32,
		argLen:  3,
		asm:     arm.AADC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SBCshiftLL",
		auxType: auxInt32,
		argLen:  3,
		asm:     arm.ASBC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SBCshiftRL",
		auxType: auxInt32,
		argLen:  3,
		asm:     arm.ASBC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SBCshiftRA",
		auxType: auxInt32,
		argLen:  3,
		asm:     arm.ASBC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "RSCshiftLL",
		auxType: auxInt32,
		argLen:  3,
		asm:     arm.ARSC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "RSCshiftRL",
		auxType: auxInt32,
		argLen:  3,
		asm:     arm.ARSC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "RSCshiftRA",
		auxType: auxInt32,
		argLen:  3,
		asm:     arm.ARSC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ADDSshiftLL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ADDSshiftRL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "ADDSshiftRA",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SUBSshiftLL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SUBSshiftRL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "SUBSshiftRA",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "RSBSshiftLL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ARSB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "RSBSshiftRL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ARSB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "RSBSshiftRA",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ARSB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "ADDshiftLLreg",
		argLen: 3,
		asm:    arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "ADDshiftRLreg",
		argLen: 3,
		asm:    arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "ADDshiftRAreg",
		argLen: 3,
		asm:    arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SUBshiftLLreg",
		argLen: 3,
		asm:    arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SUBshiftRLreg",
		argLen: 3,
		asm:    arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SUBshiftRAreg",
		argLen: 3,
		asm:    arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "RSBshiftLLreg",
		argLen: 3,
		asm:    arm.ARSB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "RSBshiftRLreg",
		argLen: 3,
		asm:    arm.ARSB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "RSBshiftRAreg",
		argLen: 3,
		asm:    arm.ARSB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "ANDshiftLLreg",
		argLen: 3,
		asm:    arm.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "ANDshiftRLreg",
		argLen: 3,
		asm:    arm.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "ANDshiftRAreg",
		argLen: 3,
		asm:    arm.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "ORshiftLLreg",
		argLen: 3,
		asm:    arm.AORR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "ORshiftRLreg",
		argLen: 3,
		asm:    arm.AORR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "ORshiftRAreg",
		argLen: 3,
		asm:    arm.AORR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "XORshiftLLreg",
		argLen: 3,
		asm:    arm.AEOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "XORshiftRLreg",
		argLen: 3,
		asm:    arm.AEOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "XORshiftRAreg",
		argLen: 3,
		asm:    arm.AEOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "BICshiftLLreg",
		argLen: 3,
		asm:    arm.ABIC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "BICshiftRLreg",
		argLen: 3,
		asm:    arm.ABIC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "BICshiftRAreg",
		argLen: 3,
		asm:    arm.ABIC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MVNshiftLLreg",
		argLen: 2,
		asm:    arm.AMVN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MVNshiftRLreg",
		argLen: 2,
		asm:    arm.AMVN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MVNshiftRAreg",
		argLen: 2,
		asm:    arm.AMVN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "ADCshiftLLreg",
		argLen: 4,
		asm:    arm.AADC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "ADCshiftRLreg",
		argLen: 4,
		asm:    arm.AADC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "ADCshiftRAreg",
		argLen: 4,
		asm:    arm.AADC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SBCshiftLLreg",
		argLen: 4,
		asm:    arm.ASBC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SBCshiftRLreg",
		argLen: 4,
		asm:    arm.ASBC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SBCshiftRAreg",
		argLen: 4,
		asm:    arm.ASBC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "RSCshiftLLreg",
		argLen: 4,
		asm:    arm.ARSC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "RSCshiftRLreg",
		argLen: 4,
		asm:    arm.ARSC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "RSCshiftRAreg",
		argLen: 4,
		asm:    arm.ARSC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "ADDSshiftLLreg",
		argLen: 3,
		asm:    arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "ADDSshiftRLreg",
		argLen: 3,
		asm:    arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "ADDSshiftRAreg",
		argLen: 3,
		asm:    arm.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SUBSshiftLLreg",
		argLen: 3,
		asm:    arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SUBSshiftRLreg",
		argLen: 3,
		asm:    arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SUBSshiftRAreg",
		argLen: 3,
		asm:    arm.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "RSBSshiftLLreg",
		argLen: 3,
		asm:    arm.ARSB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "RSBSshiftRLreg",
		argLen: 3,
		asm:    arm.ARSB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "RSBSshiftRAreg",
		argLen: 3,
		asm:    arm.ARSB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "CMP",
		argLen: 2,
		asm:    arm.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:    "CMPconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:        "CMN",
		argLen:      2,
		commutative: true,
		asm:         arm.ACMN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:    "CMNconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.ACMN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:        "TST",
		argLen:      2,
		commutative: true,
		asm:         arm.ATST,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:    "TSTconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.ATST,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:        "TEQ",
		argLen:      2,
		commutative: true,
		asm:         arm.ATEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:    "TEQconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm.ATEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:   "CMPF",
		argLen: 2,
		asm:    arm.ACMPF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "CMPD",
		argLen: 2,
		asm:    arm.ACMPD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:    "CMPshiftLL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:    "CMPshiftRL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:    "CMPshiftRA",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:    "CMNshiftLL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ACMN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:    "CMNshiftRL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ACMN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:    "CMNshiftRA",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ACMN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:    "TSTshiftLL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ATST,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:    "TSTshiftRL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ATST,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:    "TSTshiftRA",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ATST,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:    "TEQshiftLL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ATEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:    "TEQshiftRL",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ATEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:    "TEQshiftRA",
		auxType: auxInt32,
		argLen:  2,
		asm:     arm.ATEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:   "CMPshiftLLreg",
		argLen: 3,
		asm:    arm.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "CMPshiftRLreg",
		argLen: 3,
		asm:    arm.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "CMPshiftRAreg",
		argLen: 3,
		asm:    arm.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "CMNshiftLLreg",
		argLen: 3,
		asm:    arm.ACMN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "CMNshiftRLreg",
		argLen: 3,
		asm:    arm.ACMN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "CMNshiftRAreg",
		argLen: 3,
		asm:    arm.ACMN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "TSTshiftLLreg",
		argLen: 3,
		asm:    arm.ATST,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "TSTshiftRLreg",
		argLen: 3,
		asm:    arm.ATST,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "TSTshiftRAreg",
		argLen: 3,
		asm:    arm.ATST,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "TEQshiftLLreg",
		argLen: 3,
		asm:    arm.ATEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "TEQshiftRLreg",
		argLen: 3,
		asm:    arm.ATEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "TEQshiftRAreg",
		argLen: 3,
		asm:    arm.ATEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "CMPF0",
		argLen: 1,
		asm:    arm.ACMPF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "CMPD0",
		argLen: 1,
		asm:    arm.ACMPD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:              "MOVWconst",
		auxType:           auxInt32,
		argLen:            0,
		rematerializeable: true,
		asm:               arm.AMOVW,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:              "MOVFconst",
		auxType:           auxFloat64,
		argLen:            0,
		rematerializeable: true,
		asm:               arm.AMOVF,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:              "MOVDconst",
		auxType:           auxFloat64,
		argLen:            0,
		rematerializeable: true,
		asm:               arm.AMOVD,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:              "MOVWaddr",
		auxType:           auxSymOff,
		argLen:            1,
		rematerializeable: true,
		symEffect:         SymAddr,
		asm:               arm.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294975488}, // SP SB
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:           "MOVBload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:           "MOVBUload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm.AMOVBU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:           "MOVHload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:           "MOVHUload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm.AMOVHU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:           "MOVWload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:           "MOVFload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm.AMOVF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:           "MOVDload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:           "MOVBstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
		},
	},
	{
		name:           "MOVHstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
		},
	},
	{
		name:           "MOVWstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
		},
	},
	{
		name:           "MOVFstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm.AMOVF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:           "MOVDstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "MOVWloadidx",
		argLen: 3,
		asm:    arm.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "MOVWloadshiftLL",
		auxType: auxInt32,
		argLen:  3,
		asm:     arm.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "MOVWloadshiftRL",
		auxType: auxInt32,
		argLen:  3,
		asm:     arm.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "MOVWloadshiftRA",
		auxType: auxInt32,
		argLen:  3,
		asm:     arm.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MOVBUloadidx",
		argLen: 3,
		asm:    arm.AMOVBU,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MOVBloadidx",
		argLen: 3,
		asm:    arm.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MOVHUloadidx",
		argLen: 3,
		asm:    arm.AMOVHU,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MOVHloadidx",
		argLen: 3,
		asm:    arm.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MOVWstoreidx",
		argLen: 4,
		asm:    arm.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{2, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
		},
	},
	{
		name:    "MOVWstoreshiftLL",
		auxType: auxInt32,
		argLen:  4,
		asm:     arm.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{2, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
		},
	},
	{
		name:    "MOVWstoreshiftRL",
		auxType: auxInt32,
		argLen:  4,
		asm:     arm.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{2, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
		},
	},
	{
		name:    "MOVWstoreshiftRA",
		auxType: auxInt32,
		argLen:  4,
		asm:     arm.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{2, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
		},
	},
	{
		name:   "MOVBstoreidx",
		argLen: 4,
		asm:    arm.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{2, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
		},
	},
	{
		name:   "MOVHstoreidx",
		argLen: 4,
		asm:    arm.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{2, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
			},
		},
	},
	{
		name:   "MOVBreg",
		argLen: 1,
		asm:    arm.AMOVBS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MOVBUreg",
		argLen: 1,
		asm:    arm.AMOVBU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MOVHreg",
		argLen: 1,
		asm:    arm.AMOVHS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MOVHUreg",
		argLen: 1,
		asm:    arm.AMOVHU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MOVWreg",
		argLen: 1,
		asm:    arm.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:         "MOVWnop",
		argLen:       1,
		resultInArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MOVWF",
		argLen: 1,
		asm:    arm.AMOVWF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 2147483648, // F15
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "MOVWD",
		argLen: 1,
		asm:    arm.AMOVWD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 2147483648, // F15
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "MOVWUF",
		argLen: 1,
		asm:    arm.AMOVWF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 2147483648, // F15
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "MOVWUD",
		argLen: 1,
		asm:    arm.AMOVWD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 2147483648, // F15
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "MOVFW",
		argLen: 1,
		asm:    arm.AMOVFW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			clobbers: 2147483648, // F15
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MOVDW",
		argLen: 1,
		asm:    arm.AMOVDW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			clobbers: 2147483648, // F15
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MOVFWU",
		argLen: 1,
		asm:    arm.AMOVFW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			clobbers: 2147483648, // F15
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MOVDWU",
		argLen: 1,
		asm:    arm.AMOVDW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			clobbers: 2147483648, // F15
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "MOVFD",
		argLen: 1,
		asm:    arm.AMOVFD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "MOVDF",
		argLen: 1,
		asm:    arm.AMOVDF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "CMOVWHSconst",
		auxType:      auxInt32,
		argLen:       2,
		resultInArg0: true,
		asm:          arm.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:         "CMOVWLSconst",
		auxType:      auxInt32,
		argLen:       2,
		resultInArg0: true,
		asm:          arm.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "SRAcond",
		argLen: 3,
		asm:    arm.ASRA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:         "CALLstatic",
		auxType:      auxSymOff,
		argLen:       1,
		clobberFlags: true,
		call:         true,
		symEffect:    SymNone,
		reg: regInfo{
			clobbers: 4294924287, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
		},
	},
	{
		name:         "CALLclosure",
		auxType:      auxInt64,
		argLen:       3,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 128},   // R7
				{0, 29695}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 SP R14
			},
			clobbers: 4294924287, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
		},
	},
	{
		name:         "CALLinter",
		auxType:      auxInt64,
		argLen:       2,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 4294924287, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
		},
	},
	{
		name:           "LoweredNilCheck",
		argLen:         2,
		nilCheck:       true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
			},
		},
	},
	{
		name:   "Equal",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "NotEqual",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "LessThan",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "LessEqual",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "GreaterThan",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "GreaterEqual",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "LessThanU",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "LessEqualU",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "GreaterThanU",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:   "GreaterEqualU",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:           "DUFFZERO",
		auxType:        auxInt64,
		argLen:         3,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2}, // R1
				{1, 1}, // R0
			},
			clobbers: 16386, // R1 R14
		},
	},
	{
		name:           "DUFFCOPY",
		auxType:        auxInt64,
		argLen:         3,
		faultOnNilArg0: true,
		faultOnNilArg1: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4}, // R2
				{1, 2}, // R1
			},
			clobbers: 16391, // R0 R1 R2 R14
		},
	},
	{
		name:           "LoweredZero",
		auxType:        auxInt64,
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2},     // R1
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 2, // R1
		},
	},
	{
		name:           "LoweredMove",
		auxType:        auxInt64,
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		faultOnNilArg1: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4},     // R2
				{1, 2},     // R1
				{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 6, // R1 R2
		},
	},
	{
		name:      "LoweredGetClosurePtr",
		argLen:    0,
		zeroWidth: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 128}, // R7
			},
		},
	},
	{
		name:              "LoweredGetCallerSP",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:              "LoweredGetCallerPC",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:    "LoweredPanicBoundsA",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4}, // R2
				{1, 8}, // R3
			},
		},
	},
	{
		name:    "LoweredPanicBoundsB",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2}, // R1
				{1, 4}, // R2
			},
		},
	},
	{
		name:    "LoweredPanicBoundsC",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1}, // R0
				{1, 2}, // R1
			},
		},
	},
	{
		name:    "LoweredPanicExtendA",
		auxType: auxInt64,
		argLen:  4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 16}, // R4
				{1, 4},  // R2
				{2, 8},  // R3
			},
		},
	},
	{
		name:    "LoweredPanicExtendB",
		auxType: auxInt64,
		argLen:  4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 16}, // R4
				{1, 2},  // R1
				{2, 4},  // R2
			},
		},
	},
	{
		name:    "LoweredPanicExtendC",
		auxType: auxInt64,
		argLen:  4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 16}, // R4
				{1, 1},  // R0
				{2, 2},  // R1
			},
		},
	},
	{
		name:   "FlagEQ",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagLT_ULT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagLT_UGT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagGT_UGT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagGT_ULT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "InvertFlags",
		argLen: 1,
		reg:    regInfo{},
	},
	{
		name:         "LoweredWB",
		auxType:      auxSym,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymNone,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4}, // R2
				{1, 8}, // R3
			},
			clobbers: 4294918144, // R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
		},
	},

	{
		name:        "ADCSflags",
		argLen:      3,
		commutative: true,
		asm:         arm64.AADCS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				{1, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "ADCzerocarry",
		argLen: 1,
		asm:    arm64.AADC,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:        "ADD",
		argLen:      2,
		commutative: true,
		asm:         arm64.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ADDconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1878786047}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ADDSconstflags",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.AADDS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:        "ADDSflags",
		argLen:      2,
		commutative: true,
		asm:         arm64.AADDS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				{1, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "SUB",
		argLen: 2,
		asm:    arm64.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "SUBconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "SBCSflags",
		argLen: 3,
		asm:    arm64.ASBCS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				{1, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "SUBSflags",
		argLen: 2,
		asm:    arm64.ASUBS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				{1, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:        "MUL",
		argLen:      2,
		commutative: true,
		asm:         arm64.AMUL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:        "MULW",
		argLen:      2,
		commutative: true,
		asm:         arm64.AMULW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:        "MNEG",
		argLen:      2,
		commutative: true,
		asm:         arm64.AMNEG,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:        "MNEGW",
		argLen:      2,
		commutative: true,
		asm:         arm64.AMNEGW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:        "MULH",
		argLen:      2,
		commutative: true,
		asm:         arm64.ASMULH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:        "UMULH",
		argLen:      2,
		commutative: true,
		asm:         arm64.AUMULH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:        "MULL",
		argLen:      2,
		commutative: true,
		asm:         arm64.ASMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:        "UMULL",
		argLen:      2,
		commutative: true,
		asm:         arm64.AUMULL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "DIV",
		argLen: 2,
		asm:    arm64.ASDIV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "UDIV",
		argLen: 2,
		asm:    arm64.AUDIV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "DIVW",
		argLen: 2,
		asm:    arm64.ASDIVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "UDIVW",
		argLen: 2,
		asm:    arm64.AUDIVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOD",
		argLen: 2,
		asm:    arm64.AREM,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "UMOD",
		argLen: 2,
		asm:    arm64.AUREM,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MODW",
		argLen: 2,
		asm:    arm64.AREMW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "UMODW",
		argLen: 2,
		asm:    arm64.AUREMW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:        "FADDS",
		argLen:      2,
		commutative: true,
		asm:         arm64.AFADDS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:        "FADDD",
		argLen:      2,
		commutative: true,
		asm:         arm64.AFADDD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FSUBS",
		argLen: 2,
		asm:    arm64.AFSUBS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FSUBD",
		argLen: 2,
		asm:    arm64.AFSUBD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:        "FMULS",
		argLen:      2,
		commutative: true,
		asm:         arm64.AFMULS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:        "FMULD",
		argLen:      2,
		commutative: true,
		asm:         arm64.AFMULD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:        "FNMULS",
		argLen:      2,
		commutative: true,
		asm:         arm64.AFNMULS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:        "FNMULD",
		argLen:      2,
		commutative: true,
		asm:         arm64.AFNMULD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FDIVS",
		argLen: 2,
		asm:    arm64.AFDIVS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FDIVD",
		argLen: 2,
		asm:    arm64.AFDIVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:        "AND",
		argLen:      2,
		commutative: true,
		asm:         arm64.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ANDconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:        "OR",
		argLen:      2,
		commutative: true,
		asm:         arm64.AORR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ORconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.AORR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:        "XOR",
		argLen:      2,
		commutative: true,
		asm:         arm64.AEOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "XORconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.AEOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "BIC",
		argLen: 2,
		asm:    arm64.ABIC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "EON",
		argLen: 2,
		asm:    arm64.AEON,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "ORN",
		argLen: 2,
		asm:    arm64.AORN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:            "LoweredMuluhilo",
		argLen:          2,
		resultNotInArgs: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				{1, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MVN",
		argLen: 1,
		asm:    arm64.AMVN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "NEG",
		argLen: 1,
		asm:    arm64.ANEG,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "NEGSflags",
		argLen: 1,
		asm:    arm64.ANEGS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "NGCzerocarry",
		argLen: 1,
		asm:    arm64.ANGC,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "FABSD",
		argLen: 1,
		asm:    arm64.AFABSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FNEGS",
		argLen: 1,
		asm:    arm64.AFNEGS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FNEGD",
		argLen: 1,
		asm:    arm64.AFNEGD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FSQRTD",
		argLen: 1,
		asm:    arm64.AFSQRTD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "REV",
		argLen: 1,
		asm:    arm64.AREV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "REVW",
		argLen: 1,
		asm:    arm64.AREVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "REV16W",
		argLen: 1,
		asm:    arm64.AREV16W,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "RBIT",
		argLen: 1,
		asm:    arm64.ARBIT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "RBITW",
		argLen: 1,
		asm:    arm64.ARBITW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "CLZ",
		argLen: 1,
		asm:    arm64.ACLZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "CLZW",
		argLen: 1,
		asm:    arm64.ACLZW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "VCNT",
		argLen: 1,
		asm:    arm64.AVCNT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "VUADDLV",
		argLen: 1,
		asm:    arm64.AVUADDLV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:         "LoweredRound32F",
		argLen:       1,
		resultInArg0: true,
		zeroWidth:    true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:         "LoweredRound64F",
		argLen:       1,
		resultInArg0: true,
		zeroWidth:    true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FMADDS",
		argLen: 3,
		asm:    arm64.AFMADDS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FMADDD",
		argLen: 3,
		asm:    arm64.AFMADDD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FNMADDS",
		argLen: 3,
		asm:    arm64.AFNMADDS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FNMADDD",
		argLen: 3,
		asm:    arm64.AFNMADDD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FMSUBS",
		argLen: 3,
		asm:    arm64.AFMSUBS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FMSUBD",
		argLen: 3,
		asm:    arm64.AFMSUBD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FNMSUBS",
		argLen: 3,
		asm:    arm64.AFNMSUBS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FNMSUBD",
		argLen: 3,
		asm:    arm64.AFNMSUBD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "MADD",
		argLen: 3,
		asm:    arm64.AMADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{2, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MADDW",
		argLen: 3,
		asm:    arm64.AMADDW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{2, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MSUB",
		argLen: 3,
		asm:    arm64.AMSUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{2, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MSUBW",
		argLen: 3,
		asm:    arm64.AMSUBW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{2, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "SLL",
		argLen: 2,
		asm:    arm64.ALSL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "SLLconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.ALSL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "SRL",
		argLen: 2,
		asm:    arm64.ALSR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "SRLconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.ALSR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "SRA",
		argLen: 2,
		asm:    arm64.AASR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "SRAconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.AASR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "ROR",
		argLen: 2,
		asm:    arm64.AROR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "RORW",
		argLen: 2,
		asm:    arm64.ARORW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "RORconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.AROR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "RORWconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.ARORW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "EXTRconst",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AEXTR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "EXTRWconst",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AEXTRW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "CMP",
		argLen: 2,
		asm:    arm64.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:    "CMPconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:   "CMPW",
		argLen: 2,
		asm:    arm64.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:    "CMPWconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm64.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:        "CMN",
		argLen:      2,
		commutative: true,
		asm:         arm64.ACMN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:    "CMNconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.ACMN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:        "CMNW",
		argLen:      2,
		commutative: true,
		asm:         arm64.ACMNW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:    "CMNWconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm64.ACMNW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:        "TST",
		argLen:      2,
		commutative: true,
		asm:         arm64.ATST,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:    "TSTconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.ATST,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:        "TSTW",
		argLen:      2,
		commutative: true,
		asm:         arm64.ATSTW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:    "TSTWconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     arm64.ATSTW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:   "FCMPS",
		argLen: 2,
		asm:    arm64.AFCMPS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FCMPD",
		argLen: 2,
		asm:    arm64.AFCMPD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FCMPS0",
		argLen: 1,
		asm:    arm64.AFCMPS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FCMPD0",
		argLen: 1,
		asm:    arm64.AFCMPD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:    "MVNshiftLL",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.AMVN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "MVNshiftRL",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.AMVN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "MVNshiftRA",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.AMVN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "NEGshiftLL",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.ANEG,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "NEGshiftRL",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.ANEG,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "NEGshiftRA",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.ANEG,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ADDshiftLL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ADDshiftRL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ADDshiftRA",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "SUBshiftLL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "SUBshiftRL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "SUBshiftRA",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ANDshiftLL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ANDshiftRL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ANDshiftRA",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ORshiftLL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AORR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ORshiftRL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AORR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ORshiftRA",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AORR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "XORshiftLL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AEOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "XORshiftRL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AEOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "XORshiftRA",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AEOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "BICshiftLL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.ABIC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "BICshiftRL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.ABIC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "BICshiftRA",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.ABIC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "EONshiftLL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AEON,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "EONshiftRL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AEON,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "EONshiftRA",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AEON,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ORNshiftLL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AORN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ORNshiftRL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AORN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "ORNshiftRA",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.AORN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "CMPshiftLL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:    "CMPshiftRL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:    "CMPshiftRA",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:    "CMNshiftLL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.ACMN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:    "CMNshiftRL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.ACMN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:    "CMNshiftRA",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.ACMN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:    "TSTshiftLL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.ATST,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:    "TSTshiftRL",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.ATST,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:    "TSTshiftRA",
		auxType: auxInt64,
		argLen:  2,
		asm:     arm64.ATST,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:         "BFI",
		auxType:      auxInt64,
		argLen:       2,
		resultInArg0: true,
		asm:          arm64.ABFI,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				{1, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:         "BFXIL",
		auxType:      auxInt64,
		argLen:       2,
		resultInArg0: true,
		asm:          arm64.ABFXIL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				{1, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "SBFIZ",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.ASBFIZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "SBFX",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.ASBFX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "UBFIZ",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.AUBFIZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "UBFX",
		auxType: auxInt64,
		argLen:  1,
		asm:     arm64.AUBFX,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:              "MOVDconst",
		auxType:           auxInt64,
		argLen:            0,
		rematerializeable: true,
		asm:               arm64.AMOVD,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:              "FMOVSconst",
		auxType:           auxFloat64,
		argLen:            0,
		rematerializeable: true,
		asm:               arm64.AFMOVS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:              "FMOVDconst",
		auxType:           auxFloat64,
		argLen:            0,
		rematerializeable: true,
		asm:               arm64.AFMOVD,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:              "MOVDaddr",
		auxType:           auxSymOff,
		argLen:            1,
		rematerializeable: true,
		symEffect:         SymAddr,
		asm:               arm64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372037928517632}, // SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:           "MOVBload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm64.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:           "MOVBUload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm64.AMOVBU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:           "MOVHload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:           "MOVHUload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm64.AMOVHU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:           "MOVWload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:           "MOVWUload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm64.AMOVWU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:           "MOVDload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:           "FMOVSload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm64.AFMOVS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:           "FMOVDload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            arm64.AFMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "MOVDloadidx",
		argLen: 3,
		asm:    arm64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVWloadidx",
		argLen: 3,
		asm:    arm64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVWUloadidx",
		argLen: 3,
		asm:    arm64.AMOVWU,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVHloadidx",
		argLen: 3,
		asm:    arm64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVHUloadidx",
		argLen: 3,
		asm:    arm64.AMOVHU,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVBloadidx",
		argLen: 3,
		asm:    arm64.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVBUloadidx",
		argLen: 3,
		asm:    arm64.AMOVBU,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "FMOVSloadidx",
		argLen: 3,
		asm:    arm64.AFMOVS,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FMOVDloadidx",
		argLen: 3,
		asm:    arm64.AFMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "MOVHloadidx2",
		argLen: 3,
		asm:    arm64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVHUloadidx2",
		argLen: 3,
		asm:    arm64.AMOVHU,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVWloadidx4",
		argLen: 3,
		asm:    arm64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVWUloadidx4",
		argLen: 3,
		asm:    arm64.AMOVWU,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVDloadidx8",
		argLen: 3,
		asm:    arm64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:           "MOVBstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm64.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:           "MOVHstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:           "MOVWstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:           "MOVDstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:           "STP",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm64.ASTP,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:           "FMOVSstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm64.AFMOVS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:           "FMOVDstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm64.AFMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "MOVBstoreidx",
		argLen: 4,
		asm:    arm64.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:   "MOVHstoreidx",
		argLen: 4,
		asm:    arm64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:   "MOVWstoreidx",
		argLen: 4,
		asm:    arm64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:   "MOVDstoreidx",
		argLen: 4,
		asm:    arm64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:   "FMOVSstoreidx",
		argLen: 4,
		asm:    arm64.AFMOVS,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FMOVDstoreidx",
		argLen: 4,
		asm:    arm64.AFMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "MOVHstoreidx2",
		argLen: 4,
		asm:    arm64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:   "MOVWstoreidx4",
		argLen: 4,
		asm:    arm64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:   "MOVDstoreidx8",
		argLen: 4,
		asm:    arm64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:           "MOVBstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm64.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:           "MOVHstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:           "MOVWstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:           "MOVDstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:           "MOVQstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            arm64.ASTP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:   "MOVBstorezeroidx",
		argLen: 3,
		asm:    arm64.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:   "MOVHstorezeroidx",
		argLen: 3,
		asm:    arm64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:   "MOVWstorezeroidx",
		argLen: 3,
		asm:    arm64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:   "MOVDstorezeroidx",
		argLen: 3,
		asm:    arm64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:   "MOVHstorezeroidx2",
		argLen: 3,
		asm:    arm64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:   "MOVWstorezeroidx4",
		argLen: 3,
		asm:    arm64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:   "MOVDstorezeroidx8",
		argLen: 3,
		asm:    arm64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:   "FMOVDgpfp",
		argLen: 1,
		asm:    arm64.AFMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FMOVDfpgp",
		argLen: 1,
		asm:    arm64.AFMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "FMOVSgpfp",
		argLen: 1,
		asm:    arm64.AFMOVS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FMOVSfpgp",
		argLen: 1,
		asm:    arm64.AFMOVS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVBreg",
		argLen: 1,
		asm:    arm64.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVBUreg",
		argLen: 1,
		asm:    arm64.AMOVBU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVHreg",
		argLen: 1,
		asm:    arm64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVHUreg",
		argLen: 1,
		asm:    arm64.AMOVHU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVWreg",
		argLen: 1,
		asm:    arm64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVWUreg",
		argLen: 1,
		asm:    arm64.AMOVWU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "MOVDreg",
		argLen: 1,
		asm:    arm64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:         "MOVDnop",
		argLen:       1,
		resultInArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "SCVTFWS",
		argLen: 1,
		asm:    arm64.ASCVTFWS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "SCVTFWD",
		argLen: 1,
		asm:    arm64.ASCVTFWD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "UCVTFWS",
		argLen: 1,
		asm:    arm64.AUCVTFWS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "UCVTFWD",
		argLen: 1,
		asm:    arm64.AUCVTFWD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "SCVTFS",
		argLen: 1,
		asm:    arm64.ASCVTFS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "SCVTFD",
		argLen: 1,
		asm:    arm64.ASCVTFD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "UCVTFS",
		argLen: 1,
		asm:    arm64.AUCVTFS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "UCVTFD",
		argLen: 1,
		asm:    arm64.AUCVTFD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FCVTZSSW",
		argLen: 1,
		asm:    arm64.AFCVTZSSW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "FCVTZSDW",
		argLen: 1,
		asm:    arm64.AFCVTZSDW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "FCVTZUSW",
		argLen: 1,
		asm:    arm64.AFCVTZUSW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "FCVTZUDW",
		argLen: 1,
		asm:    arm64.AFCVTZUDW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "FCVTZSS",
		argLen: 1,
		asm:    arm64.AFCVTZSS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "FCVTZSD",
		argLen: 1,
		asm:    arm64.AFCVTZSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "FCVTZUS",
		argLen: 1,
		asm:    arm64.AFCVTZUS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "FCVTZUD",
		argLen: 1,
		asm:    arm64.AFCVTZUD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "FCVTSD",
		argLen: 1,
		asm:    arm64.AFCVTSD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FCVTDS",
		argLen: 1,
		asm:    arm64.AFCVTDS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FRINTAD",
		argLen: 1,
		asm:    arm64.AFRINTAD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FRINTMD",
		argLen: 1,
		asm:    arm64.AFRINTMD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FRINTND",
		argLen: 1,
		asm:    arm64.AFRINTND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FRINTPD",
		argLen: 1,
		asm:    arm64.AFRINTPD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "FRINTZD",
		argLen: 1,
		asm:    arm64.AFRINTZD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:    "CSEL",
		auxType: auxCCop,
		argLen:  3,
		asm:     arm64.ACSEL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				{1, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:    "CSEL0",
		auxType: auxCCop,
		argLen:  2,
		asm:     arm64.ACSEL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:         "CALLstatic",
		auxType:      auxSymOff,
		argLen:       1,
		clobberFlags: true,
		call:         true,
		symEffect:    SymNone,
		reg: regInfo{
			clobbers: 9223372035512336383, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
		},
	},
	{
		name:         "CALLclosure",
		auxType:      auxInt64,
		argLen:       3,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 67108864},   // R26
				{0, 1744568319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30 SP
			},
			clobbers: 9223372035512336383, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
		},
	},
	{
		name:         "CALLinter",
		auxType:      auxInt64,
		argLen:       2,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			clobbers: 9223372035512336383, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
		},
	},
	{
		name:           "LoweredNilCheck",
		argLen:         2,
		nilCheck:       true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
			},
		},
	},
	{
		name:   "Equal",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "NotEqual",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "LessThan",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "LessEqual",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "GreaterThan",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "GreaterEqual",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "LessThanU",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "LessEqualU",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "GreaterThanU",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "GreaterEqualU",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "LessThanF",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "LessEqualF",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "GreaterThanF",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "GreaterEqualF",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:           "DUFFZERO",
		auxType:        auxInt64,
		argLen:         2,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65536}, // R16
			},
			clobbers: 536936448, // R16 R30
		},
	},
	{
		name:           "LoweredZero",
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65536},     // R16
				{1, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			clobbers: 65536, // R16
		},
	},
	{
		name:           "DUFFCOPY",
		auxType:        auxInt64,
		argLen:         3,
		faultOnNilArg0: true,
		faultOnNilArg1: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 131072}, // R17
				{1, 65536},  // R16
			},
			clobbers: 604176384, // R16 R17 R26 R30
		},
	},
	{
		name:           "LoweredMove",
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		faultOnNilArg1: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 131072},    // R17
				{1, 65536},     // R16
				{2, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
			clobbers: 196608, // R16 R17
		},
	},
	{
		name:      "LoweredGetClosurePtr",
		argLen:    0,
		zeroWidth: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 67108864}, // R26
			},
		},
	},
	{
		name:              "LoweredGetCallerSP",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:              "LoweredGetCallerPC",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:   "FlagEQ",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagLT_ULT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagLT_UGT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagGT_UGT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagGT_ULT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "InvertFlags",
		argLen: 1,
		reg:    regInfo{},
	},
	{
		name:           "LDAR",
		argLen:         2,
		faultOnNilArg0: true,
		asm:            arm64.ALDAR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:           "LDARB",
		argLen:         2,
		faultOnNilArg0: true,
		asm:            arm64.ALDARB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:           "LDARW",
		argLen:         2,
		faultOnNilArg0: true,
		asm:            arm64.ALDARW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:           "STLR",
		argLen:         3,
		faultOnNilArg0: true,
		hasSideEffects: true,
		asm:            arm64.ASTLR,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:           "STLRW",
		argLen:         3,
		faultOnNilArg0: true,
		hasSideEffects: true,
		asm:            arm64.ASTLRW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
		},
	},
	{
		name:            "LoweredAtomicExchange64",
		argLen:          3,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:            "LoweredAtomicExchange32",
		argLen:          3,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:            "LoweredAtomicAdd64",
		argLen:          3,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:            "LoweredAtomicAdd32",
		argLen:          3,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:            "LoweredAtomicAdd64Variant",
		argLen:          3,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:            "LoweredAtomicAdd32Variant",
		argLen:          3,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:            "LoweredAtomicCas64",
		argLen:          4,
		resultNotInArgs: true,
		clobberFlags:    true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:            "LoweredAtomicCas32",
		argLen:          4,
		resultNotInArgs: true,
		clobberFlags:    true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:            "LoweredAtomicAnd8",
		argLen:          3,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		asm:             arm64.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:            "LoweredAtomicOr8",
		argLen:          3,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		asm:             arm64.AORR,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
			},
			outputs: []outputInfo{
				{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
			},
		},
	},
	{
		name:         "LoweredWB",
		auxType:      auxSym,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymNone,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4}, // R2
				{1, 8}, // R3
			},
			clobbers: 9223372035244163072, // R30 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
		},
	},
	{
		name:    "LoweredPanicBoundsA",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4}, // R2
				{1, 8}, // R3
			},
		},
	},
	{
		name:    "LoweredPanicBoundsB",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2}, // R1
				{1, 4}, // R2
			},
		},
	},
	{
		name:    "LoweredPanicBoundsC",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1}, // R0
				{1, 2}, // R1
			},
		},
	},

	{
		name:        "ADD",
		argLen:      2,
		commutative: true,
		asm:         mips.AADDU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:    "ADDconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     mips.AADDU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 536870910}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "SUB",
		argLen: 2,
		asm:    mips.ASUBU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:    "SUBconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     mips.ASUBU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:        "MUL",
		argLen:      2,
		commutative: true,
		asm:         mips.AMUL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			clobbers: 105553116266496, // HI LO
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:        "MULT",
		argLen:      2,
		commutative: true,
		asm:         mips.AMUL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 35184372088832}, // HI
				{1, 70368744177664}, // LO
			},
		},
	},
	{
		name:        "MULTU",
		argLen:      2,
		commutative: true,
		asm:         mips.AMULU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 35184372088832}, // HI
				{1, 70368744177664}, // LO
			},
		},
	},
	{
		name:   "DIV",
		argLen: 2,
		asm:    mips.ADIV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 35184372088832}, // HI
				{1, 70368744177664}, // LO
			},
		},
	},
	{
		name:   "DIVU",
		argLen: 2,
		asm:    mips.ADIVU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 35184372088832}, // HI
				{1, 70368744177664}, // LO
			},
		},
	},
	{
		name:        "ADDF",
		argLen:      2,
		commutative: true,
		asm:         mips.AADDF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:        "ADDD",
		argLen:      2,
		commutative: true,
		asm:         mips.AADDD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "SUBF",
		argLen: 2,
		asm:    mips.ASUBF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "SUBD",
		argLen: 2,
		asm:    mips.ASUBD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:        "MULF",
		argLen:      2,
		commutative: true,
		asm:         mips.AMULF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:        "MULD",
		argLen:      2,
		commutative: true,
		asm:         mips.AMULD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "DIVF",
		argLen: 2,
		asm:    mips.ADIVF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "DIVD",
		argLen: 2,
		asm:    mips.ADIVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:        "AND",
		argLen:      2,
		commutative: true,
		asm:         mips.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:    "ANDconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     mips.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:        "OR",
		argLen:      2,
		commutative: true,
		asm:         mips.AOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:    "ORconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     mips.AOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:        "XOR",
		argLen:      2,
		commutative: true,
		asm:         mips.AXOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:    "XORconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     mips.AXOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:        "NOR",
		argLen:      2,
		commutative: true,
		asm:         mips.ANOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:    "NORconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     mips.ANOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "NEG",
		argLen: 1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "NEGF",
		argLen: 1,
		asm:    mips.ANEGF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "NEGD",
		argLen: 1,
		asm:    mips.ANEGD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "SQRTD",
		argLen: 1,
		asm:    mips.ASQRTD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "SLL",
		argLen: 2,
		asm:    mips.ASLL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:    "SLLconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     mips.ASLL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "SRL",
		argLen: 2,
		asm:    mips.ASRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:    "SRLconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     mips.ASRL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "SRA",
		argLen: 2,
		asm:    mips.ASRA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:    "SRAconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     mips.ASRA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "CLZ",
		argLen: 1,
		asm:    mips.ACLZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "SGT",
		argLen: 2,
		asm:    mips.ASGT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:    "SGTconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     mips.ASGT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "SGTzero",
		argLen: 1,
		asm:    mips.ASGT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "SGTU",
		argLen: 2,
		asm:    mips.ASGTU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:    "SGTUconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     mips.ASGTU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "SGTUzero",
		argLen: 1,
		asm:    mips.ASGTU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "CMPEQF",
		argLen: 2,
		asm:    mips.ACMPEQF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "CMPEQD",
		argLen: 2,
		asm:    mips.ACMPEQD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "CMPGEF",
		argLen: 2,
		asm:    mips.ACMPGEF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "CMPGED",
		argLen: 2,
		asm:    mips.ACMPGED,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "CMPGTF",
		argLen: 2,
		asm:    mips.ACMPGTF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "CMPGTD",
		argLen: 2,
		asm:    mips.ACMPGTD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:              "MOVWconst",
		auxType:           auxInt32,
		argLen:            0,
		rematerializeable: true,
		asm:               mips.AMOVW,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:              "MOVFconst",
		auxType:           auxFloat32,
		argLen:            0,
		rematerializeable: true,
		asm:               mips.AMOVF,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:              "MOVDconst",
		auxType:           auxFloat64,
		argLen:            0,
		rematerializeable: true,
		asm:               mips.AMOVD,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:              "MOVWaddr",
		auxType:           auxSymOff,
		argLen:            1,
		rematerializeable: true,
		symEffect:         SymAddr,
		asm:               mips.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 140737555464192}, // SP SB
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:           "MOVBload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:           "MOVBUload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVBU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:           "MOVHload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:           "MOVHUload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVHU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:           "MOVWload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:           "MOVFload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:           "MOVDload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:           "MOVBstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
		},
	},
	{
		name:           "MOVHstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
		},
	},
	{
		name:           "MOVWstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
		},
	},
	{
		name:           "MOVFstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVF,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 35183835217920},  // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
		},
	},
	{
		name:           "MOVDstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 35183835217920},  // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
		},
	},
	{
		name:           "MOVBstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
		},
	},
	{
		name:           "MOVHstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
		},
	},
	{
		name:           "MOVWstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
		},
	},
	{
		name:   "MOVBreg",
		argLen: 1,
		asm:    mips.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "MOVBUreg",
		argLen: 1,
		asm:    mips.AMOVBU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "MOVHreg",
		argLen: 1,
		asm:    mips.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "MOVHUreg",
		argLen: 1,
		asm:    mips.AMOVHU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "MOVWreg",
		argLen: 1,
		asm:    mips.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:         "MOVWnop",
		argLen:       1,
		resultInArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:         "CMOVZ",
		argLen:       3,
		resultInArg0: true,
		asm:          mips.ACMOVZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				{1, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				{2, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:         "CMOVZzero",
		argLen:       2,
		resultInArg0: true,
		asm:          mips.ACMOVZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "MOVWF",
		argLen: 1,
		asm:    mips.AMOVWF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "MOVWD",
		argLen: 1,
		asm:    mips.AMOVWD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "TRUNCFW",
		argLen: 1,
		asm:    mips.ATRUNCFW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "TRUNCDW",
		argLen: 1,
		asm:    mips.ATRUNCDW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "MOVFD",
		argLen: 1,
		asm:    mips.AMOVFD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:   "MOVDF",
		argLen: 1,
		asm:    mips.AMOVDF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
			outputs: []outputInfo{
				{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
			},
		},
	},
	{
		name:         "CALLstatic",
		auxType:      auxSymOff,
		argLen:       1,
		clobberFlags: true,
		call:         true,
		symEffect:    SymNone,
		reg: regInfo{
			clobbers: 140737421246462, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31 F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30 HI LO
		},
	},
	{
		name:         "CALLclosure",
		auxType:      auxInt64,
		argLen:       3,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 4194304},   // R22
				{0, 402653182}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP R31
			},
			clobbers: 140737421246462, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31 F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30 HI LO
		},
	},
	{
		name:         "CALLinter",
		auxType:      auxInt64,
		argLen:       2,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
			clobbers: 140737421246462, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31 F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30 HI LO
		},
	},
	{
		name:           "LoweredAtomicLoad",
		argLen:         2,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:           "LoweredAtomicStore",
		argLen:         3,
		faultOnNilArg0: true,
		hasSideEffects: true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
		},
	},
	{
		name:           "LoweredAtomicStorezero",
		argLen:         2,
		faultOnNilArg0: true,
		hasSideEffects: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
		},
	},
	{
		name:            "LoweredAtomicExchange",
		argLen:          3,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:            "LoweredAtomicAdd",
		argLen:          3,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:            "LoweredAtomicAddconst",
		auxType:         auxInt32,
		argLen:          2,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:            "LoweredAtomicCas",
		argLen:          4,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{2, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:           "LoweredAtomicAnd",
		argLen:         3,
		faultOnNilArg0: true,
		hasSideEffects: true,
		asm:            mips.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
		},
	},
	{
		name:           "LoweredAtomicOr",
		argLen:         3,
		faultOnNilArg0: true,
		hasSideEffects: true,
		asm:            mips.AOR,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
			},
		},
	},
	{
		name:           "LoweredZero",
		auxType:        auxInt32,
		argLen:         3,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2},         // R1
				{1, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
			clobbers: 2, // R1
		},
	},
	{
		name:           "LoweredMove",
		auxType:        auxInt32,
		argLen:         4,
		faultOnNilArg0: true,
		faultOnNilArg1: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4},         // R2
				{1, 2},         // R1
				{2, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
			clobbers: 6, // R1 R2
		},
	},
	{
		name:           "LoweredNilCheck",
		argLen:         2,
		nilCheck:       true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
			},
		},
	},
	{
		name:   "FPFlagTrue",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:   "FPFlagFalse",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:      "LoweredGetClosurePtr",
		argLen:    0,
		zeroWidth: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 4194304}, // R22
			},
		},
	},
	{
		name:              "LoweredGetCallerSP",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:              "LoweredGetCallerPC",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
			},
		},
	},
	{
		name:         "LoweredWB",
		auxType:      auxSym,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymNone,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1048576}, // R20
				{1, 2097152}, // R21
			},
			clobbers: 140737219919872, // R31 F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30 HI LO
		},
	},
	{
		name:    "LoweredPanicBoundsA",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 8},  // R3
				{1, 16}, // R4
			},
		},
	},
	{
		name:    "LoweredPanicBoundsB",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4}, // R2
				{1, 8}, // R3
			},
		},
	},
	{
		name:    "LoweredPanicBoundsC",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2}, // R1
				{1, 4}, // R2
			},
		},
	},
	{
		name:    "LoweredPanicExtendA",
		auxType: auxInt64,
		argLen:  4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 32}, // R5
				{1, 8},  // R3
				{2, 16}, // R4
			},
		},
	},
	{
		name:    "LoweredPanicExtendB",
		auxType: auxInt64,
		argLen:  4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 32}, // R5
				{1, 4},  // R2
				{2, 8},  // R3
			},
		},
	},
	{
		name:    "LoweredPanicExtendC",
		auxType: auxInt64,
		argLen:  4,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 32}, // R5
				{1, 2},  // R1
				{2, 4},  // R2
			},
		},
	},

	{
		name:        "ADDV",
		argLen:      2,
		commutative: true,
		asm:         mips.AADDVU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:    "ADDVconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     mips.AADDVU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 268435454}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "SUBV",
		argLen: 2,
		asm:    mips.ASUBVU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:    "SUBVconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     mips.ASUBVU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:        "MULV",
		argLen:      2,
		commutative: true,
		asm:         mips.AMULV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 1152921504606846976}, // HI
				{1, 2305843009213693952}, // LO
			},
		},
	},
	{
		name:        "MULVU",
		argLen:      2,
		commutative: true,
		asm:         mips.AMULVU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 1152921504606846976}, // HI
				{1, 2305843009213693952}, // LO
			},
		},
	},
	{
		name:   "DIVV",
		argLen: 2,
		asm:    mips.ADIVV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 1152921504606846976}, // HI
				{1, 2305843009213693952}, // LO
			},
		},
	},
	{
		name:   "DIVVU",
		argLen: 2,
		asm:    mips.ADIVVU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 1152921504606846976}, // HI
				{1, 2305843009213693952}, // LO
			},
		},
	},
	{
		name:        "ADDF",
		argLen:      2,
		commutative: true,
		asm:         mips.AADDF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:        "ADDD",
		argLen:      2,
		commutative: true,
		asm:         mips.AADDD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "SUBF",
		argLen: 2,
		asm:    mips.ASUBF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "SUBD",
		argLen: 2,
		asm:    mips.ASUBD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:        "MULF",
		argLen:      2,
		commutative: true,
		asm:         mips.AMULF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:        "MULD",
		argLen:      2,
		commutative: true,
		asm:         mips.AMULD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "DIVF",
		argLen: 2,
		asm:    mips.ADIVF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "DIVD",
		argLen: 2,
		asm:    mips.ADIVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:        "AND",
		argLen:      2,
		commutative: true,
		asm:         mips.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:    "ANDconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     mips.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:        "OR",
		argLen:      2,
		commutative: true,
		asm:         mips.AOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:    "ORconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     mips.AOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:        "XOR",
		argLen:      2,
		commutative: true,
		asm:         mips.AXOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:    "XORconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     mips.AXOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:        "NOR",
		argLen:      2,
		commutative: true,
		asm:         mips.ANOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:    "NORconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     mips.ANOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "NEGV",
		argLen: 1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "NEGF",
		argLen: 1,
		asm:    mips.ANEGF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "NEGD",
		argLen: 1,
		asm:    mips.ANEGD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "SQRTD",
		argLen: 1,
		asm:    mips.ASQRTD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "SLLV",
		argLen: 2,
		asm:    mips.ASLLV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:    "SLLVconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     mips.ASLLV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "SRLV",
		argLen: 2,
		asm:    mips.ASRLV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:    "SRLVconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     mips.ASRLV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "SRAV",
		argLen: 2,
		asm:    mips.ASRAV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:    "SRAVconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     mips.ASRAV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "SGT",
		argLen: 2,
		asm:    mips.ASGT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:    "SGTconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     mips.ASGT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "SGTU",
		argLen: 2,
		asm:    mips.ASGTU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:    "SGTUconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     mips.ASGTU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "CMPEQF",
		argLen: 2,
		asm:    mips.ACMPEQF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "CMPEQD",
		argLen: 2,
		asm:    mips.ACMPEQD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "CMPGEF",
		argLen: 2,
		asm:    mips.ACMPGEF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "CMPGED",
		argLen: 2,
		asm:    mips.ACMPGED,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "CMPGTF",
		argLen: 2,
		asm:    mips.ACMPGTF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "CMPGTD",
		argLen: 2,
		asm:    mips.ACMPGTD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:              "MOVVconst",
		auxType:           auxInt64,
		argLen:            0,
		rematerializeable: true,
		asm:               mips.AMOVV,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:              "MOVFconst",
		auxType:           auxFloat64,
		argLen:            0,
		rematerializeable: true,
		asm:               mips.AMOVF,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:              "MOVDconst",
		auxType:           auxFloat64,
		argLen:            0,
		rematerializeable: true,
		asm:               mips.AMOVD,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:              "MOVVaddr",
		auxType:           auxSymOff,
		argLen:            1,
		rematerializeable: true,
		symEffect:         SymAddr,
		asm:               mips.AMOVV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018460942336}, // SP SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:           "MOVBload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:           "MOVBUload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVBU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:           "MOVHload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:           "MOVHUload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVHU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:           "MOVWload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:           "MOVWUload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVWU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:           "MOVVload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:           "MOVFload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:           "MOVDload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            mips.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:           "MOVBstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
		},
	},
	{
		name:           "MOVHstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
		},
	},
	{
		name:           "MOVWstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
		},
	},
	{
		name:           "MOVVstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVV,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
		},
	},
	{
		name:           "MOVFstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:           "MOVDstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:           "MOVBstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
		},
	},
	{
		name:           "MOVHstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
		},
	},
	{
		name:           "MOVWstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
		},
	},
	{
		name:           "MOVVstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            mips.AMOVV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
		},
	},
	{
		name:   "MOVBreg",
		argLen: 1,
		asm:    mips.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "MOVBUreg",
		argLen: 1,
		asm:    mips.AMOVBU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "MOVHreg",
		argLen: 1,
		asm:    mips.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "MOVHUreg",
		argLen: 1,
		asm:    mips.AMOVHU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "MOVWreg",
		argLen: 1,
		asm:    mips.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "MOVWUreg",
		argLen: 1,
		asm:    mips.AMOVWU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "MOVVreg",
		argLen: 1,
		asm:    mips.AMOVV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:         "MOVVnop",
		argLen:       1,
		resultInArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "MOVWF",
		argLen: 1,
		asm:    mips.AMOVWF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "MOVWD",
		argLen: 1,
		asm:    mips.AMOVWD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "MOVVF",
		argLen: 1,
		asm:    mips.AMOVVF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "MOVVD",
		argLen: 1,
		asm:    mips.AMOVVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "TRUNCFW",
		argLen: 1,
		asm:    mips.ATRUNCFW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "TRUNCDW",
		argLen: 1,
		asm:    mips.ATRUNCDW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "TRUNCFV",
		argLen: 1,
		asm:    mips.ATRUNCFV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "TRUNCDV",
		argLen: 1,
		asm:    mips.ATRUNCDV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "MOVFD",
		argLen: 1,
		asm:    mips.AMOVFD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:   "MOVDF",
		argLen: 1,
		asm:    mips.AMOVDF,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
			outputs: []outputInfo{
				{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
	},
	{
		name:         "CALLstatic",
		auxType:      auxSymOff,
		argLen:       1,
		clobberFlags: true,
		call:         true,
		symEffect:    SymNone,
		reg: regInfo{
			clobbers: 4611686018393833470, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31 HI LO
		},
	},
	{
		name:         "CALLclosure",
		auxType:      auxInt64,
		argLen:       3,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 4194304},   // R22
				{0, 201326590}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP R31
			},
			clobbers: 4611686018393833470, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31 HI LO
		},
	},
	{
		name:         "CALLinter",
		auxType:      auxInt64,
		argLen:       2,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
			clobbers: 4611686018393833470, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31 HI LO
		},
	},
	{
		name:           "DUFFZERO",
		auxType:        auxInt64,
		argLen:         2,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
			clobbers: 134217730, // R1 R31
		},
	},
	{
		name:           "LoweredZero",
		auxType:        auxInt64,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2},         // R1
				{1, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
			clobbers: 2, // R1
		},
	},
	{
		name:           "LoweredMove",
		auxType:        auxInt64,
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		faultOnNilArg1: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4},         // R2
				{1, 2},         // R1
				{2, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
			clobbers: 6, // R1 R2
		},
	},
	{
		name:           "LoweredAtomicLoad8",
		argLen:         2,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:           "LoweredAtomicLoad32",
		argLen:         2,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:           "LoweredAtomicLoad64",
		argLen:         2,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:           "LoweredAtomicStore32",
		argLen:         3,
		faultOnNilArg0: true,
		hasSideEffects: true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
		},
	},
	{
		name:           "LoweredAtomicStore64",
		argLen:         3,
		faultOnNilArg0: true,
		hasSideEffects: true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
		},
	},
	{
		name:           "LoweredAtomicStorezero32",
		argLen:         2,
		faultOnNilArg0: true,
		hasSideEffects: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
		},
	},
	{
		name:           "LoweredAtomicStorezero64",
		argLen:         2,
		faultOnNilArg0: true,
		hasSideEffects: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
		},
	},
	{
		name:            "LoweredAtomicExchange32",
		argLen:          3,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:            "LoweredAtomicExchange64",
		argLen:          3,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:            "LoweredAtomicAdd32",
		argLen:          3,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:            "LoweredAtomicAdd64",
		argLen:          3,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:            "LoweredAtomicAddconst32",
		auxType:         auxInt32,
		argLen:          2,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:            "LoweredAtomicAddconst64",
		auxType:         auxInt64,
		argLen:          2,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:            "LoweredAtomicCas32",
		argLen:          4,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{2, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:            "LoweredAtomicCas64",
		argLen:          4,
		resultNotInArgs: true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{2, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
			},
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:           "LoweredNilCheck",
		argLen:         2,
		nilCheck:       true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
			},
		},
	},
	{
		name:   "FPFlagTrue",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:   "FPFlagFalse",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:      "LoweredGetClosurePtr",
		argLen:    0,
		zeroWidth: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 4194304}, // R22
			},
		},
	},
	{
		name:              "LoweredGetCallerSP",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:              "LoweredGetCallerPC",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
			},
		},
	},
	{
		name:         "LoweredWB",
		auxType:      auxSym,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymNone,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1048576}, // R20
				{1, 2097152}, // R21
			},
			clobbers: 4611686018293170176, // R31 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31 HI LO
		},
	},
	{
		name:    "LoweredPanicBoundsA",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 8},  // R3
				{1, 16}, // R4
			},
		},
	},
	{
		name:    "LoweredPanicBoundsB",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4}, // R2
				{1, 8}, // R3
			},
		},
	},
	{
		name:    "LoweredPanicBoundsC",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2}, // R1
				{1, 4}, // R2
			},
		},
	},

	{
		name:        "ADD",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "ADDconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     ppc64.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:        "FADD",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AFADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:        "FADDS",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AFADDS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "SUB",
		argLen: 2,
		asm:    ppc64.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "FSUB",
		argLen: 2,
		asm:    ppc64.AFSUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FSUBS",
		argLen: 2,
		asm:    ppc64.AFSUBS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:        "MULLD",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AMULLD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:        "MULLW",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AMULLW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:        "MULHD",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AMULHD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:        "MULHW",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AMULHW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:        "MULHDU",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AMULHDU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:        "MULHWU",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AMULHWU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:            "LoweredMuluhilo",
		argLen:          2,
		resultNotInArgs: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:        "FMUL",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AFMUL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:        "FMULS",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AFMULS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FMADD",
		argLen: 3,
		asm:    ppc64.AFMADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{2, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FMADDS",
		argLen: 3,
		asm:    ppc64.AFMADDS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{2, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FMSUB",
		argLen: 3,
		asm:    ppc64.AFMSUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{2, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FMSUBS",
		argLen: 3,
		asm:    ppc64.AFMSUBS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{2, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "SRAD",
		argLen: 2,
		asm:    ppc64.ASRAD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "SRAW",
		argLen: 2,
		asm:    ppc64.ASRAW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "SRD",
		argLen: 2,
		asm:    ppc64.ASRD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "SRW",
		argLen: 2,
		asm:    ppc64.ASRW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "SLD",
		argLen: 2,
		asm:    ppc64.ASLD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "SLW",
		argLen: 2,
		asm:    ppc64.ASLW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "ROTL",
		argLen: 2,
		asm:    ppc64.AROTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "ROTLW",
		argLen: 2,
		asm:    ppc64.AROTLW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:            "LoweredAdd64Carry",
		argLen:          3,
		resultNotInArgs: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{2, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "ADDconstForCarry",
		auxType: auxInt16,
		argLen:  1,
		asm:     ppc64.AADDC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			clobbers: 2147483648, // R31
		},
	},
	{
		name:   "MaskIfNotCarry",
		argLen: 1,
		asm:    ppc64.AADDME,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "SRADconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     ppc64.ASRAD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "SRAWconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     ppc64.ASRAW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "SRDconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     ppc64.ASRD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "SRWconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     ppc64.ASRW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "SLDconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     ppc64.ASLD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "SLWconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     ppc64.ASLW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "ROTLconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     ppc64.AROTL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "ROTLWconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     ppc64.AROTLW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:         "CNTLZD",
		argLen:       1,
		clobberFlags: true,
		asm:          ppc64.ACNTLZD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:         "CNTLZW",
		argLen:       1,
		clobberFlags: true,
		asm:          ppc64.ACNTLZW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "CNTTZD",
		argLen: 1,
		asm:    ppc64.ACNTTZD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "CNTTZW",
		argLen: 1,
		asm:    ppc64.ACNTTZW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "POPCNTD",
		argLen: 1,
		asm:    ppc64.APOPCNTD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "POPCNTW",
		argLen: 1,
		asm:    ppc64.APOPCNTW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "POPCNTB",
		argLen: 1,
		asm:    ppc64.APOPCNTB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "FDIV",
		argLen: 2,
		asm:    ppc64.AFDIV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FDIVS",
		argLen: 2,
		asm:    ppc64.AFDIVS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "DIVD",
		argLen: 2,
		asm:    ppc64.ADIVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "DIVW",
		argLen: 2,
		asm:    ppc64.ADIVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "DIVDU",
		argLen: 2,
		asm:    ppc64.ADIVDU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "DIVWU",
		argLen: 2,
		asm:    ppc64.ADIVWU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "FCTIDZ",
		argLen: 1,
		asm:    ppc64.AFCTIDZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FCTIWZ",
		argLen: 1,
		asm:    ppc64.AFCTIWZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FCFID",
		argLen: 1,
		asm:    ppc64.AFCFID,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FCFIDS",
		argLen: 1,
		asm:    ppc64.AFCFIDS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FRSP",
		argLen: 1,
		asm:    ppc64.AFRSP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "MFVSRD",
		argLen: 1,
		asm:    ppc64.AMFVSRD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "MTVSRD",
		argLen: 1,
		asm:    ppc64.AMTVSRD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:        "AND",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "ANDN",
		argLen: 2,
		asm:    ppc64.AANDN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:        "ANDCC",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AANDCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:        "OR",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "ORN",
		argLen: 2,
		asm:    ppc64.AORN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:        "ORCC",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AORCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:        "NOR",
		argLen:      2,
		commutative: true,
		asm:         ppc64.ANOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:        "XOR",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AXOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:        "XORCC",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AXORCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:        "EQV",
		argLen:      2,
		commutative: true,
		asm:         ppc64.AEQV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "NEG",
		argLen: 1,
		asm:    ppc64.ANEG,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "FNEG",
		argLen: 1,
		asm:    ppc64.AFNEG,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FSQRT",
		argLen: 1,
		asm:    ppc64.AFSQRT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FSQRTS",
		argLen: 1,
		asm:    ppc64.AFSQRTS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FFLOOR",
		argLen: 1,
		asm:    ppc64.AFRIM,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FCEIL",
		argLen: 1,
		asm:    ppc64.AFRIP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FTRUNC",
		argLen: 1,
		asm:    ppc64.AFRIZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FROUND",
		argLen: 1,
		asm:    ppc64.AFRIN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FABS",
		argLen: 1,
		asm:    ppc64.AFABS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FNABS",
		argLen: 1,
		asm:    ppc64.AFNABS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FCPSGN",
		argLen: 2,
		asm:    ppc64.AFCPSGN,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:    "ORconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     ppc64.AOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "XORconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     ppc64.AXOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:         "ANDconst",
		auxType:      auxInt64,
		argLen:       1,
		clobberFlags: true,
		asm:          ppc64.AANDCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "ANDCCconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     ppc64.AANDCC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "MOVBreg",
		argLen: 1,
		asm:    ppc64.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "MOVBZreg",
		argLen: 1,
		asm:    ppc64.AMOVBZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "MOVHreg",
		argLen: 1,
		asm:    ppc64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "MOVHZreg",
		argLen: 1,
		asm:    ppc64.AMOVHZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "MOVWreg",
		argLen: 1,
		asm:    ppc64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "MOVWZreg",
		argLen: 1,
		asm:    ppc64.AMOVWZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVBZload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVBZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVHload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVHZload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVHZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVWload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVWZload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVWZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVDload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVDBRload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVDBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVWBRload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVWBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVHBRload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVHBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVBZloadidx",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVBZ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVHloadidx",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVHZloadidx",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVHZ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVWloadidx",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVWZloadidx",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVWZ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVDloadidx",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVHBRloadidx",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVHBR,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVWBRloadidx",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVWBR,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVDBRloadidx",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AMOVDBR,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "FMOVDloadidx",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AFMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:           "FMOVSloadidx",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AFMOVS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:           "MOVDBRstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVDBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVWBRstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVWBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVHBRstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVHBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "FMOVDload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AFMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:           "FMOVSload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            ppc64.AFMOVS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:           "MOVBstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVHstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVWstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVDstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "FMOVDstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AFMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{0, 1073733630},         // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "FMOVSstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AFMOVS,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{0, 1073733630},         // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVBstoreidx",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{2, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVHstoreidx",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{2, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVWstoreidx",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{2, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVDstoreidx",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{2, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "FMOVDstoreidx",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AFMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{0, 1073733630},         // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630},         // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "FMOVSstoreidx",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AFMOVS,
		reg: regInfo{
			inputs: []inputInfo{
				{2, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{0, 1073733630},         // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630},         // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVHBRstoreidx",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVHBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{2, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVWBRstoreidx",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVWBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{2, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVDBRstoreidx",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVDBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{2, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVBstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVHstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVWstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "MOVDstorezero",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            ppc64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:              "MOVDaddr",
		auxType:           auxSymOff,
		argLen:            1,
		rematerializeable: true,
		symEffect:         SymAddr,
		asm:               ppc64.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:              "MOVDconst",
		auxType:           auxInt64,
		argLen:            0,
		rematerializeable: true,
		asm:               ppc64.AMOVD,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:              "FMOVDconst",
		auxType:           auxFloat64,
		argLen:            0,
		rematerializeable: true,
		asm:               ppc64.AFMOVD,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:              "FMOVSconst",
		auxType:           auxFloat32,
		argLen:            0,
		rematerializeable: true,
		asm:               ppc64.AFMOVS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "FCMPU",
		argLen: 2,
		asm:    ppc64.AFCMPU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:   "CMP",
		argLen: 2,
		asm:    ppc64.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "CMPU",
		argLen: 2,
		asm:    ppc64.ACMPU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "CMPW",
		argLen: 2,
		asm:    ppc64.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "CMPWU",
		argLen: 2,
		asm:    ppc64.ACMPWU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "CMPconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     ppc64.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "CMPUconst",
		auxType: auxInt64,
		argLen:  1,
		asm:     ppc64.ACMPU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "CMPWconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     ppc64.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:    "CMPWUconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     ppc64.ACMPWU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "Equal",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "NotEqual",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "LessThan",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "FLessThan",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "LessEqual",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "FLessEqual",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "GreaterThan",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "FGreaterThan",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "GreaterEqual",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:   "FGreaterEqual",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:      "LoweredGetClosurePtr",
		argLen:    0,
		zeroWidth: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 2048}, // R11
			},
		},
	},
	{
		name:              "LoweredGetCallerSP",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:              "LoweredGetCallerPC",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "LoweredNilCheck",
		argLen:         2,
		clobberFlags:   true,
		nilCheck:       true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			clobbers: 2147483648, // R31
		},
	},
	{
		name:         "LoweredRound32F",
		argLen:       1,
		resultInArg0: true,
		zeroWidth:    true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:         "LoweredRound64F",
		argLen:       1,
		resultInArg0: true,
		zeroWidth:    true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
			outputs: []outputInfo{
				{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
	},
	{
		name:         "CALLstatic",
		auxType:      auxSymOff,
		argLen:       1,
		clobberFlags: true,
		call:         true,
		symEffect:    SymNone,
		reg: regInfo{
			clobbers: 576460745860964344, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29 g F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
		},
	},
	{
		name:         "CALLclosure",
		auxType:      auxInt64,
		argLen:       3,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4096}, // R12
				{1, 2048}, // R11
			},
			clobbers: 576460745860964344, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29 g F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
		},
	},
	{
		name:         "CALLinter",
		auxType:      auxInt64,
		argLen:       2,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4096}, // R12
			},
			clobbers: 576460745860964344, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29 g F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
		},
	},
	{
		name:           "LoweredZero",
		auxType:        auxInt64,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 8}, // R3
			},
			clobbers: 8, // R3
		},
	},
	{
		name:           "LoweredMove",
		auxType:        auxInt64,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		faultOnNilArg1: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 8},  // R3
				{1, 16}, // R4
			},
			clobbers: 1944, // R3 R4 R7 R8 R9 R10
		},
	},
	{
		name:           "LoweredAtomicStore32",
		auxType:        auxInt64,
		argLen:         3,
		faultOnNilArg0: true,
		hasSideEffects: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "LoweredAtomicStore64",
		auxType:        auxInt64,
		argLen:         3,
		faultOnNilArg0: true,
		hasSideEffects: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "LoweredAtomicLoad8",
		auxType:        auxInt64,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "LoweredAtomicLoad32",
		auxType:        auxInt64,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "LoweredAtomicLoad64",
		auxType:        auxInt64,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "LoweredAtomicLoadPtr",
		auxType:        auxInt64,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:            "LoweredAtomicAdd32",
		argLen:          3,
		resultNotInArgs: true,
		clobberFlags:    true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:            "LoweredAtomicAdd64",
		argLen:          3,
		resultNotInArgs: true,
		clobberFlags:    true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:            "LoweredAtomicExchange32",
		argLen:          3,
		resultNotInArgs: true,
		clobberFlags:    true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:            "LoweredAtomicExchange64",
		argLen:          3,
		resultNotInArgs: true,
		clobberFlags:    true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:            "LoweredAtomicCas64",
		auxType:         auxInt64,
		argLen:          4,
		resultNotInArgs: true,
		clobberFlags:    true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{2, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:            "LoweredAtomicCas32",
		auxType:         auxInt64,
		argLen:          4,
		resultNotInArgs: true,
		clobberFlags:    true,
		faultOnNilArg0:  true,
		hasSideEffects:  true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{2, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
			outputs: []outputInfo{
				{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "LoweredAtomicAnd8",
		argLen:         3,
		faultOnNilArg0: true,
		hasSideEffects: true,
		asm:            ppc64.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:           "LoweredAtomicOr8",
		argLen:         3,
		faultOnNilArg0: true,
		hasSideEffects: true,
		asm:            ppc64.AOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
			},
		},
	},
	{
		name:         "LoweredWB",
		auxType:      auxSym,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymNone,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1048576}, // R20
				{1, 2097152}, // R21
			},
			clobbers: 576460746931503104, // R16 R17 R18 R19 R22 R23 R24 R25 R26 R27 R28 R29 R31 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
		},
	},
	{
		name:    "LoweredPanicBoundsA",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 32}, // R5
				{1, 64}, // R6
			},
		},
	},
	{
		name:    "LoweredPanicBoundsB",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 16}, // R4
				{1, 32}, // R5
			},
		},
	},
	{
		name:    "LoweredPanicBoundsC",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 8},  // R3
				{1, 16}, // R4
			},
		},
	},
	{
		name:   "InvertFlags",
		argLen: 1,
		reg:    regInfo{},
	},
	{
		name:   "FlagEQ",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagLT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagGT",
		argLen: 0,
		reg:    regInfo{},
	},

	{
		name:         "FADDS",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AFADDS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "FADD",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AFADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "FSUBS",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AFSUBS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "FSUB",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AFSUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "FMULS",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          s390x.AFMULS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "FMUL",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		asm:          s390x.AFMUL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "FDIVS",
		argLen:       2,
		resultInArg0: true,
		asm:          s390x.AFDIVS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "FDIV",
		argLen:       2,
		resultInArg0: true,
		asm:          s390x.AFDIV,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "FNEGS",
		argLen:       1,
		clobberFlags: true,
		asm:          s390x.AFNEGS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "FNEG",
		argLen:       1,
		clobberFlags: true,
		asm:          s390x.AFNEG,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "FMADDS",
		argLen:       3,
		resultInArg0: true,
		asm:          s390x.AFMADDS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "FMADD",
		argLen:       3,
		resultInArg0: true,
		asm:          s390x.AFMADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "FMSUBS",
		argLen:       3,
		resultInArg0: true,
		asm:          s390x.AFMSUBS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "FMSUB",
		argLen:       3,
		resultInArg0: true,
		asm:          s390x.AFMSUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "LPDFR",
		argLen: 1,
		asm:    s390x.ALPDFR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "LNDFR",
		argLen: 1,
		asm:    s390x.ALNDFR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "CPSDR",
		argLen: 2,
		asm:    s390x.ACPSDR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:    "FIDBR",
		auxType: auxInt8,
		argLen:  1,
		asm:     s390x.AFIDBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:           "FMOVSload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            s390x.AFMOVS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:           "FMOVDload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            s390x.AFMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:              "FMOVSconst",
		auxType:           auxFloat32,
		argLen:            0,
		rematerializeable: true,
		asm:               s390x.AFMOVS,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:              "FMOVDconst",
		auxType:           auxFloat64,
		argLen:            0,
		rematerializeable: true,
		asm:               s390x.AFMOVD,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:      "FMOVSloadidx",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       s390x.AFMOVS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:      "FMOVDloadidx",
		auxType:   auxSymOff,
		argLen:    3,
		symEffect: SymRead,
		asm:       s390x.AFMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:           "FMOVSstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.AFMOVS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:           "FMOVDstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.AFMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:      "FMOVSstoreidx",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       s390x.AFMOVS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:      "FMOVDstoreidx",
		auxType:   auxSymOff,
		argLen:    4,
		symEffect: SymWrite,
		asm:       s390x.AFMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "ADD",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          s390x.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "ADDW",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          s390x.AADDW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "ADDconst",
		auxType:      auxInt32,
		argLen:       1,
		clobberFlags: true,
		asm:          s390x.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "ADDWconst",
		auxType:      auxInt32,
		argLen:       1,
		clobberFlags: true,
		asm:          s390x.AADDW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "ADDload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            s390x.AADD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "ADDWload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            s390x.AADDW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "SUB",
		argLen:       2,
		clobberFlags: true,
		asm:          s390x.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "SUBW",
		argLen:       2,
		clobberFlags: true,
		asm:          s390x.ASUBW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "SUBconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "SUBWconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.ASUBW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "SUBload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            s390x.ASUB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "SUBWload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            s390x.ASUBW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "MULLD",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AMULLD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "MULLW",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AMULLW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "MULLDconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AMULLD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "MULLWconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AMULLW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MULLDload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            s390x.AMULLD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MULLWload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            s390x.AMULLW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "MULHD",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AMULHD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 2048, // R11
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:         "MULHDU",
		argLen:       2,
		commutative:  true,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AMULHDU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 2048, // R11
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:         "DIVD",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.ADIVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 2048, // R11
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:         "DIVW",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.ADIVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 2048, // R11
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:         "DIVDU",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.ADIVDU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 2048, // R11
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:         "DIVWU",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.ADIVWU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 2048, // R11
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:         "MODD",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AMODD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 2048, // R11
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:         "MODW",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AMODW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 2048, // R11
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:         "MODDU",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AMODDU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 2048, // R11
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:         "MODWU",
		argLen:       2,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AMODWU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
			clobbers: 2048, // R11
			outputs: []outputInfo{
				{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
			},
		},
	},
	{
		name:         "AND",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          s390x.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "ANDW",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          s390x.AANDW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "ANDconst",
		auxType:      auxInt64,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "ANDWconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AANDW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "ANDload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            s390x.AAND,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "ANDWload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            s390x.AANDW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "OR",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          s390x.AOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "ORW",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          s390x.AORW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "ORconst",
		auxType:      auxInt64,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "ORWconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AORW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "ORload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            s390x.AOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "ORWload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            s390x.AORW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "XOR",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          s390x.AXOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "XORW",
		argLen:       2,
		commutative:  true,
		clobberFlags: true,
		asm:          s390x.AXORW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "XORconst",
		auxType:      auxInt64,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AXOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "XORWconst",
		auxType:      auxInt32,
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		asm:          s390x.AXORW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "XORload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            s390x.AXOR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "XORWload",
		auxType:        auxSymOff,
		argLen:         3,
		resultInArg0:   true,
		clobberFlags:   true,
		faultOnNilArg1: true,
		symEffect:      SymRead,
		asm:            s390x.AXORW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:        "ADDC",
		argLen:      2,
		commutative: true,
		asm:         s390x.AADDC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:    "ADDCconst",
		auxType: auxInt16,
		argLen:  1,
		asm:     s390x.AADDC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "ADDE",
		argLen:       3,
		commutative:  true,
		resultInArg0: true,
		asm:          s390x.AADDE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "SUBC",
		argLen: 2,
		asm:    s390x.ASUBC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "SUBE",
		argLen:       3,
		resultInArg0: true,
		asm:          s390x.ASUBE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "CMP",
		argLen: 2,
		asm:    s390x.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:   "CMPW",
		argLen: 2,
		asm:    s390x.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:   "CMPU",
		argLen: 2,
		asm:    s390x.ACMPU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:   "CMPWU",
		argLen: 2,
		asm:    s390x.ACMPWU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:    "CMPconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     s390x.ACMP,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:    "CMPWconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     s390x.ACMPW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:    "CMPUconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     s390x.ACMPU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:    "CMPWUconst",
		auxType: auxInt32,
		argLen:  1,
		asm:     s390x.ACMPWU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:   "FCMPS",
		argLen: 2,
		asm:    s390x.ACEBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "FCMP",
		argLen: 2,
		asm:    s390x.AFCMPU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "SLD",
		argLen: 2,
		asm:    s390x.ASLD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "SLW",
		argLen: 2,
		asm:    s390x.ASLW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:    "SLDconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     s390x.ASLD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:    "SLWconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     s390x.ASLW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "SRD",
		argLen: 2,
		asm:    s390x.ASRD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "SRW",
		argLen: 2,
		asm:    s390x.ASRW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:    "SRDconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     s390x.ASRD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:    "SRWconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     s390x.ASRW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "SRAD",
		argLen:       2,
		clobberFlags: true,
		asm:          s390x.ASRAD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "SRAW",
		argLen:       2,
		clobberFlags: true,
		asm:          s390x.ASRAW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "SRADconst",
		auxType:      auxInt8,
		argLen:       1,
		clobberFlags: true,
		asm:          s390x.ASRAD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "SRAWconst",
		auxType:      auxInt8,
		argLen:       1,
		clobberFlags: true,
		asm:          s390x.ASRAW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "RLLG",
		argLen: 2,
		asm:    s390x.ARLLG,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "RLL",
		argLen: 2,
		asm:    s390x.ARLL,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:    "RLLGconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     s390x.ARLLG,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:    "RLLconst",
		auxType: auxInt8,
		argLen:  1,
		asm:     s390x.ARLL,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "NEG",
		argLen:       1,
		clobberFlags: true,
		asm:          s390x.ANEG,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "NEGW",
		argLen:       1,
		clobberFlags: true,
		asm:          s390x.ANEGW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "NOT",
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "NOTW",
		argLen:       1,
		resultInArg0: true,
		clobberFlags: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "FSQRT",
		argLen: 1,
		asm:    s390x.AFSQRT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "MOVDEQ",
		argLen:       3,
		resultInArg0: true,
		asm:          s390x.AMOVDEQ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "MOVDNE",
		argLen:       3,
		resultInArg0: true,
		asm:          s390x.AMOVDNE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "MOVDLT",
		argLen:       3,
		resultInArg0: true,
		asm:          s390x.AMOVDLT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "MOVDLE",
		argLen:       3,
		resultInArg0: true,
		asm:          s390x.AMOVDLE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "MOVDGT",
		argLen:       3,
		resultInArg0: true,
		asm:          s390x.AMOVDGT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "MOVDGE",
		argLen:       3,
		resultInArg0: true,
		asm:          s390x.AMOVDGE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "MOVDGTnoinv",
		argLen:       3,
		resultInArg0: true,
		asm:          s390x.AMOVDGT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "MOVDGEnoinv",
		argLen:       3,
		resultInArg0: true,
		asm:          s390x.AMOVDGE,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "MOVBreg",
		argLen: 1,
		asm:    s390x.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "MOVBZreg",
		argLen: 1,
		asm:    s390x.AMOVBZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "MOVHreg",
		argLen: 1,
		asm:    s390x.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "MOVHZreg",
		argLen: 1,
		asm:    s390x.AMOVHZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "MOVWreg",
		argLen: 1,
		asm:    s390x.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "MOVWZreg",
		argLen: 1,
		asm:    s390x.AMOVWZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "MOVDreg",
		argLen: 1,
		asm:    s390x.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "MOVDnop",
		argLen:       1,
		resultInArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:              "MOVDconst",
		auxType:           auxInt64,
		argLen:            0,
		rematerializeable: true,
		asm:               s390x.AMOVD,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "LDGR",
		argLen: 1,
		asm:    s390x.ALDGR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "LGDR",
		argLen: 1,
		asm:    s390x.ALGDR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "CFDBRA",
		argLen: 1,
		asm:    s390x.ACFDBRA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "CGDBRA",
		argLen: 1,
		asm:    s390x.ACGDBRA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "CFEBRA",
		argLen: 1,
		asm:    s390x.ACFEBRA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "CGEBRA",
		argLen: 1,
		asm:    s390x.ACGEBRA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "CEFBRA",
		argLen: 1,
		asm:    s390x.ACEFBRA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "CDFBRA",
		argLen: 1,
		asm:    s390x.ACDFBRA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "CEGBRA",
		argLen: 1,
		asm:    s390x.ACEGBRA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "CDGBRA",
		argLen: 1,
		asm:    s390x.ACDGBRA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "LEDBR",
		argLen: 1,
		asm:    s390x.ALEDBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "LDEBR",
		argLen: 1,
		asm:    s390x.ALDEBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:              "MOVDaddr",
		auxType:           auxSymOff,
		argLen:            1,
		rematerializeable: true,
		symEffect:         SymRead,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295000064}, // SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:      "MOVDaddridx",
		auxType:   auxSymOff,
		argLen:    2,
		symEffect: SymRead,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295000064}, // SP SB
				{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MOVBZload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            s390x.AMOVBZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MOVBload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            s390x.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MOVHZload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            s390x.AMOVHZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MOVHload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            s390x.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MOVWZload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            s390x.AMOVWZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MOVWload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            s390x.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MOVDload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            s390x.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "MOVWBR",
		argLen: 1,
		asm:    s390x.AMOVWBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "MOVDBR",
		argLen: 1,
		asm:    s390x.AMOVDBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MOVHBRload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            s390x.AMOVHBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MOVWBRload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            s390x.AMOVWBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MOVDBRload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            s390x.AMOVDBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MOVBstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "MOVHstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "MOVWstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "MOVDstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "MOVHBRstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.AMOVHBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "MOVWBRstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.AMOVWBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "MOVDBRstore",
		auxType:        auxSymOff,
		argLen:         3,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.AMOVDBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "MVC",
		auxType:        auxSymValAndOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		faultOnNilArg1: true,
		symEffect:      SymNone,
		asm:            s390x.AMVC,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:        "MOVBZloadidx",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         s390x.AMOVBZ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:        "MOVBloadidx",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         s390x.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:        "MOVHZloadidx",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         s390x.AMOVHZ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:        "MOVHloadidx",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         s390x.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:        "MOVWZloadidx",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         s390x.AMOVWZ,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:        "MOVWloadidx",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         s390x.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:        "MOVDloadidx",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         s390x.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:        "MOVHBRloadidx",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         s390x.AMOVHBR,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:        "MOVWBRloadidx",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         s390x.AMOVWBR,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:        "MOVDBRloadidx",
		auxType:     auxSymOff,
		argLen:      3,
		commutative: true,
		symEffect:   SymRead,
		asm:         s390x.AMOVDBR,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:        "MOVBstoreidx",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         s390x.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:        "MOVHstoreidx",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         s390x.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:        "MOVWstoreidx",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         s390x.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:        "MOVDstoreidx",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         s390x.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:        "MOVHBRstoreidx",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         s390x.AMOVHBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:        "MOVWBRstoreidx",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         s390x.AMOVWBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:        "MOVDBRstoreidx",
		auxType:     auxSymOff,
		argLen:      4,
		commutative: true,
		symEffect:   SymWrite,
		asm:         s390x.AMOVDBR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "MOVBstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.AMOVB,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
		},
	},
	{
		name:           "MOVHstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.AMOVH,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
		},
	},
	{
		name:           "MOVWstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
		},
	},
	{
		name:           "MOVDstoreconst",
		auxType:        auxSymValAndOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
		},
	},
	{
		name:           "CLEAR",
		auxType:        auxSymValAndOff,
		argLen:         2,
		clobberFlags:   true,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.ACLEAR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:         "CALLstatic",
		auxType:      auxSymOff,
		argLen:       1,
		clobberFlags: true,
		call:         true,
		symEffect:    SymNone,
		reg: regInfo{
			clobbers: 4294933503, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 g R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
		},
	},
	{
		name:         "CALLclosure",
		auxType:      auxInt64,
		argLen:       3,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 4096},  // R12
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			clobbers: 4294933503, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 g R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
		},
	},
	{
		name:         "CALLinter",
		auxType:      auxInt64,
		argLen:       2,
		clobberFlags: true,
		call:         true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			clobbers: 4294933503, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 g R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
		},
	},
	{
		name:   "InvertFlags",
		argLen: 1,
		reg:    regInfo{},
	},
	{
		name:   "LoweredGetG",
		argLen: 1,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:      "LoweredGetClosurePtr",
		argLen:    0,
		zeroWidth: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 4096}, // R12
			},
		},
	},
	{
		name:              "LoweredGetCallerSP",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:              "LoweredGetCallerPC",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "LoweredNilCheck",
		argLen:         2,
		clobberFlags:   true,
		nilCheck:       true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:         "LoweredRound32F",
		argLen:       1,
		resultInArg0: true,
		zeroWidth:    true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "LoweredRound64F",
		argLen:       1,
		resultInArg0: true,
		zeroWidth:    true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:         "LoweredWB",
		auxType:      auxSym,
		argLen:       3,
		clobberFlags: true,
		symEffect:    SymNone,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4}, // R2
				{1, 8}, // R3
			},
			clobbers: 4294918144, // R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
		},
	},
	{
		name:    "LoweredPanicBoundsA",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4}, // R2
				{1, 8}, // R3
			},
		},
	},
	{
		name:    "LoweredPanicBoundsB",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2}, // R1
				{1, 4}, // R2
			},
		},
	},
	{
		name:    "LoweredPanicBoundsC",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 1}, // R0
				{1, 2}, // R1
			},
		},
	},
	{
		name:   "FlagEQ",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagLT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagGT",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:   "FlagOV",
		argLen: 0,
		reg:    regInfo{},
	},
	{
		name:           "MOVBZatomicload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            s390x.AMOVBZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MOVWZatomicload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            s390x.AMOVWZ,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MOVDatomicload",
		auxType:        auxSymOff,
		argLen:         2,
		faultOnNilArg0: true,
		symEffect:      SymRead,
		asm:            s390x.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "MOVWatomicstore",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymWrite,
		asm:            s390x.AMOVW,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "MOVDatomicstore",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymWrite,
		asm:            s390x.AMOVD,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "LAA",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            s390x.ALAA,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "LAAG",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            s390x.ALAAG,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "AddTupleFirst32",
		argLen: 2,
		reg:    regInfo{},
	},
	{
		name:   "AddTupleFirst64",
		argLen: 2,
		reg:    regInfo{},
	},
	{
		name:           "LoweredAtomicCas32",
		auxType:        auxSymOff,
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            s390x.ACS,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1},     // R0
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			clobbers: 1, // R0
			outputs: []outputInfo{
				{1, 0},
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "LoweredAtomicCas64",
		auxType:        auxSymOff,
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            s390x.ACSG,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 1},     // R0
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			clobbers: 1, // R0
			outputs: []outputInfo{
				{1, 0},
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:           "LoweredAtomicExchange32",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            s390x.ACS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 1}, // R0
			},
		},
	},
	{
		name:           "LoweredAtomicExchange64",
		auxType:        auxSymOff,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		hasSideEffects: true,
		symEffect:      SymRdWr,
		asm:            s390x.ACSG,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			outputs: []outputInfo{
				{1, 0},
				{0, 1}, // R0
			},
		},
	},
	{
		name:         "FLOGR",
		argLen:       1,
		clobberFlags: true,
		asm:          s390x.AFLOGR,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			clobbers: 2, // R1
			outputs: []outputInfo{
				{0, 1}, // R0
			},
		},
	},
	{
		name:         "POPCNT",
		argLen:       1,
		clobberFlags: true,
		asm:          s390x.APOPCNT,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
			outputs: []outputInfo{
				{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
			},
		},
	},
	{
		name:   "SumBytes2",
		argLen: 1,
		reg:    regInfo{},
	},
	{
		name:   "SumBytes4",
		argLen: 1,
		reg:    regInfo{},
	},
	{
		name:   "SumBytes8",
		argLen: 1,
		reg:    regInfo{},
	},
	{
		name:           "STMG2",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.ASTMG,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // R1
				{2, 4},     // R2
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "STMG3",
		auxType:        auxSymOff,
		argLen:         5,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.ASTMG,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // R1
				{2, 4},     // R2
				{3, 8},     // R3
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "STMG4",
		auxType:        auxSymOff,
		argLen:         6,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.ASTMG,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // R1
				{2, 4},     // R2
				{3, 8},     // R3
				{4, 16},    // R4
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "STM2",
		auxType:        auxSymOff,
		argLen:         4,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.ASTMY,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // R1
				{2, 4},     // R2
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "STM3",
		auxType:        auxSymOff,
		argLen:         5,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.ASTMY,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // R1
				{2, 4},     // R2
				{3, 8},     // R3
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "STM4",
		auxType:        auxSymOff,
		argLen:         6,
		faultOnNilArg0: true,
		symEffect:      SymWrite,
		asm:            s390x.ASTMY,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 2},     // R1
				{2, 4},     // R2
				{3, 8},     // R3
				{4, 16},    // R4
				{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
		},
	},
	{
		name:           "LoweredMove",
		auxType:        auxInt64,
		argLen:         4,
		clobberFlags:   true,
		faultOnNilArg0: true,
		faultOnNilArg1: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2},     // R1
				{1, 4},     // R2
				{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			clobbers: 6, // R1 R2
		},
	},
	{
		name:           "LoweredZero",
		auxType:        auxInt64,
		argLen:         3,
		clobberFlags:   true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 2},     // R1
				{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
			},
			clobbers: 2, // R1
		},
	},

	{
		name:      "LoweredStaticCall",
		auxType:   auxSymOff,
		argLen:    1,
		call:      true,
		symEffect: SymNone,
		reg: regInfo{
			clobbers: 12884901887, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 g
		},
	},
	{
		name:    "LoweredClosureCall",
		auxType: auxInt64,
		argLen:  3,
		call:    true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
			clobbers: 12884901887, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 g
		},
	},
	{
		name:    "LoweredInterCall",
		auxType: auxInt64,
		argLen:  2,
		call:    true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
			clobbers: 12884901887, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 g
		},
	},
	{
		name:              "LoweredAddr",
		auxType:           auxSymOff,
		argLen:            1,
		rematerializeable: true,
		symEffect:         SymAddr,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "LoweredMove",
		auxType: auxInt64,
		argLen:  3,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "LoweredZero",
		auxType: auxInt64,
		argLen:  2,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "LoweredGetClosurePtr",
		argLen: 0,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:              "LoweredGetCallerPC",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:              "LoweredGetCallerSP",
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:           "LoweredNilCheck",
		argLen:         2,
		nilCheck:       true,
		faultOnNilArg0: true,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:      "LoweredWB",
		auxType:   auxSym,
		argLen:    3,
		symEffect: SymNone,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				{1, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "LoweredRound32F",
		argLen: 1,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "LoweredConvert",
		argLen: 2,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "Select",
		argLen: 3,
		asm:    wasm.ASelect,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{2, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "I64Load8U",
		auxType: auxInt64,
		argLen:  2,
		asm:     wasm.AI64Load8U,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "I64Load8S",
		auxType: auxInt64,
		argLen:  2,
		asm:     wasm.AI64Load8S,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "I64Load16U",
		auxType: auxInt64,
		argLen:  2,
		asm:     wasm.AI64Load16U,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "I64Load16S",
		auxType: auxInt64,
		argLen:  2,
		asm:     wasm.AI64Load16S,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "I64Load32U",
		auxType: auxInt64,
		argLen:  2,
		asm:     wasm.AI64Load32U,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "I64Load32S",
		auxType: auxInt64,
		argLen:  2,
		asm:     wasm.AI64Load32S,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "I64Load",
		auxType: auxInt64,
		argLen:  2,
		asm:     wasm.AI64Load,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "I64Store8",
		auxType: auxInt64,
		argLen:  3,
		asm:     wasm.AI64Store8,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 4295032831},  // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
			},
		},
	},
	{
		name:    "I64Store16",
		auxType: auxInt64,
		argLen:  3,
		asm:     wasm.AI64Store16,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 4295032831},  // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
			},
		},
	},
	{
		name:    "I64Store32",
		auxType: auxInt64,
		argLen:  3,
		asm:     wasm.AI64Store32,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 4295032831},  // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
			},
		},
	},
	{
		name:    "I64Store",
		auxType: auxInt64,
		argLen:  3,
		asm:     wasm.AI64Store,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 4295032831},  // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
			},
		},
	},
	{
		name:    "F32Load",
		auxType: auxInt64,
		argLen:  2,
		asm:     wasm.AF32Load,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:    "F64Load",
		auxType: auxInt64,
		argLen:  2,
		asm:     wasm.AF64Load,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:    "F32Store",
		auxType: auxInt64,
		argLen:  3,
		asm:     wasm.AF32Store,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 4294901760},  // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
			},
		},
	},
	{
		name:    "F64Store",
		auxType: auxInt64,
		argLen:  3,
		asm:     wasm.AF64Store,
		reg: regInfo{
			inputs: []inputInfo{
				{1, 4294901760},  // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
			},
		},
	},
	{
		name:              "I64Const",
		auxType:           auxInt64,
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:              "F64Const",
		auxType:           auxFloat64,
		argLen:            0,
		rematerializeable: true,
		reg: regInfo{
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "I64Eqz",
		argLen: 1,
		asm:    wasm.AI64Eqz,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64Eq",
		argLen: 2,
		asm:    wasm.AI64Eq,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64Ne",
		argLen: 2,
		asm:    wasm.AI64Ne,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64LtS",
		argLen: 2,
		asm:    wasm.AI64LtS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64LtU",
		argLen: 2,
		asm:    wasm.AI64LtU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64GtS",
		argLen: 2,
		asm:    wasm.AI64GtS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64GtU",
		argLen: 2,
		asm:    wasm.AI64GtU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64LeS",
		argLen: 2,
		asm:    wasm.AI64LeS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64LeU",
		argLen: 2,
		asm:    wasm.AI64LeU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64GeS",
		argLen: 2,
		asm:    wasm.AI64GeS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64GeU",
		argLen: 2,
		asm:    wasm.AI64GeU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "F64Eq",
		argLen: 2,
		asm:    wasm.AF64Eq,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "F64Ne",
		argLen: 2,
		asm:    wasm.AF64Ne,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "F64Lt",
		argLen: 2,
		asm:    wasm.AF64Lt,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "F64Gt",
		argLen: 2,
		asm:    wasm.AF64Gt,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "F64Le",
		argLen: 2,
		asm:    wasm.AF64Le,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "F64Ge",
		argLen: 2,
		asm:    wasm.AF64Ge,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64Add",
		argLen: 2,
		asm:    wasm.AI64Add,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:    "I64AddConst",
		auxType: auxInt64,
		argLen:  1,
		asm:     wasm.AI64Add,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64Sub",
		argLen: 2,
		asm:    wasm.AI64Sub,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64Mul",
		argLen: 2,
		asm:    wasm.AI64Mul,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64DivS",
		argLen: 2,
		asm:    wasm.AI64DivS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64DivU",
		argLen: 2,
		asm:    wasm.AI64DivU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64RemS",
		argLen: 2,
		asm:    wasm.AI64RemS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64RemU",
		argLen: 2,
		asm:    wasm.AI64RemU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64And",
		argLen: 2,
		asm:    wasm.AI64And,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64Or",
		argLen: 2,
		asm:    wasm.AI64Or,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64Xor",
		argLen: 2,
		asm:    wasm.AI64Xor,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64Shl",
		argLen: 2,
		asm:    wasm.AI64Shl,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64ShrS",
		argLen: 2,
		asm:    wasm.AI64ShrS,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64ShrU",
		argLen: 2,
		asm:    wasm.AI64ShrU,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "F64Neg",
		argLen: 1,
		asm:    wasm.AF64Neg,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "F64Add",
		argLen: 2,
		asm:    wasm.AF64Add,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "F64Sub",
		argLen: 2,
		asm:    wasm.AF64Sub,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "F64Mul",
		argLen: 2,
		asm:    wasm.AF64Mul,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "F64Div",
		argLen: 2,
		asm:    wasm.AF64Div,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "I64TruncSatF64S",
		argLen: 1,
		asm:    wasm.AI64TruncSatF64S,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64TruncSatF64U",
		argLen: 1,
		asm:    wasm.AI64TruncSatF64U,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "F64ConvertI64S",
		argLen: 1,
		asm:    wasm.AF64ConvertI64S,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "F64ConvertI64U",
		argLen: 1,
		asm:    wasm.AF64ConvertI64U,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "I64Extend8S",
		argLen: 1,
		asm:    wasm.AI64Extend8S,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64Extend16S",
		argLen: 1,
		asm:    wasm.AI64Extend16S,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64Extend32S",
		argLen: 1,
		asm:    wasm.AI64Extend32S,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "F64Sqrt",
		argLen: 1,
		asm:    wasm.AF64Sqrt,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "F64Trunc",
		argLen: 1,
		asm:    wasm.AF64Trunc,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "F64Ceil",
		argLen: 1,
		asm:    wasm.AF64Ceil,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "F64Floor",
		argLen: 1,
		asm:    wasm.AF64Floor,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "F64Nearest",
		argLen: 1,
		asm:    wasm.AF64Nearest,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "F64Abs",
		argLen: 1,
		asm:    wasm.AF64Abs,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "F64Copysign",
		argLen: 2,
		asm:    wasm.AF64Copysign,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
			outputs: []outputInfo{
				{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
	},
	{
		name:   "I64Ctz",
		argLen: 1,
		asm:    wasm.AI64Ctz,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64Clz",
		argLen: 1,
		asm:    wasm.AI64Clz,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64Rotl",
		argLen: 2,
		asm:    wasm.AI64Rotl,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},
	{
		name:   "I64Popcnt",
		argLen: 1,
		asm:    wasm.AI64Popcnt,
		reg: regInfo{
			inputs: []inputInfo{
				{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
			},
			outputs: []outputInfo{
				{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
			},
		},
	},

	{
		name:        "Add8",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Add16",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Add32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Add64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "AddPtr",
		argLen:  2,
		generic: true,
	},
	{
		name:        "Add32F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Add64F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Sub8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Sub16",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Sub32",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Sub64",
		argLen:  2,
		generic: true,
	},
	{
		name:    "SubPtr",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Sub32F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Sub64F",
		argLen:  2,
		generic: true,
	},
	{
		name:        "Mul8",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul16",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul32F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul64F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Div32F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div64F",
		argLen:  2,
		generic: true,
	},
	{
		name:        "Hmul32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Hmul32u",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Hmul64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Hmul64u",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul32uhilo",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul64uhilo",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul32uover",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Mul64uover",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Avg32u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Avg64u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div8u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div16u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div32u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div64u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Div128u",
		argLen:  3,
		generic: true,
	},
	{
		name:    "Mod8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Mod8u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Mod16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Mod16u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Mod32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Mod32u",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Mod64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Mod64u",
		argLen:  2,
		generic: true,
	},
	{
		name:        "And8",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "And16",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "And32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "And64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Or8",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Or16",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Or32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Or64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Xor8",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Xor16",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Xor32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Xor64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Lsh8x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh8x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh8x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh8x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh16x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh16x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh16x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh16x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh32x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh32x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh32x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh32x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh64x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh64x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh64x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Lsh64x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64x8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64x16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64x32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64x64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8Ux8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8Ux16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8Ux32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh8Ux64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16Ux8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16Ux16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16Ux32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh16Ux64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32Ux8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32Ux16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32Ux32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh32Ux64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64Ux8",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64Ux16",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64Ux32",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:    "Rsh64Ux64",
		auxType: auxBool,
		argLen:  2,
		generic: true,
	},
	{
		name:        "Eq8",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Eq16",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Eq32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Eq64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "EqPtr",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "EqInter",
		argLen:  2,
		generic: true,
	},
	{
		name:    "EqSlice",
		argLen:  2,
		generic: true,
	},
	{
		name:        "Eq32F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Eq64F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Neq8",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Neq16",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Neq32",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Neq64",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "NeqPtr",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "NeqInter",
		argLen:  2,
		generic: true,
	},
	{
		name:    "NeqSlice",
		argLen:  2,
		generic: true,
	},
	{
		name:        "Neq32F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Neq64F",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Less8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less8U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less16",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less16U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less32",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less32U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less64",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less64U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less32F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Less64F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq8U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq16",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq16U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq32",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq32U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq64",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq64U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq32F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Leq64F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Greater8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Greater8U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Greater16",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Greater16U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Greater32",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Greater32U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Greater64",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Greater64U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Greater32F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Greater64F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Geq8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Geq8U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Geq16",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Geq16U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Geq32",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Geq32U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Geq64",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Geq64U",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Geq32F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Geq64F",
		argLen:  2,
		generic: true,
	},
	{
		name:    "CondSelect",
		argLen:  3,
		generic: true,
	},
	{
		name:        "AndB",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "OrB",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "EqB",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "NeqB",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Not",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Neg8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Neg16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Neg32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Neg64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Neg32F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Neg64F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Com8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Com16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Com32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Com64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz8NonZero",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz16NonZero",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz32NonZero",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ctz64NonZero",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitLen8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitLen16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitLen32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitLen64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Bswap32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Bswap64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitRev8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitRev16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitRev32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "BitRev64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "PopCount8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "PopCount16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "PopCount32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "PopCount64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "RotateLeft8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "RotateLeft16",
		argLen:  2,
		generic: true,
	},
	{
		name:    "RotateLeft32",
		argLen:  2,
		generic: true,
	},
	{
		name:    "RotateLeft64",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Sqrt",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Floor",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Ceil",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Trunc",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Round",
		argLen:  1,
		generic: true,
	},
	{
		name:    "RoundToEven",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Abs",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Copysign",
		argLen:  2,
		generic: true,
	},
	{
		name:      "Phi",
		argLen:    -1,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:    "Copy",
		argLen:  1,
		generic: true,
	},
	{
		name:         "Convert",
		argLen:       2,
		resultInArg0: true,
		zeroWidth:    true,
		generic:      true,
	},
	{
		name:    "ConstBool",
		auxType: auxBool,
		argLen:  0,
		generic: true,
	},
	{
		name:    "ConstString",
		auxType: auxString,
		argLen:  0,
		generic: true,
	},
	{
		name:    "ConstNil",
		argLen:  0,
		generic: true,
	},
	{
		name:    "Const8",
		auxType: auxInt8,
		argLen:  0,
		generic: true,
	},
	{
		name:    "Const16",
		auxType: auxInt16,
		argLen:  0,
		generic: true,
	},
	{
		name:    "Const32",
		auxType: auxInt32,
		argLen:  0,
		generic: true,
	},
	{
		name:    "Const64",
		auxType: auxInt64,
		argLen:  0,
		generic: true,
	},
	{
		name:    "Const32F",
		auxType: auxFloat32,
		argLen:  0,
		generic: true,
	},
	{
		name:    "Const64F",
		auxType: auxFloat64,
		argLen:  0,
		generic: true,
	},
	{
		name:    "ConstInterface",
		argLen:  0,
		generic: true,
	},
	{
		name:    "ConstSlice",
		argLen:  0,
		generic: true,
	},
	{
		name:      "InitMem",
		argLen:    0,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:      "Arg",
		auxType:   auxSymOff,
		argLen:    0,
		zeroWidth: true,
		symEffect: SymRead,
		generic:   true,
	},
	{
		name:      "Addr",
		auxType:   auxSym,
		argLen:    1,
		symEffect: SymAddr,
		generic:   true,
	},
	{
		name:      "LocalAddr",
		auxType:   auxSym,
		argLen:    2,
		symEffect: SymAddr,
		generic:   true,
	},
	{
		name:      "SP",
		argLen:    0,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:      "SB",
		argLen:    0,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:    "Load",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Store",
		auxType: auxTyp,
		argLen:  3,
		generic: true,
	},
	{
		name:    "Move",
		auxType: auxTypSize,
		argLen:  3,
		generic: true,
	},
	{
		name:    "Zero",
		auxType: auxTypSize,
		argLen:  2,
		generic: true,
	},
	{
		name:    "StoreWB",
		auxType: auxTyp,
		argLen:  3,
		generic: true,
	},
	{
		name:    "MoveWB",
		auxType: auxTypSize,
		argLen:  3,
		generic: true,
	},
	{
		name:    "ZeroWB",
		auxType: auxTypSize,
		argLen:  2,
		generic: true,
	},
	{
		name:      "WB",
		auxType:   auxSym,
		argLen:    3,
		symEffect: SymNone,
		generic:   true,
	},
	{
		name:    "PanicBounds",
		auxType: auxInt64,
		argLen:  3,
		generic: true,
	},
	{
		name:    "PanicExtend",
		auxType: auxInt64,
		argLen:  4,
		generic: true,
	},
	{
		name:    "ClosureCall",
		auxType: auxInt64,
		argLen:  3,
		call:    true,
		generic: true,
	},
	{
		name:      "StaticCall",
		auxType:   auxSymOff,
		argLen:    1,
		call:      true,
		symEffect: SymNone,
		generic:   true,
	},
	{
		name:    "InterCall",
		auxType: auxInt64,
		argLen:  2,
		call:    true,
		generic: true,
	},
	{
		name:    "SignExt8to16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SignExt8to32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SignExt8to64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SignExt16to32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SignExt16to64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SignExt32to64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ZeroExt8to16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ZeroExt8to32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ZeroExt8to64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ZeroExt16to32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ZeroExt16to64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ZeroExt32to64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Trunc16to8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Trunc32to8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Trunc32to16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Trunc64to8",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Trunc64to16",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Trunc64to32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32to32F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32to64F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64to32F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64to64F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32Fto32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32Fto64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64Fto32",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64Fto64",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32Fto64F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64Fto32F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Round32F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Round64F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "IsNonNil",
		argLen:  1,
		generic: true,
	},
	{
		name:    "IsInBounds",
		argLen:  2,
		generic: true,
	},
	{
		name:    "IsSliceInBounds",
		argLen:  2,
		generic: true,
	},
	{
		name:    "NilCheck",
		argLen:  2,
		generic: true,
	},
	{
		name:      "GetG",
		argLen:    1,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:    "GetClosurePtr",
		argLen:  0,
		generic: true,
	},
	{
		name:    "GetCallerPC",
		argLen:  0,
		generic: true,
	},
	{
		name:    "GetCallerSP",
		argLen:  0,
		generic: true,
	},
	{
		name:    "PtrIndex",
		argLen:  2,
		generic: true,
	},
	{
		name:    "OffPtr",
		auxType: auxInt64,
		argLen:  1,
		generic: true,
	},
	{
		name:    "SliceMake",
		argLen:  3,
		generic: true,
	},
	{
		name:    "SlicePtr",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SliceLen",
		argLen:  1,
		generic: true,
	},
	{
		name:    "SliceCap",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ComplexMake",
		argLen:  2,
		generic: true,
	},
	{
		name:    "ComplexReal",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ComplexImag",
		argLen:  1,
		generic: true,
	},
	{
		name:    "StringMake",
		argLen:  2,
		generic: true,
	},
	{
		name:    "StringPtr",
		argLen:  1,
		generic: true,
	},
	{
		name:    "StringLen",
		argLen:  1,
		generic: true,
	},
	{
		name:    "IMake",
		argLen:  2,
		generic: true,
	},
	{
		name:    "ITab",
		argLen:  1,
		generic: true,
	},
	{
		name:    "IData",
		argLen:  1,
		generic: true,
	},
	{
		name:    "StructMake0",
		argLen:  0,
		generic: true,
	},
	{
		name:    "StructMake1",
		argLen:  1,
		generic: true,
	},
	{
		name:    "StructMake2",
		argLen:  2,
		generic: true,
	},
	{
		name:    "StructMake3",
		argLen:  3,
		generic: true,
	},
	{
		name:    "StructMake4",
		argLen:  4,
		generic: true,
	},
	{
		name:    "StructSelect",
		auxType: auxInt64,
		argLen:  1,
		generic: true,
	},
	{
		name:    "ArrayMake0",
		argLen:  0,
		generic: true,
	},
	{
		name:    "ArrayMake1",
		argLen:  1,
		generic: true,
	},
	{
		name:    "ArraySelect",
		auxType: auxInt64,
		argLen:  1,
		generic: true,
	},
	{
		name:    "StoreReg",
		argLen:  1,
		generic: true,
	},
	{
		name:    "LoadReg",
		argLen:  1,
		generic: true,
	},
	{
		name:      "FwdRef",
		auxType:   auxSym,
		argLen:    0,
		symEffect: SymNone,
		generic:   true,
	},
	{
		name:    "Unknown",
		argLen:  0,
		generic: true,
	},
	{
		name:      "VarDef",
		auxType:   auxSym,
		argLen:    1,
		zeroWidth: true,
		symEffect: SymNone,
		generic:   true,
	},
	{
		name:      "VarKill",
		auxType:   auxSym,
		argLen:    1,
		symEffect: SymNone,
		generic:   true,
	},
	{
		name:      "VarLive",
		auxType:   auxSym,
		argLen:    1,
		zeroWidth: true,
		symEffect: SymRead,
		generic:   true,
	},
	{
		name:      "KeepAlive",
		argLen:    2,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:    "InlMark",
		auxType: auxInt32,
		argLen:  1,
		generic: true,
	},
	{
		name:    "Int64Make",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Int64Hi",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Int64Lo",
		argLen:  1,
		generic: true,
	},
	{
		name:        "Add32carry",
		argLen:      2,
		commutative: true,
		generic:     true,
	},
	{
		name:        "Add32withcarry",
		argLen:      3,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Sub32carry",
		argLen:  2,
		generic: true,
	},
	{
		name:    "Sub32withcarry",
		argLen:  3,
		generic: true,
	},
	{
		name:        "Add64carry",
		argLen:      3,
		commutative: true,
		generic:     true,
	},
	{
		name:    "Sub64borrow",
		argLen:  3,
		generic: true,
	},
	{
		name:    "Signmask",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Zeromask",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Slicemask",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32Uto32F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32Uto64F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32Fto32U",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64Fto32U",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64Uto32F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64Uto64F",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt32Fto64U",
		argLen:  1,
		generic: true,
	},
	{
		name:    "Cvt64Fto64U",
		argLen:  1,
		generic: true,
	},
	{
		name:      "Select0",
		argLen:    1,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:      "Select1",
		argLen:    1,
		zeroWidth: true,
		generic:   true,
	},
	{
		name:    "AtomicLoad8",
		argLen:  2,
		generic: true,
	},
	{
		name:    "AtomicLoad32",
		argLen:  2,
		generic: true,
	},
	{
		name:    "AtomicLoad64",
		argLen:  2,
		generic: true,
	},
	{
		name:    "AtomicLoadPtr",
		argLen:  2,
		generic: true,
	},
	{
		name:    "AtomicLoadAcq32",
		argLen:  2,
		generic: true,
	},
	{
		name:           "AtomicStore32",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicStore64",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicStorePtrNoWB",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicStoreRel32",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicExchange32",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicExchange64",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicAdd32",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicAdd64",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicCompareAndSwap32",
		argLen:         4,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicCompareAndSwap64",
		argLen:         4,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicCompareAndSwapRel32",
		argLen:         4,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicAnd8",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicOr8",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicAdd32Variant",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:           "AtomicAdd64Variant",
		argLen:         3,
		hasSideEffects: true,
		generic:        true,
	},
	{
		name:      "Clobber",
		auxType:   auxSymOff,
		argLen:    0,
		symEffect: SymNone,
		generic:   true,
	},
}

func (o Op) Asm() obj.As          { return opcodeTable[o].asm }
func (o Op) Scale() int16         { return int16(opcodeTable[o].scale) }
func (o Op) String() string       { return opcodeTable[o].name }
func (o Op) UsesScratch() bool    { return opcodeTable[o].usesScratch }
func (o Op) SymEffect() SymEffect { return opcodeTable[o].symEffect }
func (o Op) IsCall() bool         { return opcodeTable[o].call }

var registers386 = [...]Register{
	{0, x86.REG_AX, 0, "AX"},
	{1, x86.REG_CX, 1, "CX"},
	{2, x86.REG_DX, 2, "DX"},
	{3, x86.REG_BX, 3, "BX"},
	{4, x86.REGSP, -1, "SP"},
	{5, x86.REG_BP, 4, "BP"},
	{6, x86.REG_SI, 5, "SI"},
	{7, x86.REG_DI, 6, "DI"},
	{8, x86.REG_X0, -1, "X0"},
	{9, x86.REG_X1, -1, "X1"},
	{10, x86.REG_X2, -1, "X2"},
	{11, x86.REG_X3, -1, "X3"},
	{12, x86.REG_X4, -1, "X4"},
	{13, x86.REG_X5, -1, "X5"},
	{14, x86.REG_X6, -1, "X6"},
	{15, x86.REG_X7, -1, "X7"},
	{16, 0, -1, "SB"},
}
var gpRegMask386 = regMask(239)
var fpRegMask386 = regMask(65280)
var specialRegMask386 = regMask(0)
var framepointerReg386 = int8(5)
var linkReg386 = int8(-1)
var registersAMD64 = [...]Register{
	{0, x86.REG_AX, 0, "AX"},
	{1, x86.REG_CX, 1, "CX"},
	{2, x86.REG_DX, 2, "DX"},
	{3, x86.REG_BX, 3, "BX"},
	{4, x86.REGSP, -1, "SP"},
	{5, x86.REG_BP, 4, "BP"},
	{6, x86.REG_SI, 5, "SI"},
	{7, x86.REG_DI, 6, "DI"},
	{8, x86.REG_R8, 7, "R8"},
	{9, x86.REG_R9, 8, "R9"},
	{10, x86.REG_R10, 9, "R10"},
	{11, x86.REG_R11, 10, "R11"},
	{12, x86.REG_R12, 11, "R12"},
	{13, x86.REG_R13, 12, "R13"},
	{14, x86.REG_R14, 13, "R14"},
	{15, x86.REG_R15, 14, "R15"},
	{16, x86.REG_X0, -1, "X0"},
	{17, x86.REG_X1, -1, "X1"},
	{18, x86.REG_X2, -1, "X2"},
	{19, x86.REG_X3, -1, "X3"},
	{20, x86.REG_X4, -1, "X4"},
	{21, x86.REG_X5, -1, "X5"},
	{22, x86.REG_X6, -1, "X6"},
	{23, x86.REG_X7, -1, "X7"},
	{24, x86.REG_X8, -1, "X8"},
	{25, x86.REG_X9, -1, "X9"},
	{26, x86.REG_X10, -1, "X10"},
	{27, x86.REG_X11, -1, "X11"},
	{28, x86.REG_X12, -1, "X12"},
	{29, x86.REG_X13, -1, "X13"},
	{30, x86.REG_X14, -1, "X14"},
	{31, x86.REG_X15, -1, "X15"},
	{32, 0, -1, "SB"},
}
var gpRegMaskAMD64 = regMask(65519)
var fpRegMaskAMD64 = regMask(4294901760)
var specialRegMaskAMD64 = regMask(0)
var framepointerRegAMD64 = int8(5)
var linkRegAMD64 = int8(-1)
var registersARM = [...]Register{
	{0, arm.REG_R0, 0, "R0"},
	{1, arm.REG_R1, 1, "R1"},
	{2, arm.REG_R2, 2, "R2"},
	{3, arm.REG_R3, 3, "R3"},
	{4, arm.REG_R4, 4, "R4"},
	{5, arm.REG_R5, 5, "R5"},
	{6, arm.REG_R6, 6, "R6"},
	{7, arm.REG_R7, 7, "R7"},
	{8, arm.REG_R8, 8, "R8"},
	{9, arm.REG_R9, 9, "R9"},
	{10, arm.REGG, -1, "g"},
	{11, arm.REG_R11, -1, "R11"},
	{12, arm.REG_R12, 10, "R12"},
	{13, arm.REGSP, -1, "SP"},
	{14, arm.REG_R14, 11, "R14"},
	{15, arm.REG_R15, -1, "R15"},
	{16, arm.REG_F0, -1, "F0"},
	{17, arm.REG_F1, -1, "F1"},
	{18, arm.REG_F2, -1, "F2"},
	{19, arm.REG_F3, -1, "F3"},
	{20, arm.REG_F4, -1, "F4"},
	{21, arm.REG_F5, -1, "F5"},
	{22, arm.REG_F6, -1, "F6"},
	{23, arm.REG_F7, -1, "F7"},
	{24, arm.REG_F8, -1, "F8"},
	{25, arm.REG_F9, -1, "F9"},
	{26, arm.REG_F10, -1, "F10"},
	{27, arm.REG_F11, -1, "F11"},
	{28, arm.REG_F12, -1, "F12"},
	{29, arm.REG_F13, -1, "F13"},
	{30, arm.REG_F14, -1, "F14"},
	{31, arm.REG_F15, -1, "F15"},
	{32, 0, -1, "SB"},
}
var gpRegMaskARM = regMask(21503)
var fpRegMaskARM = regMask(4294901760)
var specialRegMaskARM = regMask(0)
var framepointerRegARM = int8(-1)
var linkRegARM = int8(14)
var registersARM64 = [...]Register{
	{0, arm64.REG_R0, 0, "R0"},
	{1, arm64.REG_R1, 1, "R1"},
	{2, arm64.REG_R2, 2, "R2"},
	{3, arm64.REG_R3, 3, "R3"},
	{4, arm64.REG_R4, 4, "R4"},
	{5, arm64.REG_R5, 5, "R5"},
	{6, arm64.REG_R6, 6, "R6"},
	{7, arm64.REG_R7, 7, "R7"},
	{8, arm64.REG_R8, 8, "R8"},
	{9, arm64.REG_R9, 9, "R9"},
	{10, arm64.REG_R10, 10, "R10"},
	{11, arm64.REG_R11, 11, "R11"},
	{12, arm64.REG_R12, 12, "R12"},
	{13, arm64.REG_R13, 13, "R13"},
	{14, arm64.REG_R14, 14, "R14"},
	{15, arm64.REG_R15, 15, "R15"},
	{16, arm64.REG_R16, 16, "R16"},
	{17, arm64.REG_R17, 17, "R17"},
	{18, arm64.REG_R18, -1, "R18"},
	{19, arm64.REG_R19, 18, "R19"},
	{20, arm64.REG_R20, 19, "R20"},
	{21, arm64.REG_R21, 20, "R21"},
	{22, arm64.REG_R22, 21, "R22"},
	{23, arm64.REG_R23, 22, "R23"},
	{24, arm64.REG_R24, 23, "R24"},
	{25, arm64.REG_R25, 24, "R25"},
	{26, arm64.REG_R26, 25, "R26"},
	{27, arm64.REGG, -1, "g"},
	{28, arm64.REG_R29, -1, "R29"},
	{29, arm64.REG_R30, 26, "R30"},
	{30, arm64.REGSP, -1, "SP"},
	{31, arm64.REG_F0, -1, "F0"},
	{32, arm64.REG_F1, -1, "F1"},
	{33, arm64.REG_F2, -1, "F2"},
	{34, arm64.REG_F3, -1, "F3"},
	{35, arm64.REG_F4, -1, "F4"},
	{36, arm64.REG_F5, -1, "F5"},
	{37, arm64.REG_F6, -1, "F6"},
	{38, arm64.REG_F7, -1, "F7"},
	{39, arm64.REG_F8, -1, "F8"},
	{40, arm64.REG_F9, -1, "F9"},
	{41, arm64.REG_F10, -1, "F10"},
	{42, arm64.REG_F11, -1, "F11"},
	{43, arm64.REG_F12, -1, "F12"},
	{44, arm64.REG_F13, -1, "F13"},
	{45, arm64.REG_F14, -1, "F14"},
	{46, arm64.REG_F15, -1, "F15"},
	{47, arm64.REG_F16, -1, "F16"},
	{48, arm64.REG_F17, -1, "F17"},
	{49, arm64.REG_F18, -1, "F18"},
	{50, arm64.REG_F19, -1, "F19"},
	{51, arm64.REG_F20, -1, "F20"},
	{52, arm64.REG_F21, -1, "F21"},
	{53, arm64.REG_F22, -1, "F22"},
	{54, arm64.REG_F23, -1, "F23"},
	{55, arm64.REG_F24, -1, "F24"},
	{56, arm64.REG_F25, -1, "F25"},
	{57, arm64.REG_F26, -1, "F26"},
	{58, arm64.REG_F27, -1, "F27"},
	{59, arm64.REG_F28, -1, "F28"},
	{60, arm64.REG_F29, -1, "F29"},
	{61, arm64.REG_F30, -1, "F30"},
	{62, arm64.REG_F31, -1, "F31"},
	{63, 0, -1, "SB"},
}
var gpRegMaskARM64 = regMask(670826495)
var fpRegMaskARM64 = regMask(9223372034707292160)
var specialRegMaskARM64 = regMask(0)
var framepointerRegARM64 = int8(-1)
var linkRegARM64 = int8(29)
var registersMIPS = [...]Register{
	{0, mips.REG_R0, -1, "R0"},
	{1, mips.REG_R1, 0, "R1"},
	{2, mips.REG_R2, 1, "R2"},
	{3, mips.REG_R3, 2, "R3"},
	{4, mips.REG_R4, 3, "R4"},
	{5, mips.REG_R5, 4, "R5"},
	{6, mips.REG_R6, 5, "R6"},
	{7, mips.REG_R7, 6, "R7"},
	{8, mips.REG_R8, 7, "R8"},
	{9, mips.REG_R9, 8, "R9"},
	{10, mips.REG_R10, 9, "R10"},
	{11, mips.REG_R11, 10, "R11"},
	{12, mips.REG_R12, 11, "R12"},
	{13, mips.REG_R13, 12, "R13"},
	{14, mips.REG_R14, 13, "R14"},
	{15, mips.REG_R15, 14, "R15"},
	{16, mips.REG_R16, 15, "R16"},
	{17, mips.REG_R17, 16, "R17"},
	{18, mips.REG_R18, 17, "R18"},
	{19, mips.REG_R19, 18, "R19"},
	{20, mips.REG_R20, 19, "R20"},
	{21, mips.REG_R21, 20, "R21"},
	{22, mips.REG_R22, 21, "R22"},
	{23, mips.REG_R24, 22, "R24"},
	{24, mips.REG_R25, 23, "R25"},
	{25, mips.REG_R28, 24, "R28"},
	{26, mips.REGSP, -1, "SP"},
	{27, mips.REGG, -1, "g"},
	{28, mips.REG_R31, 25, "R31"},
	{29, mips.REG_F0, -1, "F0"},
	{30, mips.REG_F2, -1, "F2"},
	{31, mips.REG_F4, -1, "F4"},
	{32, mips.REG_F6, -1, "F6"},
	{33, mips.REG_F8, -1, "F8"},
	{34, mips.REG_F10, -1, "F10"},
	{35, mips.REG_F12, -1, "F12"},
	{36, mips.REG_F14, -1, "F14"},
	{37, mips.REG_F16, -1, "F16"},
	{38, mips.REG_F18, -1, "F18"},
	{39, mips.REG_F20, -1, "F20"},
	{40, mips.REG_F22, -1, "F22"},
	{41, mips.REG_F24, -1, "F24"},
	{42, mips.REG_F26, -1, "F26"},
	{43, mips.REG_F28, -1, "F28"},
	{44, mips.REG_F30, -1, "F30"},
	{45, mips.REG_HI, -1, "HI"},
	{46, mips.REG_LO, -1, "LO"},
	{47, 0, -1, "SB"},
}
var gpRegMaskMIPS = regMask(335544318)
var fpRegMaskMIPS = regMask(35183835217920)
var specialRegMaskMIPS = regMask(105553116266496)
var framepointerRegMIPS = int8(-1)
var linkRegMIPS = int8(28)
var registersMIPS64 = [...]Register{
	{0, mips.REG_R0, -1, "R0"},
	{1, mips.REG_R1, 0, "R1"},
	{2, mips.REG_R2, 1, "R2"},
	{3, mips.REG_R3, 2, "R3"},
	{4, mips.REG_R4, 3, "R4"},
	{5, mips.REG_R5, 4, "R5"},
	{6, mips.REG_R6, 5, "R6"},
	{7, mips.REG_R7, 6, "R7"},
	{8, mips.REG_R8, 7, "R8"},
	{9, mips.REG_R9, 8, "R9"},
	{10, mips.REG_R10, 9, "R10"},
	{11, mips.REG_R11, 10, "R11"},
	{12, mips.REG_R12, 11, "R12"},
	{13, mips.REG_R13, 12, "R13"},
	{14, mips.REG_R14, 13, "R14"},
	{15, mips.REG_R15, 14, "R15"},
	{16, mips.REG_R16, 15, "R16"},
	{17, mips.REG_R17, 16, "R17"},
	{18, mips.REG_R18, 17, "R18"},
	{19, mips.REG_R19, 18, "R19"},
	{20, mips.REG_R20, 19, "R20"},
	{21, mips.REG_R21, 20, "R21"},
	{22, mips.REG_R22, 21, "R22"},
	{23, mips.REG_R24, 22, "R24"},
	{24, mips.REG_R25, 23, "R25"},
	{25, mips.REGSP, -1, "SP"},
	{26, mips.REGG, -1, "g"},
	{27, mips.REG_R31, 24, "R31"},
	{28, mips.REG_F0, -1, "F0"},
	{29, mips.REG_F1, -1, "F1"},
	{30, mips.REG_F2, -1, "F2"},
	{31, mips.REG_F3, -1, "F3"},
	{32, mips.REG_F4, -1, "F4"},
	{33, mips.REG_F5, -1, "F5"},
	{34, mips.REG_F6, -1, "F6"},
	{35, mips.REG_F7, -1, "F7"},
	{36, mips.REG_F8, -1, "F8"},
	{37, mips.REG_F9, -1, "F9"},
	{38, mips.REG_F10, -1, "F10"},
	{39, mips.REG_F11, -1, "F11"},
	{40, mips.REG_F12, -1, "F12"},
	{41, mips.REG_F13, -1, "F13"},
	{42, mips.REG_F14, -1, "F14"},
	{43, mips.REG_F15, -1, "F15"},
	{44, mips.REG_F16, -1, "F16"},
	{45, mips.REG_F17, -1, "F17"},
	{46, mips.REG_F18, -1, "F18"},
	{47, mips.REG_F19, -1, "F19"},
	{48, mips.REG_F20, -1, "F20"},
	{49, mips.REG_F21, -1, "F21"},
	{50, mips.REG_F22, -1, "F22"},
	{51, mips.REG_F23, -1, "F23"},
	{52, mips.REG_F24, -1, "F24"},
	{53, mips.REG_F25, -1, "F25"},
	{54, mips.REG_F26, -1, "F26"},
	{55, mips.REG_F27, -1, "F27"},
	{56, mips.REG_F28, -1, "F28"},
	{57, mips.REG_F29, -1, "F29"},
	{58, mips.REG_F30, -1, "F30"},
	{59, mips.REG_F31, -1, "F31"},
	{60, mips.REG_HI, -1, "HI"},
	{61, mips.REG_LO, -1, "LO"},
	{62, 0, -1, "SB"},
}
var gpRegMaskMIPS64 = regMask(167772158)
var fpRegMaskMIPS64 = regMask(1152921504338411520)
var specialRegMaskMIPS64 = regMask(3458764513820540928)
var framepointerRegMIPS64 = int8(-1)
var linkRegMIPS64 = int8(27)
var registersPPC64 = [...]Register{
	{0, ppc64.REG_R0, -1, "R0"},
	{1, ppc64.REGSP, -1, "SP"},
	{2, 0, -1, "SB"},
	{3, ppc64.REG_R3, 0, "R3"},
	{4, ppc64.REG_R4, 1, "R4"},
	{5, ppc64.REG_R5, 2, "R5"},
	{6, ppc64.REG_R6, 3, "R6"},
	{7, ppc64.REG_R7, 4, "R7"},
	{8, ppc64.REG_R8, 5, "R8"},
	{9, ppc64.REG_R9, 6, "R9"},
	{10, ppc64.REG_R10, 7, "R10"},
	{11, ppc64.REG_R11, 8, "R11"},
	{12, ppc64.REG_R12, 9, "R12"},
	{13, ppc64.REG_R13, -1, "R13"},
	{14, ppc64.REG_R14, 10, "R14"},
	{15, ppc64.REG_R15, 11, "R15"},
	{16, ppc64.REG_R16, 12, "R16"},
	{17, ppc64.REG_R17, 13, "R17"},
	{18, ppc64.REG_R18, 14, "R18"},
	{19, ppc64.REG_R19, 15, "R19"},
	{20, ppc64.REG_R20, 16, "R20"},
	{21, ppc64.REG_R21, 17, "R21"},
	{22, ppc64.REG_R22, 18, "R22"},
	{23, ppc64.REG_R23, 19, "R23"},
	{24, ppc64.REG_R24, 20, "R24"},
	{25, ppc64.REG_R25, 21, "R25"},
	{26, ppc64.REG_R26, 22, "R26"},
	{27, ppc64.REG_R27, 23, "R27"},
	{28, ppc64.REG_R28, 24, "R28"},
	{29, ppc64.REG_R29, 25, "R29"},
	{30, ppc64.REGG, -1, "g"},
	{31, ppc64.REG_R31, -1, "R31"},
	{32, ppc64.REG_F0, -1, "F0"},
	{33, ppc64.REG_F1, -1, "F1"},
	{34, ppc64.REG_F2, -1, "F2"},
	{35, ppc64.REG_F3, -1, "F3"},
	{36, ppc64.REG_F4, -1, "F4"},
	{37, ppc64.REG_F5, -1, "F5"},
	{38, ppc64.REG_F6, -1, "F6"},
	{39, ppc64.REG_F7, -1, "F7"},
	{40, ppc64.REG_F8, -1, "F8"},
	{41, ppc64.REG_F9, -1, "F9"},
	{42, ppc64.REG_F10, -1, "F10"},
	{43, ppc64.REG_F11, -1, "F11"},
	{44, ppc64.REG_F12, -1, "F12"},
	{45, ppc64.REG_F13, -1, "F13"},
	{46, ppc64.REG_F14, -1, "F14"},
	{47, ppc64.REG_F15, -1, "F15"},
	{48, ppc64.REG_F16, -1, "F16"},
	{49, ppc64.REG_F17, -1, "F17"},
	{50, ppc64.REG_F18, -1, "F18"},
	{51, ppc64.REG_F19, -1, "F19"},
	{52, ppc64.REG_F20, -1, "F20"},
	{53, ppc64.REG_F21, -1, "F21"},
	{54, ppc64.REG_F22, -1, "F22"},
	{55, ppc64.REG_F23, -1, "F23"},
	{56, ppc64.REG_F24, -1, "F24"},
	{57, ppc64.REG_F25, -1, "F25"},
	{58, ppc64.REG_F26, -1, "F26"},
	{59, ppc64.REG_F27, -1, "F27"},
	{60, ppc64.REG_F28, -1, "F28"},
	{61, ppc64.REG_F29, -1, "F29"},
	{62, ppc64.REG_F30, -1, "F30"},
	{63, ppc64.REG_F31, -1, "F31"},
}
var gpRegMaskPPC64 = regMask(1073733624)
var fpRegMaskPPC64 = regMask(576460743713488896)
var specialRegMaskPPC64 = regMask(0)
var framepointerRegPPC64 = int8(1)
var linkRegPPC64 = int8(-1)
var registersS390X = [...]Register{
	{0, s390x.REG_R0, 0, "R0"},
	{1, s390x.REG_R1, 1, "R1"},
	{2, s390x.REG_R2, 2, "R2"},
	{3, s390x.REG_R3, 3, "R3"},
	{4, s390x.REG_R4, 4, "R4"},
	{5, s390x.REG_R5, 5, "R5"},
	{6, s390x.REG_R6, 6, "R6"},
	{7, s390x.REG_R7, 7, "R7"},
	{8, s390x.REG_R8, 8, "R8"},
	{9, s390x.REG_R9, 9, "R9"},
	{10, s390x.REG_R10, -1, "R10"},
	{11, s390x.REG_R11, 10, "R11"},
	{12, s390x.REG_R12, 11, "R12"},
	{13, s390x.REGG, -1, "g"},
	{14, s390x.REG_R14, 12, "R14"},
	{15, s390x.REGSP, -1, "SP"},
	{16, s390x.REG_F0, -1, "F0"},
	{17, s390x.REG_F1, -1, "F1"},
	{18, s390x.REG_F2, -1, "F2"},
	{19, s390x.REG_F3, -1, "F3"},
	{20, s390x.REG_F4, -1, "F4"},
	{21, s390x.REG_F5, -1, "F5"},
	{22, s390x.REG_F6, -1, "F6"},
	{23, s390x.REG_F7, -1, "F7"},
	{24, s390x.REG_F8, -1, "F8"},
	{25, s390x.REG_F9, -1, "F9"},
	{26, s390x.REG_F10, -1, "F10"},
	{27, s390x.REG_F11, -1, "F11"},
	{28, s390x.REG_F12, -1, "F12"},
	{29, s390x.REG_F13, -1, "F13"},
	{30, s390x.REG_F14, -1, "F14"},
	{31, s390x.REG_F15, -1, "F15"},
	{32, 0, -1, "SB"},
}
var gpRegMaskS390X = regMask(23551)
var fpRegMaskS390X = regMask(4294901760)
var specialRegMaskS390X = regMask(0)
var framepointerRegS390X = int8(-1)
var linkRegS390X = int8(14)
var registersWasm = [...]Register{
	{0, wasm.REG_R0, 0, "R0"},
	{1, wasm.REG_R1, 1, "R1"},
	{2, wasm.REG_R2, 2, "R2"},
	{3, wasm.REG_R3, 3, "R3"},
	{4, wasm.REG_R4, 4, "R4"},
	{5, wasm.REG_R5, 5, "R5"},
	{6, wasm.REG_R6, 6, "R6"},
	{7, wasm.REG_R7, 7, "R7"},
	{8, wasm.REG_R8, 8, "R8"},
	{9, wasm.REG_R9, 9, "R9"},
	{10, wasm.REG_R10, 10, "R10"},
	{11, wasm.REG_R11, 11, "R11"},
	{12, wasm.REG_R12, 12, "R12"},
	{13, wasm.REG_R13, 13, "R13"},
	{14, wasm.REG_R14, 14, "R14"},
	{15, wasm.REG_R15, 15, "R15"},
	{16, wasm.REG_F0, -1, "F0"},
	{17, wasm.REG_F1, -1, "F1"},
	{18, wasm.REG_F2, -1, "F2"},
	{19, wasm.REG_F3, -1, "F3"},
	{20, wasm.REG_F4, -1, "F4"},
	{21, wasm.REG_F5, -1, "F5"},
	{22, wasm.REG_F6, -1, "F6"},
	{23, wasm.REG_F7, -1, "F7"},
	{24, wasm.REG_F8, -1, "F8"},
	{25, wasm.REG_F9, -1, "F9"},
	{26, wasm.REG_F10, -1, "F10"},
	{27, wasm.REG_F11, -1, "F11"},
	{28, wasm.REG_F12, -1, "F12"},
	{29, wasm.REG_F13, -1, "F13"},
	{30, wasm.REG_F14, -1, "F14"},
	{31, wasm.REG_F15, -1, "F15"},
	{32, wasm.REGSP, -1, "SP"},
	{33, wasm.REGG, -1, "g"},
	{34, 0, -1, "SB"},
}
var gpRegMaskWasm = regMask(65535)
var fpRegMaskWasm = regMask(4294901760)
var specialRegMaskWasm = regMask(0)
var framepointerRegWasm = int8(-1)
var linkRegWasm = int8(-1)
