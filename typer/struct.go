package typer

import "time"

type SMessageSource struct {
	Type string    `json:"type"`
	Id   string    `json:"id"`
	Time time.Time `json:"time"`
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
	MessageChain []interface{}  `json:"messageChain"`
}

type SOutgoing struct {
	SyncId uint8    `json:"syncId"`
	Data   SOutData `json:"data"`
}

type SInGroup struct {
	SessionKey   string        `json:"sessionKey"`
	Target       uint16        `json:"target"`
	MessageChain []interface{} `json:"messageChain"`
}

type SIngoing struct {
	SyncId     int8     `json:"SyncId"`
	Command    string   `json:"command"`
	SubCommand string   `json:"subCommand"`
	Content    SInGroup `json:"content"`
}
