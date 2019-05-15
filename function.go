package main

import "fmt"

func max1(x,y string)(string,string)  {
	return x,y
}
func main()  {
	b,a:=max1("Good","bad")
	fmt.Println(a,b)
}
