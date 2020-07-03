## 前言
> go，代码组织像c语言，语法和一些特性 像 js

## TODO
> https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/06.4.md
> https://github.com/adonovan/gopl.io Go程序设计语言源码
## go命令

- go fmt 官方格式化代码
- go doc 提取首行注释
- go install 安装go包
- go fix 旧版更新和修改到新版
- go test 轻量级的单元测试框架
## 笔记
- `{`不能单独一行
```gotemplate
	func main()
	{   //错误
	
	}

```
- `:= `只能在函数内使用
- go文件以小写字母组成，多个部分，以 `_`分割
### 学习链接
> http://c.biancheng.net/view/2.html 例子错误太多之外，还比较适合新手

> http://tour.studygolang.com/basics/3 可以学习到更多的例子加深使用，适合新手

> https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/03.9.md 《Go 语言入门指南》

> https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter01/01.1.html 《Go语言中文网》

## 安装


- windows
	- 安装地址：
```cmd

> go run hello.go  # 解析器执行运行

> go build hello.go # 编程成二进制.exe文件，再运行exe


```


- linux
 
## 疑问
1. 无法使用单引号 ？
2. %d ？ 应该是数值？ 数字 double 
3. %T ？？ 打印出来数据类型
4. %V ？？ 打印出来数据类型 Printf()来决定
5. 数组不是一个常量
	- 不能以 const 来声明
6. go如何声明 一串数组 [5454,"xxx",true] 
	- 在go 中，这是一个interface
	
7. for迭代中 可以使用 空白符 "_"省略索引或者值
```gotemplate
for _,num :=range [545,545,5]{
	fmt.Println(num)
}
``` 

8. 如何比较 "2019"==2019？
9. 查看数据的 类型 类似， js 中的typeof
10. 怎么打印不出来常量 的位操作
> //fmt.Printf("%i",Big)//怎 
> var i int=999 //期待 int

11. 发查询反而没有单线程顺序查询快？？
12. 在某些场景下，互斥锁要比读写锁更快！！！

13. 如何从这个字符串解析自己想要的数据
> &{GET / HTTP/1.1 1 1 map[Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3] Accept-Encoding:[gzip, deflate, br] Accept-Language:[zh-CN,zh;q=0.9,es;q=0.8,es-ES;q=0.7] Cache-Control:[max-age=0] Connection:[keep-alive] Upgrade-Insecure-Requests:[1] User-Agent:[Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36]] {} <nil> 0 [] false 127.0.0.1:8888 map[] map[] <nil> map[] 127.0.0.1:56865 / <nil> <nil> <nil> 0xc000018200}
> https://studygolang.com/articles/14107#reply1


解决：
```text
fmt.Println("xxxxx====xxx")
fmt.Println("Method：", r.Method)
fmt.Println("URL：", r.URL)
fmt.Println("Proto：", r.Proto)
fmt.Println("ProtoMajor ：", r.ProtoMajor)
fmt.Println("ProtoMinor ：", r.ProtoMinor)
fmt.Println("Header ：", r.Header)
fmt.Println("Accept: ：", r.Header["Accept"])
fmt.Println("Body  ：", r.Body)
fmt.Println("ContentLength   ：", r.ContentLength)
fmt.Println("TransferEncoding    ：", r.TransferEncoding)
fmt.Println("Close     ：", r.Close)
fmt.Println("Host ：", r.Host)
fmt.Println("Form ：", r.Form)
fmt.Println("PostForm  ：", r.PostForm)
fmt.Println("MultipartForm  ：", r.MultipartForm)
fmt.Println("Trailer   ：", r.Trailer)
fmt.Println("RemoteAddr   ：", r.RemoteAddr)
fmt.Println("RequestURI   ：", r.RequestURI)
fmt.Println("MultipartForm  ：", r.MultipartForm)
fmt.Println("TLS   ：", r.TLS)
fmt.Println("Cancel   ：", r.Cancel)
fmt.Println("Response   ：", r.Response)
```
14. 语法提示 main redeclared in this block
> 同一个目录下面不能有多个package main

15. 为什么如何单个return？
```gotemplate
//必须要给值
func mongo() {
	const x=8;
	return  x
}
```

16. 如何引入不同目录的其他函数go文件
> 需要设置$GOROOT,好编译器有关
17. 同目录的文件如何引用
18. go  语言如何返回json接口数据,map 转json

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Name struct {
	Job string
}
func main()  {
	method1()
	method2()

}

/*方法一*/
func method1(){
	name :=Name{Job:"web前端"}

	fmt.Println("name:",name)
	buf,err:=json.MarshalIndent(name,"","  ")
	if err!=nil{
		return
	}
	fmt.Println("方法一：",string(buf))
}
/*方法二*/
func method2()  {
	name :=Name{"Go后端开发"}
	jsonData,err:=json.Marshal(name)
	if err!=nil{
		return
	}
	fmt.Println("方法二：",string(jsonData))
}

```

19. go struct 转为map 输出

```gotemplate
type Object struct {
	X int
	Y int
}
m := make(map[string]interface{})
i:=Object{11,33}
j,_:=json.Marshal(i)
_ = json.Unmarshal(j, &m)
fmt.Println(m)

```

20. 怎么才可以缩减为11位呢？json后莫名加了两个冒号导致手机加冒号多了两位

> json转换过程中出现的 


21. json 如何处理url 转为对象？两个值存在才数组？
```json
{"code":2500,"data":{"error":["bad_verification_code"],"error_description":["The code passed is incorrect or expired."],"error_uri":["https://developer.github.com/apps/managing-oauth-apps/troubleshooting-oauth-app-access-token-request-errors/#bad-verification-code"]},"message":"The code passed is incorrect or expired."}
```

22. go的函数是堆还是栈？
应该是堆吧？

## 主要特性
- 自动垃圾回收
- 丰富的内置类型
- 函数多返回值
- 错误处理
- 匿名函数和闭包
- 类型和接口
- 并发编程
- 反射
- 语言交互性

## 初试go,hello world

> go run hello.go  直接执行

> go build hello.go



```go
package main //?

import "fmt"//?

//一定要main开头
func main()  {
	fmt.Printf("hello,worl11d \n");
}
```
## 函数包
### 
### 函数执行顺序

- init()
- main()
- customer()


### len(a) 只能是string 吗？

```go
package main
import "fmt"
func main(){
	var numbers =make([]int,3,5)
	printSlice(numbers)
	
}

func printSlice(ints []int) {
	fmt.Println("len=%d cap=%d slice=%v\n",len(ints),cap(ints),ints)
}
```

### cap()
测量切片最大长度可以达到多少

### fmt 

|函数|入参|描述|用例|
|-|-|-|-|
|println(t)||||
|Scan()||||
|Printlh(a,b)||||
|Printf("%d",a)||||


### unsafe.Sizeof() 是干嘛的这个函数

|函数|入参|描述|用例|
|-|-|-|-|
|Sizeof()||||

### fmt.printf()

|符号|描述|
|---|---|
|%d|数字,十进制整数|
|%x|十六进制整数|
|%o|八进制整数|
|%b|二进制整数|
|%f|浮点数 3.14149265|
|%g|浮点数 3.165659599656599|
|%e|浮点数 3.21656559e+00|
|%t|布尔型 true，false|
|%c|字符串 unicode 码点|
|%s|字符串|
|%q|带引号字符串 "abc" 或字符 'c'|
|%v|内置格式的任何值|
|%T|任何值的类型|
|%%|百分号本身（无操作数）|

### append()函数
- 增加切片的容量，必须创建一个新的更大的切片并把原来的分片的内容拷贝进来

```gotemplate
var n []int
n=append(n,9859,556,6)//[9859,556,6]

```
### copy()函数
> copy(n1,n)//n 拷贝到 n1

### fmt.Sprintf()
> fmt.Sprintf("%s:%v",header,rand.Int31())

### rand.Int31()

### time.Sleep(time.Second)

### time.Second

### string->float strconv.ParseFloat() 
### string->int strconv.ParseInt() 
### string->bool strconv.ParseBool() 

### int->string


```gotemplate
str1 :=strconv.Itoa(i) //1 

str2 :=fmt.Srpintf("%d",i)//2

str3 :=strconv.formatInt()//3

```

### int->float float(i)

```gotemplate
	float(i)

```

### int->bool bool(i)

### float->string

```gotemplate
str1 :=strconv.Itof(f) //1 Itoa方法
str2 :=fmt.Sprintf("%f",f) Sprintf 方法
str3 :=strconv.FormatFloat()//FormatFloat 转换
```

### float->int int(i)

### float->bool bool(i)

### bool->string  
```gotemplate
str1 :=fmt.Sprintf("%d",b)
str2 :=strconv.FormatBool()
```
### bool->int int(i)

### bool->float float(i)

### byte->string string(byte)

### byte->int/bool/float encoding/binary

### ini/bool/float -> byte encoding/binary

### log.Fatal 似乎是一个打印错日志

### defer 
用途
- defer file.Close()关闭文件流
- defer mu.Unlock() 解锁一个加锁的资源
>mu.Lock()
>defer mu.Unlock()

- 打印最终的报告
>defer EndPrint()

- 关闭数据库链接
>defer disconnectDB()

## 基础代码结构

|描述|用例|
|----|----|
|包声明| |
|引入包| |
|函数| |
|变量| |
|语句&表达式| |
|注释| |
|// -> 标记| |
|关键字| |
|常量| |
|字符串| |
|符号| |
|fm| |
|.| |
|Println| |
|(| |
|"hello world"| |
|)| |

## 保留字
### 关键字

|关键字|描述|
|----|----|
|break ||
|default ||
|func | |
|interface | |
|select | |
|case | |
|defer | |
|go | |
|map | |
|struct | |
|chan | |
|else | |
|goto | |
|package | |
|switch | |
|const | |
|fallthrougt | |
|if | |
|range | |
|type | |
|continue | |
|for | |
|imort | |
|return | |
|var | |

### 预定义标识符

|符号|描述|
|----|----|
|append | |
|bool | |
|byte | |
|cap ||
|close | |
|complex | |
|complex64 | |
|complex128 | |
|uint | |
|uint8 | |
|uint16 | |
|uint32 | |
|uint64 | |
|uintptr | |
|copy | |
|false | |
|float32 | |
|float64 | |
|imag | |
|int ||
|int8 | |
|int16 | |
|int32 | |
|iota | |
|len ||
|make | |
|new ||
|nil ||
|panic | |
|print | |
|println | |
|real | |
|recover | |
|string | |
|true | |

## 声明&&变量



### 声明变量

- 指定变量类型，如果没有初始值，变量默认为零值
> var s ini;

- 根据值自行判定变量类型
> var f= true

- 省略var `:=` 左侧如果没有声明新的变量，会产生编译错误
- 全局声明的变量允许不被使用
- 局部声明的变量，必须要使用，否则会造成编译器报错
```gotemplate
var val int
val := //错误，左侧没有声明新的变量
var,t :=1,2//不会错误

```

```gotemplate
var age int
var str string = "I love the world"//go 中 这个string 类型是可以省略的

```

- 必须空格区分
- 被声明的变量，必须使用，对于局部作用域来说，但全局是允许声明可不使用。
- 可在一行声明多个变量。var e,f =123,"hello" => g,h:=123,"world"，并行赋值|同时赋值
- 不能对变量，重复声明
- 交换两个变量的值，前提是类型相同 a,b=b,a
- 空白标识符，_ 用于抛弃值 _,b=5,7. _ 
- 不能以数字开头
- 不能使用关键字
- 不能使用运算符

### 变量作用域

- 函数内定义的变量：局部变量
- 函数外定义的变量：全局变量
- 函数定义中的变量：形式参数

### 初始化局部和全局变量
- int 初始化 0
- float32 初始化 0
- pointer 初始化 nil
### `:=` 短变量声明，省略var 关键字
- 不能在函数外使用，只能在函数内使用
- 必须不能被声明过

### 常量
const 声明

### 零值
- `0` 数值0
- `false` 布尔 false
- `""` 空字符串

## 值类型和引用类型

### 值类型（*）

	值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。

> (基本类型，存储在栈中，常量)  j=i 将i的值进行拷贝,&i 获取i的内存地址。每次的地址可能不一样


- 以下基本类型都属于值类型
	- int
	- float
	- bool
	- string
> 且变量都直接指向存在内存中的值

- `j = i`,是内存中将i的值进行拷贝，可通过`&i`获取变量 `i`的内存地址,每次内存地址可能不一样， 
- 值类型的变量的值都存在 `栈` 中

|类型|描述|
|----|----|
|int||
|int| |
|float| |
|bool| |
|string| |

- 常量
常量用于枚举
```gotemplate
const {
    Unkown=0
    Female=1
    male=2
}
```
- iota(特殊常量，可以被编译器修改的常量)???	
    
|||||||||
|-|-|-|-|-|-|-|-|
|const (|||\|||const (||
||a=iota||\||||a=iota||
||b=iota||=>|||b||
||c=iota||\||||c||
|)|||\|||)|||

```gotemplate
const (  			
	a=iota				
	b=iota     =>		
	c=iota						
								
```



- iota 用法

```gotemplate
package main
import "fmt"
func main(){
    const (
        a=iota
        b
        c
        d ="haha"
        e
        h
        i
        f=100
        g=iota
        i
    )
}
```


### 引用类型（*）

	引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

> (复杂的数据结构，)r2=r1 只有引用被拷贝

## 数据类型


### 布尔类型
- true 
- false

### 数字类型
- 整形int
    - uint8  `0 ~ 255`
    - uint16 `0 ~ 65535`
    - uint32 `0 ~ 4294967295`
    - uint64 `0 ~ 188446744073709551615`
    - int8   `-128 ~ 127`
    - int16  `-32768 ~ 32767`
    - int32  `-2147483648 ~ 2147483647`
    - int64  `-9223372036854775808 ~ 9223372036854775807`
- 浮点
	- float32 IEEE-754 32
    - float64 IEEE-754 64
    - complex64 32位实数和虚数
    - complex128 64位实数和虚数 (实数：有理数+无理数) （复数：实数+虚数）	
- 其他
	- byte uint8
    - rune ini32
    - uint 32/64
    - int 与uint
    - uintptr 无符号整形，存放一个指针pointer
    
### 字符串类型
- UTF-8编码标识的 
- unicode文本

### 派生类型
- 指针类型 Pointer
- 数组类型
- 结构化类型 struct
- Channel 类型
- 函数类型
- 切片类型
- 接口类型 interface
- Map 类型

## 运算符
- :=  (赋值操作符)
```gotemplate
	var a =5;  => a:=5
    var b =false; => b:=false
```
### 算数运算符
- +
- -
- *
- /
- %
- ++
- --

### 关系运算符(都返回：布尔值) -> 

|符号|用例|
|----|----|
|==  |a==b|
|!=  |a!=b|
|>   |a>b |
|<   |a<b |
|>=  |a>=b|
|<=  |a<=b|

### 逻辑运算符
- %% 逻辑 AND 运算符，两边都是true 才true ，否则false  (a&&b)
- || 逻辑 OR  有一个true 则true。否则false			(a||b)
- |  逻辑 NOT 条件为true，则逻辑为false 否则true		!(a&&b)
### 位运算符(都是操作二进制)
- &	与运算符，双目运算符。(a&b) a与b的二进位 相与  (0-0 、1-0、0-1 为0；1-1为1) 

| 8 | 7 | 6 | 5 |  4| 3 | 2 | 1 |
|---|---|---|---|---|---|---|---|
| 0 | 0 | 1 | 1 | 1 | 1 | 0 | 0 |
| 0 | 0 | 0 | 0 | 1 | 1 | 0 | 1 |
| 0 | 0 | 0 | 0 | 1 | 1 | 0 | 0 |

- |	或运算符，双目运算符。(a|b) a与b的二进位 相或 （0-0 为,1-1、1-0、0-1 为1）

| 8 | 7 | 6 | 5 | 4 | 3 | 2 | 1 |
|---|---|---|---|---|---|---|---|
| 0 | 0 | 1 | 1 | 1 | 1 | 0 | 0 |
| 0 | 0 | 0 | 0 | 1 | 1 | 0 | 1 |
| 0 | 0 | 1 | 1 | 1 | 1 | 0 | 1 |

- ^ 	异或运算符，双目运算符。(a^b) a与b的二进位 想异或。（0-0,1-1 都为0,1-0,0-1 为1）

| 8 | 7 | 6 | 5 | 4 | 3 | 2 | 1 |
|---|---|---|---|---|---|---|---|
| 0 | 0 | 1 | 1 | 1 | 1 | 0 | 0 |
| 0 | 0 | 0 | 0 | 1 | 1 | 0 | 1 |
| 0 | 0 | 1 | 1 | 0 | 0 | 0 | 1 |

- <<	左移运算符，双目运算符。(a<<2) a的二进位，整体向左边移动2个位

| 8 | 7 | 6 | 5 | 4 | 3 | 2 | 1 |
|---|---|---|---|---|---|---|---|
| 0 | 0 | 1 | 1 | 1 |1 | 0  | 0 |
| 1 | 1 | 1 | 1 | 0 |0 | 0  | 0 |

- \>>	右移运算符，双目运算符。(a>>2) a的二进位，整体向右边移动，补0

| 8 | 7 | 6 | 5 | 4 | 3 | 2 | 1 |
|---|---|---|---|---|---|---|---|
| 0 | 0 | 1 | 1 | 1 | 1 | 0 | 0 |
| 0 | 0 | 0 | 0 | 1 | 1 | 1 | 1 |

### 赋值运算符

|符号 | 用例           |
|----|----------------|
|=   | c=a+b          |
|+=	 | c+=a   => c=c+a|
|-=	 | c-=a   => c=c-a|
|*=	 | c*=a   => c=c*a|
|/=	 | c/=a   => c=c/a|
|%=	 | c%=a   => =c%a|
|<<= |	c<<=2 => c=c<<2|
|>>= |	c>>=2 => c=c>>2|
|&=	 | c&=2   => c=c&2|
|^=	 | c^=2   => c=c^2|
|｜= | c\｜=2   => c=c\｜2|
### 其他运算符

|符号 |描述|       用例           |
|----|----|-----------------------|
|&|返回变量存储地址|  &a 给出实际地址|
|*|指针变量       |*a 是一个指针变量|

### 运算符优先级
> (使用括号提升表达式运算符优先级)

|优先级|运算符|
|----|-------------|
|7|^ !             |
|6|* / % << >> & &^|
|5|+ - ^           |
|4|== != < <= >= > |
|3|<-              |
|2|&&              |
|1|    |           |
    
## 语句
### 条件语句

- if
```gotemplate
if a>20{		
							
}	
```

- if ..else
```gotemplate
if a>20{

}else{

}
```

- if ..if

```gotemplate 
if a>20{
    if a==30 { 

    }
}
```
- switch

```gotemplate

package main

import "fmt"

fnc main(){
	switch var1 {		
	case var1:			
        ...				
	case var2:			
        ...				
	case var 3			
        ...				
	default:			
        ....			
	}	
}

```
				
- select 
> !!!select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。

```gotemplate
select{						
    case communication clause:	
        statement(s);			
    case communication clause:	
        statement(s);			
}								
```

1. case 都必须是一个通信？？！！
2. channel表达式都会求值？？
3. 所有被发送的表达式都会被求值？？
4. 如果任意某个通信都可以运行，它就会执行，其他被忽略
5. 多个case都可以运行，select随机公平的选出一个执行，其他不会执行
6. 否则 如果有default，则执行
7. 如果没有default 语句，select将会阻塞，直到某个通信可以运行，go不会重新对channel或值进行求值
	
### 循环控制语句
- 只有一种循环结构，for 循环
- for
- for ...for

```gotemplate
package main
import "fmt"

func main(){
	for i,x:= range numbers {
        fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
    }		
}

```

- break	-> 中断for循环或跳出switch
- continue	->	跳过当前循环剩下语句，进行下一轮循环
- goto	->	将控制转移到被标记的语句
- 无限循环，条件语句永远不为false则无限循环	

### 判断字符串为空
- if str == "" {}
- if len(str)==0{}
## 数组
- 如何声明 一串数组 [5454,"xxx",true]
### 声明数组
	var variable_name [SIZE] variable_type //语法
	
	var balance[10] float // 定义一个长度为10的 float32 类型
	
### 初始化数组
	var balance = [5] float32{9.6,65.5,6.6,.6,9.5,9.66}//err
	var balance = [5] float32{9.6,65.5,6.6,.6,9.5}//success
- 初始化数组中 `{}` 的元素个数不能大于 `[]` 的数字
- `[]`没有设置数字，则自动设置	
- 如果`{}`里面的元素个数小于 `[]`定义的，则会自动补全0.

### 多维数组（省略）

## 指针
- 一个指针变量指向一个值的内存地址

## 结构体

- @TODO 数组可以存储同一个类型的数组
- @TODO 数组如何支持结构体 arr=[999,"899",false]
- 那只是属性而已，怎么去调用方法？？

```gotemplate
package main

import (
	"fmt"
	"math"
)

type Class struct {
	x,y float64
}

// * 地址
func (v *Class) Woo() float64  {
	fmt.Println(Class{})
	return math.Sqrt(v.x*v.x+v.y*v.y)
}

func main()  {
	v :=&Class{3,4}
	fmt.Println(v.Woo())
	fmt.Println(v)
}

```

- 结构体实例
```go
package main
import "fmt"

type Books struct {
	title string
	author string
	page int
}

func main()  {
 fmt.Println(Books{"三百六十行","孙悟空",999})
 
 // 当然也可以使用key value
 fmt.Println(Books{title:"哇哈哈",author:"宗庆后",page:666})
 
 // 忽略字段为 0 或 空


 fmt.Println(Books{title:"演员的自我修养",author:"无名氏"})
}


```

- 结构体语法
```gotemplate
//
 type struct_variable_type struct{
    member definition
    member definition
    ...
    member definition
 }
 
 ```
 
### 结构体指针

## 切片(slice)
[slice切片](./slice.go)

- 对数组的抽象
- “动态数组”——切片
- 长度不限制
- make函数创建切片

### 切片如何转数组

### 定义切片
> var xx []type //不需要说明长度

### 切片初始化
> s:=[] int {1,2,3}

## 语言范围 TODO
- range 关键字用于for循环中 迭代 数组、切片、通道、集合的TODO?
- range在数组、切片中它返回元素和索引和索引对应的值
- range在集合中，返回key-value对 的key值
> https://www.runoob.com/go/go-range.html

## interface 接口


## goroutine 

问题:
- goroutine 还可以继续goroutine 吗？

### 概念与总结
- 线程(Thread):
  - 有称：轻量级进程(Lightwight process,LWP)
  - 程序执行流的最小单元
  - 标准线程包括：线程ID，指令指针（PC），寄存器集合，堆栈
  - 线程是进程的一个实体，系统独立调度分配的单位
  - 无法拥有系统资源，仅个别必备资源
  - 可与同属一个进程的其他线程共享进程的全部资源
- 协程(Coroutine):
  - 又称：微线程与子例程（函数）
  - 程序组件
  - 协程比较灵活，实践中没有比子例程广泛
  - 共享堆，不共享栈


## 并发concurrent
![concurrent](./src/images/concurrent.jpg)

### go并发模型：CSP(communicating sequential processes)

- go`不要以共享内存的方式来通信，相反，要通过通信来共享内存。`


## 通道channel


## 函数

```gotemplate
package main
import "fmt"

func hello([参数list])[return_types]{
            //functions body
}
```

1. func 声明函数
2. hello 函数名称，
3. 函数签名= 函数名+参数列表
4. 实参，参数类型、顺序、参数个数，可选
5. return_types 返回类型，返回返回一列值，return_types 是该列值的数据类型，不是必须，有些功能不需要返回值
6. 函数体 里面搞什么
7. 返回多个值

```go
package main
import "fmt"
func main()  {
    world("9","2")
}
func world(x,y string)(string,string){
    return x,y
}
```

8. 函数如果使用参数，则该变量成为函数的形参
9. 形参类似定义在函数体内的 局部变量
10. 值传递，在函数体内去修改，不会影响，因为是属于基本类型的的拷贝值
11. 引用传递参数的话，在函数体内去修改，会影响外部参数
12. 函数可作为值，被赋值给其他o变量
13. 最少要有一个main函数
### 函数闭包

```go
package main 
import "fmt"
func getSequence() func() int{
    i:=0//var i int = 0
    return  func() int{
        i+=1
        return i
    }
}
func main(){
    nextNumber :=getSequence()//nextNumber 为一个函数，值为0
    fmt.Println(nextNumber())
    fmt.Println(nextNumber())
    fmt.Println(nextNumber())
    preNumber :=getSequence()//新的函数
    fmt.Println(preNumber())
    fmt.Println(preNumber())
}
```
### 递归函数
go 实现 斐波那契数列函数

```go
/**
@desc 递归 recursion

*/
package main

import (
	"fmt"
)

//
func recursion(num int) int {
	if num <= 1 {
		return 1
	}
	return recursion(num-1) + recursion(num-2)
}

func main() {
	var a = recursion(8)
	fmt.Println(a)
}

```

go 通过闭包实现斐波那契数列

```go
package main

import "fmt"

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	x,y:=0,1
	return func() int {
		temp := x
		x, y = y, (x + y)
		return temp
	}
}
```

###  传递变长参数

> func myFunc(a, b, arg ...int) {}

```gotemplate
func Greeting(prefix string, who ...string)
Greeting("hello:", "Joe", "Anna", "Eileen")
```
### 内置函数

来源：https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/06.6.md

|名称|说明|
|---|---|
|close|用于管道通信|
|len、cap|len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）|
|new、make|new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针（详见第 10.1 节）。它也可以被用于基本类型：`v := new(int)`。make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作（详见第 7.2.3/4 节、第 8.1.1 节和第 14.2.1 节）**new() 是一个函数，不要忘记它的括号**|
|copy、append|用于复制和连接切片|
|panic、recover|两者均用于错误处理机制|
|print、println|底层打印函数，在部署环境中建议使用 fmt 包|
|complex、real imag|用于创建和操作复数|



## 时间Date
更多 http://docs.studygolang.com/pkg/time/


## 字符

- 字符串链接使用 `+`来实现
- `\n` 换行符
- `\r` 回车符
- `\t` tab键
- `\u` 或 `\U` unicode字符
- `\\`：反斜杠自身

### 方法
更多：http://docs.studygolang.com/pkg/strings/

- strings.HasPrefix(s,prefix string) //bool prefix 开头
- strings.HasSuffix(s,suffix string) //bool suffix 结尾
- strings.Contains(s,substr string) //bool 判断字符串包含
- strings.Index(s,str string) int //索引，-1不包含
- strings.LastIndex(s,str string) int 从最后出现的位置
- strings.IndexRune(s string,r rune) int 如果需要查询非 ASCII 编码的字符在父字符串中的位置，建议使用以下函数来对字符进行定位：
- strings.Replace(str,old,new,n) `str`字符串，前`n`个字符串` old`替换为`new`，并返回一新的字符串,n=-1则替换 所有`old` 为 `new`
- strings.Count(s,str,string) //int 计算字符串 str 在字符串 s 中出现的非重叠次数：
- strings.Repeat(s,count int) //string 重复字符串
- strings.ToLower(s) 小写
- strings.ToUpper(s) 大写
- strings.TrimSpace(s)提出字符串开头和结尾的空白符号
- strings.Trim(s,"cut") 将结尾和开头的`cut`去除
- strings.TrimLeft(s)
- strings.TrimRight(s)
- strings.Join(sl [] string,sep string)//"GO1 - The ABC of Go - 25 - "-> "GO1;The ABC of Go;25"
- strings.Fields(str) //todo? "The quick brown fox jumps over the lazy dog"-> ["The quick brown fox jumps over the lazy dog"]
- strings.Split(s,str)//"GO1|The ABC of Go|25"->[GO1 The ABC of Go 25]

### 字符串与其他类型转换，通过 `strconv` 包实现
更多 http://docs.studygolang.com/pkg/strconv/

## 标准库包
|包名   |描述   |
|---    |---   |
|bufio	|带缓冲的 I/O 操作|
|bytes	|实现字节操作|
|container	|封装堆、列表和环形列表等容器|
|crypto	|加密算法|
|database	|数据库驱动和接口|
|debug	|各种调试文件格式访问及调试功能|
|encoding	|常见算法如 JSON、XML、Base64 等|
|flag	|命令行解析|
|fmt	|格式化操作|
|go|	Go 语言的词法、语法树、类型等。可通过这个包进行代码信息提取和修改|
|html	|HTML 转义及模板系统|
|image	|常见图形格式的访问及生成|
|io	|实现 I/O 原始访问接口及访问封装|
|math|	数学库|
|net|	网络库，支持 Socket、HTTP、邮件、RPC、SMTP 等|
|os|	操作系统平台不依赖平台操作封装|
|path	|兼容各操作系统的路径操作实用函数|
|plugin|	Go 1.7 加入的插件系统。支持将代码编译为插件，按需加载|
|reflect|	语言反射支持。可以动态获得代码中的类型信息，获取和修改变量的值|
|regexp	|正则表达式封装|
|runtime|	运行时接口|
|sort	|排序接口|
|strings|	字符串转换、解析及实用函数|
|time	|时间接口|
|text	|文本模板及 Token 词法器|

## http

- w http.ResponseWriter 返回的消息是byte类型

>w.Write([]byte("hello world""))


```go
/**

@desc go 建立 http 服务器

*/
package main

import (
	"io"
	"log"
	"net/http"
)

func main()  {
		goServer()
}

func goServer() {
	http.HandleFunc("/",httpServer)
	err := http.ListenAndServe(":8888",nil)
	if err !=nil {
		log.Fatal("ListenAndServe",err)
	}
}

func httpServer(w http.ResponseWriter,req *http.Request)  {
	_, _ = io.WriteString(w, "hello,world11\n")
}

```


## 内置依赖包

> 摘录自 https://studygolang.com/static/pkgdoc/main.html

|main|sub|描述|
| --- | --- | --- |
|archive|       | 档案| 
|       | tar | tar包实现tar格式压缩文件存取| 
|       | zip | zip提取zip档案文件读写 | 
| bufio  | | 带缓存I/O操作 | 
| builtin | | 为Go预声明标识符提供了文档 | 
| bytes  | | 操作[]byte的常用函数 | 
| compress  | | | 
|   | bzip3| | 
|   | flate| | 
|   | gzip | | 
|   | lzw| | 
|   | zlib| | 
| container  | | | 
|   | heap| | 
|   | list| | 
|   | ring| | 
| context  | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 
|   | | | 

## Http
- https://godoc.org/net/http#Request.Method
> http - GoDoc
## 算法

### 倍数算法，此时for 也类似while
```go
package main

import "fmt"

func main()  {
	sum :=1
	for sum <64{
		fmt.Println(sum)
		sum+=sum
	}
	fmt.Println("end:",sum)
}


/*
  1
  2
  4
  8
  16
  32
  sum:64
  */
```

### 牛顿法实现开平方函数

简单10次循环找差值

```go

package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)

	for i:=0;i<10 ; i++{
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(8))
}

```

- 完善型牛顿法
```go

package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	const E = 0.000001
	z := float64(1)
	k := float64(0)
	for ; ; z -= (z*z - x) / (2 * z) {
		if z-k <= E&&z-k >= -E {
			return z
		}
		k = z
	}
}

func main() {
	fmt.Println(Sqrt(8))
}

```

