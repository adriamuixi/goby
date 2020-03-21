package logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/goby/config"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Error   *log.Logger
	Warning *log.Logger
)

func InitLogger(configuration *config.LoggerConfig) {

	var traceHandle, infoHandle, warningHandle, errorHandle io.Writer
	if configuration.Error {
		errorHandle = os.Stderr //TODO: How to send errors to slack from platform?
	} else {
		errorHandle = ioutil.Discard
	}
	if configuration.Warning {
		warningHandle = os.Stderr
	} else {
		warningHandle = ioutil.Discard
	}
	if configuration.Info {
		infoHandle = os.Stdout
	} else {
		infoHandle = ioutil.Discard
	}
	if configuration.Trace {
		traceHandle = os.Stdout
	} else {
		traceHandle = ioutil.Discard
	}

	Trace = log.New(traceHandle,
		"Web.TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle,
		"Web.INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle,
		"Web.WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle,
		"Web.ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}
