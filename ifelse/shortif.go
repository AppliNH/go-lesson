package ifelse

import (
	"fmt"
	"strconv"
)

func ExecLesson2() {

	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Printf("No error, n is %v and its type is now %T", n,n)
	} else {
		fmt.Println("An error has occured : ", err)
	}

}
