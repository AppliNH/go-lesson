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
}
