package functions

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("You gave me a negative number")
	}
	return math.Sqrt(x), nil
}

func ExecLesson2() {

	result, err := sqrt(16)

	if(err !=nil) {
		fmt.Println(result)
	} else{
		fmt.Println(err)
	}
}
