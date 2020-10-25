package ws

import "time"

var config Config

func init() {
	config = Config{
		WS: WsConfig{
			Port:            80,
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			Cors:            true,
			InChanLength:    500,
			OutChanLength:   500,
			MessageSize:     1024,
			PingPeriod:      6 * 9 * time.Second,
			PongWait:        60 * time.Second,
		},
		Redis: RedisConfig{
			Host: "127.0.0.1:2379",
			Db:   0,
		},
	}
}

func SetConfig(c Config) {
	config = c
}
