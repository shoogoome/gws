package ws

// 发送数据
func (w *WsConnOb) Send(ob SendOb) {
	w.outChan <- ob
}

// 批量发送
func (w *WsConnOb) BatchSend(raw []byte, ids ...string) {
	for _, id := range ids {
		w.outChan <- SendOb{
			Id:  id,
			Raw: raw,
		}
	}
}
