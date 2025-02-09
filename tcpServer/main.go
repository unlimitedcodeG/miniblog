package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// 在本机 8080 端口监听 TCP 连接
	listener, err := net.Listen("tcp", ":8088")
	if err != nil {
		fmt.Println("监听错误:", err)
		return
	}
	defer listener.Close()
	fmt.Println("服务端正在 8088 端口监听...")

	// 循环等待客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("连接接收错误:", err)
			continue
		}
		// 每个连接交给一个 goroutine 处理
		go handleConnection(conn)
	}
}

// 处理每个客户端连接的函数
func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		// 以换行符为结束符读取客户端发送的数据
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取错误:", err)
			break
		}
		fmt.Printf("收到消息: %s", message)

		// 回显数据给客户端
		_, err = conn.Write([]byte("回显: " + message))
		if err != nil {
			fmt.Println("发送错误:", err)
			break
		}
	}
}
