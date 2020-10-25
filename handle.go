package ws

func (w *WsConnOb) SetHandle(handle func([]byte, interface{})) {
	w.handle = handle
}
