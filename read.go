package ws

func (w *WsConnOb) Read() []byte {

	select {
	case data := <-w.inChan:
		return data
	case <-w.closeChan:
		return nil
	}
}
