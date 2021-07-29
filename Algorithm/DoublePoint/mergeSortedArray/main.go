package main

import "fmt"

/*
Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
*/

func merge(nums1 []int, m int, nums2 []int, n int)  {
	index1, index2, indexMerge := m-1, n-1, m+n-1
	for index1 >= 0 || index2 >= 0 {
		if index1 < 0 {
			nums1[indexMerge] = nums2[index2]
			index2--
		} else if index2 < 0 {
			nums1[indexMerge] = nums1[index1]
			index1--
		} else if nums1[index1] > nums2[index2] {
			nums1[indexMerge] = nums1[index1]
			index1--
		} else {
			nums1[indexMerge] = nums2[index2]
			index2--
		}
		indexMerge--
	}
}

func main() {
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

