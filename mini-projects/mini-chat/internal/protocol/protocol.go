package protocol

import (
	"errors"
	"strings"
)

var (
	ErrEmptyLine  = errors.New("empty line")
	ErrBadCommand = errors.New("bad command")
)

type Kind int

const (
	KindMessage Kind = iota
	KindCommand
)

type Parsed struct {
	Kind Kind
	Text string
	Cmd  string
	Arg  string
}

func ParseLine(line string) (Parsed, error) {
	str := strings.TrimSpace(line)
	if str == "" {
		return Parsed{}, ErrEmptyLine
	}

	if !strings.HasPrefix(str, "/") {
		return Parsed{Kind: KindMessage, Text: str}, nil
	}
	cmd := strings.TrimPrefix(str, "/")
	cmd = strings.TrimSpace(cmd)
	if cmd == "" {
		return Parsed{}, ErrBadCommand
	}
	cmd, arg, _ := strings.Cut(cmd, " ")

	arg = strings.TrimSpace(arg)
	return Parsed{Kind: KindCommand, Cmd: strings.ToLower(cmd), Arg: arg}, nil
}
