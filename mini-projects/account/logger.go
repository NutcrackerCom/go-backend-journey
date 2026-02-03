package account

import (
	"fmt"
	"os"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
}

type Writer struct{}

func (w Writer) Info(str string) {
	fmt.Fprintf(os.Stdout, "Msg: %s", str)
}

func (w Writer) Error(err string) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
}

type Collector struct {
	info   []string
	errors []string
}

func (c *Collector) Info(str string) {
	c.info = append(c.info, str)
}

func (c *Collector) Error(str string) {
	c.errors = append(c.errors, str)
}

type NopLogger struct{}

func (NopLogger) Info(string)  {}
func (NopLogger) Error(string) {}

type StdoutWriter struct{}
