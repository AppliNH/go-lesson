package rune_string_byte

import (
	"fmt"
	"strings"
)

func ExecLesson1(startLetter rune, stopLetter rune) {
	// rune = 1 byte
	// utf8 encoded rune = 2 bytes
	start, stop := startLetter,stopLetter

	fmt.Printf("%-10s %-10s %-10s %-14s \n%s\n","letter","dec","hex","encoded",strings.Repeat("-",50))

	for n:=start; n <=stop;n++ {
		
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x \n",n,string(n)) //%c prints a character depending on the give code point
	}

}