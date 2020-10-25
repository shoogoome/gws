package ws

// 发送数据
func (w *WsConnOb) Send(ob SendOb) {
	w.outChan <- ob
}
