package logger

type ILogger interface {
	Info(string)
	Error(string)
	Panic(string)
}

func NewLoggerInstace(fileName string) ILogger {
	return NewLoggerLogrusInstace(fileName)
}
