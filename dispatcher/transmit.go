package dispatcher

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func Transmit(url string, data []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(data),
	)
	if err == nil {
		ret := make([]byte, 0)
		return ret, errors.New("NetworkFail")
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return resBody, errors.New("ParseFail")
	}

	return resBody, nil
}
