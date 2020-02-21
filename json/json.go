package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}
type ServerSlice struct {
	Servers []Server
}

func main() {
	var s ServerSlice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},
            {"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`

	fmt.Println(&s)
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	fmt.Println(s.Servers[0].ServerIP)
}
