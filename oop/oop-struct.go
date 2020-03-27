package oop

import "fmt"

type car struct {
	wheels   int
	doors     int
	bigTrunk bool
	brand    string
}

func ExecLesson() {
	myCar := car{doors: 5, wheels: 4, bigTrunk: true}
	fmt.Println("My car has ", myCar.doors, " doors !")
}
