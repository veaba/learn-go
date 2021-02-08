package main

import (
	"fmt"
	"time"

	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

func main() {
	str := createToken2(110, "dddddddddddddd")
	fmt.Println("str=>", str)
}
func createToken2(appId int32, secret string) string {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{

		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 10).Unix(), // max 10 minute
		"iss": appId,
	})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
	return tokenString
}
