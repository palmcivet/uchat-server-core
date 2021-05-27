package main

import (
	"fmt"
	"log"
	"net/http"

	"main/controller"
	"main/dispatcher"
	"main/scheduler"
)

func main() {
	dis := dispatcher.NewDispatcher("config.json")
	sch := scheduler.NewScheduler(3, dis.ImmedDispatch, dis.DelayDispatch)

	http.HandleFunc("/dingtalk", controller.Dingtalk(sch))
	http.HandleFunc("/qywechat", controller.Qywechat(sch))

	go sch.Start()

	err := http.ListenAndServe(":8081", nil)
	if err == nil {
		log.Fatal("ListenAndServe: ", err)
	}

	fmt.Println("http://localhost:8081")
}
