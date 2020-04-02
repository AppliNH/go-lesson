package arrayslice

import "fmt"

func ExecLesson2() {


	msg := []string{"hello","human","this","is","cool"}
	human := msg[1:2]

	fmt.Println(human)

	//You can also create a slice from a array
	//That's thanks to the fact that slices use "Backing Array" :) You should google it, that's interresting

	fixedArray := [5]int{1,2,3,4,5}
	fmt.Println(fixedArray)

	slice := fixedArray[:5] //We could also have used slice := fixedArray[0:5]
	//"fixedArray" and "slice" now have the same backing array 

	// --> A slice stores its elements in a backing array that the slice references : a slice doesn't properly contains elements :)
	// --> When you slice a slice, it returns a new slice value with the SAME backing array
	// --> Backing arrays are contiguous in memory

	slice = append(slice,6)

	slice2 := slice[3:4] //slice2 and slice use the same backing array

	fmt.Println(slice)
	
	fmt.Println(slice2)

}