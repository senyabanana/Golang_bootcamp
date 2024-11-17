package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "application.h"
#include "window.h"
*/
import "C"
import "unsafe"

func main() {
	C.InitApplication()

	title := C.CString("School 21")
	defer C.free(unsafe.Pointer(title))

	wndPtr := C.Window_Create(100, 100, 300, 200, title)

	C.Window_MakeKeyAndOrderFront(wndPtr)

	C.RunApplication()
}
