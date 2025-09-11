package metrics

import (
	"context"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Gauge for pod readiness: 1 = Ready, 0 = NotReady
var PodReadyStatus = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pod_ready_status",
		Help: "Pod readiness status (1 = Ready, 0 = NotReady)",
	},
	[]string{"namespace", "pod"},
)

// Gauge for total container restarts per pod
var PodRestartCount = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pod_restart_count",
		Help: "Total container restarts per pod",
	},
	[]string{"namespace", "pod"},
)

// Register metrics with Prometheus
func InitPodMetrics() {
	prometheus.MustRegister(PodReadyStatus)
	prometheus.MustRegister(PodRestartCount)
}

// Poll Kubernetes API and update metrics
func RecordPodMetrics(clientset *kubernetes.Clientset) {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error fetching pods: %v", err)
		return
	}

	for _, pod := range pods.Items {
		ready := 0.0
		restarts := 0.0

		for _, cs := range pod.Status.ContainerStatuses {
			if cs.Ready {
				ready = 1.0 // At least one container is ready
			}
			restarts += float64(cs.RestartCount)
		}

		PodReadyStatus.WithLabelValues(pod.Namespace, pod.Name).Set(ready)
		PodRestartCount.WithLabelValues(pod.Namespace, pod.Name).Set(restarts)
	}
}