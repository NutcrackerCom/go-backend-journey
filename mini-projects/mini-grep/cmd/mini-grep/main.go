package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/mini-grep/internal/grep"
)

func main() {
	pattern := flag.String("pattern", "", "substring to search (required)")
	file := flag.String("file", "", "file path (required)")
	dir := flag.String("dir", "", "dir path (required)")
	flag.Parse()

	if *pattern == "" || (*file == "" && *dir == "") {
		flag.Usage()
		os.Exit(2)
	}
	if *file != "" && *dir != "" {
		flag.Usage()
		os.Exit(2)
	}

	if *file != "" {
		matches, err := grep.FindInFile(*file, *pattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		name := filepath.Base(*file)
		for _, m := range matches {
			fmt.Printf("%s:%d:%s\n", name, m.Line, m.Text)
		}
	}

	if *dir != "" {
		matches, err := grep.FindInDir(*dir, *pattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		for _, dirMatch := range matches {
			for _, m := range dirMatch.Matches {
				fmt.Printf("%s:%d:%s\n", dirMatch.Path, m.Line, m.Text)
			}
		}
	}
}
