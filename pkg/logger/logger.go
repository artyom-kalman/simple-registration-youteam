package logger

import (
	"fmt"
	"log/slog"
	"os"
)

var logger slog.Logger

func InitLogger() {
	opts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger = *slog.New(slog.NewTextHandler(os.Stdout, &opts))

	logger.Info("Initialized logger")
}

func Info(mes string, args ...any) {
	logger.Info(fmt.Sprintf(mes, args...))
}

func Debug(mes string, args ...any) {
	logger.Debug(fmt.Sprintf(mes, args...))
}

func Error(mes string, args ...any) {
	logger.Error(fmt.Sprintf(mes, args...))
}
