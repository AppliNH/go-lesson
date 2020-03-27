package _switch

import "fmt"

func ExecLesson2() {

	i := 142

	switch {
	case i > 100:
		fmt.Println("big")
		fallthrough //allow to execute the next statement in the nearest case block, without checking the case condition
	case i > 0:
		fmt.Println("positive")
	default:
		fmt.Println(i)

	}
}
