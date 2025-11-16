package main

import (
	"log"
	"net/http"
	"time"

	"github.com/amolkhetan/k8s-health-monitor/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// Use in-cluster config for Minikube
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Failed to get in-cluster config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create Kubernetes client: %v", err)
	}

	metrics.InitNodeMetrics()
	metrics.InitPodMetrics()

	// Poll metrics every 30 seconds
	go func() {
		for {
			metrics.RecordNodeMetrics(clientset)
			metrics.RecordPodMetrics(clientset)
			time.Sleep(30 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Println("Serving metrics on :2112/metrics")
	log.Fatal(http.ListenAndServe(":2112", nil))
}