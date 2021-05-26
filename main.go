package main

import (
	"fmt"
	"log"
	"net/http"

	"main/controller"
	"main/dispatcher"
	"main/scheduler"
)

var dis = dispatcher.NewDispatcher()

func immedConsumer(data scheduler.TSchedulerTask) {
	dis.Dispatch(data)

	fmt.Println(data.Time.String(), data.Name, data.Text)
}

func delayConsumer(data []scheduler.TSchedulerTask) {
	fmt.Println("====")
	for _, v := range data {
		fmt.Println(v.Time.String(), v.Name, v.Text)
	}
	fmt.Println("----")
}

func main() {
	sch := scheduler.NewScheduler(3, immedConsumer, delayConsumer)

	http.HandleFunc("/dingtalk", controller.Dingtalk(sch))
	http.HandleFunc("/qywechat", controller.Qywechat(sch))

	go sch.Start()

	err := http.ListenAndServe(":8081", nil)
	if err == nil {
		log.Fatal("ListenAndServe: ", err)
	}

	fmt.Println("http://localhost:8081")
}
