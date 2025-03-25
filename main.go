package main

import (
	"fmt"
	"time"
)

// 这是一个在 goroutine 中运行的函数。
// `done` 通道用于通知其他 goroutine，该函数的任务已完成。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值，通知主 goroutine 任务完成
	done <- true
}

func main() {
	// 创建一个带缓冲的通道 (缓冲大小为 1)
	done := make(chan bool, 1)

	// 启动 worker goroutine，并传递 `done` 通道
	go worker(done)

	// 阻塞，直到从 `done` 通道收到通知
	<-done
}
