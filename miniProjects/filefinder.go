package miniProjects

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func menu(path string) {
	var userSelection int
	fmt.Println()
	fmt.Println("What do you want to do here ?")
	fmt.Println(" 1 - Write to a file")
	fmt.Println(" 2 - Read a file")

	fmt.Scanln(&userSelection)

	switch userSelection {
	case 1:
		writeFile(path)
	case 2:
		readFile(path)
	}

}

func FileFinderMain(path string) {
	fmt.Println()

	if path == "" {
		fmt.Println("Hey I need a path please, just add it as a last argument")
	} else {

		// var icon_go_file rune = '🦠'
		// var icon_file rune = '📄'
		// var icon_folder rune = '📁'

		fmt.Println("_______ FILE FINDER ______________")
		fmt.Println("Directory : ", path)
		fmt.Println()
		if files, err := ioutil.ReadDir(path); err == nil {
			for _, f := range files {
				if f.IsDir() {
					fmt.Println("DIR", " | ", f.Name(), " - ", f.Size(), " Bytes")
				} else if strings.Contains(f.Name(), ".go") {
					fmt.Println("GO ", " | ", f.Name(), " - ", f.Size(), " Bytes")
				} else {
					fmt.Println("     | ", f.Name(), " - ", f.Size(), " Bytes")
				}
			}
		}
		menu(path)

	}

}

func readFile(path string) {
	var fileName string
	fmt.Println("Name of the file, in the current path : ")
	fmt.Scanln(&fileName)

	data, err := ioutil.ReadFile(path + fileName)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Content of ",fileName," : ")
	fmt.Println(string(data))

}

func writeFile(path string) {
	var fileName string
	var content *bufio.Reader
	var line []rune
	fmt.Println("Full name of your file : ")
	fmt.Scanln(&fileName)
	fmt.Println("Content of your file : ")

	content = bufio.NewReader(os.Stdin)

	for {
		c, _, err := content.ReadRune()
		if err == io.EOF || c == '\n' {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading standard input\n")
			os.Exit(1)
		}
		line = append(line, c)
	}

	if err := ioutil.WriteFile(path+fileName, []byte(string(line[:len(line)])), 0644); err == nil {
		fmt.Println("Done :)")

	} else {
		fmt.Println("Error :(")
		fmt.Println(err)
	}

}
