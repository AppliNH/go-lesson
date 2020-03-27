package _switch

import "fmt"

func ExecLesson() {

	i := 5

	switch { //switch true { ... }
	case i > 0:
		fmt.Println(" i > 0") //if you fall in this case, Go wille execute the statements and exit the switch.
		// You don't need to write the break here !
	case i < 0:
		fmt.Println("i < 0")
	default:
		fmt.Println(" i = 0")

	}

}
