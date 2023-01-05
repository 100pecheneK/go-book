package main

// #cgo CFLAGS: -I${SRCDIR}/callClib
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Going to call a C function!")
	C.helloFromC()
	fmt.Println("Going to call another C function with arguments")
	message := C.CString("This is Misha")
	defer C.free(unsafe.Pointer(message))
	C.printMessage(message)
	fmt.Println("All perfectly done!")
}
