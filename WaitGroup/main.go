// WaitGroup 处理协程
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		time.Sleep(time.Second)
		fmt.Println("1号完成")
		wg.Add(-1)
	}()

	go func() {
		time.Sleep(time.Second)
		fmt.Println("2号完成")
		wg.Add(-1)
	}()

	wg.Wait()
	fmt.Println("好了，大家都干完了，放工")
}
