package main

import (
	"time"

	"github.com/allen13/kubepacket/pkg/controller"
	"github.com/allen13/kubepacket/pkg/packet"
	"github.com/allen13/kubepacket/pkg/prom"
)

const (
	SNAPLEN = 65536
)

func main() {
	go prom.StartPrometheusEndpoint()
	flush, _ := time.ParseDuration("2m")
	go packet.Capture("lo", "", SNAPLEN, flush)
	controller.Start()
}
