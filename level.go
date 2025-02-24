package slogprometheus

import (
	"context"
	"log/slog"
	"strings"
)

var logLevelMap = map[slog.Leveler]string{
	slog.LevelDebug: "debug",
	slog.LevelInfo:  "info",
	slog.LevelWarn:  "warn",
	slog.LevelError: "error",
}

func prepareLogLevel(lvl slog.Leveler) string {
	name, ok := logLevelMap[lvl]
	if ok {
		return name
	}

	return strings.ToLower(lvl.Level().String())
}

func calcCurrentLogLevel() slog.Leveler {
	return calcLogLevel(slog.Default().Handler())
}

func calcLogLevel(handler slog.Handler) slog.Leveler {
	lvlOrder := []slog.Level{
		slog.LevelDebug,
		slog.LevelInfo,
		slog.LevelWarn,
		slog.LevelError,
	}
	ctx := context.Background()

	for _, lvl := range lvlOrder {
		if handler.Enabled(ctx, lvl) {
			return lvl
		}
	}

	return slog.LevelDebug
}
