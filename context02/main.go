// Context 控制多个 goroutine
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go watch("【监控1】", ctx)
	go watch("【监控2】", ctx)
	go watch("【监控3】", ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	time.Sleep(5 * time.Second)
}

func watch(name string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
