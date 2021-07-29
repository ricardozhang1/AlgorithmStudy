package main

import "fmt"

/*
Input: "abcba"
Output: True
Explanation: You could delete the character 'c'.
*/

func validPalindrome(s string) bool {
	if s == "" { return false }
	//sSlice := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		si, sj := s[i], s[j]
		if si != sj {
			return isPalindrome(s, i, j-1) || isPalindrome(s,i+1, j)
		}
		i += 1
		j -= 1
	}
	return true
}

func isPalindrome(s string, i, j int) bool {
	for i<j {
		if s[i] != s[j] {
			return false
		}
		i += 1
		j -= 1
	}
	return true
}

func main() {
	s := "abdfgba"
	result := validPalindrome(s)
	fmt.Println(result)
}
