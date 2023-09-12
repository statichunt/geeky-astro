package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/exp/slog"
)

type CustomMiddleWare struct {
	logger *slog.Logger
}

func (mw *CustomMiddleWare) WithLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ip, _ := getIP(r)
		mw.logger.Debug("Request", "method", r.Method, "url", r.URL, "ip", ip)
		next.ServeHTTP(w, r)
		mw.logger.Debug("Response", "method", r.Method, "url", r.URL, "time", fmt.Sprintf("%d µs", time.Since(start).Microseconds()))
	})
}

func (mw *CustomMiddleWare) WithMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		req := statusRecorder{w, http.StatusOK}
		next.ServeHTTP(&req, r)
		webRequestHistogramMetric.WithLabelValues(r.Method, r.URL.Path, fmt.Sprintf("%d", req.status)).Observe(float64(time.Since(start).Microseconds()))
	})
}
