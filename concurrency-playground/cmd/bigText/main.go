package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

func readText(addr string) []byte {
	data, err := os.ReadFile(addr)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	return data
}

func worker(chars []rune, result chan map[rune]int, wg *sync.WaitGroup) {
	defer wg.Done()
	var res map[rune]int = make(map[rune]int)
	for _, elem := range chars {
		res[elem]++
	}
	result <- res
}

func workerWithMtx(chars []rune, m map[rune]int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for _, elem := range chars {
		mu.Lock()
		m[elem]++
		mu.Unlock()
	}
}

func calcRes(m map[rune]int) int {
	res := 0
	for _, val := range m {
		res += val
	}
	return res
}

func main() {

	dataByte := readText("/home/chernyshev/local/go/go-backend-journey/concurrency-playground/cmd/bigText/Text.txt")
	dataRune := bytes.Runes(dataByte)
	var brutForse map[rune]int = make(map[rune]int)

	startTime := time.Now()
	for _, b := range dataRune {
		brutForse[b]++
	}
	duration := time.Since(startTime)
	fmt.Println("Brute Force duration", duration)
	fmt.Println("Brute Force calc", "Uniq chars:", len(brutForse), "Number chars:", calcRes(brutForse))

	var numElem int = len(dataRune)
	var numGoroutine int = 10

	ch := make(chan map[rune]int, numGoroutine)

	wg := &sync.WaitGroup{}
	step := numElem/numGoroutine + 1

	startTime = time.Now()
	for ind := range numGoroutine {
		wg.Add(1)
		start := ind * step
		stop := min(start+step, len(dataRune))
		go worker(dataRune[start:stop], ch, wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var resGoroutine map[rune]int = make(map[rune]int)
	for val := range ch {
		for key, val := range val {
			resGoroutine[key] += val
		}
	}
	duration = time.Since(startTime)
	fmt.Println("Goroutine duration", duration)
	fmt.Println("Goroutine calc:", "Uniq chars:", len(resGoroutine), "Number chars:", calcRes(resGoroutine))

	wg = &sync.WaitGroup{}
	mu := &sync.Mutex{}

	var resMutex map[rune]int = make(map[rune]int)

	step = numElem/numGoroutine + 1

	startTime = time.Now()
	for ind := range numGoroutine {
		wg.Add(1)
		start := ind * step
		stop := min(start+step, len(dataRune))
		go workerWithMtx(dataRune[start:stop], resMutex, wg, mu)
	}

	wg.Wait()
	duration = time.Since(startTime)
	fmt.Println("Goroutine with mtx duration", duration)
	fmt.Println("Goroutine with mtx calc:", "Uniq chars:", len(resMutex), "Number chars:", calcRes(resMutex))
}
