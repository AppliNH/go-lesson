package labels

import "fmt"

func someFunc(label int) {

	switch label {
	case 1:
		goto step1
	case 2:
		goto step2
	default:
		fmt.Println("This isn't a valid step dude")
	}

step1:
	fmt.Println("Gosh this is step 1")

step2:
	fmt.Println("Gosh this is step 2")

}

func ExecLesson3() {

	someFunc(2)

}
