package ws

import (
	"fmt"
	"time"
)

// log
func log(msg ...interface{}) {
	fmt.Println(append([]interface{}{time.Now().Format("2006-01-02 15:04:05")}, msg...)...)
}
