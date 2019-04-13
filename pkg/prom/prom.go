package prom

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	PacketCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "packet_count",
			Help: "Track counts of packets for certain source and destinations",
		},
		[]string{"proto", "src_ip", "src_port", "dst_ip", "dst_port"},
	)
)

func init() {
	prometheus.MustRegister(PacketCount)
}

// StartPrometheusEndpoint - Setup and Run the prometheus endpoint
func StartPrometheusEndpoint() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
