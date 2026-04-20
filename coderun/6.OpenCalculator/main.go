package main

import "fmt"

func main() {
	var x, y, z, num int

	fmt.Scan(&x, &y, &z, &num)

	var m map[int]struct{} = make(map[int]struct{})

	for num > 0 {
		n := num % 10
		if n != x && n != y && n != z {
			m[n] = struct{}{}
		}
		num /= 10
	}
	fmt.Println(len(m))
}
