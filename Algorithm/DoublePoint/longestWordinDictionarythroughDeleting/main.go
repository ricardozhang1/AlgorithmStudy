package main

import (
	"fmt"
	"sort"
)

/*
Input:
s = "abpcplea", d = ["ale","apple","monkey","plea"]

Output:
"apple"
*/

func findLongestWord(s string, dictionary []string) string {
	var longestWord string
	sort.Strings(dictionary)
	for _, target := range dictionary {
		l1 := len(longestWord)
		l2 := len(target)
		if l1 > l2 || (l1==l2 && longestWord != target) {
			continue
		}
		if isSubstr(s, target) {
			longestWord = target
		}
	}
	return longestWord
}

func isSubstr(s ,target string) bool {
	i, j := 0, 0
	for i<len(s) && j<len(target) {
		if s[i] == target[j] {
			j++
		}
		i++
	}
	return j == len(target)
}

func main() {
	s := "abpcplea"
	d := []string{"a","b","c"}
	result := findLongestWord(s, d)
	fmt.Println(result)
}
