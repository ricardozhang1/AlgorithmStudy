package main

import "fmt"

/*
Input: numbers={2, 7, 11, 15}, target=9
Output: index1=1, index2=2

题目描述：在有序数组中找出两个数，使它们的和为 target。
*/



func toSum(numbers []int, target int) []int {
	if len(numbers) == 0 { return nil }
	i := 0; j := len(numbers)-1
	for i<j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{numbers[i], numbers[j]}
		} else if sum > target {
			j -= 1
		} else {
			i += 1
		}
	}
	return nil
}

func main() {
	var numbers = []int{2, 7, 11, 15}
	var target = 18
	result := toSum(numbers, target)
	fmt.Println(result)
}

