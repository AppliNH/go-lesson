package shadowing

import (
	"fmt"
	"strconv"
)

func ExecLesson() {

	fmt.Println("Here's an example of shadowing")
	var a int

	if a,err:= strconv.Atoi("97"); err == nil{ 
		fmt.Println("Everything when fine without errors, look at 'a' : ",a)
		//a is only equal to this value in this context, because of ":="
	}

	fmt.Println("Let's check out how our 'a' variable is doing : ",a)
	fmt.Println("HOW MY GOD IT GOT SHADOWED CAN YOU BELIEVE IT !")

	fmt.Println("Open ./shadowing/shadowing.go please.")
	
	
	//"FIX"
	var (
		n int 
		_err error
	)

	if n,_err = strconv.Atoi("97"); _err == nil{ 
		fmt.Println("Everything when fine without errors, look at 'n' : ",n)
	}
	fmt.Println("Let's check out how our 'n' variable is doing : ",n)





}