package themap

import "fmt"

func ExecLesson() {

	// A map is quite similar to Python's dict

	myDict := make(map[string]int)
	var mySecondDict map[string]string
	myDict["wheels"] = 4
	myDict["doors"] = 5
	myDict["people"] = 5
	delete(myDict, "people")

	fmt.Println(myDict)

	// Detect if a key exists in a map :
	key := "doors"

	if value, ok := myDict[key]; ok {
		fmt.Println(key, " exists in myDict, its value is ", value)
	} else {
		fmt.Println(key, " doesn't exist in myDict")
	}

	// You can compare maps with Sprintf if maps are [string]string

	myStringDict := map[string]string {"ok":"ok","notok":"meh"}
	mySecondDict = myStringDict

	first := fmt.Sprintf("%s",myStringDict)
	second := fmt.Sprintf("%s",mySecondDict)

	if first == second {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	// When you create a map, Go creates a Map Header
	// So a map variable is only a pointer to a map header value
	// And the map header value points out to the map data
	// The map header has a lot of properties.. you should check it out


	// Destroy a map
	for k := range myDict {
		delete(myDict,k)
	}

	// Btw, looping through map keys in order to get a element is not a good idea
	// If you wanna do that, you should use a slice.

	// Maps allow direct access to a value

	// The value can be of any type, but the key should be int, bool, string, rune => simple types
	// so you can do a map[string]map[int][bool]
}
