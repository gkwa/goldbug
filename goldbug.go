package goldbug

import (
	"log/slog"
	"os"
)

func SetDefaultLogger() {
	logLevel := &slog.LevelVar{} // INFO
	logLevel.Set(slog.LevelDebug)
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey && len(groups) == 0 {
				return slog.Attr{}
			}
			return a
		},
	}
	handler1 := slog.NewTextHandler(os.Stderr, &opts)

	slog.SetDefault(slog.New(handler1))
}
