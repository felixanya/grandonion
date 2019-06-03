package responsibility_chain

import (
	"errors"
	"fmt"
)

const (
	INFO = iota
	DEBUG
	WARN
	ERROR
	FATAL
)

type ILogger interface {
	setNextLogger(ILogger) error
	logMessage(int, string)
	write(string)
}

// console logger
type ConsoleLogger struct {
	Level 	int
	Next 	ILogger
}

func NewConsoleLogger(level int) ILogger {
	return &ConsoleLogger{
		Level: level,
	}
}

func (cl *ConsoleLogger) setNextLogger(logger ILogger) error {
	if logger == nil {
		return errors.New("nil logger error")
	}
	cl.Next = logger
	return nil
}

func (cl *ConsoleLogger) write(msg string) {
	fmt.Println("Standard Console Logger::", msg)
}

func (cl *ConsoleLogger) logMessage(level int, msg string) {
	if cl.Level <= level {
		cl.write(msg)
	} else {
		cl.Next.logMessage(level, msg)
	}
}

// Error Logger
type ErrorLogger struct {
	Level 	int
	Next 	ILogger
}

func NewErrorLogger(level int) ILogger {
	return &ErrorLogger{
		Level: level,
	}
}

func (cl *ErrorLogger) setNextLogger(logger ILogger) error {
	if logger == nil {
		return errors.New("nil logger error")
	}
	cl.Next = logger
	return nil
}

func (cl *ErrorLogger) write(msg string) {
	fmt.Println("Standard Error Logger::", msg)
}

func (cl *ErrorLogger) logMessage(level int, msg string) {
	if cl.Level <= level {
		cl.write(msg)
	} else {
		cl.Next.logMessage(level, msg)
	}
}

// File Logger
type FileLogger struct {
	Level 	int
	Next 	ILogger
}

func NewFileLogger(level int) ILogger {
	return &FileLogger{
		Level: level,
	}
}

func (cl *FileLogger) setNextLogger(logger ILogger) error {
	if logger == nil {
		return errors.New("nil logger error")
	}
	cl.Next = logger
	return nil
}

func (cl *FileLogger) write(msg string) {
	fmt.Println("Standard File Logger::", msg)
}

func (cl *FileLogger) logMessage(level int, msg string) {
	if cl.Level <= level {
		cl.write(msg)
	} else {
		cl.Next.logMessage(level, msg)
	}
}

// get logger chain
type LoggerChain struct {
}

func (lc *LoggerChain) GetLogger() ILogger {
	errorLogger := NewErrorLogger(ERROR)
	fileLogger := NewFileLogger(DEBUG)
	consoleLogger := NewConsoleLogger(INFO)

	errorLogger.setNextLogger(fileLogger)
	fileLogger.setNextLogger(consoleLogger)

	return errorLogger
}