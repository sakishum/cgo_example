package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lfoo
#include "foo.h"
*/
import "C"

func main() {
	C.foo()
}
