package variables

import "fmt"

func ExecLesson() {
	var number int = 0
	otherNumber := 5
	const myConstant = "ok this last one isn't a variable"
	//btw declarations here doesn't need to be used/called for the whole stuff to compile without erros :)

	fmt.Println(number, otherNumber)
}
