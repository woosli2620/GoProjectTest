package main

import (
	"fmt"
)

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool 

	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水：")
	fmt.Scanln(&sal)
	fmt.Println("是否通过考试:")
	fmt.Scanln(&isPass)
	fmt.Printf("名字是：%v\n年龄是：%v\n薪水是：%v\n是否通过考试：%v\n",name,age,sal,isPass)

	fmt.Println("**************************************************")

	//格式化输入
	fmt.Println("请输入你的姓名、年龄、薪水，是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t",&name,&age,&sal,&isPass)
	fmt.Printf("名字是：%v\n年龄是：%v\n薪水是：%v\n是否通过考试：%v\n",name,age,sal,isPass)
}