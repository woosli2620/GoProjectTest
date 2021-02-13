package main

import (
	"fmt"
)

func test(n1 int){
	n1 = n1 + 1
	fmt.Println("n1 = ",n1)
}
func test2(n1 int,n2 int) int {
	sum := n1 + n2
	return sum
}

func test3(n1 int)(int,int,int) {
	return n1+1,n1+2,n1+3
}

func main() {
	var n1 int = 0
	test(n1)
	fmt.Println("n1 = ",n1)

	sum := test2(1,2)
	fmt.Println("test2 sum = ",sum)

	a1,a2,a3 := test3(20)
	fmt.Println("test3 = ",a1,a2,a3)

	_,_,a4 := test3(30)
	fmt.Println("test3 = ",a4)
}