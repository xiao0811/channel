package main

import (
	"fmt"
	"sync"
	"time"
)

// Data struct
type Data struct {
	X int
	Y int
}

/*
一个标准的协程+信道实现
*/
func main() {
	start := time.Now()
	taskChan := make(chan *Data)
	TCount := 10
	var wg sync.WaitGroup
	// 产生任务
	go func() {
		for i := 0; i < 1000; i++ {
			taskChan <- &Data{
				X: i,
				Y: 2 * i,
			}
		} // 全部任务都输入后关闭信道，告诉工作者进程没有新任务了。
		close(taskChan)
	}() // 启动 TCount 个协程执行任务
	wg.Add(TCount)
	for i := 0; i < TCount; i++ { // 注意：如果协程内使用了 i，必须有这一步，或者选择通过参数传递进协程。
		// 否则 i 会被 for 所在的协程修改，协程实际使用时值并不确定。
		go func() { // 协程结束时报告当前协程执行完毕。
			defer func() { wg.Done() }()

			// fmt.Printf("工作者 %v 启动...\r\n", i)
			for task := range taskChan { // 建立匿名函数执行任务的目的是为了捕获单个任务崩溃，防止造成整个工作者、系统崩溃。
				fmt.Println(task)
			}

			// fmt.Printf("工作者 %v 结束。\r\n", i)
		}()
	} //等待所有任务完成
	wg.Wait()
	fmt.Println("全部任务结束", time.Now().Sub(start))
}
