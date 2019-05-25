/**
@desc map 一种无序的键值对的集合，通过key来快速检索数据，key 索引，指向数据的值
@map hash表来实现的
@如果不初始化，则会创建一个nil map，nil map 无法存放键值对
*/
package main

import "fmt"

func main()  {
	var aMap  map[string]string //创建集合
	fmt.Println(aMap)//map[]

	aMap=make(map[string]string)
	fmt.Println(aMap)//map[] //map[]

	aMap["2019"]="MSI2018，季中冠军赛"
	aMap["2018"]="屌丝之年"
	aMap["2020"]="希望之年"
	fmt.Println(aMap)

	//for k,v:=range aMap{
	//	fmt.Println(k)//key
	//	fmt.Println(aMap[k])//value
	//	fmt.Println(v)//value
	//}

	//删除集合 delete函数
	_delete(aMap)

	fmt.Println(aMap)

	checkMap(aMap)


}

func checkMap(strings map[string]string) {
	_,ok := strings["2009"]
	if ok{
		fmt.Println("有值")
	}
}


func _delete(aMap map[string]string) {
	fmt.Println(aMap)
	delete(aMap,"2018")
	fmt.Println(aMap)

	for k,v:=range aMap{
		fmt.Println(k,v)
	}
}

