package main

import (
	"encoding/json"
	"fmt"
)

type Name struct {
	Job string
}
func main()  {
	method1()
	method2()

}

/*方法一*/
func method1(){
	name :=Name{Job:"web前端"}

	fmt.Println("name:",name)
	buf,err:=json.MarshalIndent(name,"","  ")
	if err!=nil{
		return
	}
	fmt.Println("方法一：",string(buf))
}
/*方法二*/
func method2()  {
	name :=Name{"Go后端开发"}
	jsonData,err:=json.Marshal(name)
	if err!=nil{
		return
	}
	fmt.Println("方法二：",string(jsonData))
}
