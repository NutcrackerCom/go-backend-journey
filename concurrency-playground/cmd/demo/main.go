package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(nums []int64, result chan int64, wg *sync.WaitGroup) {
	defer wg.Done()
	var res int64 = 0
	for _, elem := range nums {
		res += elem
	}
	result <- res
}

func main() {
	var numElem int64 = 100_000_000
	var numGoroutine int64 = 100

	ch := make(chan int64, numGoroutine+1)

	wg := &sync.WaitGroup{}

	var bigDataSlice []int64 = make([]int64, numElem+1)
	var i int64
	for i = 0; i <= numElem; i++ {
		bigDataSlice[i] = i
	}

	timeStart := time.Now()

	var resForce int64 = 0

	for _, elem := range bigDataSlice {
		resForce += elem
	}

	duration := time.Since(timeStart)
	fmt.Println("Brute Force calc", duration, resForce)

	timeStart = time.Now()
	step := numElem / numGoroutine
	for ind := range numGoroutine + 1 {
		wg.Add(1)
		start := ind * step
		stop := min(start+step, int64(len(bigDataSlice)))
		go worker(bigDataSlice[start:stop], ch, wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	var resGoroutine int64 = 0
	for val := range ch {
		resGoroutine += val
	}
	duration = time.Since(timeStart)
	fmt.Println("Goroutine calc", duration, resGoroutine)

}
