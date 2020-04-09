package scanner

import (
	"bufio"
	"fmt"
	"os"
)

func ExecLesson() {
	fmt.Println("Say something : ")

	scannit := bufio.NewScanner(os.Stdin)
	var lines int

	for scannit.Scan() { // get out of the loop = nothing to read ==> equals false
		
		fmt.Println(lines,":", scannit.Text())
		lines ++
	}

	if err := scannit.Err() ;err !=nil {
		fmt.Println("ERROR : ",err)
	}

}
