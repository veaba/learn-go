package main
package main

import (
"github.com/kataras/iris/v12"
)

/**
* 先声明结构体
 */
type User struct {
	Name string `json:name`
	Job  string `json:"job"`
}

func main() {
	app := iris.New()
	app.Handle("GET", "/", func(ctx iris.Context) {
		//ctx.WriteString("Hello world")
		user := &User{Name: "hah"}
		ctx.JSON(user)
	})
	app.Run(iris.Addr(":8081"))
}

