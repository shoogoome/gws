package ws

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second // 允许等待的写入时间
)

// 系统配置
type Config struct {
	WS    WsConfig    // ws配置
	Redis RedisConfig // redis配置
}

type WsConfig struct {
	Port            int
	ReadBufferSize  int
	WriteBufferSize int
	Cors            bool
	InChanLength    int
	OutChanLength   int
	MessageSize     int64
	PingPeriod      time.Duration
	PongWait        time.Duration
	Persistence     bool
	PersistenceKey  string
	HostKey         string
}

type RedisConfig struct {
	Host     string
	Password string
	Db       int
}

// 发送数据对象
type SendOb struct {
	Id  string
	Raw []byte
	pid string // 持久化id
}

type sendResponse struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

// ws连接对象
type WsConnOb struct {
	Id        string
	connect   *websocket.Conn
	inChan    chan []byte
	outChan   chan SendOb
	closeChan chan struct{}
	mutex     sync.Mutex
	isClosed  bool
}
