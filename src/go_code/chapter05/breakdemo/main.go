package main

import (
	"fmt"
)

func main() {
	
	//100以内的数求和，求出第一次大于20的当前数

	for i := 1;i<=100;i++ {

		for j := 1;j<100;j++ {
			if i + j > 20 {
				fmt.Printf("第一次大于20的两个数之和位：%d + %d \n",i,j)
				break
			}
		}
	}

	//提示登录验证，有三次机会，用户名是：张无忌，密码888登录成功
	var name string
	var pwd string
	var i int = 1

	for ;i<=3;i++ {
		fmt.Println("请输入用户名：")
		fmt.Scanln(&name)
		fmt.Println("请输入密码：")
		fmt.Scanln(&pwd)

		if name == "张无忌" && pwd == "888" {
			fmt.Println("登录成功！")
			break
		} else {
			fmt.Printf("你还有%d此登录机会！",3 - i )
		}
	}

	if i > 3 {
		fmt.Println("登录失败！")
	}

}