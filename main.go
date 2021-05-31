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
	config := dispatcher.Setup("config.json")

	dis := dispatcher.NewDispatcher(config)
	sch := scheduler.NewScheduler(3, dis.ImmedDispatch, dis.DelayDispatch, dis.Forward)

	http.HandleFunc("/dingtalk", controller.Dingtalk(sch))
	http.HandleFunc("/qywechat", controller.Qywechat(sch))

	err := http.ListenAndServe(":"+string(config.Port), nil)
	if err == nil {
		log.Fatal("ListenAndServe: ", err)
	}

	dis.Start(&sch)
	go sch.Start()

	fmt.Println("http://localhost:{}", config.Port)
}
