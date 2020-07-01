package main

import (
	"fmt"
	"net/url"
)

func main() {
	var str = "a=b&c=d"
	// TODO 如何解析上面为json字符串
	query, err := url.ParseQuery(str)
	if err != nil {
		panic(err)
	}
	fmt.Println(query)
}
