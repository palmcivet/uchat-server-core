package dispatcher

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type sToken struct {
	Dingtalk string `json:"dingtalk"`
	Qywechat string `json:"qywechat"`
}

type sMirai struct {
	Groupid json.Number `json:"groupid"`
	Account json.Number `json:"account"`
	Authkey string      `json:"authkey"`
	Http    string      `json:"http"`
	Ws      string      `json:"ws"`
}

type sConfig struct {
	Token   sToken      `json:"token"`
	Mirai   sMirai      `json:"mirai"`
	Weekend bool        `json:"weekend"`
	Google  bool        `json:"google"`
	Port    json.Number `json:"port"`
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
