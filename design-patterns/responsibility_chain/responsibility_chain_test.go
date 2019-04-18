package responsibility_chain

import "testing"

func TestNewConsoleLogger(t *testing.T) {
	loggerChain := new(LoggerChain)

	logger :=loggerChain.GetLogger()
	logger.logMessage(ERROR, "test ERROR")
	logger.logMessage(INFO, "Info test")


}