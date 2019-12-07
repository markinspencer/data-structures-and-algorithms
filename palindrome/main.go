// --- Directions
// Given a string, return ture if the string is a palindrome
// or false if it is not . Palinfromes are strings that
// form the same word if it is reversed. *Do* include spaces
// and punctuation in determining if the string is a palindrome
// --- Examples
// palindrome("abba") == true
// palindrome("abcdefg") == false

package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	rs := make([]string, 0)

	for _, ch := range strings.Split(s, "") {
		rs = append([]string{ch}, rs...)
	}

	return strings.Join(rs, "")
}

func palindrome(s string) bool {
	return s == reverse(s)
}

func main() {
	fmt.Println(palindrome("abba"))
	fmt.Println(palindrome("abcdefg"))
}
