package typer

import "time"

type SMessage struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	Url  string `json:"url,omitempty"`
}

type SResponseData struct {
	Code    uint8  `json:"code"`
	Session string `json:"session"`
}

type SResponse struct {
	SyncId uint8         `json:"syncId"`
	Data   SResponseData `json:"data"`
}

type SOutDataGroup struct {
	Id         uint16 `json:"id"`
	Name       string `json:"name"`
	Nermission string `json:"permission"`
}

type SOutDataSender struct {
	Id                 uint16
	MemberName         string        `json:"memberName"`
	SpecialTitle       string        `json:"specialTitle"`
	Permission         string        `json:"permission"`
	JoinTimestamp      time.Time     `json:"joinTimestamp"`
	LastSpeakTimestamp time.Time     `json:"lastSpeakTimestamp"`
	MuteTimeRemaining  time.Time     `json:"muteTimeRemaining"`
	Group              SOutDataGroup `json:"group"`
}

type SOutData struct {
	Type         string         `json:"type"`
	Sender       SOutDataSender `json:"sender"`
	MessageChain []SMessage     `json:"messageChain"`
}

type SOutgoing struct {
	SyncId uint8    `json:"syncId"`
	Data   SOutData `json:"data"`
}

type SInGroup struct {
	SessionKey   string     `json:"sessionKey"`
	Target       uint16     `json:"target"`
	MessageChain []SMessage `json:"messageChain"`
}

type SIngoing struct {
	SyncId     int8     `json:"SyncId"`
	Command    string   `json:"command"`
	SubCommand string   `json:"subCommand"`
	Content    SInGroup `json:"content"`
}
