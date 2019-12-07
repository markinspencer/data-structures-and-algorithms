// --- Directions
// Given a string, return a new string with the reverse
// order of characters
// --- Examples
// reverse("apple") == "leppa"
// reverse("hello") == "olleh"
// reverse("Greetings!") == "!sgniteerG"

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

func main() {
	fmt.Println(reverse("apple"))
	fmt.Println(reverse("hello"))
	fmt.Println(reverse("Greetings!"))
}
