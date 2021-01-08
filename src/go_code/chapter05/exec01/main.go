package main

import (
	"fmt"
)

func main () {

	var n1 int32 = 10
	var n2 int32 = 60

	if n1 + n2 > 50 {
		fmt.Println("hello world")
	} else {
		fmt.Println("小于50")
	}

	var n3 float64 = 11.0
	var n4 float64 = 17.0
	if n3 > 10.0 && n4 < 20.0 {
		fmt.Println("两数之和是：",n3 + n4)
	}

	var n5 int32 = 10
	var n6 int32 = 5
	if (n6 + n5 ) % 3 == 0 && (n6 + n5) % 5 == 0 {
		fmt.Println("能被3和5整除！")
	}

	var year int32
	fmt.Println("请输入年：")
	//fmt.Scanf("%d",&year);
	fmt.Scanln(&year)

	if (year % 4 == 0 && year % 100 != 0 ) ||  year % 400 == 0 {
		fmt.Printf("%d是闰年！",year)
	} else {
		fmt.Printf("%d 不是闰年！",year)
	}
}