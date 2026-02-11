package main

import (
	"fmt"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/mini-grep/internal/grep"
)

func main() {
	//r := strings.NewReader("hello\ngo\nworld\ngo go\n")
	matches, err := grep.FindInDir("", "go", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", matches)
}
