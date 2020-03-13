package oop

import "fmt"

type car struct {
	wheels   int
	door     int
	bigTrunk bool
	brand    string
}

func ExecLesson() {
	myCar := car{door: 4, wheels: 4, bigTrunk: true}
	fmt.Println("My car has ", myCar.door, " doors !")
}
