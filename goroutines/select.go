package goroutines

import (
	"fmt"
	"time"
)

func fibonacci5(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
			case c <- x: // If I can send to c
				x, y = y, x+y
			case <-quit: //If I can receive from quit
				fmt.Println("quit")
				return
		}
	}
}

func bomb() {
	tick := time.Tick(100 * time.Millisecond) //pass current date to channel "tick" every 100ms
	boom := time.After(500 * time.Millisecond) //pass current date to channel "boom" after 500ms
	for {
		select {
			case <-tick: //If something comes from "tick"
				fmt.Println("tick.")
			case <-boom: //If something comes from "boom"
				fmt.Println("BOOM!")
				return
			default: // default
				fmt.Println("    .")
				time.Sleep(50 * time.Millisecond)
		}
	}
}

func ExecLesson5() {
	bomb()
	c := make(chan int)
	quit := make(chan int)

	go func() {

		for i := 0; i < 10; i++ {
			fmt.Println(<-c) //reads from c
		}
		quit <- 0 //kills the process

	}()

	fibonacci5(c, quit)
}
