package _json

import (
	"encoding/json"
	"fmt"
)

type AdditionnalInfos map[string]string

type car struct {
	Wheels   int `json:"wheels"` // Properties need to start with a capital letter to be exported and used by json.Marshal 
	Doors    int `json:"doors"` // renames json field. It's also called a field tag
	BigTrunk bool `json:"bigTrunk"`
	Brand    string `json:"brand"`
	secretcodeOMG int // won't be displayed by json. You can also put `json:"-"`
	AdditionnalInfos `json:"other,omitempty"` // omitempty = omitted by json.Marshal if empty
}

func ExecLesson() {

	myCar1 := car{Doors: 5, Wheels: 4,secretcodeOMG:1, BigTrunk: true, Brand: "Ford", AdditionnalInfos: AdditionnalInfos{"engine": "Diesel"}}
	myCar2 := car{Doors: 5, Wheels: 4,secretcodeOMG:2, BigTrunk: true, Brand: "Tesla", AdditionnalInfos: AdditionnalInfos{"engine": "Electric", "superCool": "true"}}
	myCar3 := car{Doors: 3, Wheels: 4,secretcodeOMG:3, BigTrunk: false, Brand: "Renault", AdditionnalInfos: AdditionnalInfos{"engine": "Diesel"}}

	myCars := []car{myCar1, myCar2, myCar3}

	out, _ := json.Marshal(myCars)
	out, _ = json.MarshalIndent(myCars, "", "\t") //You don't need the above if you use this one

	fmt.Println(string(out))

}
