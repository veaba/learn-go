/**
@desc error 错误处理
@TODO 不是很懂。
*/
package main

import "fmt"

type DivideRrror struct {
	dividee int
	divider int
}

// 实现error接口
func (de *DivideRrror) Error() string {
	strFormat := `
		Connot proceed,the divider is zero.
	dividee:%d
	divider:0
	`
	return fmt.Sprintf(strFormat,de.dividee)
}

// 定义 'int' 类型除法运算的函数

func Divide(varDividee int,varDivider int)(result int,errorMsg string)  {
	if varDividee==0{
		dData := DivideRrror{
			dividee:varDividee,
			divider:varDivider,
		}
		errorMsg = dData.Error()
		return
	}else {
		return  varDividee/varDivider,""
	}
}
func main()  {
	//正常情况下
	if result,errorMsg :=Divide(100,10);errorMsg==""{
		fmt.Println("100/10=",result)
	}

	//当被除数为0的时候，则会返回错误信息
	if _,errorMsg := Divide(100,0);errorMsg!=""{
		fmt.Println("errorMsg is :",errorMsg)
	}
}
