// --- Directions
// Write a program that console logs the numbers
// from 1 to n, but for multiples of three print
// "fizz" instead of the number and for the multiples
// of five print "buzz". For numbers which are multiples
// of both three and five print "fizzbuzz"
// --- Example
// fizzbuzz(5)
// 1
// 2
// fizz
// 4
// buzz

package main

import "fmt"

func fizzBuzz(r int) {
	for i := 1; i <= r; i++ {
		switch {
		case i%(3*5) == 0:
			fmt.Println("fizzbuzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}

func main() {
	fizzBuzz(15)
}
