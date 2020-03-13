package functions

import (
	"fmt"
)

func inc(x *int) {
	*x++
}

func ExecLesson3() {

	i := 0
	inc(&i)
	fmt.Println(i)
}
