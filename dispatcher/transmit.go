package dispatcher

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func Transmit(url string, v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, errors.New("NetworkFail")
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return resBody, errors.New("ParseFail")
	}

	return resBody, nil
}
