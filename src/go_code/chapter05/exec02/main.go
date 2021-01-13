package main

import (
	"fmt"
)

func main () {

	// var ch byte
	// fmt.Println("请输入一个字符（a,b,c,d,e）")
	// fmt.Scanf("%c",&ch)

	// switch ch {
	// case 'a':
	// 	fmt.Println("A")
	// case 'b':
	// 	fmt.Println("B")
	// case 'c':
	// 	fmt.Println("C")
	// case 'd':
	// 	fmt.Println("D")
	// default:
	// 	fmt.Println("other")
	// }

	var score float32
	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)

	switch int( score / 60 ) {
	case 1:
		fmt.Println("合格")
	case 0:
		fmt.Println("不合格")
	default:
		fmt.Println("输入有误")
	}
}