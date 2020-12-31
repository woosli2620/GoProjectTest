package main

import (
    "fmt"
    "unsafe"
)


//演示golang 布尔类型的使用
func main() {
   var b = false

   //占用空间是1个字节，值只能是true或者false
   fmt.Println("b = ",b)
   fmt.Println("b占用的空间：",unsafe.Sizeof(b))
}
