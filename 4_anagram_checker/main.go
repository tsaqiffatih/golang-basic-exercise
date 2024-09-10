package main

import "strings"

func isAnagram(string1, string2 string) bool {
	string1 = strings.ToLower(string1)
	string2 = strings.ToLower(string2)

	if len(string1) != len(string2) {
		return false
	}

	runesString1 := []rune(string1)
	runesString2 := []rune(string2)

	
}

func main() {

}