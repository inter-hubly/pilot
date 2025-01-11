package hlog

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/inter-hubly/pilot/hctx"
)

var logger *slog.Logger

func init() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	logger = slog.New(handler)
}

type LogStruct struct {
	Callfunc string
	Message  string
}

func NewLogStruct(svc, sms string) *LogStruct {
	return &LogStruct{
		Callfunc: svc,
		Message:  sms,
	}
}

func (l *LogStruct) toString(ctx context.Context) string {
	return fmt.Sprintf("{Service: %s.%s, Message: %s}", hctx.Tenant.Get(ctx), l.Callfunc, l.Message)
}

func Debug(ctx context.Context, svc, msg string) {
	logger.Debug(NewLogStruct(svc, msg).toString(ctx))
}

func Info(ctx context.Context, svc, msg string) {
	logger.Info(NewLogStruct(svc, msg).toString(ctx))
}

func Warn(ctx context.Context, svc, msg string) {
	logger.Warn(NewLogStruct(svc, msg).toString(ctx))
}

func Error(ctx context.Context, svc, msg string) {
	logger.Error(NewLogStruct(svc, msg).toString(ctx))
}
