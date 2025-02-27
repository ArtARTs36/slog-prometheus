package slogprometheus

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"log/slog"
)

type Collectors struct {
	logCount   *prometheus.CounterVec
	loggerInfo *loggerInfoCollector
}

func NewCollectors(namespace string) *Collectors {
	return &Collectors{
		logCount: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name:      "logs_count",
			Namespace: namespace,
			Help:      "Logs: count of logs per level",
		}, []string{"level"}),
		loggerInfo: newLoggerInfoCollector(namespace),
	}
}

func (c *Collectors) IncLogCount(lvl slog.Level) {
	c.logCount.WithLabelValues(prepareLogLevel(lvl)).Inc()
}

func (c *Collectors) Map() map[string]prometheus.Collector {
	return map[string]prometheus.Collector{
		"logs_count":  c.logCount,
		"logger_info": c.loggerInfo,
	}
}

func (c *Collectors) Register(registerer prometheus.Registerer) error {
	for name, collector := range c.Map() {
		if err := registerer.Register(collector); err != nil {
			return fmt.Errorf("failed to register collector %q: %w", name, err)
		}
	}

	return nil
}
