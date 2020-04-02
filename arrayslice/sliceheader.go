package arrayslice

import (
	"fmt"
	"reflect"
	"unsafe"
)

type collection []string

func ExecLesson3() {

	data := collection{"book1", "book2", "book3"}

	fmt.Println(data)
	change(data)
	fmt.Println(data)

	lost := data[1:2] //lost slice won't be able to access previous elements
	// BTW you CAN'T extend it back using lost := lost[0:3]
	otherSlice := data[0:1] //otherSlice is able to access next elements

	fmt.Printf("Variable address : %p \n", data) //same

	fmt.Printf("Data slice capacity : %v \n", cap(data))

	fmt.Printf("Lost slice capacity : %v \n", cap(lost)) //..so it affects its capacity

	fmt.Printf("otherSlice slice capacity : %v \n", cap(otherSlice))

	//fmt.Printf("lost variable address : %p \n",lost)

	fmt.Println(lost)

	change2(data)

	fmt.Println(lost) //Here you can see that changes on the slice "data" also happen on the slice "lost"
	// That's because they have the same backing array :)

	// A slice header has a fixed size. It will never change, even if you have a billion of elements
	// Slice operations are cheap
	// cause passing a slice to a function only passes its header
	// but don't forget that a slice also creates a backing array :)
	// but it won't be passed to a function. Just the slice header (length, capacity, pointer)
	// Using "append" creates a new slice header, with a larger capacity
	fmt.Println("___________________")
	fmt.Println("In general : ")
	fmt.Printf("Slice data : %+v \n", (*reflect.SliceHeader)(unsafe.Pointer(&data)))
	fmt.Printf("Slice lost : %+v \n", (*reflect.SliceHeader)(unsafe.Pointer(&lost)))
	fmt.Printf("Slice otherSlice : %+v \n", (*reflect.SliceHeader)(unsafe.Pointer(&otherSlice)))

	fmt.Println("___________________")

	nums := []int{}
	nums = append(nums, 1, 3, 2, 4)
	fmt.Println("nums : ", nums)
	fmt.Printf("Slice nums : %+v \n", (*reflect.SliceHeader)(unsafe.Pointer(&nums)))

	nums = append(nums,nums[2:]...) //add nums[2:] to nums, so it duplicates 2 and 4
	// ++ when you use append, the capacity doubles if needed
	// so here, because we have now 6 elements, the capacity of the backing array goes from 4 to 8
	fmt.Println("nums : ", nums)
	fmt.Printf("Slice nums : %+v \n", (*reflect.SliceHeader)(unsafe.Pointer(&nums)))

	nums = append(nums[:2], 7, 9) //overwrites 2 and 4, but keeps the capacity
	// even if it's an append, it doesn't increase the capacity, because we overwrite
	fmt.Println("nums : ", nums)
	fmt.Printf("Slice nums : %+v \n", (*reflect.SliceHeader)(unsafe.Pointer(&nums)))

	nums = nums[0:8] //made possible because the capacity = 8
	//In order to just add the hidden other values we could have done :
	// --> nums = append(nums, nums[len(nums)+1:cap(nums)]...)
	// However, there's 6 elemets in the slice, the two other blocks are equal to 0
	fmt.Println("nums : ", nums)
	fmt.Printf("Slice nums : %+v \n", (*reflect.SliceHeader)(unsafe.Pointer(&nums)))

	//Let's replace these zeros by numbers, cause I don't like zeros
	nums = append(nums[:6],4,9)
	fmt.Println("nums : ", nums)
	fmt.Printf("Slice nums : %+v \n", (*reflect.SliceHeader)(unsafe.Pointer(&nums)))

	//Let's just add one more
	nums = append(nums, 1)
	fmt.Println("nums : ", nums)
	fmt.Printf("Slice nums : %+v \n", (*reflect.SliceHeader)(unsafe.Pointer(&nums))) // capacity doubles each time it's full


}

func change(_data collection) {
	_data[1] = "Woops I lost it"
	fmt.Printf("Function's parameter address : %p \n", _data) //same
}
func change2(_data collection) {
	_data[1] = "Found it !"

}
