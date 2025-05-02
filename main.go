package main

import (
	"fmt"
)

func main() {
	str := "Hi"
	b := []byte(str)

	fmt.Println(b)
	fmt.Println(string(b))
}
