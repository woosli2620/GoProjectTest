package main

import (
    "fmt"
)


//演示golang小数类型的使用
func main() {
    var price float32 = 0.3333
    fmt.Println("price = ",price)

    var num1  float32 = -1.0000078
    var num2  float64 = -78973932794.09
    fmt.Println("num1 = ",num1,"num2 = ",num2)

    //用浮点数，精度有可能有损失
    var num3  float32 = -123.0000901
    var num4  float64 = -123.0000901
    num5 := -123.0000901
    fmt.Println("num3 = ",num3,"num4 = ",num4,"num5 = ",num5)

    //golang 的浮点类型默认是floag64类型
    var num6 = 1.1
    fmt.Printf("num6 的数据类型是 %T：",num6)

    //十进制形式：如5.12  .512（必须有小数点）
    num7 := 512
    num8 := .123
    fmt.Println("num7 = ",num7,"num8 = ",num8)


    //科学计数法形式
    num9 := 5.1234e2 //5.1234 * 10 的2次方
    num10 := 5.1234E2 //5.1234 * 10 的2次方
    num11 := 5.1234E-2 //5.1234 / 10 的2次方
    fmt.Println("num9 = ",num9,"num10 = ",num10,"num11 = ",num11)
}
