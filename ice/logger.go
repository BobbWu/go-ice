package ice

type Logger interface {
	Print(message string)
	Trace(category string, message string)
	Warning(message string)
	Error(message string)
	GetPrefix() string
	CloneWithPrefix(prefix string) Logger
}
