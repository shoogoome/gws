package ws

// 发送数据
func (w *WsConnOb) Send(ob SendOb, dbRaw ...[]byte) {
	if config.WS.Persistence && len(dbRaw) > 0 {
		ob.pid = getRandomString(6)
		persistence[ob.pid] = dbRaw[0]
	}
	w.outChan <- ob
}

// 批量发送
func (w *WsConnOb) BatchSend(raw []byte, ids []string, dbRaw ...[]byte) {

	first := true
	for _, id := range ids {
		s := SendOb{
			Id:  id,
			Raw: raw,
		}
		if config.WS.Persistence && len(dbRaw) > 0 && first {
			s.pid = getRandomString(6)
			persistence[s.pid] = dbRaw[0]
			first = false
		}
		w.outChan <- s
	}
}
