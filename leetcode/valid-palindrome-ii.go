package leetcode

import "unicode"

// https://leetcode-cn.com/problems/valid-palindrome-ii/
var errCount = 1

func validPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left <= right {
		if !isValid(s[left]) {
			left++
		} else if !isValid(s[right]) {
			right--
		} else if unicode.ToUpper(rune(s[left])) != unicode.ToUpper(rune(s[right])) {
			if errCount > 0 {
				errCount--
				return validPalindrome(s[left:right]) || validPalindrome(s[left+1:right+1])
			}
			return false
		} else {
			left++
			right--
		}
	}
	return true
}
