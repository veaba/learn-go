package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main() {
	var str = "access_token=xxx&expires_in=28800&refresh_token=r1.xxxx&refresh_token_expires_in=15638400&scope=&token_type=bearer"

	q, _ := url.ParseQuery(str)
	fmt.Println("====>", q)

	a, _ := json.Marshal(q) // a is byte[]
	fmt.Println("json byte=>", a)
	fmt.Println("json string=>", string(a))

	fmt.Println("json get property=>", json.RawMessage(a)) // byte
	// fmt.Println("json get property=>", json.Unmarshal(a)) // byte
	// 步骤 url 解析=> map => type =>json

}
