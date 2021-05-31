package controller

import (
	"net/http"

	"main/scheduler"
)

func Qywechat(sch scheduler.Scheduler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
