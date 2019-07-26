package logger_format

import "testing"

func TestDebugF(t *testing.T) {
    DebugF("this is a test of %s", "debug logger")
}
