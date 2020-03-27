package _types

import "fmt"

type MyType int64 //Underlying type of MyType is int64
//every Go type has an underlying type


func ExecLesson1() {

	fmt.Println("Please open ./_types/underlying-types.go to see some interesting stuff.")

	fmt.Println("Here I create my own type, which I call 'MyType'.")
	fmt.Println("Underlying type of MyType is int64.")
	fmt.Println("every Go type has an underlying type.")
	fmt.Println("A type can also use himself as its underlying type.")
	fmt.Println("You can see that with the type int64. ")
	
	fmt.Println("But defining a type this way doesn't do anything special. It's juste int64 with another name.")
	fmt.Println("This only give to MyType the representativity and the size of int64, its underlying type.")
	fmt.Println("In order to change the game, we need to attach a structure to this.")
	var ms int64 = 1000
	var ns MyType

	ns = MyType(ms)
	ms = int64(ns)

	//but ns = ms or ms = ns won't work !
	// however newVar := ns * ms is fine !


}