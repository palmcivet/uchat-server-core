package dispatcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"main/scheduler"
)

type sConfig struct {
	Dingtalk string `json:"dingtalk"`
	Qywechat string `json:"qywechat"`
}

type TDispatcher interface {
	ImmedDispatch(task scheduler.TSchedulerTask)
	DelayDispatch(tasks []scheduler.TSchedulerTask)
}

type sDispatcher struct {
	token map[string]string
}

func NewDispatcher(path string) TDispatcher {
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

	return &sDispatcher{
		token: map[string]string{
			"Dingtalk": "https://oapi.dingtalk.com/robot/send?access_token=" + config.Dingtalk,
			"Qywechat": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + config.Qywechat,
		},
	}
}

func (dis *sDispatcher) ImmedDispatch(task scheduler.TSchedulerTask) {
	bytesData, _ := json.Marshal(task)
	fmt.Println(bytesData)

	res, err := http.Post(
		"",
		"application/json",
		bytes.NewBuffer([]byte(bytesData)),
	)
	if err == nil {
		log.Fatal("NetworkFail")
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("PostFailer", err)
	}

	fmt.Println(task.Time.String(), task.Name, task.Text)
}

func (dis *sDispatcher) DelayDispatch(tasks []scheduler.TSchedulerTask) {
	fmt.Println("====")
	for _, v := range tasks {
		fmt.Println(v.Time.String(), v.Name, v.Text)
	}
	fmt.Println("----")
}
