package typer

import (
	"encoding/json"
)

/* MessageChain 类型 */

type SMessageSource struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Time int64  `json:"time"`
}

type SMessageQuote struct {
	Type     string        `json:"type"`
	Id       string        `json:"id"`
	Groupid  string        `json:"groupId"`
	SenderId string        `json:"senderId"`
	TargetId string        `json:"targetId"`
	Origin   []interface{} `json:"origin"`
}

type SMessageImage struct {
	Type    string `json:"type"`
	ImageId string `json:"imageId"`
	Url     string `json:"url"`
	Path    string `json:"path"`
	Base64  string `json:"base64"`
}

type SMessagePlain struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

/* WS 连接建立时的响应体 */

type SResponseData struct {
	Code    json.Number `json:"code"`
	Session string      `json:"session"`
}

type SResponse struct {
	SyncId string        `json:"syncId"`
	Data   SResponseData `json:"data"`
}

/* 从 mirai 推送的消息格式 */

type SOutDataGroup struct {
	Id         json.Number `json:"id"`
	Name       string      `json:"name"`
	Nermission string      `json:"permission"`
}

type SOutDataSender struct {
	Id                 json.Number   `json:"id"`
	MemberName         string        `json:"memberName"`
	SpecialTitle       string        `json:"specialTitle"`
	Permission         string        `json:"permission"`
	JoinTimestamp      int           `json:"joinTimestamp"`
	LastSpeakTimestamp int           `json:"lastSpeakTimestamp"`
	MuteTimeRemaining  int           `json:"muteTimeRemaining"`
	Group              SOutDataGroup `json:"group"`
}

type SOutData struct {
	Type         string                   `json:"type"`
	Sender       SOutDataSender           `json:"sender"`
	MessageChain []map[string]interface{} `json:"messageChain"`
}

type SOutgoing struct {
	SyncId string   `json:"syncId"`
	Data   SOutData `json:"data"`
}

/* 推送到 mirai 的消息格式 */

type SInGroup struct {
	SessionKey   string        `json:"sessionKey"`
	Target       json.Number   `json:"target"`
	MessageChain []interface{} `json:"messageChain"`
}

type SIngoing struct {
	SyncId     int8     `json:"syncId"`
	Command    string   `json:"command"`
	SubCommand string   `json:"subCommand,omitempty"`
	Content    SInGroup `json:"content"`
}
