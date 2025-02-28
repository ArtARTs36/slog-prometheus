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

	http.ListenAndServe("localhost:8080", promhttp.Handler())
}
```

---

Usage with JSON Handler

```go
package main

import (
	slogprometheus "github.com/artarts36/slog-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slogprometheus.Option{
		Level: slog.LevelDebug,
	}.WrapHandler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	logger.Debug("debug log msg")
	logger.Info("info log msg")
	logger.Warn("warn log msg")
	logger.Error("error log msg")

	http.ListenAndServe("localhost:8080", promhttp.Handler())
}
```

## Exposing metrics

```text
# HELP slog_logger_info Logger info
# TYPE slog_logger_info gauge
slog_logger_info{level="DEBUG"} 1
# HELP slog_logs_count Logs: count of logs per level
# TYPE slog_logs_count counter
slog_logs_count{level="DEBUG"} 1
slog_logs_count{level="ERROR"} 1
slog_logs_count{level="INFO"} 1
slog_logs_count{level="WARN"} 1
```
