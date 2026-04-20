package main

import "fmt"

/*
1	3	6	10	15
7	16	27	40	55
18	39	63	90	120
34	72	114	160	210
55	115	180	250	325

*/

func Amount(rec [][]int, x1, x2, y1, y2 int) int {
	return 0
}

func main() {
	var N, M int
	var K int
	var x1, y1 int
	var x2, y2 int
	var rec = make([][]int, N)
	fmt.Scan(&N, &M, &K)

	for i := range N {
		rec[i] = make([]int, M)
		for j := range M {
			fmt.Scan(&rec[i][j])
		}
	}
	fmt.Println(Amount(rec, x1, x2, y1, y2))
}
