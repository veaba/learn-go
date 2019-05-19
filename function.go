package main

import "fmt"

// (x string ,y string) 缩写-> (x,y string)
func max1(x,y string)(string,string)  {
	return x,y
}
func main()  {
	b,a:=max1("Good","bad")
	fmt.Println(a,b)
}
