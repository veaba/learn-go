package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	arrayObject := map[string]string{
		"高一":  "高一",
		"高二":  "高二",
		"高三":  "高三",
		"初一":  "初一",
		"初二":  "初二",
		"初三":  "初三",
		"四年级": "小四",
		"小四":  "小四",
		"五年级": "小五",
		"小五":  "小五",
		"六年级": "小六",
		"小六":  "小六",
		"一年级": "小一",
		"二年级": "小一",
		"三年级": "小一",
	}
	arrayObject["dd"] = "ddd"
	arrayObjectData, _ := json.Marshal(arrayObject)
	println("string ===>", string(arrayObjectData))
	println("byte ===>", arrayObjectData)

	input := [][]string{
		[]string{"b", "3", "abc", "5.3"},
		[]string{"a", "4", "efg", "9.1"},
		[]string{"b", "4", "abc", "5.3"},
		[]string{"c", "3", "hij", "5.5"},
		[]string{"a", "2", "abc", "9.2"},
	}

	input1, _ := json.Marshal(input)

	fmt.Println("input==>", input)
	fmt.Println("input1==>", input1)
}
