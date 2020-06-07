// go 实现对象数组的返回

func main()  {
	var rbody []map[string]interface{}
	t := make(map[string]interface{})
	t["hahda"] = "hahdas"
	t["bbbb"] = "bsbds"
	rbody = append(rbody, t)

	
	//cnnJson := make(map[string]interface{})
	//cnnJson["body"]=rbody
	a, _ := json.Marshal(rbody)
	println("&&&&=>", string(a))

}