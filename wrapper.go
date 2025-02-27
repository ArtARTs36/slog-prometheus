package slogprometheus

import (
	"context"
	"log/slog"
)

type WrappedHandler struct {
	base   slog.Handler
	second slog.Handler
}

func newWrappedHandler(base slog.Handler, our slog.Handler) slog.Handler {
	return &WrappedHandler{
		base:   base,
		second: our,
	}
}

func (h *WrappedHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.base.Enabled(ctx, level)
}

func (h *WrappedHandler) Handle(ctx context.Context, rec slog.Record) error {
	err := h.base.Handle(ctx, rec)

	if h.second.Enabled(ctx, rec.Level) {
		_ = h.second.Handle(ctx, rec)
	}

	return err
}

func (h *WrappedHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h.modify(h.base.WithAttrs(attrs))
}

func (h *WrappedHandler) WithGroup(name string) slog.Handler {
	return h.modify(h.base.WithGroup(name))
}

func (h *WrappedHandler) modify(base slog.Handler) slog.Handler {
	return &WrappedHandler{
		base:   base,
		second: h.second,
	}
}
