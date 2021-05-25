package main

import (
	"fmt"
	"log"
	"net/http"

	"main/controller"
	"main/scheduler"
)

func immed(data scheduler.TSchedulerTask) {
	fmt.Println(data.Time.String(), data.Name, data.Text)
}

func delay(data []scheduler.TSchedulerTask) {
	fmt.Println("====")
	for _, v := range data {
		fmt.Println(v.Time.String(), v.Name, v.Text)
	}
	fmt.Println("----")
}

func main() {
	sch := scheduler.NewScheduler(3, immed, delay)

	http.HandleFunc("/dingtalk", controller.Dingtalk(sch))
	http.HandleFunc("/wechat", controller.Wechat(sch))

	go sch.Start()

	err := http.ListenAndServe(":8081", nil)

	if err == nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		fmt.Println("http://localhost:8081")
	}
}
