// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains a check for -0.0.

package main

import (
	"go/ast"
	"go/constant"
	"go/token"
)

func init() {
	register("minuszero",
		"check for -0.0",
		checkMinusZero,
		unaryExpr)
}

func checkMinusZero(f *File, n ast.Node) {
	u := n.(*ast.UnaryExpr)
	if u.Op != token.SUB {
		return
	}
	v := f.pkg.types[u.X].Value
	if v == nil || v.Kind() != constant.Float || v.String() != "0" {
		return
	}
	f.Badf(u.Pos(), "minus zero: %s", f.gofmt(u))
}
