package main

import (
	"net/http"
	"time"

	"github.com/amolkhetan/k8s-health-monitor/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, _ := rest.InClusterConfig()
	clientset, _ := kubernetes.NewForConfig(config)

	metrics.InitNodeMetrics()
	metrics.InitPodMetrics()

	go func() {
		for {
			metrics.RecordNodeMetrics(clientset)
			metrics.RecordPodMetrics(clientset)
			time.Sleep(30 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}