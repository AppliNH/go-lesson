package main

// please run go install module_name before if you're not in your go workspace (defined by GOPATH env variable)
// export function must start with capital letter
// pkg name = folder name
import (
	"fmt"
	"os"

	_iota "primitivo.com/applinh/learngo/_iota"
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
		case "array_slice":
			array_slice.ExecLesson()
		case "for_loops":
			for_loops.ExecLesson()
			for_loops.ExecLesson2()
			for_loops.ExecLesson3()
		case "functions":
			functions.ExecLesson1()
			functions.ExecLesson2()
			functions.ExecLesson3()
		case "ifelse":
			ifelse.ExecLesson()
			ifelse.ExecLesson2()
		case "_map":
			_map.ExecLesson()
		case "oop":
			oop.ExecLesson()
		case "goRoutines":
			goRoutines.ExecLesson()
			goRoutines.ExecLesson2()
			goRoutines.ExecLesson3()
			goRoutines.ExecLesson4()
			goRoutines.ExecLesson5()
		case "variables":
			vars.ExecLesson()
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
				default:
					fmt.Println("No project found for ", a[2])
				}
			} else {
				fmt.Println("Please, pick a project to execute")
			}
		case "labels":
			labels.ExecLesson()
			labels.ExecLesson2()
			labels.ExecLesson3()
		default:
			fmt.Println("Please pass as a correct argument in the command line")
		}
	} else {
		fmt.Println("You gotta choose a topic ! Check out ./main.go :)")
	}

}
