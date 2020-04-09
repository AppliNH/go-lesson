package main

// please run go install module_name before if you're not in your go workspace (defined by GOPATH env variable)
// export function must start with capital letter
// pkg name = folder name
import (
	"fmt"
	"os"
	"strings"

	_iota "primitivo.com/applinh/learngo/_iota"
	_json "primitivo.com/applinh/learngo/_json"
	_switch "primitivo.com/applinh/learngo/_switch"
	typ "primitivo.com/applinh/learngo/_types"
	array_slice "primitivo.com/applinh/learngo/arrayslice"
	for_loops "primitivo.com/applinh/learngo/forloops"
	functions "primitivo.com/applinh/learngo/functions"
	goRoutines "primitivo.com/applinh/learngo/goroutines"
	ifelse "primitivo.com/applinh/learngo/ifelse"
	labels "primitivo.com/applinh/learngo/labels"
	me "primitivo.com/applinh/learngo/me"
	miniProjects "primitivo.com/applinh/learngo/miniProjects"
	oop "primitivo.com/applinh/learngo/oop"
	pointers "primitivo.com/applinh/learngo/pointers"
	rune_string_byte "primitivo.com/applinh/learngo/rune_string_byte"
	scanner "primitivo.com/applinh/learngo/scanner"
	shadowing "primitivo.com/applinh/learngo/shadowing"
	_map "primitivo.com/applinh/learngo/themap"
	vars "primitivo.com/applinh/learngo/variables"
)

/*
main_function
This goes to go doc
*/
func main() {

	fmt.Println(me.Say_hi())

	if a := os.Args; len(a) >= 2 {
		fmt.Println("You choose : ", a[1])
		switch a[1] {
		//oh btw, an OR multi case is
		/* case "blabla","blibli":
		do things */
		case "pointers":
			pointers.ExecLesson()
		case "json":
			_json.ExecLesson()
			fmt.Println(strings.Repeat("_", 25))
			_json.ExecLesson2() //covid Tracker
		case "scanner":
			scanner.ExecLesson()
		case "rune_string_byte":
			rune_string_byte.ExecLesson1(9984,10175)
			fmt.Println(strings.Repeat("_", 25))
			rune_string_byte.ExecLesson2()
			fmt.Println(strings.Repeat("_", 25))
			rune_string_byte.ExecLesson3()
			fmt.Println(strings.Repeat("_", 25))
			rune_string_byte.ExecLesson4()
		case "array_slice":
			array_slice.ExecLesson()
			fmt.Println(strings.Repeat("_", 25))
			array_slice.ExecLesson2()
			fmt.Println(strings.Repeat("_", 25))
			array_slice.ExecLesson3()
			fmt.Println(strings.Repeat("_", 25))
			array_slice.ExecLesson4()
			fmt.Println(strings.Repeat("_", 25))
			array_slice.ExecLesson5()
			fmt.Println(strings.Repeat("_", 25))
			array_slice.ExecLesson6()
		case "for_loops":
			for_loops.ExecLesson()
			fmt.Println(strings.Repeat("_", 25))
			for_loops.ExecLesson2()
			fmt.Println(strings.Repeat("_", 25))
			for_loops.ExecLesson3()
			fmt.Println(strings.Repeat("_", 25))
		case "functions":
			functions.ExecLesson1()
			fmt.Println(strings.Repeat("_", 25))
			functions.ExecLesson2()
			fmt.Println(strings.Repeat("_", 25))
			functions.ExecLesson3()
		case "ifelse":
			ifelse.ExecLesson()
			fmt.Println(strings.Repeat("_", 25))
			ifelse.ExecLesson2()
		case "_map":
			_map.ExecLesson()
		case "oop":
			oop.ExecLesson()
			fmt.Println(strings.Repeat("_", 25))
			oop.ExecLesson2()
			fmt.Println(strings.Repeat("_", 25))
			oop.ExecLesson3()
		case "goRoutines":
			goRoutines.ExecLesson()
			fmt.Println(strings.Repeat("_", 25))
			goRoutines.ExecLesson2()
			fmt.Println(strings.Repeat("_", 25))
			goRoutines.ExecLesson3()
			fmt.Println(strings.Repeat("_", 25))
			goRoutines.ExecLesson4()
			fmt.Println(strings.Repeat("_", 25))
			goRoutines.ExecLesson5()
		case "variables":
			vars.ExecLesson()
			fmt.Println(strings.Repeat("_", 25))
			vars.ExecLesson2()
		case "_types":
			typ.ExecLesson1()
			fmt.Println("Have a look at ./_types/custom-const.go too !")
		case "_iota":
			_iota.ExecLesson1()
		case "shadowing":
			shadowing.ExecLesson()
		case "_switch":
			_switch.ExecLesson()
			fmt.Println(strings.Repeat("_", 25))
			_switch.ExecLesson2()
		case "mini_projects":
			if len(a) >= 3  {
				switch a[2] {
				case "MultiTable":
					miniProjects.MultiTable(10)
				case "WordFinder":
					miniProjects.WordFinder("Ok that's cool dude !", "cool")
				case "RetroClock":
					miniProjects.RetroClock()
				case "FileFinder":
					miniProjects.FileFinderMain(a[3])
				case "BouncingBall":
					miniProjects.BouncingBall(10,10)
				case "SpamMask":
					miniProjects.SpamMask()
				case "GoGrep":
					if (len(a) == 4) {
						miniProjects.GoGrep(a[3])
					} else {
						fmt.Println("I need a word to search in the poem")
					}
				default:
					fmt.Println("No project found for ", a[2])
				}
			} else {
				fmt.Println("Please, pick a project to execute")
			}
		case "labels":
			labels.ExecLesson()
			fmt.Println(strings.Repeat("_", 25))
			labels.ExecLesson2()
			fmt.Println(strings.Repeat("_", 25))
			labels.ExecLesson3()
		default:
			fmt.Println("Please pass as a correct argument in the command line")
		}
	} else {
		fmt.Println("You gotta choose a topic ! Check out ./main.go :)")
	}

}
