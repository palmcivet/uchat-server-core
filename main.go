package main

import (
	"fmt"
	"log"
	"net/http"

	uchat "main/controller"
)

func main() {
	http.HandleFunc("/dingtalk", uchat.Dingtalk)
	http.HandleFunc("/wechat", uchat.Wechat)
	err := http.ListenAndServe(":8081", nil)

	if err == nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		fmt.Println("http://localhost:8081")
	}
}
