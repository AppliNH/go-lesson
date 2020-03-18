package goroutines

import "fmt"

func ExecLesson3() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//I can't write ch <- 89 because my channel buffer size is 2 !
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
