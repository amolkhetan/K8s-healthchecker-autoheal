package main

import (
	"log"
	"time"

	"github.com/amolkhetan/k8s-health-checker/pkg/client"
)

func main() {
	for {
		if err := client.TestKubeAPI(); err != nil {
			log.Printf("❌ API test failed: %v", err)
		} else {
			log.Println("✅ API test succeeded")
		}

		time.Sleep(30 * time.Second)
	}
}
