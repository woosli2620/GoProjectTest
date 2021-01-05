package main

import (
    "fmt"
    "strconv"
)


//演示golang 基本数据类型转string
func main() {
    var num1 int = 99
    var num2 float64 = 23.456
    var b bool = true
    var myChar byte = 'h'

    var str string

    //第一种方式
    str = fmt.Sprintf("%d",num1)
    fmt.Printf("str type = %T num1 = %v \n",str,str)

    str = fmt.Sprintf("%f",num2)
    fmt.Printf("str type = %T num2 = %v \n",str,str)

    str = fmt.Sprintf("%t",b)
    fmt.Printf("str type = %T b = %q \n",str,str)

    str = fmt.Sprintf("%c",myChar)
    fmt.Printf("str type = %T b = %q \n",str,str)

    ///第二种方式
    var num3 int = 99
    var num4 float64 = 23.456
    var b2 bool = true

    str = strconv.FormatInt(int64(num3),10)
    fmt.Printf("str type = %T b = %q \n",str,str)

    str = strconv.FormatFloat(num4,'f',10,64)
    fmt.Printf("str type = %T b = %q \n",str,str)

    str = strconv.FormatBool(b2)
    fmt.Printf("str type = %T b = %q \n",str,str)


    var num5 int64 = 4567
    str = strconv.Itoa(int(num5))
    fmt.Printf("str type = %T b = %q \n",str,str)







}

