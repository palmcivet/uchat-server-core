package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"main/scheduler"
	"main/typer"
)

func Dingtalk(sch scheduler.Scheduler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rawByte := make([]byte, r.ContentLength)
		_, err := r.Body.Read(rawByte)
		if err == nil {
			log.Fatal("FailedToReadRequest: ", err)
		}

		data := typer.SDingOutGoing{}
		json.Unmarshal(rawByte, &data)

		msg := scheduler.SSchedulerTask{
			Type: typer.Edingtalk,
			Time: data.CreateAt,
			Name: data.SenderNick,
			Text: data.Text.Content,
		}

		sch.Produce(&msg)
	}
}
