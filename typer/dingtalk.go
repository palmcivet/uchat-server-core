package typer

/* 转入钉钉的消息格式 */

type SDingInAt struct {
	AtMobiles []string `json:"atMobiles"`
	AtUserIds []string `json:"atUserIds"`
	IsAtAll   bool     `json:"isAtAll"`
}

type SDingInText struct {
	Content string `json:"content"`
}

type SDingInMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type SDingIngoingText struct {
	At      SDingInAt   `json:"at"`
	Text    SDingInText `json:"text"`
	Msgtype string      `json:"msgtype"`
}

type SDingIngoingMarkdown struct {
	At       SDingInAt       `json:"at"`
	Markdown SDingInMarkdown `json:"markdown"`
	Msgtype  string          `json:"msgtype"`
}
