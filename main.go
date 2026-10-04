package main

import (
	"fmt"
)

func facto(n int) int {
	var sum int = 1

	if n <= 0 {
		fmt.Println("erro!")
		return 0
	}

	for i := 1; i <= n; i++ {
		sum *= i
	}
	return sum
}

func average(n, cout int) float64 {
	end := float64(n) / float64(cout)
	return end
}

func main() {
	fmt.Println("第一个任务：hello!,rocky qqID 3533079387")

	const pi = 3.1415926
	r := 5.0

	fmt.Println("第二个任务 area =", pi*r*r)

	sum := 0
	for i := 1; i < 1001; i++ {
		sum += i
	}
	fmt.Println("第三个任务 正整数1到1000的sum =", sum)

	fmt.Println("第四个任务，请输入正整数n")
	n := 0
	_, err := fmt.Scanln(&n)
	if err != nil {
		{
			fmt.Println("输入错误！情输入正整数！")
		}
	}
	fmt.Println("n的阶乘为", facto(n))

	fmt.Println("第五个任务")
	var i int = -1
	var cout int = 0
	var sum_2 int = 0
	for ; i != 0; cout++ {

		fmt.Println("请输入一个整数（输入以0结束）")

		_, err := fmt.Scanln(&i)
		if err != nil {
			fmt.Println("读取失败！录入终止！")
			break
		}
		sum_2 += i
	}

	end := average(sum_2, cout)
	if end >= 60 {
		fmt.Println("平均成绩为", end, "合格")
	} else {
		fmt.Println("平均成绩为", end, "不合格")
	}
	if end == 20071114 {
		fmt.Println("Happy Birthday! Mr Bear")
	}
	return
}
