/*
69. Sqrt(x)
Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.
You must not use any built-in exponent function or operator.
For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
*/

package main

import (
	"fmt"
)

func mySqrt(x int) int {
	l := 0
	r := x + 1
	for r-l > 1 {
		m := l + (r-l)/2
		if m*m > x {
			r = m
		} else {
			l = m
		}
	}
	return l
}

func main() {
	x := 8
	fmt.Println(mySqrt(x))

}

