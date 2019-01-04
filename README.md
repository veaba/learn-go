## 前言
> go，代码组织像c语言，语法和一些特性 像 js
## 疑问
1. 无法使用单引号 ？
2. %d ？ 应该是数值？
3. %T ？？ 打印出来数据类型
## 函数包
### len(a)
### fmt 
|函数|入参|描述|用例|
|-|-|-|-|
|println(t)||||
|Scan()||||
|Printlh(a,b)||||
|Printf("%d",a)||||
### unsafe
|函数|入参|描述|用例|
|-|-|-|-|
|Sizeof()||||

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
## 声明
### 声明变量
```js
var age int
```
- 必须空格区分
- 被声明的变量，必须使用，对于局部作用域来说，但全局是允许声明可不使用。
- 可在一行声明多个变量。var e,f =123,"hello" => g,h:=123,"world"，并行赋值|同时赋值
- 不能对变量，重复声明
- 交换两个变量的值，前提是类型相同 a,b=b,a
- 空白标识符，_ 用于抛弃值 _,b=5,7. _ 

## 值类型和引用类型
### 值类型
> (基本类型，存储在栈中，常量)  j=i 将i的值进行拷贝,&i 获取i的内存地址。每次的地址可能不一样
|类型|描述|
|----|----|
|int||
|int| |
|float| |
|bool| |
|string| |

- 常量
    - 常量用于枚举
    ```js
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

```js
const (  			
	a=iota				
	b=iota     =>		
	c=iota						
								
```



    - iota 用法
    ```js
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
### 引用类型
> (复杂的数据结构，)r2=r1 只有引用被拷贝
## 数据类型

### 布尔类型
- true 
- false
### 数字类型
- 整形int
    - uint8 0-255
    - uint16 0-65535
    - uint32 0-4294967295
    - uint64 0-188446744073709551615
    - int8 -128-127
    - int16 -32768-32767
    - int32 -2147483648-2147483647
    - int64 -9223372036854775808-9223372036854775807
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
```js
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
    |8|7|6|5|4|3|2|1|
    |-|-|-|-|-|-|-|-|
    |0| |0| |1| 1| 1| 1| 0| 0|
    |0| |0| |0| 0| 1| 1| 0| 1|
    |0| |0| |0| 0| 1| 1| 0| 0|


- |	或运算符，双目运算符。(a|b) a与b的二进位 相或 （0-0 为,1-1、1-0、0-1 为1）
    |8|7|6|5|4|3|2|1|
    |-|-|-|-|-|-|-|-|
    |0|0|1|1|1|1|0|0
    |0|0|0|0|1|1|0|1
    |0|0|1|1|1|1|0|1
- ^ 	异或运算符，双目运算符。(a^b) a与b的二进位 想异或。（0-0,1-1 都为0,1-0,0-1 为1）
    |8|7|6|5|4|3|2|1|
    |-|-|-|-|-|-|-|-|
    |0| 0| 1| 1| 1| 1| 0| 0|
    |0| 0| 0| 0| 1| 1| 0| 1|
    |0| 0| 1| 1| 0| 0| 0| 1|
- <<	左移运算符，双目运算符。(a<<2) a的二进位，整体向左边移动2个位
   |8|7|6|5|4|3|2|1||-|
    |0| 0| 1| 1| 1| 1| 0| 0|
    |1| 1| 1| 1| 0| 0| 0| 0|
- \>>	右移运算符，双目运算符。(a>>2) a的二进位，整体向右边移动，补0

    |8|7|6|5|4|3|2|1|
    |-|-|-|-|-|-|-|-|
    |0| 0| 1| 1| 1| 1| 0| 0|
    |0| 0| 0| 0| 1| 1| 1| 1|

### 赋值运算符
|符号|用例|
|----|----|
|=   | c=a+b|
|+=	 | c+=a   => c=c+a|
|-=	 | c-=a   => c=c-a|
|*=	 | c*=a   => c=c*a|
|/=	 | c/=a   => c=c/a|
|%=	 | c%=a   => =c%a|
|<<= |	c<<=2 => c=c<<2|
|>>= |	c>>=2 => c=c>>2|
|&=	 | c&=2   => c=c&2|
|^=	 | c^=2   => c=c^2|
||=	 | c|=2   => c=c|2|
### 其他运算符
|符号|描述|用例|
|----|----|----|
|&|返回变量存储地址|&a 给出实际地址|
|*|指针变量|*a 是一个指针变量|
### 运算符优先级
> (使用括号提升表达式运算符优先级)
|优先级|运算符|
|----|----|
|7|^ !|
|6|* / % << >> & &^|
|5|+ - | ^|
|4|== != < <= >= >|
|3|<-|
|2|&&|
|1| |||
    
## 语句
### 条件语句
- if
```js
if a>20{		
							
}	
```
- if ..else
```js
if a>20{

}else{

}
```
- if ..if
```js
if a>20{
    if a==30 { 

    }
}
```
- switch

```js
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
```	    
				
- select 
> !!!select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
```js
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
- for
- for ...for
```js
for i,x:= range numbers {
    fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
}		

```

- break	-> 中断for循环或跳出switch
- continue	->	跳过当前循环剩下语句，进行下一轮循环
- goto	->	将控制转移到被标记的语句
- 无限循环，条件语句永远不为false则无限循环	
## 函数
```js
func hello([参数list])[return_types]{
            //functions body
}
```
1、func 声明函数
2、hello 函数名称，
3、函数签名= 函数名+参数列表
4、实参，参数类型、顺序、参数个数，可选
5、return_types 返回类型，返回返回一列值，return_types 是该列值的数据类型，不是必须，有些功能不需要返回值
6、函数体 里面搞什么
7、返回多个值

    func world(x,y string)(string,string){
        return x,y
    }
8、函数如果使用参数，则该变量成为函数的形参
9、形参类似定义在函数体内的 局部变量
10、值传递，在函数体内去修改，不会影响，因为是属于基本类型的的拷贝值
11、引用传递参数的话，在函数体内去修改，会影响外部参数
12、函数可作为值，被赋值给其他变量
### 函数闭包
```js
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