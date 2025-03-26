/*
 获得指定的行的字符串
*/
package main

import (
	"fmt"
	"strings"
)

const content = "\n## 使用插件\n\n在使用 `createApp()` 初始化 Vue 应用程序后，你可以通过调用 `use()` 方法将插件添加到你的应用程序中。\n\n我们将使用在[编写插件](#编写插件)部分中创建的 `i18nPlugin` 进行演示。\n\n`use()` 方法有两个参数。第一个是要安装的插件，在这种情况下为 `i18nPlugin`。\n\n它还会自动阻止你多次使用同一插件，因此在同一插件上多次调用只会安装一次该插件。\n\n第二个参数是可选的，并且取决于每个特定的插件。在演示 `i18nPlugin` 的情况下，它是带有转换后的字符串的对象。\n"

func speRaw() {
	// return
}
func main() {
	ret := strings.Split(content, "\n")
	// fmt.Println("name:", content)
	fmt.Println("len=>", len(ret))
	fmt.Println("ret=>:", ret[11])
	// TODO  如果指定 11 行
	// 查找 第 10个 /n  —— 第 12个 /n 并返回内容
	for i, item := range ret {
		fmt.Println(i, "==>", item)
	}

}
