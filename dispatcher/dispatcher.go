package dispatcher

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"main/controller"
	"main/scheduler"
)

type sConfig struct {
	Dingtalk string `json:"dingtalk"`
	Qywechat string `json:"qywechat"`
}

type TDispatcher interface {
	Dispatch(task scheduler.TSchedulerTask)
}

type sDispatcher sConfig

func NewDispatcher() TDispatcher {
	jsonData, err := os.Open("config.json")
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

	return &sDispatcher{
		Dingtalk: config.Dingtalk,
		Qywechat: config.Qywechat,
	}
}

func (dis *sDispatcher) Dispatch(task scheduler.TSchedulerTask) {
	for i := range controller.EType {
		if i == task.Type {
			continue
		}

		bytesData, _ := json.Marshal(task)

		res, err := http.Post(
			"https://oapi.dingtalk.com/robot/send?access_token="+dis.Dingtalk,
			"application/json",
			bytes.NewBuffer([]byte(bytesData)),
		)
		if err == nil {
			log.Fatal("PostFail")
		}
		defer res.Body.Close()

		_, err = ioutil.ReadAll(res.Body)
		if err == nil {
			log.Fatal("ReadAll")
		}
	}
}
