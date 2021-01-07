package main

import (
	"fmt"
)

//演示；赋值运算
func main() {
	a := 9
	b := 2

	//a b 交换
	var c int
	c = a
	a = b
	b = c
	fmt.Printf("changed a = %v,b = %v \n",a,b)

	a += 7
	fmt.Println("a = ",a)

	var d int 
	d = a
	d = 8 + 2 * 8

	fmt.Println(d)
}