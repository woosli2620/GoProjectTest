package main

import (
    "fmt"
)


//演示golang 字符串类型的使用
func main() {
    var address string = "北京长城 110 hello world!"
    address = "aaa"
    fmt.Println("address = ",address)

    //字符串一旦赋值不能修改
    var str = "hello"
    //str[0] = 'a' //不能修改str的内容
    fmt.Println("str = ",str)

    //反引号以源代码效果输出
    str2 := "abc\ndef"
    fmt.Println("str2 = ",str2)

    str3 :=  `数字水印的主要特征有：
    ①透明性。水印与原始数据紧密结合并隐藏其中，水印的存在不能破坏原数据的欣赏价值和使用价值。
    ②鲁棒性。在经过有损压缩、录制、打印、扫描、旋转、平移等常规处理操作后仍能检测到水印的能力。`

    fmt.Println("str3 = ",str3)

    //字符串拼接方式
    var str4 = "hello" + "world"
    str4 += "haha"
    fmt.Println("str4 = ",str4)

    var  str5 = "hello" + "haha" + 
    "hello" + "haha" + "hello" + 
    "haha" + "hello" + "haha" + 
    "hello" + "haha"
    fmt.Println("str5 = ",str5)


    var a int  //0
    var b float32//0
    var c float64//0
    var isMarried bool//false
    var name string//""

    fmt.Printf("\na = %d,b=%v,c=%v,isMarried=%v,name=%v",a,b,c,isMarried,name)

}
