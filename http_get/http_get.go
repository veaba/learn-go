package http_get

import (
	"fmt"
	"net/http"
)

func main() {
	//http.Handle("",http.FileServer(http.Dir(".")))
	//http.ListenAndServe("0.0.0.0:8080", nil)
	getHttpRes2()
}

func getHttpRes2() {
	res, err := http.Get("http://www.baidu.com")
	fmt.Println(res)
	fmt.Println(err)
}
