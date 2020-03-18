package goroutines

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func power(number int, exposant int) {
	var result int = number
	for j := 0; j < exposant; j++ {
		result = result * number
		time.Sleep(100 * time.Millisecond)
		fmt.Println(result)
	}
	fmt.Println("result : ", result)
}
func ExecLesson() {
	go power(2, 5)
	say("world")
}
