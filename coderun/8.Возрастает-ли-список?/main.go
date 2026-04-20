package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	rase := 0
  num := 0
  fmt.Scan(&num)
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	line, _ := reader.ReadString('\n')
	for _, r := range line {
    if r != ' ' {
      i, _ := strconv.Atoi(string(r))
      if num < i {
        rase++
      }
    }

	}
  if 
}
