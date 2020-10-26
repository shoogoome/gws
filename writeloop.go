package ws

import (
	"time"

	"github.com/gorilla/websocket"
)

func (w *WsConnOb) writeLoop() {
	ticker := time.NewTicker(config.WS.PingPeriod)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case sendOb := <-w.outChan:
			if wc, ok := conn[sendOb.Id]; ok {
				sendLocal(wc, sendOb)
			} else {
				sendNetwork(sendOb, true)
			}
		case <-w.closeChan:
			return
		case <-ticker.C:
			w.connect.SetWriteDeadline(time.Now().Add(writeWait))
			if err := w.connect.WriteMessage(websocket.PingMessage, nil); err != nil {
				w.close()
				return
			}
		}
	}
}

// 本地通讯
func sendLocal(w *WsConnOb, sendOb SendOb) {
	if err := w.connect.WriteMessage(websocket.TextMessage, sendOb.Raw); err != nil {
		// TODO log
	}
}

// 跨主机通讯
func sendNetwork(sendOb SendOb, first bool) {
	if host, ok := dns[sendOb.Id]; ok {
		response, err := requests(host, sendOb)
		if err != nil {
			// TODO log
		}
		if !response.Status {
			// TODO response.Msg
		}
		// 正常结束
		return
	}
	// 查不到或者上面过程出现问题则从redis
	// 更新dns信息并重新请求
	if first {
		updateDns(sendOb.Id)
		sendNetwork(sendOb, false)
	} else {
		// TODO log
	}
}

// 更新dns记录
func updateDns(id string) {
	host, ok := redisCtl.GetHost(id)
	if ok {
		dns[id] = host
	} else {
		// TODO log
	}
}