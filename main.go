package main

import (
	"context"
	"fmt"
	"time"
)

func doTask(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second): // 任务完成
		fmt.Println("Task completed")
	case <-ctx.Done(): // 监听取消信号
		fmt.Println("Task cancelled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go doTask(ctx)

	time.Sleep(4 * time.Second)
}
