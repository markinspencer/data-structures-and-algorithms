// --- Directions
// Given a string, reutn the character that is most
// commonly used in the string
// --- Examples
// maxChar("abccccccd") == "c"
// maxChar("apple 1231111") == "1"
package main

import (
	"fmt"
	"strings"
)

func maxChar(s string) string {
	chars := make(map[string]int)

	for _, c := range strings.Split(s, "") {
		if _, ok := chars[c]; ok {
			chars[c]++
		} else {
			chars[c] = 1
		}
	}

	var maxChar string
	max := 0

	for k, v := range chars {
		if v > max {
			maxChar, max = k, v
		}
	}

	return string(maxChar)
}

func main() {
	fmt.Println(maxChar("abccccccd"))     // "c"
	fmt.Println(maxChar("apple 1231111")) // "1"
}
