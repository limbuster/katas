package main

// LongestPalindrome will return the longest substring palindrome
func LongestPalindrome(s string) string {
	result := ""

	if len(s) <= 1 {
		return s
	}

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j <= len(s); j++ {
			if len(result) < len(s[i:j]) && IsPalindrome(s[i:j]) {
				result = s[i:j]
			}
		}
	}

	return result
}

// IsPalindrome takes a string and returns true if it's a palindrome
func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
