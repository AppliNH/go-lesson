package _iota

import "fmt"

func ExecLesson1() {

	// a constant can repeat the type and the expression of the previous constant
	// iota increases each time, starts with 0
	const (
		zero = iota
		one
		two
		three
		four
		five
	)
	//it restarts from the begining for each new scope
	const (
		Mezero = iota 	//0
		_				//1
		Metwo 			//2
		_				//3
		Mefour 			//4
	) //with blank identifier, we can make the IOTA increase

	fmt.Println(one, two, three, four, five, Mezero, Metwo, Mefour)

}
