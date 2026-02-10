package grep

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func compare(answer []Match, ideal []Match) bool {
	if len(answer) != len(ideal) {
		return false
	}
	for i := range answer {
		if (answer[i].Line != ideal[i].Line) || (answer[i].Text != ideal[i].Text) {
			return false
		}
	}
	return true
}

func TestEmptyPattern(t *testing.T) {
	patern := ""
	r := strings.NewReader("hello\ngo\nworld\ngo go\n")
	matches, err := FindAll(r, patern)
	if !errors.Is(err, ErrEmptyPattern) {
		t.Fatalf("expected %v error, got %v", ErrEmptyPattern, err)
	}
	if matches != nil {
		t.Fatalf("expected nil matches")
	}
}

func TestEmptyPatternFile(t *testing.T) {
	patern := ""
	matches, err := FindInFile("", patern)
	if !errors.Is(err, ErrEmptyPattern) {
		t.Fatalf("expected %v error, got %v", ErrEmptyPattern, err)
	}
	if matches != nil {
		t.Fatalf("expected nil matches")
	}
}

func TestNormalPattern(t *testing.T) {
	patern := "go"
	r := strings.NewReader("hello\ngo\nworld\ngo go\n")
	answer := []Match{{2, "go"}, {4, "go go"}}
	matches, err := FindAll(r, patern)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !compare(matches, answer) {
		t.Fatal("got wrong Mathes")
	}
}

func TestNormalFile(t *testing.T) {
	patern := "go"
	file, err := os.CreateTemp("", "mini-grep-*.txt")
	if err != nil {
		t.Fatalf("got error in openning file %v", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()
	file.Write([]byte("hello\ngo\nworld\ngo go\n"))
	answer := []Match{{2, "go"}, {4, "go go"}}
	matches, err := FindInFile(file.Name(), patern)
	if err != nil {
		t.Fatalf("expected nil error? got %v", err)
	}
	if !compare(matches, answer) {
		t.Fatal("got wrong Mathes", matches)
	}
}

func TestEmptyAnswer(t *testing.T) {
	patern := "test"
	r := strings.NewReader("hello\ngo\nworld\ngo go\n")
	answer := []Match{}
	matches, err := FindAll(r, patern)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !compare(matches, answer) {
		t.Fatal("got wrong Mathes")
	}
}

func TestEmptyFile(t *testing.T) {
	patern := "go"
	file, err := os.CreateTemp("", "mini-grep-*.txt")
	if err != nil {
		t.Fatalf("got error in openning file %v", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()
	answer := []Match{}
	matches, err := FindInFile(file.Name(), patern)
	if err != nil {
		t.Fatalf("expected nil error got %v", err)
	}
	if !compare(matches, answer) {
		t.Fatal("got wrong Mathes", matches)
	}
}

//to do
/*
func TestEmptyDir(t *testing.T){
	patern := "go"
	dir, err := os.MkdirTemp("", "mini-grep-*")
	if err != nil {
        t.Fatalf("got error in creating dir %v", err)
    }
    defer os.RemoveAll(dir)

	file, err := os.CreateTemp(dir, "mini-grep-*.txt")
	if err != nil {
		t.Fatalf("got error in openning file %v", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()
	answer := []Match{}
}
*/
