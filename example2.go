package main

/*
#include <stdlib.h>

const char* thus(const char* s) {
	return s;
}
*/
import "C"

import "unsafe"

func main() {
	mC := C.CString("Hello World!")
	defer C.free(unsafe.Pointer(mC))
	mTGo := C.GoString(C.thus(mC))
	print(mTGo)
}
