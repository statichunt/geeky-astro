package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var rebuildHistogramMetric = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Namespace: "dizzyPulse",
	Name:      "rebuild_duration_seconds",
	Help:      "The total number of rebuilds",
	Buckets:   prometheus.DefBuckets,
}, []string{"method", "url", "status"})

var webRequestHistogramMetric = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Namespace: "dizzyPulse",
	Name:      "web_request_duration_micro_seconds",
	Help:      "The total number of rebuilds",
	Buckets:   prometheus.DefBuckets,
}, []string{"method", "url", "status"})
