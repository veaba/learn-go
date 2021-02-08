package main

import (
	"fmt"
	"net/url"
	"reflect"
)

func main() {
	var a = "http://127.0.0.1"
	var b = a + "?code=2021"

	fmt.Println("a===>", a)
	fmt.Println("b===>", b)

	aRes, _ := url.Parse(a)
	bRes, _ := url.Parse(b)
	// cRes, _ := url.ParseRequestURI(a)
	// dRes, _ := url.ParseRequestURI(b)
	// eRes, _ := url.Parse(b)
	fmt.Println("url params a===>", aRes) //  map[http://127.0.0.1:[]]
	fmt.Println("url params b===>", bRes) //  map[http://127.0.0.1:[]]

	// cRes := aRes.Query()
	// cRes.Set("name", "lii")
	// fmt.Println("cRes=>", cRes)
	// fmt.Println("cRes encode=>", cRes.Encode())
	// fmt.Println("aRes=>", aRes)
	// aRes.RawQuery = cRes.Encode()
	// fmt.Println("aRes=>", aRes)
	// fmt.Println("================>")

	dRes := bRes.Query()
	dRes.Set("domain", "badu.com")
	fmt.Println("dRes=>", dRes)
	fmt.Println("dRes encode=>", dRes.Encode())
	fmt.Println("bRes=>", bRes)
	bRes.RawQuery = dRes.Encode()
	fmt.Println("bRes=>", bRes)
	fmt.Println("bRes=>", *bRes)

	fmt.Println(reflect.TypeOf(bRes))                   // *url.URL
	fmt.Println("*url.URL to string =>", bRes.String()) // http://127.0.0.1?code=2021&domain=badu.com

}
