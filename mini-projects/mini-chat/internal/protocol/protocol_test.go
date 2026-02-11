package protocol

import (
	"errors"
	"testing"
)

func TestNormalCmd(t *testing.T) {
	reference := Parsed{Kind: KindCommand, Cmd: "name", Arg: "Bob"}

	answer, err := ParseLine("/name Bob")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if answer != reference {
		t.Fatal("Got wrong Parsed", answer.Cmd, answer.Arg)
	}
}

func TestNormalText(t *testing.T) {
	line := "Even the smallest person can change the course of the future"
	reference := Parsed{Kind: KindMessage, Text: line}
	answer, err := ParseLine(line)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if answer != reference {
		t.Fatal("Got wrong Parsed", answer.Cmd, answer.Arg)
	}
}

func TestEmptyLine(t *testing.T) {
	_, err := ParseLine("")
	if !errors.Is(err, ErrEmptyLine) {
		t.Fatalf("expected error: %v, got %v", ErrEmptyLine, err)
	}
}

func TestSpaceLine(t *testing.T) {
	_, err := ParseLine("    ")
	if !errors.Is(err, ErrEmptyLine) {
		t.Fatalf("expected error: %v, got %v", ErrEmptyLine, err)
	}
}

func TestBadCommand(t *testing.T) {
	_, err := ParseLine("/")
	if !errors.Is(err, ErrBadCommand) {
		t.Fatalf("expected error: %v, got %v", ErrBadCommand, err)
	}
}
