package arrayslice

import "fmt"

func ExecLesson() {
	var fixedSizeArr [5]int //Its length is fixed, you can't add more than 5 elements. It creates [0,0,0,0,0]
	//You add elements using "fixedSizeArr[0] = 4"
	secondArr := [5]int{1, 2, 3, 4, 5}

	sliceArr := []int{} //Its length is dynamic
	sliceArr = append(sliceArr, 1) //You add elements using this

	//slice := []int creates a slice which value is nil, it's a nil slice
	//slice := []int{} creates an empty slice, a no-nil slice


	// Recommended : prefer declaring a nil slice instead of empty slice, it will save memory

	//However, in both cases, len(slice) will work, and return 0

	//You CANNNOT add elements using "sliceArr[0] = 4"
	//But you can init the slice with "sliceArr := []int{ 2, 4 }"

	//You can't compare slices and arrays. Because they don't have the same length
	//unless the slice's length is fixed, which would turn it into an array :)

	//A slice can only be compared to nil

	//You can compare two slices by using a loop
	

	fmt.Println(fixedSizeArr)
	fmt.Println(secondArr)
	fmt.Println("sliceArr : ", sliceArr)

	//Caution !
	//You can test the equality between two arrays
	// but :
	// - Arrays' size need to be equal
	// - elements types need to the same
	// - elements need to be equal and in the same order
	// Ex [3,9,6] == [3,9,6] but [3,9,6] == [3,6,9]. Order matters !
}
