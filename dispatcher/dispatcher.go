package dispatcher

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"

	"main/scheduler"
	"main/typer"
)

type sUrl struct {
	dingtalk string
	qywechat string
	feishu   string
	cli      string
	web      string
	qq       string
}

type sTransfer struct {
	conn    *websocket.Conn
	session string
	sMirai
}

type TDispatcher interface {
	HandleCore(fun func(p []byte))
	QQDispatch(task scheduler.TSchedulerTask)
	ImmedDispatch(task scheduler.TSchedulerTask)
	DelayDispatch(tasks []scheduler.TSchedulerTask)
}

type sDispatcher struct {
	urls     sUrl
	transfer sTransfer
	weekend  bool
	google   bool
}

func NewDispatcher(config sConfig) TDispatcher {
	return &sDispatcher{
		urls: sUrl{
			dingtalk: fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", config.Token.Dingtalk),
			qywechat: fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", config.Token.Qywechat),
			qq:       fmt.Sprintf("%s/message?verifyKey=%s&qq=%d", config.Mirai.Ws, config.Mirai.Authkey, config.Mirai.Account),
		},
		transfer: sTransfer{
			sMirai: sMirai{
				Groupid: config.Mirai.Groupid,
				Account: config.Mirai.Account,
				Authkey: config.Mirai.Authkey,
				Http:    config.Mirai.Http,
				Ws:      config.Mirai.Ws,
			},
		},
		weekend: config.Weekend,
		google:  config.Google,
	}
}

func (dis *sDispatcher) HandleCore(receiver func(p []byte)) {
	var err error

	dis.transfer.conn, _, err = websocket.DefaultDialer.Dial(dis.urls.qq, nil)
	if err != nil {
		log.Fatal("WebSocketFail", err)
	}
	defer dis.transfer.conn.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, msg, err := dis.transfer.conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
			}
			receiver(msg)
		}
	}()
}

func (dis sDispatcher) QQDispatch(task scheduler.TSchedulerTask) {
	dis.transfer.conn.WriteJSON(&typer.SOutgoing{
		SyncId:  0,
		Command: "sendGroupMessage",
		Content: typer.SOutGroup{
			SessionKey: dis.transfer.session,
			Target:     dis.transfer.Groupid,
			MessageChain: []typer.SOutMessage{
				{
					Type: "Plain",
					Text: task.Text,
				},
			},
		},
	})
}

/*
 * Transmit 到各种平台
 */
func (dis *sDispatcher) ImmedDispatch(task scheduler.TSchedulerTask) {
	bytesData, _ := json.Marshal(task)

	fmt.Println(task.Time.String(), task.Name, task.Text)

	// Dingtalk
	if task.Type != typer.Edingtalk {
		_, err := Transmit(dis.urls.dingtalk, bytesData)
		if err != nil {
			log.Println("PostFail", err)
		}
	}

	// QYWechat
	if task.Type != typer.Eqywechat {
		_, err := Transmit(dis.urls.qywechat, bytesData)
		if err != nil {
			log.Println("PostFail", err)
		}
	}

	// Feishu
	if task.Type != typer.Efeishu {
		_, err := Transmit(dis.urls.feishu, bytesData)
		if err != nil {
			log.Println("PostFail", err)
		}
	}

	// CLI
	if task.Type != typer.Ecli {
		_, err := Transmit(dis.urls.cli, bytesData)
		if err != nil {
			log.Println("PostFail", err)
		}
	}

	// Web
	if task.Type != typer.Eweb {
		_, err := Transmit(dis.urls.web, bytesData)
		if err != nil {
			log.Println("PostFail", err)
		}
	}
}

func (dis *sDispatcher) DelayDispatch(tasks []scheduler.TSchedulerTask) {
	fmt.Println("====")
	for _, v := range tasks {
		fmt.Println(v.Time.String(), v.Name, v.Text)
	}
	fmt.Println("----")
}
