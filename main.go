package main

import (
	"fmt"
)

func main() {
	// #短变量声明；
	str := "Hi"
	b := []byte(str)

	fmt.Println(b)
	fmt.Println(string(b))
}
