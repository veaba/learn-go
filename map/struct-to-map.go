package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Object struct {
		X int
		Y int
	}
	m := make(map[string]interface{})
	i := Object{11, 33}
	fmt.Println("i==>", i) // i==> {11 33}

	j, _ := json.Marshal(i)
	fmt.Println("j==>", j) // j==> [123 34 88 34 58 49 49 44 34 89 34 58 51 51 125]

	_ = json.Unmarshal(j, &m)
	fmt.Println("m==>", m) // m==> map[X:11 Y:33]
	fmt.Println("i==>", i) // i==> {11 33}
	fmt.Println("j==>", j) //j==> [123 34 88 34 58 49 49 44 34 89 34 58 51 51 125]
	fmt.Println(string(j)) //buf =>= string {"X":11,"Y":33}

}
