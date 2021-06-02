package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"main/controller"
	"main/dispatcher"
	"main/scheduler"
)

func init() {
	file, err := os.OpenFile("main.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func main() {
	config := dispatcher.Setup("target/config.json")

	dis := dispatcher.NewDispatcher(config)
	sch := scheduler.NewScheduler(3, dis.ImmedDispatch, dis.DelayDispatch, dis.Forward)

	http.HandleFunc("/dingtalk", controller.Dingtalk(sch))
	http.HandleFunc("/qywechat", controller.Qywechat(sch))

	sch.Start()
	dis.Start(&sch)

	fmt.Println("http://localhost:" + config.Port)

	err := http.ListenAndServe(":"+string(config.Port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
