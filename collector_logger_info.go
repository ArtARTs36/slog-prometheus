package slogprometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type loggerInfoCollector struct {
	collector *prometheus.GaugeVec

	valuesReady bool
}

func newLoggerInfoCollector() *loggerInfoCollector {
	return &loggerInfoCollector{
		collector: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "slog_logger_info",
			Help: "Logger info",
		}, []string{"level"}),
	}
}

func (c *loggerInfoCollector) Describe(desc chan<- *prometheus.Desc) {
	c.collector.Describe(desc)
}

func (c *loggerInfoCollector) Collect(metric chan<- prometheus.Metric) {
	c.prepareValues()

	c.collector.Collect(metric)
}

func (c *loggerInfoCollector) prepareValues() {
	if c.valuesReady {
		return
	}

	c.collector.WithLabelValues(
		prepareLogLevel(calcCurrentLogLevel()),
	).Set(1)
	c.valuesReady = true
}
