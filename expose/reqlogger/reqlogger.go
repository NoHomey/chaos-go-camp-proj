package reqlogger

import (
	"time"

	"go.uber.org/zap"
)

//Request represents request data to be logged.
type Request struct {
	Method string
	URL    string
	IP     string
}

//Error represents the data for and error to be logged.
type Error struct {
	Name       string
	Message    string
	Data       interface{}
	StatusCode int
	Err        error
}

//Logger abstracts request logger.
type Logger interface {
	Request(req Request)
	Response(req Request, d time.Duration, status int)
	Error(req Request, d time.Duration, err Error)
}

//New constructs new request logger.
func New(logger *zap.Logger) Logger {
	return impl{logger}
}

type impl struct {
	logger *zap.Logger
}

func (l impl) Request(req Request) {
	l.logger.Info(
		"handling request",
		zap.String("url", req.URL),
		zap.String("method", req.Method),
		zap.String("ip", req.IP),
		zap.Time("time", time.Now()),
	)
}

func (l impl) Response(req Request, d time.Duration, status int) {
	l.logger.Info(
		"sending result",
		zap.String("url", req.URL),
		zap.String("method", req.Method),
		zap.String("ip", req.IP),
		zap.Time("time", time.Now()),
		zap.Duration("duration", d),
		zap.Int("statusCode", status),
	)
}

func (l impl) Error(req Request, d time.Duration, err Error) {
	l.logger.Error(
		"sending error",
		zap.String("url", req.URL),
		zap.String("method", req.Method),
		zap.String("ip", req.IP),
		zap.Time("time", time.Now()),
		zap.Duration("duration", d),
		zap.String("errorName", err.Name),
		zap.String("errorMessage", err.Message),
		zap.Reflect("errorData", err.Data),
		zap.Reflect("statusCode", err.StatusCode),
		zap.Error(err.Err),
	)
}
