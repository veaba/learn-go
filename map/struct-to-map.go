package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	type Object struct {
		X int
		Y int
	}
	m := make(map[string]interface{})
	i:=Object{11,33}
	fmt.Println("i==>",i)
	
	j,_:=json.Marshal(i)
	fmt.Println("j==>",j)

	_ = json.Unmarshal(j, &m)
	fmt.Println("m==>",m)
	fmt.Println("i==>",i)
	fmt.Println("j==>",j)
}