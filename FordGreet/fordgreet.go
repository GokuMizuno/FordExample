//Gordon Stangler
//cgo example, compiled with gcc, with CFLAGS
package main

//#cgo CFLAGS:  -g -Wall
//#include <stdlib.h>
//#include "FordGreet.h"

import "C"
import(
	"fmt"
	"unsafe"
)

func main() {
	name:= C.CString("Ford team")
	defer C.free(unsafe.Pointer(name))

	year := C.int(2019)

	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))

	size := C.greet(name, year, (*C.char)(ptr))

	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))
}
