package miniProjects

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func GoGrep(query string) {

	rx := regexp.MustCompile(`[^a-z]+`) // if the regexp is incorrect, program won't compile

	scannit := bufio.NewScanner(os.Stdin)
	scannit.Split(bufio.ScanWords)

	words := make(map[string]bool)

	for scannit.Scan() {
		word := strings.ToLower(scannit.Text())
		word = rx.ReplaceAllString(word, "") // deletes punctuation and numbers
		//fmt.Println(word)
		words[word] = true
	}

	if err := scannit.Err(); err != nil {
		fmt.Println("ERROR : ", err)
	}

	if words[query] {
		fmt.Println("The text contains ", query)
		return
	}

	fmt.Println("The text doesn't contain ", query)

}
