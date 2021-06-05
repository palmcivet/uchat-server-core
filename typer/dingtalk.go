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

/* Dingtalk 推送的消息格式 */

type SDingOutAt struct {
	DingtalkId string `json:"dingtalkId"`
	StaffId    string `json:"staffId"`
}

type SDingOutText struct {
	Content string `json:"content"`
}

type SDingOutGoing struct {
	ConversationId            string       `json:"conversationId"`
	AtUsers                   []SDingOutAt `json:"atUsers"`
	ChatbotCorpId             string       `json:"chatbotCorpId"`
	ChatbotUserId             string       `json:"chatbotUserId"`
	MsgId                     string       `json:"msgId"`
	SenderNick                string       `json:"senderNick"`
	IsAdmin                   bool         `json:"isAdmin"`
	SenderStaffId             string       `json:"senderStaffId"`
	SessionWebhookExpiredTime int16        `json:"sessionWebhookExpiredTime"`
	CreateAt                  int64        `json:"createAt"`
	SenderCorpId              string       `json:"senderCorpId"`
	ConversationType          string       `json:"conversationType"`
	SenderId                  string       `json:"senderId"`
	ConversationTitle         string       `json:"conversationTitle"`
	IsInAtList                bool         `json:"isInAtList"`
	SessionWebhook            string       `json:"sessionWebhook"`
	Text                      SDingOutText `json:"text"`
	MsgType                   string       `json:"msgType"`
}
