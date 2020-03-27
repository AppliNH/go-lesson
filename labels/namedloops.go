package labels

import (
	"fmt"
	"time"
)


func ExecLesson() {
	now := time.Now().Local().Unix()
	endTime := time.Now().Local().Add(time.Second * time.Duration(4)).Unix()
	fmt.Println(time.Unix(now,0), time.Unix(endTime,0))
	fmt.Printf("[")

	// you can also name an if

	waiting:
		for {

			if time.Now().Local().Unix() >= endTime {
				fmt.Print("] Done !")
				break waiting //you can also do continue "labelName"
			}
			fmt.Printf("=")
			
			time.Sleep(100 * time.Millisecond)
	
			//btw if you wanna do a loading bar like this, you better have a look at ./goroutines/select.go

		}

}
