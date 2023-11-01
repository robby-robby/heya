package internal

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	ERROR
)

type Logger struct {
	Level    *LogLevel
	StdError io.Writer
	StdOut   io.Writer
}

type Lgg interface {
	Info(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})

	Infof(format string, a ...interface{})
	Debugf(format string, a ...interface{})
	Errorf(format string, a ...interface{})
	Panicf(format string, a ...interface{})
}

func (l *Logger) logOutput(level, prefix string, args []interface{}) {
	if *l.Level <= FromString(level) {
		str := prefix + ": " + stringify(args...) + "\n"
		l.StdOut.Write([]byte(str))
	}
}

func (l *Logger) Info(args ...interface{}) {
	l.logOutput("INFO", "INFO", args)
}

func (l *Logger) Debug(args ...interface{}) {
	l.logOutput("DEBUG", "DEBUG", args)
}

func (l *Logger) Error(args ...interface{}) {
	l.logOutput("ERROR", "ERROR", args)
}

func stringify(args ...interface{}) string {

	var builder strings.Builder
	for _, arg := range args {
		switch v := arg.(type) {
		case fmt.Stringer:
			builder.WriteString(v.String())
			builder.WriteRune(' ')
		case string:
			builder.WriteString(v)
			builder.WriteRune(' ')
		default:
			builder.WriteString(fmt.Sprintf("%v", v))
			builder.WriteRune(' ')
		}
	}

	return strings.TrimSuffix(builder.String(), " ")
}
func (l *Logger) Panic(args ...interface{}) {
	panic(stringify(args...))
}

func (l *Logger) Infof(format string, a ...interface{}) {
	l.Info(fmt.Sprintf(format, a...))
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	l.Debug(fmt.Sprintf(format, a...))
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	l.Error(fmt.Sprintf(format, a...))
}

func (l *Logger) Panicf(format string, a ...interface{}) {
	l.Panic(fmt.Sprintf(format, a...))
}

func (ll *LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "ERROR"}[*ll]
}

// Takes either DEBUG, INFO, or ERROR
func FromString(s string) LogLevel {
	switch s {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "ERROR":
		return ERROR
	default:
		return ERROR
	}
}

func NewLogger(level LogLevel, stdOut, stdError io.Writer) *Logger {
	if stdError == nil {
		stdError = os.Stderr
	}
	if stdOut == nil {
		stdOut = os.Stdout
	}
	return &Logger{
		Level:    &level,
		StdError: stdError,
		StdOut:   stdOut,
	}
}
