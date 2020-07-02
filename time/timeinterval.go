/**
* 定时器
*/
package main
import (
	"fmt"
	"time"
)
func main()  {
	fmt.Println("1111=>")
	for {
		fmt.Println("forever ==>",time.Now())
		time.Sleep(time.Second*2)
	}
	fmt.Println("End")
}