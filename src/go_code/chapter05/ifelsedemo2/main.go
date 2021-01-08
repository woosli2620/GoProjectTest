package main
import (
	"fmt"
)

//流程控制
func main() {
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("大于18")
	} else {
		fmt.Println("小于18")
	}
}