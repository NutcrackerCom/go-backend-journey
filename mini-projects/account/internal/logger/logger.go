package logger

import (
	"fmt"
	"os"
)

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

type ConsoleLogger struct{}

func (w ConsoleLogger) Info(str string, args ...any) {
	fmt.Fprintf(os.Stdout, "INFO: "+str+"\n", args...)
}

func (w ConsoleLogger) Error(err string, args ...any) {
	fmt.Fprintf(os.Stderr, "Error: "+err+"\n", args...)
}

type BufferLogger struct {
	Infos  []string
	Errors []string
}

func (c *BufferLogger) Info(str string, args ...any) {
	c.Infos = append(c.Infos, fmt.Sprintf(str, args...))
}

func (c *BufferLogger) Error(str string, args ...any) {
	c.Errors = append(c.Errors, fmt.Sprintf(str, args...))
}

type NopLogger struct{}

func (NopLogger) Info(_ string, args ...any)  {}
func (NopLogger) Error(_ string, args ...any) {}
