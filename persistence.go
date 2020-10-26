package ws

type _persistence map[string][]byte

var persistence _persistence

func init() {
	persistence = make(map[string][]byte)
}

// 落地数据
func (p _persistence) do(id string) {

	v, ok := persistence[id]
	if !ok {
		// TODO log
		return
	}

	redisCtl.Persistence(v)
}
