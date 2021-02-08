package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

func main4() {
	// str := "GET https://api.github.com/repos/veaba/vuepress-plugin-editable/git/ref/heads/: 401 Bad credentials []"

	str := "PUT https://api.github.com/repos/veaba/vuepress-plugin-editable/contents//docs/README.md: 422 path cannot start with a slash [{Resource:Commit Field:path Code:invalid Message:}]"

	// TODO  access token失效
	// TODO 路径错误的情况
	reg := regexp.MustCompile(`: (?s:(.*?)) `)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}
	//提取关键信息
	// result := reg.FindAllStringSubmatch(str, -1)
	//过滤<></>
	// for _, text := range result {
	// 	fmt.Println("text[1] = ", text[1])
	// }
	// fmt.Println("ooo=>", result[0][0])
	// fmt.Println(len(result[0][0]))

	x := getGithubAPICode(str)
	a, b := strconv.ParseInt("401", 10, 64)

	fmt.Println(" 字符转数字==> ", a, b)

	fmt.Println("x===>", x)
}

func getGithubAPICode(str string) int64 {
	reg := regexp.MustCompile(`: (?s:(.*?)) `)
	if reg == nil {
		fmt.Println("MustCompile err")
		return 0
	}
	result := reg.FindAllStringSubmatch(str, -1)

	if len(result) == 1 {
		intStr := result[0]
		fmt.Println("intStr=>", intStr, intStr[0], intStr[1])
		ret, _ := strconv.ParseInt(intStr[1], 10, 64) // string 转 int 失败

		fmt.Println("ret=>", ret)
		fmt.Println("typeof=>", typeof(ret))    // int64
		fmt.Println("intStr=>", typeof(intStr)) // string
		return ret
	}
	return 1
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
