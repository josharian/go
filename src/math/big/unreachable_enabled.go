// +build unreachable

package big

import "unsafe"

func unsafeUnreachable() {
	unsafe.Unreachable()
}
