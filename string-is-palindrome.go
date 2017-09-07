package StringIsPalindrome

// Test if string is equal to itself, if itself were reversed
func StringIsPalindrome(testStr string) bool {
	return testStr == ReverseString(testStr)
}

// Reverse string
func ReverseString(input string) (output string) {
	for _, b := range input {
		output = string(b) + output
	}
	return
}
