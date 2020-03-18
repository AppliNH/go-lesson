package main

// please run go install module_name before if you're not in your go workspace (defined by GOPATH env variable)
// export function must start with capital letter
// pkg name = folder name
import (
	"fmt"

	array_slice "primitivo.com/applinh/learngo/arrayslice"
	for_loops "primitivo.com/applinh/learngo/forloops"
	functions "primitivo.com/applinh/learngo/functions"
	goRoutines "primitivo.com/applinh/learngo/goroutines"
	ifelse "primitivo.com/applinh/learngo/ifelse"
	me "primitivo.com/applinh/learngo/me"
	oop "primitivo.com/applinh/learngo/oop"
	_map "primitivo.com/applinh/learngo/themap"
)


func main() {

	fmt.Println(me.Say_hi())
	array_slice.ExecLesson()
	for_loops.ExecLesson()
	functions.ExecLesson1()
	functions.ExecLesson2()
	functions.ExecLesson3()
	ifelse.ExecLesson()
	_map.ExecLesson()
	oop.ExecLesson()
	goRoutines.ExecLesson()
	goRoutines.ExecLesson2()
	goRoutines.ExecLesson3()
	goRoutines.ExecLesson4()
	goRoutines.ExecLesson5()
}






