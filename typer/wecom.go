package typer

type SWecomInText struct {
	Content string `json:"content"`
}

type SWecomIngoingText struct {
	Touser  string       `json:"touser"`
	Msgtype string       `json:"msgtype"`
	Agentid string       `json:"agentid"`
	Text    SWecomInText `json:"text"`
}
