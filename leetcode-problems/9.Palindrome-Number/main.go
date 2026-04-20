package main

import (
	"fmt"
)

/*
Given an integer x, return true if x is a Palindrome, and false otherwise.
*/

func isPalindrome(x int) bool {
	buf := x
	if x < 0 {
		return false
	}
	y := 0
	for buf > 0 {
		y *= 10
		y += buf % 10
		buf /= 10
	}
	return y == x
}

func main() {
	fmt.Println(isPalindrome(121))
}
