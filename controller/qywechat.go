package controller

import (
	"net/http"

	"main/dispatcher"
	"main/scheduler"
)

func Qywechat(sch scheduler.TScheduler, dis dispatcher.TDispatcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
