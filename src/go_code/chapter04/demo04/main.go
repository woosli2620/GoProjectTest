package main

import (
	"fmt"
)

//声明测试函数
func test() bool {
	fmt.Println("test...")
	return true
}

//演示；逻辑运算
func main() {
	var i int = 10

	// if i > 9 && test() {
	// 	fmt.Println("ok...")
	// }

	if i > 9 || test() {
		fmt.Println("ok...")
	}

	/*
	var age int = 40

	//与的使用
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}

	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}

	//或的使用
	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}

	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}

	//非的使用

	if age > 30 {
		fmt.Println("ok5")
	}

	if !(age > 30)  {
		fmt.Println("ok6")
	}
	*/
}