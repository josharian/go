package gc

import (
	"bufio"
	"cmd/internal/obj"
	"cmd/internal/obj/x86"
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()

	arch := &x86.Linkamd64
	Ctxt = obj.Linknew(arch)
	Ctxt.Bso = bufio.NewWriter(os.Stdout)

	Widthint = arch.IntSize
	Widthptr = arch.PtrSize
	Widthreg = arch.RegSize

	initUniverse()

	os.Exit(m.Run())
}
