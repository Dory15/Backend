package logger

import (
	"sync"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type LoggerLogrus struct {
	ExternalLogger *logrus.Logger
}

var loggerLogrusInstance LoggerLogrus
var lockLogger = &sync.Mutex{}

func NewLoggerLogrusInstace(fileName string) ILogger {
	if loggerLogrusInstance.ExternalLogger == nil {
		lockLogger.Lock()
		defer lockLogger.Unlock()
		if loggerLogrusInstance.ExternalLogger == nil {

			loggerLogrusInstance.ExternalLogger = getInstanceLogger()
			loggerLogrusInstance.Info("Creando Instancia de Logger Ahora: " + fileName)
		} else {
			loggerLogrusInstance.Info("Una Instancia de Logger ya habia sido creada: " + fileName)
		}
	} else {
		loggerLogrusInstance.Info("Una Instancia de Logger ya habia sido creada: " + fileName)
	}
	return loggerLogrusInstance
}

func (logger LoggerLogrus) Info(infoMessage string) {
	logger.ExternalLogger.Info(infoMessage)
}

func (logger LoggerLogrus) Error(errorMessage string) {
	logger.ExternalLogger.Error(errorMessage)
}

func (logger LoggerLogrus) Panic(panicMessage string) {
	logger.ExternalLogger.Panic(panicMessage)
}

func getInstanceLogger() *logrus.Logger {

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "./var/log/info_access.log",
		logrus.ErrorLevel: "./var/log/error_access.log",
		logrus.DebugLevel: "./var/log/debug_access.log",
		logrus.WarnLevel:  "./var/log/warning_access.log",
		logrus.FatalLevel: "./var/log/fatal_access.log",
		logrus.PanicLevel: "./var/log/panic_access.log",
		logrus.TraceLevel: "./var/log/trace_access.log",
	}

	Log := logrus.New()

	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))

	return Log
}
