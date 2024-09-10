package main

import (
	"fmt"
	"unicode"
)

// IsPalindrome memeriksa apakah string adalah palindrome
func IsPalindrome(s string) bool {

	var cleaned []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, unicode.ToLower(r))
		}
	}

	left, right := 0, len(cleaned)-1
	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	// Contoh penggunaan
	fmt.Println(IsPalindrome("racecar"))                                               // Output: true
	fmt.Println(IsPalindrome("rararararararararararararararararararararararaRararar")) // Output: true
	fmt.Println(IsPalindrome("Able was I ere I saw Elba"))                             // Output: true
	fmt.Println(IsPalindrome("A man a plan a canal Panama"))                           // Output: true
	fmt.Println(IsPalindrome("KaBar"))                                                 // Output: false
	fmt.Println(IsPalindrome("hello"))                                                 // Output: false
	fmt.Println(IsPalindrome("madam"))                                                 // Output: true
	fmt.Println(IsPalindrome("a"))                                                     // Output: true
	fmt.Println(IsPalindrome("ab"))                                                    // Output: false
}
