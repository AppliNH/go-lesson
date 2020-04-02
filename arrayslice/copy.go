package arrayslice

import (
	"fmt"
	"reflect"
	"unsafe"
)

func ExecLesson6() {

	fmt.Println("_________________________")
	mySlice := []int{1, 2, 3, 4}

	destSlice := []int{0, 0, 0, 0, 5, 6}

	copy(destSlice, mySlice)
	fmt.Println("destSlice : ",destSlice)
	fmt.Printf("Slice destSlice : %+v \n", (*reflect.SliceHeader)(unsafe.Pointer(&destSlice)))


	// WARNING !
	// The copy function overwrites the dest slice based on index =>

	destSlice2 := []int{3,4}
	// You could use make() to dynamically handle this stuff
	copy(destSlice2,mySlice)
	fmt.Println("destSlice2 : ",destSlice2)
	fmt.Printf("Slice destSlice2 : %+v \n", (*reflect.SliceHeader)(unsafe.Pointer(&destSlice2)))


}
