package main

import (
	"fmt"
	"html/template"
)

func main() {
	var templateContent = template.Must(template.New("test").Parse(`
	<h1> Hello world html template! </h1>	
`))
	fmt.Println(templateContent)
}
