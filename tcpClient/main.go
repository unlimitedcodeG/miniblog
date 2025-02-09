package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 连接到本机 8080 端口的 TCP 服务端
	conn, err := net.Dial("tcp", "localhost:8088")
	if err != nil {
		fmt.Println("连接错误:", err)
		return
	}
	defer conn.Close()
	fmt.Println("已连接到服务端，输入文本并回车发送...")

	// 创建标准输入读取器
	stdinReader := bufio.NewReader(os.Stdin)
	for {
		// 从标准输入读取数据
		fmt.Print("请输入: ")
		input, err := stdinReader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入错误:", err)
			continue
		}

		// 发送数据到服务端
		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Println("发送数据错误:", err)
			break
		}

		// 从服务端读取回显数据
		connReader := bufio.NewReader(conn)
		reply, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Println("读取服务端回复错误:", err)
			break
		}
		fmt.Printf("服务端回复: %s", reply)
	}
}
