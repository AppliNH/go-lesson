package forloops

func ExecLesson2() {

	i := 0
	sum := 0

	for {
		if i > 10 {
			break
		}
		if i %2 != 0 {
			i++
			continue // goes back to the beginning of the for loop
			// continue statement only can continue a loop !
		}
		sum += i
		i++

	}

}
