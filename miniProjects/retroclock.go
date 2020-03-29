package miniProjects

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clockNumbers = map[int][5]string{
	0: [5]string{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	},
	1: [5]string{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	},
	2: [5]string{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	},
	3: [5]string{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	},
	4: [5]string{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	},
	5: [5]string{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	},
	6: [5]string{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	},
	7: [5]string{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	},
	8: [5]string{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	},
	9: [5]string{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}}

var separator = [5]string{

	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}


func clearScreen() {

	switch myOS := runtime.GOOS ; myOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") 
        cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
	}

}

func RetroClock() {

	for {
		fmt.Println("\f")
		hours, minutes, seconds := time.Now().Local().Clock()
		
		rClock := [][5]string{
			clockNumbers[hours/10],clockNumbers[hours%10],
			separator,
			clockNumbers[minutes/10],clockNumbers[minutes%10],
			separator,
			clockNumbers[seconds/10],clockNumbers[seconds%10],
		}
	

		for a := range rClock[0] {
			for b := range rClock {
				fmt.Print(rClock[b][a], "    ")
			}
			fmt.Println()
		}

		fmt.Println()

		time.Sleep(1 * time.Second)
		clearScreen()
	}
}
