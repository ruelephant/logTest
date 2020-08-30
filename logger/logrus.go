package logger

import (
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	lr *logrus.Logger
}

func (ll *LogrusLogger) Sum(a, b int) {
	ll.lr.Print("Library calc", a, b)
}

func (ll *LogrusLogger) Connect(host string) {
	ll.lr.Print("Library connecting to " + host)
}

func NewLogrusLogger() *LogrusLogger {
	lr := logrus.New()
	// Настраивем логгер по вкусу, или нет

	return &LogrusLogger{
		lr: lr,
	}
}