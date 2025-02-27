package slogprometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

var defaultCollectors = NewCollectors("slog")

func RegisterDefault() error {
	return defaultCollectors.Register(prometheus.DefaultRegisterer)
}

func Register(
	namespace string,
	registerer prometheus.Registerer,
) error {
	return NewCollectors(namespace).Register(registerer)
}
