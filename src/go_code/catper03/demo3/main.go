package main

import "fmt"

//全局变量
var gIn = 100
var name = "aaa"

//改写成一次声明
var (
	n30 = 300
	n40 = 900
	n50 = 1000
)

//一次性声明多个变量
func main() {

	//1
	var n1,n2,n3 int
	fmt.Println(n1,n2,n3)

	//2
	var n11,name,n13 = 100,"tom",888
	fmt.Println(n11,name,n13)

	//3
	n21,name2,n23 := 200,"tom",999
	fmt.Println(n21,name2,n23)

	fmt.Println("gIn = ",gIn)
	fmt.Println("n30,n40,n50 = ",n30,n40,n50)
}