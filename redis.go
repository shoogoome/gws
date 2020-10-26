package ws

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var redisCtl redisPool

type redisPool struct {
	pool *redis.Pool
}

const dnsKey = "websocket:dns:%s"

func (r *redisPool) GetHost(id string) (string, bool) {
	po := r.pool.Get()
	defer po.Close()

	_, _ = po.Do("select", config.Redis.Db)

	host, err := redis.String(po.Do("get", fmt.Sprintf(dnsKey, id)))
	if err != nil || len(host) == 0 {
		// TODO log
		return "", false
	}
	return host, true
}

func (r *redisPool) SetHost(id string, host string) {
	po := r.pool.Get()
	defer po.Close()

	_, _ = po.Do("select", config.Redis.Db)

	po.Do("set", fmt.Sprintf(dnsKey, id), host)
}

func (r *redisPool) Persistence(raw []byte) {
	po := r.pool.Get()
	defer po.Close()

	_, _ = po.Do("select", config.Redis.Db)

	po.Do("rpush", config.WS.PersistenceKey, raw)
}

func initRedis() {
	redisCtl = redisPool{
		pool: &redis.Pool{
			MaxIdle:     3,
			MaxActive:   50,
			IdleTimeout: 240 * time.Second,
			Dial:        dial,
		},
	}
}

func dial() (redis.Conn, error) {
	c, err := redis.Dial("tcp", config.Redis.Host)
	if err != nil {
		// TODO log
		time.Sleep(time.Second * 5)
		return dial()
	}

	if _, err := c.Do("AUTH", config.Redis.Password); err != nil {
		c.Close()
		panic(err)
	}
	return c, err
}
