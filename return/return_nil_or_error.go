package main

import "fmt"

func main()  {
	a,b :=returnWhat()
	fmt.Println(a,b)
}
type Values map[string][]string

func returnWhat() (string, error)  {
	query:= ""
	if len(query) > 1 {
		return query, nil
	}
	var err error
	return query, err
}

