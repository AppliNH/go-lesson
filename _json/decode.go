package _json

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type numbers struct {
	Confirmed int `json:"confirmed"`
	Deaths    int `json:"deaths"`
	Recovered int `json:"recovered"`
}

type covidApiInfo struct {
	Count  int                `json:"count"`
	Result map[string]numbers `json:"result"` // or map[string] map[string]int
}

func ExecLesson2() {
	var input []byte

	var covid19Results covidApiInfo

	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	if err := json.Unmarshal(input, &covid19Results); err != nil {
		fmt.Println(err)
		return
	}
	date := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	fmt.Println(strings.Repeat("_", 40))
	fmt.Println("Covid19 Stats for your country at ", date)
	fmt.Println(strings.Repeat("_", 40))
	fmt.Println("Confirmed : ", covid19Results.Result[date].Confirmed)
	fmt.Println("Deaths : ", covid19Results.Result[date].Deaths)
	fmt.Println("Recovered : ", covid19Results.Result[date].Recovered)
	fmt.Println(strings.Repeat("_", 40))
	fmt.Println("Please stay at home")
}
