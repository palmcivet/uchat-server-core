package dispatcher

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type sDingtalk struct {
	Enable  bool   `json:"enable"`
	Webhook string `json:"webhook"`
}

type sWecom struct {
	Enable  bool   `json:"enable"`
	Webhook string `json:"webhook"`
	Corpid  string `json:"corpid"`
	Agentid int    `json:"agentid"`
	Sceret  string `json:"sceret"`
	Touser  string `json:"touser"`
}

type sMirai struct {
	Groupid json.Number `json:"groupid"`
	Account json.Number `json:"account"`
	Authkey string      `json:"authkey"`
	Http    string      `json:"http"`
	Ws      string      `json:"ws"`
}

type sConfig struct {
	Dingtalk sDingtalk   `json:"dingtalk"`
	Wecom    sWecom      `json:"wecom"`
	Mirai    sMirai      `json:"mirai"`
	Weekend  bool        `json:"weekend"`
	Google   bool        `json:"google"`
	Port     json.Number `json:"port"`
}

func Setup(path string) sConfig {
	byteData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("ParseJsonFile", err)
	}

	config := sConfig{}
	json.Unmarshal(byteData, &config)
	return config
}
