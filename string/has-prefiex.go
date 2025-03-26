package main

import (
	"fmt"
	"strings"
)

func main() {
	var i = "docs/README.me"
	var j = "/docs/README.me"

	fmt.Println(strings.HasPrefix(i, "/"))
	fmt.Println(strings.HasPrefix(j, "/"))

}
