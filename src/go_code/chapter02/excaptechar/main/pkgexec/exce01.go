package main

import (
	"fmt"
	"runtime"
)

// 格式化命令：gofmt -w exce01.go
func main() {
	fmt.Println("姓名\t年龄\t籍贯\t住址\njoin\t12\t河北\t北京")

	var num = 2 + 4 * 5
	fmt.Println(num)

	fmt.Println("HelloWord!HelloWord!HelloWord!HelloWord!",
	"HelloWord!HelloWord!HelloWord!",
	"HelloWord!HelloWord!HelloWord!",
	"HelloWord!HelloWord!HelloWord!")

	fmt.Println(runtime.Version())
	fmt.Println(runtime.NumCPU())
}
