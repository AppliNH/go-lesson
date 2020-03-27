package miniProjects

import (
	"fmt"
	"strings"
)

func WordFinder(sentence string, word string) {

	fmt.Println("Welcome !")
	fmt.Println("Your sentence is : ",sentence)
	fmt.Println("Your word is : ",word)

	var corpus string = strings.ToLower(sentence)
	var wordToFind string = strings.ToLower(word)

	var index = 0

	words := strings.Fields(corpus)
	
	for i, v := range words {
		index = i
		if v == wordToFind {
			fmt.Println("Position is : ",index)
			break
		}
		if i+1 == len(words) {
			fmt.Println("Word not found")
			break
		}
	}
	

}
