package goldbug

import (
	"log/slog"
	"os"
	"path/filepath"
)

func SetDefaultLoggerText(level slog.Level) {
	logLevel := &slog.LevelVar{} // INFO
	logLevel.Set(level)
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey && len(groups) == 0 {
				return slog.Attr{}
			}
			if a.Key == slog.SourceKey {
				source, _ := a.Value.Any().(*slog.Source)
				if source != nil {
					fileName := filepath.Base(source.File)
					parentDir := filepath.Base(filepath.Dir(source.File))
					source.File = filepath.Join(parentDir, fileName)
				}
			}
			return a
		},
	}
	handler := slog.NewTextHandler(os.Stderr, &opts)

	slog.SetDefault(slog.New(handler))
}

func SetDefaultLoggerJson(level slog.Level) {
	logLevel := &slog.LevelVar{} // INFO
	logLevel.Set(level)
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				source, _ := a.Value.Any().(*slog.Source)
				if source != nil {
					fileName := filepath.Base(source.File)
					parentDir := filepath.Base(filepath.Dir(source.File))
					source.File = filepath.Join(parentDir, fileName)
				}
			}
			return a
		},
	}
	handler := slog.NewJSONHandler(os.Stderr, &opts)

	slog.SetDefault(slog.New(handler))
}
