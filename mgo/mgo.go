/**
@desc一段查询mongodb的代码
*/
package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	//"gopkg.in/mgo.v2/bson"
	"io"
	"log"
	"mongo"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	createHttpServer()
}

type People struct {
	_id  bson.ObjectId
	Name string `bson:"name"`
}

func createHttpServer() {
	session, err := mgo.Dial("127.0.0.1:27017")
	//session, err := mgo.Dial("127.0.0.1:27017/hello")
	fmt.Println("session:", session)
	if err != nil {
		return
	}
	//result:=Users{}
	var result []People
	hasError := session.DB("hello").C("worlds").Find(nil).All(&result)
	if hasError != nil {
		return
	}
	fmt.Println("result:", result)

	http.HandleFunc("/work", work)
	http.HandleFunc("/manage", manage)
	mongo.Mongo()
	log.Fatal(http.ListenAndServe(":8888", nil))
}

/*
@desc 路由
*/

func work(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	type WorkData struct {
		Hello string
	}
	//workData := `{ww:"workd"}`
	_, _ = io.WriteString(w, "xx")
}

func manage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	_, _ = io.WriteString(w, "{hello:'manage'}")
}
