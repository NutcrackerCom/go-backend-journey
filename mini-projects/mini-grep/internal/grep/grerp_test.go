package grep

import (
	"errors"
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

func TestEmptyPatern(t *testing.T) {
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

func TestNormalPatern(t *testing.T) {
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
