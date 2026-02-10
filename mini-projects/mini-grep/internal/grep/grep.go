package grep

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

var (
	ErrEmptyPattern = errors.New("empty pattern")
)

type Match struct {
	Line int
	Text string
}

func FindAll(r io.Reader, pattern string) ([]Match, error) {
	if pattern == "" {
		return nil, ErrEmptyPattern
	}
	var match []Match
	scanner := bufio.NewScanner(r)
	var line int = 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), pattern) {
			match = append(match, Match{Line: line, Text: scanner.Text()})
		}
		line++
	}
	if err := scanner.Err(); err != nil {
		return match, err
	}
	return match, nil
}

func FindInFile(path string, pattern string) ([]Match, error) {
	if pattern == "" {
		return nil, ErrEmptyPattern
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	matches, err := FindAll(file, pattern)
	if err != nil {
		return nil, err
	}
	return matches, nil
}
