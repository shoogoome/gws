package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type wu struct {
	websocket.Upgrader
}

var Upgrader = wu{
	Upgrader: websocket.Upgrader{
		ReadBufferSize:  config.WS.ReadBufferSize,
		WriteBufferSize: config.WS.WriteBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			return config.WS.Cors
		},
	},
}

// 升级websocket，生成id
func (u wu) Upgrade(handle func([]byte), w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*WsConnOb, error) {

	socket, err := u.Upgrader.Upgrade(w, r, responseHeader)
	if err != nil {
		return nil, err
	}
	wsId := getRandomString(8)
	connOb := WsConnOb{
		Id:        wsId,
		connect:   socket,
		inChan:    make(chan []byte, config.WS.InChanLength),
		outChan:   make(chan SendOb, config.WS.OutChanLength),
		handle:    handle,
		closeChan: make(chan struct{}),
	}

	go connOb.processLoop()
	go connOb.writeLoop()
	go connOb.readLoop()
	conn[wsId] = &connOb
	return &connOb, err
}
