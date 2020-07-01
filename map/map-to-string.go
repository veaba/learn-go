package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]string{"a": "aa", "b": "bb"}
	mJsonStr, _ := json.Marshal(m)
	fmt.Println(string(mJsonStr))
}
