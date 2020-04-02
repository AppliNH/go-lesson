package rune_string_byte

import "fmt"

func ExecLesson2() {

	str := "Et voilà, confinés 🤒" //string litterals are automatically utf8 encoded

	for i, r := range str {
		fmt.Printf("str[%2d] = % -12x = %q \n", i, string(r), r)
		// non english characters use more than one byte
	}

	// Only 1 byte characters are easily indexable in a slice
	// But 2 bytes characters are not

	// Note that => len("é") = 2 bc it counts the bytes
	// but utf8.Runecount("é") = 1 bc it counts the runes

	fmt.Println("==> à character is not : ", string(str[7]), " but here it is : ", string(str[7:9]))

	// You can convert a string to a rune slice tho
	// But it has an efficiency cost

	runes := []rune(str)

	for _, v := range runes {
		fmt.Printf("%c \n", v)
	}

}
