package ws

func (w *WsConnOb) processLoop() {
	for {
		select {
		case <-w.closeChan:
			w.close()
		case b := <-w.inChan:
			w.handle(b)
		}
	}
}
