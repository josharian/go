// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gofuzz

package fuzzcompile

import (
	"cmd/compile/internal/amd64"
	"cmd/compile/internal/arm"
	"cmd/compile/internal/arm64"
	"cmd/compile/internal/gc"
	"cmd/compile/internal/mips"
	"cmd/compile/internal/mips64"
	"cmd/compile/internal/ppc64"
	"cmd/compile/internal/s390x"
	"cmd/compile/internal/x86"
	"cmd/internal/objabi"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var archInits = map[string]func(*gc.Arch){
	"386":      x86.Init,
	"amd64":    amd64.Init,
	"amd64p32": amd64.Init,
	"arm":      arm.Init,
	"arm64":    arm64.Init,
	"mips":     mips.Init,
	"mipsle":   mips.Init,
	"mips64":   mips64.Init,
	"mips64le": mips64.Init,
	"ppc64":    ppc64.Init,
	"ppc64le":  ppc64.Init,
	"s390x":    s390x.Init,
}

func Fuzz(data []byte) (rc int) {
	f, err := ioutil.TempFile("", "compile")
	if err != nil {
		panic(err)
	}
	if _, err := f.Write(data); err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}

	archInit, ok := archInits[objabi.GOARCH]
	if !ok {
		fmt.Fprintf(os.Stderr, "compile: unknown architecture %q\n", objabi.GOARCH)
		os.Exit(2)
	}

	gc.InputFilename = f.Name()

	defer func() {
		err := recover()
		s, ok := err.(string)
		if ok && s == "controlled exit" {
			rc = 0
		}
		os.Remove(f.Name())
		os.Remove(filepath.Base(f.Name()) + ".o")
	}()
	gc.Main(archInit)
	return 1
}
