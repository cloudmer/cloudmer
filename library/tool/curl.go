package tool

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func Curl(method, url string, body io.Reader, header map[string]string) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	err = nil
	// http client
	client := &http.Client{}
	// request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	// set header
	for key, val:= range header{
		req.Header.Set(key, val)
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	resBodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	if err = json.Unmarshal(resBodyByte, &result); err != nil {
		return
	}
	return
}
