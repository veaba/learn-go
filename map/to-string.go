package main

import "fmt"

func main() {

	data := map[string]string{
		"access_token": "111-accessToken",
		"status":       "222-status",
	}

	fmt.Println("====>", data)
	// fmt.Println("string====>", string(data))
	fmt.Println("====>", data["status"])
	data["status"] = "wahahah"
	fmt.Println("====>", data["status"])
	println(data["status"])
}
