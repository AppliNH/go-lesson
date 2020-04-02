package arrayslice

import (
	"fmt"
	"reflect"
	"unsafe"
)

func ExecLesson4() {
	fmt.Println("_____________________")
	sliceable := []string{"f","u","l","l"}

	fmt.Println(sliceable)
	fmt.Printf("Slice sliceable : %+v \n", (*reflect.SliceHeader)(unsafe.Pointer(&sliceable)))

	sliceable = sliceable[0:3:3] //sets capacity to 3
	// [low:high:max]` => `length = high - max` and `capacity = max - low`
	// the last "l" will be hidden and not accessible anymore
	// sliceable[:3:3] also works

	fmt.Println(sliceable)
	fmt.Printf("Slice sliceable : %+v \n", (*reflect.SliceHeader)(unsafe.Pointer(&sliceable)))

	// THIS will cause an error :)
	// sliceable = sliceable[0:4]
	// fmt.Println(sliceable)
	// fmt.Printf("Slice sliceable : %+v \n", (*reflect.SliceHeader)(unsafe.Pointer(&sliceable)))



}