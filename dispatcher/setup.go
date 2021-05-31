package dispatcher

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type sToken struct {
	Dingtalk string `json:"dingtalk"`
	Qywechat string `json:"qywechat"`
}

type sMirai struct {
	Groupid uint16 `json:"groupid"`
	Account uint16 `json:"account"`
	Authkey string `json:"authkey"`
	Http    string `json:"http"`
	Ws      string `json:"ws"`
}

type sConfig struct {
	Token   sToken `json:"token"`
	Mirai   sMirai `json:"mirai"`
	Weekend bool   `json:"weekend"`
	Google  bool   `json:"google"`
	Port    uint8  `json:"port"`
}

func Setup(path string) sConfig {
	jsonData, err := os.Open(path)
	if err != nil {
		log.Fatal("ReadConfigFile", err)
	}
	defer jsonData.Close()

	byteData, err := ioutil.ReadAll(jsonData)
	if err != nil {
		log.Fatal("ParseJsonFile", err)
	}

	config := sConfig{}
	json.Unmarshal([]byte(byteData), &config)
	return config
}
