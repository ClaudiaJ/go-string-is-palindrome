package StringIsPalindrome

import "testing"

func TestStringIsPalindrome(t *testing.T) {
	palindromes := []string{
		"mom",
		"dad",
		"\u0939\u093f\u0939",
		"हिह",
		"t",
		"😲😿🙀😿😲",
		"a man a plan aca nalp a nam a",
	}
	notPalindromes := []string{
		"papa",
		"The zesty maroon fox leaps far over the head of the excited doggo",
		"A man, a plan, a canal: Panama",
		"Alpaca in the cold moonlight",
		"請問你哪位？",
		"\u0939\u093f",
		"हि",
	}

	for _, v := range palindromes {
		if check := StringIsPalindrome(v); check != true {
			t.Errorf("String \"%s\" expected to be Palindrome, but returns %v", v, check)
		}
	}

	for _, v := range notPalindromes {
		if check := StringIsPalindrome(v); check == true {
			t.Errorf("String %s not expected to be Palindrome, but returns %v", v, check)
		}
	}
}

func TestReverseString(t *testing.T) {
	reversed := map[string]string{
		"Test":         "tseT",
		"\u0939\u093f": "\u093f\u0939",
		"請問你哪位？":       "？位哪你問請",
	}

	for x, v := range reversed {
		if check := ReverseString(x); check != v {
			t.Errorf("ReverseString failed, got %s, want %s", check, v)
		}
	}
}
