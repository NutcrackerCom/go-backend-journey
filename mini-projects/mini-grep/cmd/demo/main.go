package main

import (
	"fmt"
	"strings"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/mini-grep/internal/grep"
)

func main() {
	r := strings.NewReader("hello\ngo\nworld\ngo go\n")
	matches, err := grep.FindAll(r, "gдщ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(matches)
}
