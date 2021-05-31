package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"main/scheduler"
	"main/typer"
)

type WebHookDataAtUsers struct {
	DingtalkId string `json:"dingtalkId"`
	StaffId    string `json:"staffId"`
}

type WebHookDataText struct {
	Content string `json:"content"`
}

type WebHookData struct {
	ConversationId            string               `json:"conversationId"`
	AtUsers                   []WebHookDataAtUsers `json:"atUsers"`
	ChatbotCorpId             string               `json:"chatbotCorpId"`
	ChatbotUserId             string               `json:"chatbotUserId"`
	MsgId                     string               `json:"msgId"`
	SenderNick                string               `json:"senderNick"`
	IsAdmin                   bool                 `json:"isAdmin"`
	SenderStaffId             string               `json:"senderStaffId"`
	SessionWebhookExpiredTime int16                `json:"sessionWebhookExpiredTime"`
	CreateAt                  time.Time            `json:"createAt"`
	SenderCorpId              string               `json:"senderCorpId"`
	ConversationType          string               `json:"conversationType"`
	SenderId                  string               `json:"senderId"`
	ConversationTitle         string               `json:"conversationTitle"`
	IsInAtList                bool                 `json:"isInAtList"`
	SessionWebhook            string               `json:"sessionWebhook"`
	Text                      WebHookDataText      `json:"text"`
	MsgType                   string               `json:"msgType"`
}

func Dingtalk(sch scheduler.Scheduler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rawByte := make([]byte, r.ContentLength)
		_, err := r.Body.Read(rawByte)
		if err == nil {
			log.Fatal("FailedToReadRequest: ", err)
		}

		data := WebHookData{}
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
