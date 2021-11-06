package main

import (
	"bytes"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"net/http"
)

type Metrics struct {
	RequestTotal *prometheus.CounterVec
	RequestDuration *prometheus.HistogramVec
}

var metrics Metrics

func main(){


http.Handle("/metrics",promhttp.Handler())
http.ListenAndServe(":2112",nil)
}

func createMetrics(){
	metrics.RequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "metrics_app_requests_total",
		Help: "Кол-во запросов к серверу",
	},[]string{"method","handler","code"},)
metrics.RequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{...},
[]string{"method","handler","code"},)
prometheus.MustRegister(metrics.RequestTotal)
prometheus.MustRegister(metrics.RequestTotal)
}