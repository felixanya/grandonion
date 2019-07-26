package runtime_demo

import (
    "fmt"
    "log"
    "path/filepath"
    "runtime"
    "strings"
)

func logFormatting(level string, formatting string, args ...interface{}) {
    fileName, line, funcName := "???", -1, "???"
    pc, fileName, line, ok := runtime.Caller(2)

    if ok {
        funcName = runtime.FuncForPC(pc).Name()
        funcName = filepath.Ext(funcName)
        funcName = strings.TrimPrefix(funcName, ".")

        fileName = filepath.Base(fileName)
    }

    log.Printf("%s:%d:%s: %s: %s \n", fileName, line, funcName, level, fmt.Sprintf(formatting, args...))
}
