package crud_module

import (
	"fmt"
	"log"
	"os"
)

const (
	ERROR LoggLevel = iota
	DEBUG
	DEVELOP
)
const DEFAULT_LOG_ADDRESS = "crud_module.log"

var LoggerAddress string = DEFAULT_LOG_ADDRESS
var LogLevelSetup LoggLevel = ERROR

func SetLogMode(mode LoggLevel) {
	LogLevelSetup = mode
}

func SetLogFile(fileURL string) {
	LoggerAddress = fileURL
}

func LogVersion() {
	fmt.Println("@mecuate/crud_module.package.version: " + Version)
}

func Log(message string) {
	conformedMessage := message
	switch LogLevelSetup {
	case DEBUG:
		conformedMessage = "[DEBUG] " + message
	case DEVELOP:
		conformedMessage = "[DEVELOPMENT] " + message
	}
	if err := os.WriteFile(LoggerAddress, []byte(conformedMessage), 0666); err != nil {
		log.Fatal(err)
	}
}
