package arrayslice

import "fmt"

func ExecLesson5() {

	// The make function allows to pre-allocate a backing array of a slice with the given length and capacity
	// However, these two parameters won't be fixed ! You still can dynamically appen elements

	wow := make([]int, 4) // length & capacity = 4
	// because the length is 4, it will be initialized with [0,0,0,0]
	// which is not very handy btw bc if you add 2 and 4, it will be [0,0,0,0,2,4]

	fmt.Println(wow)
	// Why is it useful ?
	// Well, when using the appen function, increasing the capacity of the backing array (by in fact creating another one)
	// is an expensive operation.
	// Here, if you already pre-allocate a backing array which is large enough, you prevent the creation of a larger backing array
	// for the next statements.

	// You can use it in order to optimize your programm :)
	// Define the capacity wisely !

	// You can also do :
	hey := make([]int, 0, 4) //length = 0 and capacity = 4
	// because the length is 0, it won't be initialized with values !
	fmt.Println(hey)

}
