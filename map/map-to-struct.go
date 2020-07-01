// map 转 struct
package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// 定义结构
	type ArrayObjStruct struct {
		Name string
		Job  string
	}
	// arrayObj map
	arrayObj := map[string]string{
		"name": "Lisi",
		"job":  "老师",
	}

	var arrayObjStruct ArrayObjStruct

	// 1 先将map 转为字符串
	arrayObjJsonStr, _ := json.Marshal(arrayObj)
	fmt.Println("第一步map转json 字符串=========>\n", string(arrayObjJsonStr))

	// 2 第二步，再将字符串对接到struct
	json.Unmarshal(arrayObjJsonStr, &arrayObjStruct)
	fmt.Println("第二步，再将字符串对接到struct==>\n", arrayObjJsonStr) // 这是一个buf
	fmt.Println("Name=======================>\n", arrayObjStruct.Name)
	fmt.Println("Job========================>\n", arrayObjStruct.Job)
}
