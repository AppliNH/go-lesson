package goroutines

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func ExecLesson2() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) //create a channel
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println("First goRoutine sum : ", x)
	fmt.Println("Second goRoutine sum : ", y)
	fmt.Println("Total : ", x+y)
}
