package ws

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 跨主机通讯handle
func PassHandle(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var sendOb SendOb
	response := sendResponse{Status: true}
	json.Unmarshal(body, &sendOb)

	if wc, ok := conn[sendOb.Id]; ok {
		sendLocal(wc, sendOb)
	} else {
		response.Status = false
		response.Msg = fmt.Sprintf("no found id:%s connect", sendOb.Id)
	}

	raw, _ := json.Marshal(response)
	w.Write(raw)
}
