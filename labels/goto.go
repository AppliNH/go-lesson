package labels

import "fmt"

func ExecLesson2() {

	var i int

	loop:
		if i < 2 {
			fmt.Println("Whoop")
			i++
			goto loop
		}
		fmt.Println("BLOODY BEETROOTS")
}