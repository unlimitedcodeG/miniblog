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
	fmt.Println("sent all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("recevied more jobs:", ok)
}
