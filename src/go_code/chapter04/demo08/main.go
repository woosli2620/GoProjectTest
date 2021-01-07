package main

import (
	"fmt"
)

//位运算测试
func main() {
	fmt.Println(2&3)
	fmt.Println(2|3)
	fmt.Println(2^3)
	fmt.Println(-2^2)

	fmt.Println(1 >> 2)
	fmt.Println(1 << 2)
}