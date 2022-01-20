package palindrome

import "unicode"

// IsPalindromeはsが前からでも後ろからでも同じように読めるかどうかを報告します。
// 大文字小文字の違いは無視され、非文字も無視されます。
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
