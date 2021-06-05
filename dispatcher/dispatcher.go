package dispatcher

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/gorilla/websocket"

	"main/scheduler"
	"main/typer"
)

type sUrl struct {
	dingtalk string
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
	urls := sUrl{
		qq: fmt.Sprintf("%s/message?verifyKey=%s&qq=%s", config.Mirai.Ws, config.Mirai.Authkey, config.Mirai.Account),
	}

	if config.Dingtalk.Enable {
		urls.dingtalk = fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", config.Dingtalk.Webhook)
	}

	return &sDispatcher{
		urls: urls,
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
	if err := json.Unmarshal(msg, &data); err != nil {
		log.Println("ParseFail", err)
	}
	dis.transfer.session = data.Data.Session
	fmt.Printf(data.Data.Session)

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
	patternXml, _ := regexp.Compile(`url="[-a-zA-Z0-9@:%_+.~#?&//=]{2,256}.[a-z]{2,4}\b(/[-a-zA-Z0-9@:%_+.~#?&//=]*)?"`)

	data := typer.SOutgoing{}
	if err := json.Unmarshal(p, &data); err != nil {
		log.Println("ParseFail", err)
	}

	if data.Data.Type != "GroupMessage" {
		return
	}

	// 命令处理

	// 指定群聊转发，使用传入的 handler
	if data.Data.Sender.Group.Id == dis.transfer.Groupid {
		text := ""
		isImage := false

		for _, v := range data.Data.MessageChain {
			switch v["type"] {
			case "Image":
				isImage = true
				text += fmt.Sprintf("\n![图片](%s)", v["url"].(string))
			case "Plain":
				text += v["text"].(string)
			case "Face":
				text += fmt.Sprintf("[%s]", v["name"].(string))
			case "At":
				text += v["display"].(string)
			case "Xml":
				t := patternXml.FindString(v["xml"].(string))
				text += t[5 : len(t)-1]
			}
		}

		f(&scheduler.SSchedulerTask{
			Type:  typer.Enon,
			Time:  int64(data.Data.MessageChain[0]["time"].(float64)),
			Name:  data.Data.Sender.MemberName,
			Text:  text,
			Image: isImage,
		})
	}
}

/*
 * QQ 即时转发
 */
func (dis sDispatcher) Forward(task *scheduler.SSchedulerTask) {
	data, _ := json.Marshal(&typer.SIngoing{
		SyncId:  -1,
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

	if err := dis.transfer.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Println(err)
	}
}

/*
 * 快转发
 */
func (dis *sDispatcher) ImmedDispatch(task scheduler.SSchedulerTask) {
	bytesData, _ := json.Marshal(task)

	// Dingtalk
	if dis.urls.dingtalk != "" && task.Type != typer.Edingtalk {
		text := fmt.Sprintf(`QQ-%s-%s

%s`, task.Name, time.Unix(task.Time, 0).Format("01-02 15:04:05"), task.Text)

		var data interface{}
		if task.Image {
			data = typer.SDingIngoingMarkdown{
				Msgtype: "markdown",
				Markdown: typer.SDingInMarkdown{
					Title: "图文消息",
					Text:  text,
				},
			}
		} else {
			data = typer.SDingIngoingText{
				Msgtype: "text",
				Text: typer.SDingInText{
					Content: text,
				},
			}
		}

		if _, err := Transmit(dis.urls.dingtalk, data); err != nil {
			log.Println("PostFail", err)
		}
	}

	// Feishu
	if task.Type != typer.Efeishu {
		if _, err := Transmit(dis.urls.feishu, bytesData); err != nil {
			log.Println("PostFail", err)
		}
	}
}

/*
 * 慢转发
 */
func (dis *sDispatcher) DelayDispatch(tasks []scheduler.SSchedulerTask) {
	msg := ""
	for _, v := range tasks {
		msg += fmt.Sprintf(`## QQ-%s-%s
%s
`, v.Name, time.Unix(v.Time, 0).Format("01-02 15:04:05"), v.Text)
	}

	data := typer.SDingIngoingMarkdown{
		Msgtype: "markdown",
		Markdown: typer.SDingInMarkdown{
			Title: "消息列表",
			Text:  msg,
		},
	}

	if _, err := Transmit(dis.urls.dingtalk, data); err != nil {
		log.Println("PostFail", err)
	}
}
