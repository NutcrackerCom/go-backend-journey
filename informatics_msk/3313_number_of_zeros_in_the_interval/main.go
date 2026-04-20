package main

import "fmt"

func NumberOfZeros(n []int, x, y int) int {
	if x < 1 || y > len(n) || y < 1 || x > len(n) {
		return 0
	}

	var prefSum = make([]int, len(n))
	sumZeros := 0
	for ind, elem := range n {
		if elem == 0 {
			sumZeros++
		}
		prefSum[ind] = sumZeros
	}
	if x == 1 {
		return prefSum[y-1]
	}
	return prefSum[y-1] - prefSum[x-2]
}

func main() {
	var x, y int
	var numQ, lenN int

	fmt.Scan(&lenN)

	var n = make([]int, lenN)
	for i := range lenN {
		fmt.Scan(&n[i])
	}
	fmt.Scan(&numQ)

	for range numQ {
		fmt.Scan(&x, &y)
		fmt.Println(NumberOfZeros(n, x, y))
	}
}
