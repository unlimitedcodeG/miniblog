package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue)
	// ✅ 即使通道关闭，仍然可以继续接收已发送但尚未读取的数据。
	// ✅ 使用 range 遍历通道时，通道关闭后会自动停止迭代，不会导致错误或死锁。
	// ✅ 只有通道为空且已关闭时，range 才会结束遍历，否则会一直阻塞等待数据。
	for {
		elem, ok := <-queue
		if !ok {
			break
		}
		fmt.Println(elem)
	}
	// for elem:=range queue{
	// 	fmt.Println(elem)
	// }
}
