package logger_format

import (
    "fmt"
    "log"
    "path/filepath"
    "runtime"
    "strings"
)

func DebugF(format string, args ...interface{}) {
    logFormatter("DEBUG", format, args...)
}

func ErrorF(format string, args ...interface{}) {
    logFormatter("ERROR", format, args...)
}

func logFormatter(level string, formatting string, args ...interface{}) {
    fileName, line, funcName := "???", -1, "???"
    pc, fileName, line, ok := runtime.Caller(2)

    if ok {
        funcName = runtime.FuncForPC(pc).Name()
        fmt.Println(">>>1>>>", funcName)
        funcName = filepath.Ext(funcName)
        fmt.Println(">>>2>>>", funcName)
        funcName = strings.TrimPrefix(funcName, ".")
        fmt.Println(">>>3>>>", funcName)

        fileName = filepath.Base(fileName)
    }

    log.Printf("%s:%d:%s: %s: %s \n", fileName, line, funcName, level, fmt.Sprintf(formatting, args...))
}

