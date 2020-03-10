package main

import (
	"net/http"

	"druid-exporter/collector"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.

	metric := collector.Collector()
	prometheus.MustRegister(metric)

	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
