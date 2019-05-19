/**

@desc go 建立 http 服务器 错误的例子

*/
package main

import "net/http"

func main()  {
	http.Handle("",http.FileServer(http.Dir(".")))
	http.ListenAndServe("0.0.0.0:8080", nil)
}
