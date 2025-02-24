package slogprometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"log/slog"
)

type Collectors struct {
	logCount   *prometheus.CounterVec
	loggerInfo *loggerInfoCollector
}

func NewCollector(namespace string) *Collectors {
	return &Collectors{
		logCount: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name:      "logs_count",
			Namespace: namespace,
			Help:      "Slog: logs logCount per level",
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
