package miniProjects

import "fmt"

func MultiTable(max int) {

	fmt.Printf("%5s", "X")

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i) //%5d handles the espacement
	}
	fmt.Println()

	for j := 0; j <= max; j++ {

		fmt.Printf("%5d", j)
		for k := 0; k <= max; k++ {
			fmt.Printf("%5d", k*j)
		}
		fmt.Println()

	}

}
