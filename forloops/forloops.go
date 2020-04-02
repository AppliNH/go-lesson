package forloops

import "fmt"

func ExecLesson() {

	myDict := make(map[string]int)
	myDict["wheels"] = 4
	myDict["doors"] = 5
	myDict["people"] = 5

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	stringArr := []string{"a", "b", "c"}


	for index, value := range stringArr {
		fmt.Println(index, " : ", value)
	}
	
	//WARNING !
	//Go won't execute a range loop if the given slice is nil !
	//But the program won't crash

	for key, value := range myDict {
		fmt.Println(key, " : ", value)
	}


	//for { ... } is an infinite loop
	//for [condition] { ... } is a while
	// for { ... break } allows to break out of the "for" loop
}
