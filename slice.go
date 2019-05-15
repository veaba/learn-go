/**
@desc 切片
*/

package main
import "fmt"
func main(){
	var numbers =make([]int,3,5)

	fmt.Println(numbers)//[0 0 0]

	//len 和 cap
	//printSlice(numbers)

	//空切片
	//printNil()

	//切片截取
	//_slice()

	// append 追加
	_append()
	// copy 拷贝

}

func _append() {
	var n []int
	pSlice(n)

	n=append(n,0) //[0]
	pSlice(n)

	n=append(n,1) //[0 1]
	pSlice(n)

	n=append(n,9859,556,6) //len=5 cap=6 slice=[0 1 9859 556 6]
	pSlice(n)

	n1 := make([]int,len(n),(cap(n))*2)//len=5 cap=12 slice=[0 1 9859 556 6]
	copy(n1,n)//n 拷贝到 n1
	pSlice(n1)


}

func _slice() {
	n := []int{6526,652,65,65,65,65,6655}
	pSlice(n)

	//原始切片
	fmt.Println(n)

	//索引1 - 4
	fmt.Println("n[1:4]",n[1:4])

	// 默认下限为0 :4
	fmt.Println("n[:4]",n[:4])

	//默认上限为len
	fmt.Println("n len",n[3:])

	//开始切片
	n1 := make([]int,0,10)
	pSlice(n1)

	n2 := n[:4]
	pSlice(n2)

	n3 := n[1:4]
	pSlice(n3)

	n4 :=n[3:]
	pSlice(n4)


}

func pSlice(n []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(n),cap(n),n)
}

func printNil() {
	var n []int
	_nil(n)

	if (n==nil) {
		fmt.Printf("切片为空")
	}
}

func _nil(ints []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(ints),cap(ints),ints) // len=0 cap=0 splice=[]
}



func printSlice(ints []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(ints),cap(ints),ints)// len=3 cap=5 splice=[0 0 0]
}
