package main

import (
    "fmt"
    "../model"
)



//演示golang 标识符使用
func main() {
   var num int = 10
   var Num int = 20

   //区分大小写
   fmt.Printf("num = %v,Num = %v\n",num,Num)

   //_是空标识符，用于占位
   //var _ int = 40
   //fmt.Println(_)

   //语法可以通过，不推荐使用
   var int int = 20
   fmt.Println("int = ",int)

   fmt.Println(model.HeroName)
}

