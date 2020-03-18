package goroutines

import (
	"fmt"
)

func fibonacci4(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func ExecLesson4() {
	c := make(chan int, 10)
	go fibonacci4(cap(c), c)
	for i := range c { //will receive values from channel c until the channel closes
		fmt.Println(i)
	}
}
