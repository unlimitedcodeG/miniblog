package main

import "fmt"

func main() {
	// 带缓冲的 channel，容量为 5
	jobs := make(chan int, 5)
	// 用于通知主 goroutine 所有工作已完成
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			// j 就是value 就是job的值  more 只不过是通道是否关闭 如果还在传递 就是true
			// 反之就是false
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("recevied all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	// 	关闭 jobs 通道，表示不会再有新任务发送。

	// 被关闭的通道仍然可以接收数据，直到所有数据读取完毕。
	fmt.Println("sent all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("recevied more jobs:", ok)
}
