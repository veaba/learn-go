package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	str, _ := createToken(64623, "ddddddddddddd")
	fmt.Println("str=>", str)
}

func createToken(appId int32, secret string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		//at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 10).Unix(), // max 10 minute
		"iss": appId,
	})
	fmt.Println("at=>", at)
	token, err := at.SignedString([]byte(secret))

	fmt.Println("token=>")
	if err != nil {
		return "", err
	}
	return token, nil
}

func LoadRSAPrivateKeyFromDisk(location string) *rsa.PrivateKey {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}
