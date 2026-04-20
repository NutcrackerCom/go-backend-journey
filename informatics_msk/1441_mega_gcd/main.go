package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func MegaGcd(n []int) int {
	if len(n) < 1 {
		return 0
	}
	gcdM := n[0]

	for _, elem := range n {
		gcdM = gcd(elem, gcdM)
	}
	return gcdM
}

func main() {
	var length int
	fmt.Scan(&length)
	n := make([]int, length)
	for i := range length {
		fmt.Scan(&n[i])
	}
	fmt.Println(MegaGcd(n))
}
