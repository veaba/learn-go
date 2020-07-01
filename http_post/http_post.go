// go 发起 post请求
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ExchangeCodeStruct struct {
	Code         string `json:"code"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func main() {
	exChangeCode("dsadda")
}
func exChangeCode(code string) {
	fmt.Println("准备交换Token，发起post请求", code)
	RobotClientId := "RobotClientId"
	RobotClientSecret := "RobotClientSecret"
	RobotAccessTokenUrl := "RobotAccessTokenUrl"
	exchangeStruct := ExchangeCodeStruct{code, RobotClientId, RobotClientSecret}
	//buf, errBuf := json.MarshalIndent(exchangeStruct, "", "")
	jsonData,errBuf :=json.Marshal(exchangeStruct)
	if errBuf != nil {
		panic(errBuf)
	}
	resp, err := http.Post(RobotAccessTokenUrl,
		"application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("token_access兑换结果===>", body)
}
