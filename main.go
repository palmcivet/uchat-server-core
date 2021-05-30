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
	sch := scheduler.NewScheduler(3, dis.ImmedDispatch, dis.DelayDispatch)

	http.HandleFunc("/dingtalk", controller.Dingtalk(sch, dis))
	http.HandleFunc("/qywechat", controller.Qywechat(sch, dis))

	dis.HandleCore(controller.Qq(sch))

	go sch.Start()

	err := http.ListenAndServe(":"+string(config.Port), nil)
	if err == nil {
		log.Fatal("ListenAndServe: ", err)
	}

	fmt.Println("http://localhost:{}", config.Port)
}
