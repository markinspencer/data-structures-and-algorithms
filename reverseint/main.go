// --- Directions
// Given an integer, return an integer that is the reverse
// ordering of numbers
// --- Examples
// reverseInt(15) == 51
// reverseInt(981) == 189
// reverseInt(5) == 5
// reverseInt(-15) == -51
// reverseInt(90) == -9

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func reverseString(s string) string {
	rs := make([]string, 0)

	for _, ch := range strings.Split(s, "") {
		rs = append([]string{ch}, rs...)
	}

	return strings.Join(rs, "")
}

func reverseInt(i int) int {
	isNegative := i < 0

	abs := int(math.Abs(float64(i)))
	s := fmt.Sprintf("%d", abs)
	rs := reverseString(s)
	ri, err := strconv.ParseInt(rs, 10, 0)
	check(err)

	if isNegative {
		return int(ri) * -1
	}

	return int(ri)
}

func main() {
	fmt.Println(reverseInt(15))  // 51
	fmt.Println(reverseInt(981)) // 189
	fmt.Println(reverseInt(5))   // 5
	fmt.Println(reverseInt(-15)) // -51
	fmt.Println(reverseInt(-90)) // -9
}
