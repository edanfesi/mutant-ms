package context

import (
	"time"

	"golang.org/x/net/context"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"mutant-ms/utils/logger"
)

type key string

const (
	loggerKey  = key("logger")
	echoLogger = "logger"
)

func GetContext(c echo.Context) context.Context {
	return context.WithValue(c.Request().Context(), loggerKey, GetEchoContextLogger(c))
}

func GeContextWithTimeout(c echo.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(GetContext(c), timeout)
}

func GetLogger(ctx context.Context) *logrus.Entry {
	log, ok := ctx.Value(loggerKey).(*logrus.Entry)
	if !ok {
		return logger.New("-")
	}

	return log
}

func GetEchoContextLogger(ctx echo.Context) *logrus.Entry {
	log, ok := ctx.Get(echoLogger).(*logrus.Entry)
	if !ok {
		return logger.New("-")
	}

	return log
}
