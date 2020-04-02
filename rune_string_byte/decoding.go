package rune_string_byte

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func ExecLesson3() {

	poem := `Au-dessus des étangs, au-dessus des vallées,
	Des montagnes, des bois, des nuages, des mers,
	Par delà le soleil, par delà les ésthers,
	Par delà les confins des sphères étoilées,
	
	Mon esprit, tu te meus avec agilité,
	Et, comme un bon nageur qui se pâme dans l'onde,
	Tu sillonnes gayement l'immensité profonde
	Avec une indicible et mâle volupté.`

	for i := 0; i < len(poem); {
		r, size := utf8.DecodeRuneInString(poem[i:])

		fmt.Printf("%c", r)
		i += size
	}
	fmt.Println()
	fmt.Println()

	//or automatic
	for _, r := range poem {
		fmt.Printf("%c", r)
	}

	// Oh btw, a for range loop loops over the runes of a string :)

	fmt.Println()
	fmt.Println()

	poemByte := []byte("été, automne, hiver")
	fmt.Printf("%s = % [1]x\n", poemByte)
	

	// THIS IS COSTFUL
	// var size int
	// for i := range string(poemByte) {
	// 	if i > 0 {
	// 		size = i
	// 		break
	// 	}
	// }
	// copy(poemByte[:size], bytes.ToUpper(poemByte[:size]))
	// fmt.Printf("%s = % [1]x \n", poemByte)

	// This is better
	_, size := utf8.DecodeRune(poemByte)
	copy(poemByte[:size], bytes.ToUpper(poemByte[:size]))
	fmt.Printf("%s = % [1]x \n", poemByte)
}
