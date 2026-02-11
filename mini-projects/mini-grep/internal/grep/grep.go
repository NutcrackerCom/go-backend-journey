package grep

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrEmptyPattern = errors.New("empty pattern")
)

type FileMatch struct {
	Path    string
	Matches []Match
}

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

func FindInDir(root string, pattern string, ext string) ([]FileMatch, error) {
	if pattern == "" {
		return nil, ErrEmptyPattern
	}
	var fileMatch []FileMatch
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if ext != "" && filepath.Ext(path) != ext {
			return nil
		}

		if d.Type().IsRegular() {
			match, err := FindInFile(path, pattern)
			if err != nil {
				return err
			}
			if len(match) > 0 {
				fileMatch = append(fileMatch, FileMatch{Path: path, Matches: match})
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return fileMatch, nil
}
