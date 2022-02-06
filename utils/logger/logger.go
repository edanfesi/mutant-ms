package logger

import (
	"os"

	"github.com/labstack/gommon/random"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

const prefixKey = "prefix"

func New(prefix string) *logrus.Entry {
	logger := logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.InfoLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02T15:04:05.1283386-00:00",
			LogFormat:       "[%time%][%lvl%][%prefix%] %msg%\n",
		},
	}

	if prefix == "-" {
		prefix = random.New().String(32, random.Alphabetic)
	}

	return logger.WithFields(
		logrus.Fields{
			prefixKey: prefix,
		},
	)
}
