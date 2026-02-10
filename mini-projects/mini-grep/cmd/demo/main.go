package main

import (
	"fmt"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/mini-grep/internal/grep"
)

func main() {
	//r := strings.NewReader("hello\ngo\nworld\ngo go\n")
	matches, err := grep.FindInFile("/home/chernyshev/sd.txt", "go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(matches)
}
