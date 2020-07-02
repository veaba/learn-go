/**

@desc go 建立 http 服务器
@todo 获取取得IP等req数据

*/
package http_TODO

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct {
	w      string
	Accept string
}

func (hello Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a, b := fmt.Fprint(w, "hello world!")
	fmt.Println(a, b)
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(r)

}

func main() {
	var h Hello
	h.w = "wwwww"
	err := http.ListenAndServe("localhost:8888", h)
	if err != nil {
		log.Fatal(err)
	}
}
