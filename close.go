package ws

func (w *WsConnOb) close() {

	w.connect.Close()
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if !w.isClosed {
		w.isClosed = true

		delete(conn, w.Id)
		close(w.closeChan)
	}

	log("id: ", w.Id, "连接关闭")
}
