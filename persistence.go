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
		log("数据落地：差无该数据: ", id)
		return
	}

	redisCtl.Persistence(v)
}
