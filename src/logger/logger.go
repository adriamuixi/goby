package logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"bitbucket.org/goby/config"
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
		"brand-service.TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle,
		"brand-service.INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle,
		"brand-service.WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle,
		"brand-service.ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}
