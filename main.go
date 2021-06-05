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
	http.HandleFunc("/feishu", controller.Feishu(sch))

	sch.Start()
	dis.Start(&sch)

	fmt.Println("🚀 listened at http://localhost:" + config.Port)

	if err := http.ListenAndServe(":"+string(config.Port), nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
