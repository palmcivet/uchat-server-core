package typer

type SResponseData struct {
	Code    uint8  `json:"code"`
	Session string `json:"session"`
}

type SResponse struct {
	SyncId uint8         `json:"syncId"`
	Data   SResponseData `json:"data"`
}

type SOutMessage struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	Url  string `json:"url,omitempty"`
}

type SOutGroup struct {
	SessionKey   string        `json:"sessionKey"`
	Target       int16         `json:"target"`
	MessageChain []SOutMessage `json:"messageChain"`
}

type SOutgoing struct {
	SyncId  int8      `json:"SyncId"`
	Command string    `json:"command"`
	Content SOutGroup `json:"content"`
}
