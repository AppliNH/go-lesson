package pointers

import "fmt"

func ExecLesson() {

	var number int = 22

	fmt.Println("number's memory address : ",&number)

	pnumber := &number // pointer value, its type is *int

	fmt.Println(*pnumber," looks, here's our variable :)")
	clonedNumber := *pnumber // clonedNumber stores a copy of the variable value
	
	fmt.Println(clonedNumber)

	var list []string
	addToList(&list)
	fmt.Println(list)

	// Oh btw you cannot address Map's elements, they are unadressable
}

func addToList(list *[]string) {
	*list = append(*list, "yo")
}