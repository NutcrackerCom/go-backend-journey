package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func Reduction(a, b int) (int, int) {
	gcd := gcd(a, b)
	return a / gcd, b / gcd
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(Reduction(a, b))
}
