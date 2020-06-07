// struct 引用
package main

import (
	"fmt"
)
func main(){

	type A struct{
		Name string 
	}
	
	var a =A{}
	a.Name="jasjdjda"

	temp :=a.Name
	
	fmt.Println("1temp==>",temp)
	a.Name="dsadda"

	fmt.Println("2temp==>",temp)

	fmt.Println("==>",a)
}