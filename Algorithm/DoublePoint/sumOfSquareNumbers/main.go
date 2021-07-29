package main

import (
	"fmt"
	"math"
)

/*
Input: 5
Output: True
Explanation: 1 * 1 + 2 * 2 = 5

题目描述：判断一个非负整数是否为两个整数的平方和。
*/

func judgeSquareSum(c int) bool {
	if c < 0 { return false }
	i := 0; j := int(math.Sqrt(float64(c)))
	for i <= j {
		sum := i*i + j*j
		if sum == c {
			return true
		} else if sum > c {
			j -= 1
		} else {
			i += 1
		}
	}
	return false
}

func main() {
	//result := judgeSquareSum(2)
	//fmt.Println(result)

	s := 10
	fmt.Println(s)
	fmt.Println(int(math.Sqrt(float64(s))))
}

