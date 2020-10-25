package ws

// 连接信息
var conn map[string]*WsConnOb

// dns信息 跨主机ws对象记录
// 若内存记录dns通讯中失效则通过redis更新
var dns map[string]string

func init() {
	conn = make(map[string]*WsConnOb)
	dns = make(map[string]string)
}
