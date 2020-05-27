package leetcode

import "unicode"

// https://leetcode-cn.com/problems/valid-palindrome/

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left <= right {
		if !isValid(s[left]) {
			left++
		} else if !isValid(s[right]) {
			right--
		} else if unicode.ToUpper(rune(s[left])) != unicode.ToUpper(rune(s[right])) {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}

func isValid(v uint8) bool {
	m := rune(v)
	return unicode.IsDigit(m) || unicode.IsLetter(m)
}
