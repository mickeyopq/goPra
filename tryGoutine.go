package tryGoroutine

import (
	"fmt"
	"time"
)

func say(s string, 牛喔 int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(牛喔) * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("哈", 1200)
	go say("幹", 700)
	go say("啊斯", 500)
	say("hello", 200)
}
