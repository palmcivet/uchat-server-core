package controller

import (
	"log"
	"net/http"

	"main/scheduler"
)

func Feishu(sch scheduler.Scheduler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rawByte := make([]byte, r.ContentLength)
		_, err := r.Body.Read(rawByte)
		if err == nil {
			log.Fatal("FailedToReadRequest: ", err)
		}
	}
}
