package main

import "fmt"

//变量声明举例
func main() {

	//第一种
	var i int
	fmt.Println("i=",i)

	//第二种:类型推导
	var num = 10.11
	fmt.Println("num=",num)

	//第三种：省略var
	num2 := "Hello World~!";
	fmt.Println("num2=",num2)



}