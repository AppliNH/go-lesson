package variables

import "fmt"

func ExecLesson2() {
	//DEFERENCING
	lyrics := "Moments so dear"
	pointerForStr := &lyrics

	*pointerForStr = "Journeys to plan"

	fmt.Println(lyrics) // Prints: Journeys to plan
}
