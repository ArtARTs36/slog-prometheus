package slogprometheus

import (
	"context"
	"log/slog"
)

type Handler struct {
	option Option
}

type Option struct {
	// Level sets the minimum log level to capture and collect.
	// Logs at this level and above will be processed. The default level is debug.
	Level slog.Leveler

	// If not provided, the collectors with namespace 'slog' is used by default.
	Collector *Collector
}

func (o Option) NewHandler() *Handler {
	if o.Level == nil {
		o.Level = slog.LevelDebug
	}

	if o.Collector == nil {
		o.Collector = defaultCollector
	}

	return &Handler{
		option: o,
	}
}

func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.option.Level.Level()
}

func (h *Handler) Handle(ctx context.Context, rec slog.Record) error {
	h.HandleSilent(ctx, rec)
	return nil
}

func (h *Handler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *Handler) WithGroup(_ string) slog.Handler {
	return h
}

func (h *Handler) HandleSilent(_ context.Context, rec slog.Record) {
	h.option.Collector.IncLogCount(rec.Level)
}
