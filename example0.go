package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func main() {
	message := C.CString("Hello World!\n")
	defer C.free(unsafe.Pointer(message)) // free 函数定义在 stdlib.h 内
	C.puts(message)                       // puts 函数定义在 stdio.h 内
}
