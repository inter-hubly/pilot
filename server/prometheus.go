package server

import (
	"context"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (e *environment) configPrometheus(ctx context.Context) {
	registry := prometheus.NewRegistry()

	// Inicialize as métricas
	httpRequestsTotal := e.newHTTPRequestsTotal()
	httpRequestDuration := e.newHTTPRequestDuration()
	memoryUsage := e.newMemoryUsage()
	goroutinesCount := e.newGoroutinesCount()
	errorCounter := e.newErrorCounter()
	dbQueryDuration := e.newDBQueryDuration()
	appUptime := e.newAppUptime()

	// Registre as métricas no registry
	registry.MustRegister(
		httpRequestsTotal,
		httpRequestDuration,
		memoryUsage,
		goroutinesCount,
		errorCounter,
		dbQueryDuration,
		appUptime,
	)
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
}

// Total HTTP Requests
func (e *environment) newHTTPRequestsTotal() *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "path", "status_code"},
	)
}

// HTTP Request Duration
func (e *environment) newHTTPRequestDuration() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
}

// Memory Usage
func (e *environment) newMemoryUsage() prometheus.Gauge {
	return prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "app_memory_usage_bytes",
			Help: "Current memory usage of the application in bytes.",
		},
	)
}

// Goroutines Count
func (e *environment) newGoroutinesCount() prometheus.Gauge {
	return prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "go_goroutines",
			Help: "Number of goroutines currently running.",
		},
	)
}

// Error Counter
func (e *environment) newErrorCounter() *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "app_errors_total",
			Help: "Total number of application errors.",
		},
		[]string{"type", "severity"},
	)
}

// Database Query Duration
func (e *environment) newDBQueryDuration() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_query_duration_seconds",
			Help:    "Duration of database queries in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"query_type"},
	)
}

// Uptime Counter
func (e *environment) newAppUptime() prometheus.Counter {
	return prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "app_uptime_seconds",
			Help: "Total uptime of the application in seconds.",
		},
	)
}
