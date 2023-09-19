package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/exp/slog"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

type Webserver struct {
	server *http.ServeMux
	logger *slog.Logger
	port   int
	mw     *CustomMiddleWare
}

func (web *Webserver) RegisterMetrics() {
	web.logger.Debug("Registering prometheus metrics handler")
	web.server.Handle("/metrics", promhttp.Handler())
}

func (web *Webserver) RegisterStaticFileServer() {
	web.logger.Debug("Registering static file server")
	fs := http.FileServer(http.Dir("../website/dist"))

	web.server.Handle("/", web.mw.WithMetricsMiddleware(web.mw.WithLoggingMiddleware(fs)))
}

func (web *Webserver) RegisterApiEndpoints() {
	web.logger.Debug("Registering api endpoints")
	apiHandler := http.HandlerFunc(rebuildHandler)
	web.server.Handle("/api/rebuild", web.mw.WithMetricsMiddleware(web.mw.WithLoggingMiddleware(apiHandler)))
}

func (web *Webserver) Start() error {
	web.logger.Info("Webserver started", "port", web.port)
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", web.port), web.server)
}
