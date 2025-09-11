package metrics

import (
	"context"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var NodeReadyStatus = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "node_ready_status",
		Help: "Node readiness status (1 = Ready, 0 = NotReady)",
	},
	[]string{"node"},
)

func InitNodeMetrics() {
	prometheus.MustRegister(NodeReadyStatus)
}

func RecordNodeMetrics(clientset *kubernetes.Clientset) {
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error fetching nodes: %v", err)
		return
	}

	for _, node := range nodes.Items {
		ready := 0.0
		for _, cond := range node.Status.Conditions {
			if cond.Type == v1.NodeReady && cond.Status == v1.ConditionTrue {
				ready = 1.0
				break
			}
		}
		NodeReadyStatus.WithLabelValues(node.Name).Set(ready)
	}
}