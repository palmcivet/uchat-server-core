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
	selfhost string
	feishu   string
	qq       string
}

type sTransfer struct {
	conn    *websocket.Conn
	session string
	sMirai
}

type Dispatcher interface {
	Start(sch *scheduler.Scheduler)
	Forward(task *scheduler.SSchedulerTask)
	ImmedDispatch(task scheduler.SSchedulerTask)
	DelayDispatch(tasks []scheduler.SSchedulerTask)
}

type sDispatcher struct {
	urls     sUrl
	transfer sTransfer
	weekend  bool
	google   bool
}

func NewDispatcher(config sConfig) Dispatcher {
	return &sDispatcher{
		urls: sUrl{
			dingtalk: fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", config.Token.Dingtalk),
			qywechat: fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", config.Token.Qywechat),
			qq:       fmt.Sprintf("%s/message?verifyKey=%s&qq=%s", config.Mirai.Ws, config.Mirai.Authkey, config.Mirai.Account),
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

/*
 * 建立 WS 连接
 */
func (dis *sDispatcher) Start(sch *scheduler.Scheduler) {
	var err error

	dis.transfer.conn, _, err = websocket.DefaultDialer.Dial(dis.urls.qq, nil)
	if err != nil {
		log.Fatal("WebSocketFail", err)
	}

	_, msg, err := dis.transfer.conn.ReadMessage()
	if err != nil {
		log.Println("RecevieFail", err)
	}

	data := typer.SResponse{}
	err = json.Unmarshal(msg, &data)
	if err != nil {
		log.Println("ParseFail", err)
	}
	dis.transfer.session = data.Data.Session

	fmt.Println("连接到 Mirai")

	go func() {
		defer dis.transfer.conn.Close()

		for {
			_, msg, err := dis.transfer.conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
			}

			dis.receive(msg, (*sch).Produce)
		}
	}()
}

/*
 * QQ 消息网关
 * 处理指令的核心构件
 */
func (dis sDispatcher) receive(p []byte, f func(*scheduler.SSchedulerTask)) {
	data := typer.SOutgoing{}
	err := json.Unmarshal(p, &data)
	if err != nil {
		fmt.Println("ParseFail", err)
	}

	if data.Data.Type != "GroupMessage" {
		return
	}

	// 命令处理

	// 指定群聊转发，使用传入的 handler
	if data.Data.Sender.Group.Id == dis.transfer.Groupid {
		f(&scheduler.SSchedulerTask{
			Type: typer.Enon,
			Time: int64(data.Data.MessageChain[0]["time"].(float64)) * 1000,
			Name: data.Data.Sender.MemberName,
			Text: data.Data.MessageChain[1]["text"].(string),
		})
	}
}

/*
 * QQ 即时转发
 */
func (dis sDispatcher) Forward(task *scheduler.SSchedulerTask) {
	dis.transfer.conn.WriteJSON(&typer.SIngoing{
		SyncId:  0,
		Command: "sendGroupMessage",
		Content: typer.SInGroup{
			SessionKey: dis.transfer.session,
			Target:     dis.transfer.Groupid,
			MessageChain: []interface{}{
				typer.SMessagePlain{
					Type: "Plain",
					Text: task.Text,
				},
			},
		},
	})
}

/*
 * 快转发
 */
func (dis *sDispatcher) ImmedDispatch(task scheduler.SSchedulerTask) {
	bytesData, _ := json.Marshal(task)

	fmt.Println(task.Time, task.Name, task.Text)

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

	// CLI/Web
	if task.Type != typer.ESelfhost {
		_, err := Transmit(dis.urls.selfhost, bytesData)
		if err != nil {
			log.Println("PostFail", err)
		}
	}
}

/*
 * 慢转发
 */
func (dis *sDispatcher) DelayDispatch(tasks []scheduler.SSchedulerTask) {
	fmt.Println("====")
	for _, v := range tasks {
		fmt.Println(v.Time, v.Name, v.Text)
	}
	fmt.Println("----")
}
