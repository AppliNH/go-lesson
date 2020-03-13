package themap

import "fmt"

func ExecLesson() {
	myDict := make(map[string]int)
	myDict["wheels"] = 4
	myDict["doors"] = 5
	myDict["people"] = 5
	delete(myDict, "people")

	fmt.Println(myDict)
}
