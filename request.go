package ws

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://%s:%d/ws/pass"

func requests(host string, ob SendOb) (sendResponse, error) {
	// 跳过认证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	r, _ := json.Marshal(ob)
	reader := bytes.NewReader(r)

	client := &http.Client{Transport: tr}
	request, _ := http.NewRequest(
		"POST",
		fmt.Sprintf(url, host, config.WS.Port),
		reader,
	)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)

	if err != nil {
		return sendResponse{Status: false}, err
	}

	_body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var re sendResponse
	err = json.Unmarshal(_body, &re)

	if err != nil {
		return sendResponse{Status: false}, err
	}
	return re, nil
}
