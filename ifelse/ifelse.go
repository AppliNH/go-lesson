package ifelse

import "fmt"

func ExecLesson() {
	var number int = 0
	otherNumber := 5

	if otherNumber > number {
		fmt.Println("otherNumber > number")

	} else {
		fmt.Println("otherNumber !> number")
	}

}
