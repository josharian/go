// +build go1.7

package gc

import (
	"fmt"
	"os"
	"runtime"
)

func whence() {
	if os.Getenv("J") == "" {
		return
	}
	pc := make([]uintptr, 2)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	var frame runtime.Frame
	x := 2
	for more := true; more && n != 0; {
		frame, more = frames.Next()
		fmt.Print(frame.Function, " ")
		x--
		if x == 0 {
			break
		}
	}
	fmt.Println()
}
