package dispatcher

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func Transmit(url string, data []byte) (int, error) {
	res, err := http.Post(url, "application/json",
		bytes.NewBuffer([]byte(data)),
	)
	if err == nil {
		return res.StatusCode, errors.New("NetworkFail")
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, errors.New("")
	}

	return res.StatusCode, nil
}
