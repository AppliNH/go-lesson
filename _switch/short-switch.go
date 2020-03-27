package _switch

import "fmt"

func ExecLesson3() {

	switch i := 5; true { //don't forget the semicolon ; ! 
	case i > 0:
		fmt.Println(" i > 0")
	case i < 0:
		fmt.Println("i < 0")
	default:
		fmt.Println(" i = 0")

	}

}
