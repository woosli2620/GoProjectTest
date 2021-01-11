package main
import (
	"fmt"
)


//流程控制
func main() {
	var n1 int32

	fmt.Println("请输入一个整数（10,20）：")
	fmt.Scanln(&n1)

	switch n1 {
	case 10:
		fmt.Println("输出~~~10")
		//fallthrough //穿透
	case 20:
		fmt.Println("输出~~~20")
		//fallthrough //穿透
	default:
		fmt.Println("其他。。。")
	}
}