package slogprometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

var DefaultCollectors = NewCollectors()

func init() {
	prometheus.MustRegister(DefaultCollectors)
}
