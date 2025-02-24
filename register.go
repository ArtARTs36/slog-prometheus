package slogprometheus

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var defaultCollector = NewCollector("slog")

func RegisterDefault() error {
	return register(defaultCollector, prometheus.DefaultRegisterer)
}

func Register(
	namespace string,
	registerer prometheus.Registerer,
) error {
	collectors := NewCollector(namespace)

	return register(collectors, registerer)
}

func register(collectors *Collector, registerer prometheus.Registerer) error {
	for name, collector := range collectors.Map() {
		if err := registerer.Register(collector); err != nil {
			return fmt.Errorf("failed to register collector %q: %w", name, err)
		}
	}

	return nil
}
