package main

import (
	"fmt"
	"time"
)

type Owen struct {
	name string
}
type Login struct {
	name string
}

// get owner repo info
func getOwnerHttp(c chan interface{}) {
	fmt.Println("获取 owner http", time.Now())
	time.Sleep(time.Second * 6)
	fmt.Println("owner time now=>", time.Now())
	c <- Owen{"owner"}

}

// get login repo info
func getLoginHttp(c chan interface{}) {
	fmt.Println("获取 login http", time.Now())
	time.Sleep(time.Second * 4)
	fmt.Println("login time now=>", time.Now())
	c <- Login{"login"}
}

func main() {
	channel := make(chan interface{})
	fmt.Println("开始执行==>")
	go getLoginHttp(channel)
	go getOwnerHttp(channel)

        // Disorder callback
	loginInfo, ownerInfo := <-channel, <-channel

	fmt.Println("loginInfo=>", loginInfo)
	fmt.Println("ownerInfo=>", ownerInfo)
	fmt.Println("ending=>", time.Now())

}
