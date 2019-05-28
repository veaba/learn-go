/**

@desc go 建立 http 服务器

*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	goServer()
}

func goServer() {
	http.HandleFunc("/", httpServer)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func httpServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	_, _ = io.WriteString(w, "hello,world11\n")
}

/*
这串数据，怎么解析？

&{GET /favicon.ico HTTP/1.1 1 1 map[Accept:[image/webp,image/apng,image/,/*;q=0.8] Accept-Encoding:[gzip, deflate, br] Accept-Language:[zh-CN,zh;q=0.9,en;q=0.8] Cache-Control:[no-cache] Connection:[keep-alive] Pragma:[no-cache] Referer:[http://127.0.0.1:8888/] User-Agent:[Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36]] {} <nil> 0 [] false 127.0.0.1:8888 map[] map[] <nil> map[] 127.0.0.1:49668 /favicon.ico <nil> <nil> <nil> 0xc00005c500}

*/
