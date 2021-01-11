package main

import (
	"fmt"
)

func main () {

	var ch byte
	fmt.Println("请输入一个字符（a,b,c,d,e）")
	fmt.Scanf("%c",&ch)

	switch ch {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	default:
		fmt.Println("other")
	}

}