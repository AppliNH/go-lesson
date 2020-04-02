package miniProjects

import (
	"fmt"
	"strings"
)

func hideLink(s string, http bool) string {

	switch http {
	case true:
		withoutHttp := s[7:]
		fmt.Println(withoutHttp)

		hidden := strings.Repeat("*", len(withoutHttp))

		return "http://" + hidden
	case false:
		withoutHttps := s[8:]
		fmt.Println(withoutHttps)

		hidden := strings.Repeat("*", len(withoutHttps))

		return "https://" + hidden

	}
	return ""

}

func SpamMask() {

	fmt.Println()

	text := "Hey this is a spam hahah https://google.fr you filthy guy ! Aufait, je suis français 🥖"
	words := strings.Fields(text)

	for i, v := range words {
		if strings.Contains(v, "http://") {
			words[i] = hideLink(words[i], true)
		} else if strings.Contains(v, "https://") {
			words[i] = hideLink(words[i], false)
		}
	}
	text = strings.Join(words, " ")

	fmt.Println(text)

}
