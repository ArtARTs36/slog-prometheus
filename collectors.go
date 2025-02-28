package slogprometheus

import (
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
)

type Collectors struct {
	logCount   *prometheus.CounterVec
	loggerInfo *loggerInfoCollector
}

func NewCollectors() *Collectors {
	return &Collectors{
		logCount: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "slog_logs_count",
			Help: "Logs: count of logs per level",
		}, []string{"level"}),
		loggerInfo: newLoggerInfoCollector(),
	}
}

func (c *Collectors) IncLogCount(lvl slog.Level) {
	c.logCount.WithLabelValues(prepareLogLevel(lvl)).Inc()
}

func (c *Collectors) Describe(desc chan<- *prometheus.Desc) {
	c.logCount.Describe(desc)
	c.loggerInfo.Describe(desc)
}

func (c *Collectors) Collect(metric chan<- prometheus.Metric) {
	c.logCount.Collect(metric)
	c.loggerInfo.Collect(metric)
}
