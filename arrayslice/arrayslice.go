package arrayslice

import "fmt"

func ExecLesson() {
	var fixedSizeArr [5]int
	secondArr := [5]int{1, 2, 3, 4, 5}

	sliceArr := []int{}
	sliceArr = append(sliceArr, 1)

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
