package slogprometheus

import (
	"log/slog"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct {
	count *prometheus.CounterVec
}

var logLevelMap = map[slog.Leveler]string{
	slog.LevelDebug: "debug",
	slog.LevelInfo:  "info",
	slog.LevelWarn:  "warn",
	slog.LevelError: "error",
}

func NewCollector(namespace string) *Collector {
	return &Collector{
		count: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name:      "logs_count",
			Namespace: namespace,
			Help:      "Slog: logs count per level",
		}, []string{"level"}),
	}
}

func (c *Collector) IncLogCount(lvl slog.Level) {
	c.count.WithLabelValues(c.prepareLogLevel(lvl)).Inc()
}

func (c *Collector) Map() map[string]prometheus.Collector {
	return map[string]prometheus.Collector{
		"logs_count": c.count,
	}
}

func (c *Collector) prepareLogLevel(lvl slog.Leveler) string {
	name, ok := logLevelMap[lvl]
	if ok {
		return name
	}

	return strings.ToLower(lvl.Level().String())
}
