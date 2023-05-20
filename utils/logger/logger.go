package logger

import (
	"fmt"
	"go.uber.org/zap"
)

type CtxLogger struct {
	Zap *zap.Logger
}

func (ctxlog *CtxLogger) NewPrefix(prefix string) *CtxLogger {
	return &CtxLogger{ctxlog.Zap.Named(prefix)}
}

func (ctxlog *CtxLogger) Info(logText string, a ...any) {
	ctxlog.Zap.Info(fmt.Sprintf(logText, a...))
}

func (ctxlog *CtxLogger) InfofJSON(logText string, a ...any) {
	ctxlog.Zap.Info(fmt.Sprintf(logText+": %+v", a...))
}
