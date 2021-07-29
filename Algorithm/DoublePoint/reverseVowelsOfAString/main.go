package main

import "fmt"

/*
Given s = "leetcode", return "leotcede".
*/

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func reverseVowels(s string) string {
	if s == "" { return "" }
	i, j := 0, len(s)-1
	newString := []byte(s)
	for i < j {
		si, sj := s[i], s[j]
		if isVowel(si) && isVowel(sj) {
			newString[i], newString[j] = sj, si
			i += 1
			j -= 1
		} else if isVowel(si) && !isVowel(sj) {
			j -= 1
		} else if !isVowel(si) && isVowel(sj) {
			i += 1
		} else {
			i += 1
			j -= 1
		}
	}
	return string(newString)
}

func main() {
	result := reverseVowels("leetcode")
	fmt.Println(result)
}
