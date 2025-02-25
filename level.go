package slogprometheus

import (
	"context"
	"log/slog"
	"strings"
)

var logLevelMap = map[slog.Leveler]string{
	slog.LevelDebug: slog.LevelDebug.String(),
	slog.LevelInfo:  slog.LevelInfo.String(),
	slog.LevelWarn:  slog.LevelWarn.String(),
	slog.LevelError: slog.LevelError.String(),
}

func prepareLogLevel(lvl slog.Leveler) string {
	name, ok := logLevelMap[lvl]
	if ok {
		return name
	}

	return strings.ToUpper(lvl.Level().String())
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
