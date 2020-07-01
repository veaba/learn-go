package main

import (
	"encoding/json"
	"fmt"
)

type School struct {
	Job  string
	Name string
}

func main() {
	school := School{"老师", "LiLei"}
	schoolStr, _ := json.Marshal(school)
	fmt.Println("struct to string <buf> ==>", schoolStr) //buf

	fmt.Println("struct to string <string> ==>", string(schoolStr))
}
