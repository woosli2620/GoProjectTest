package main

import (
	"fmt"
)

func main () {
	//统计3个班成绩情况，每个班有5名同学
	//求出各个班的平均分和所有班级的平均分【学生成绩从键盘输入】
	var classNum int = 3
	var sutNum int = 5
	var totalscore float32
	var pertotalscore float32
	var score float32
	var passCount int

	for i := 1;i<=classNum;i++ {

		pertotalscore = 0
		for j:= 1;j<=sutNum;j++ {
			fmt.Printf("请输入%d班的第%d个学生的成绩：",i,j)
			fmt.Scanln(&score)
			pertotalscore += score

			if score >= 60 {
				passCount++
			}
		}

		fmt.Printf("第 %d 个班的平均分是：%v\n",i,pertotalscore/float32(sutNum))
		totalscore += pertotalscore
	}

	fmt.Printf("所有班的平均分是：%v\n",totalscore/float32(classNum * sutNum))
	fmt.Printf("及格人数是：%v\n",passCount)
}