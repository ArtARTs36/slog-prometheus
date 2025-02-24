# slog: Prometheus handler

[![License](https://img.shields.io/github/license/artarts36/slog-prometheus)](./LICENSE)

A Prometheus handler for [slog](https://pkg.go.dev/log/slog) go Library

```
go get github.com/artarts36/slog-prometheus
```

## Usage

```go
package main

import (
	slogprometheus "github.com/artarts36/slog-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log/slog"
	"net/http"
)

func main() {
	logger := slog.New(slogprometheus.Option{
		Level: slog.LevelDebug,
	}.NewHandler())

	logger.Debug("debug log msg")
	logger.Info("info log msg")
	logger.Warn("warn log msg")
	logger.Error("error log msg")

	slogprometheus.RegisterDefault()

	http.ListenAndServe("localhost:8080", promhttp.Handler())
}
```

## Exposing metrics

```text
# HELP slog_logger_info Logger info
# TYPE slog_logger_info gauge
slog_logger_info{level="debug"} 1
# HELP slog_logs_count Slog: logs logCount per level
# TYPE slog_logs_count counter
slog_logs_count{level="debug"} 1
slog_logs_count{level="error"} 1
slog_logs_count{level="info"} 1
slog_logs_count{level="warn"} 1
```
