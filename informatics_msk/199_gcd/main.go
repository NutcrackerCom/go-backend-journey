package main

import "fmt"

func Gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return Gcd(y, x%y)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(Gcd(a, b))
}
