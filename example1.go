package main

/*
#include <stdio.h>
#include <stdlib.h>

static void cprintf(const char* msg) {
    printf(msg);
}
*/
import "C"
import "unsafe"

func main() {
	message := C.CString("Hello World!\n")
	defer C.free(unsafe.Pointer(message))
	C.cprintf(message)
}
