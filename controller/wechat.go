package controller

import (
	"net/http"

	"main/scheduler"
)

func Wechat(sch scheduler.TScheduler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
