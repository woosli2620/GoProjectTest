package main

import (
	"fmt"
)

func main() {

	// var i int = 8
	// var a int

	// //a = i++ 错误，++ -- 智能单独使用
	// //a = i--
	// fmt.Println("i = ",i);
	// fmt.Println("a = ",a);


	var i int = 1
	i++
	//++i  错误
	fmt.Println("i = ",i)

	var w int = 97
	fmt.Printf("剩余 %d 天，还有 %d 个星期和 %d 天放假\n",w,w/7,w%7)


	var huashi float32 = 134.2
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("%v对应摄氏温度 = %v\n",huashi,sheshi)

}