package rune_string_byte

import (
	"fmt"
	"unsafe"
)

type StringHeader struct {
	pointer uintptr
	length  int
}

func ExecLesson4() {

	// A string is a slice of bytes
	// A slice refers to a backing array thanks to a structure
	// So, a string is also a structure.

	hello := "hello"
	//these two will refer to the same backing array
	showStringHeader(hello)
	showStringHeader("hello")

	//but this one will not
	showStringHeader("hello!")

	//neither this one
	showStringHeader("hel")

	// This quite uses the same backing array as hello
	// You'll see the pointer value is  slightly different but
	// That's because it only takes the "o" of the backing array
	// So the address of the "o" is different from the address of the backing array begining

	// Behind the curtains, ctrings are being shared :)
	showStringHeader(hello[4:5])

	// A string is read-only
	fmt.Println()

	//Have a look here :
	for i := range hello {
		showStringHeader(hello[i : i+1])
	}
	fmt.Println()
	//slicing a string is cheap and efficient. Because it doesn't re-create a backing array.
	// However, a conversion uses a different backing array

	showStringHeader(string([]byte(hello)))
	showStringHeader(string([]byte(hello)))

	// ==> Do not convert string, unless it's really needed..
}

func showStringHeader(s string) {

	ptr := *(*StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%q: %+v \n", s, ptr)
}
