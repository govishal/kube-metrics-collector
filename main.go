package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"kube-metrics-collector/config"
	"kube-metrics-collector/kubernetes"
	"kube-metrics-collector/metrics"
)

var kubeClient kubernetes.KubernetesClient

func main() {

	cfg := config.Config{}
	cfg.UseFake = false

	// Initialize Kubernetes client based on the config
	var err error
	if cfg.UseFake {
		fmt.Println("Using Fake Kubernetes Client")
		kubeClient = kubernetes.NewFakeKubeClient()
	} else {
		fmt.Println("Using Real Kubernetes Client")
		kubeClient, err = kubernetes.NewRealKubeClient()
		if err != nil {
			log.Fatalf("Error creating Kubernetes client: %v", err)
		}
	}

	metricsService := metrics.NewMetricsService(kubeClient)

	// // Register the routes with the config passed to handlers
	http.HandleFunc("/print-metrics", func(w http.ResponseWriter, r *http.Request) {
		err := metricsService.PrintMetrics()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Metrics printed in server logs ✅"))
	})

	http.HandleFunc("/pods", handleListPods)
	http.HandleFunc("/nodes", handleListNodes)
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		handleMetrics(w, r, &cfg) // Pass the cfg to the handler
	})

	port := getPort()
	fmt.Printf("Server starting at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleListPods(w http.ResponseWriter, r *http.Request) {
	namespace := r.URL.Query().Get("namespace")
	if namespace == "" {
		namespace = "default"
	}

	// Fetch pods from the client (real or fake)
	pods, err := kubeClient.ListPods(namespace)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Format and return the pod data as JSON
	response := make([]map[string]interface{}, len(pods))
	for i, pod := range pods {
		response[i] = map[string]interface{}{
			"podName":   pod.Name,
			"namespace": pod.Namespace,
		}
	}
	writeJSON(w, response)
}

func handleListNodes(w http.ResponseWriter, r *http.Request) {
	// Fetch nodes from the client (real or fake)
	nodes, err := kubeClient.ListNodes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Format and return the node data as JSON
	response := make([]map[string]interface{}, len(nodes))
	for i, node := range nodes {
		response[i] = map[string]interface{}{
			"nodeName": node.Name,
		}
	}
	writeJSON(w, response)
}

// Handle Metrics request
func handleMetrics(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	var metricsData interface{}
	fmt.Println(cfg.UseFake)

	if cfg.UseFake {
		// Return mock metrics for the fake client
		metricsData = metrics.MockMetrics()
	} else {
		// Fetch real metrics from the cluster (e.g., from metrics-server)
		nodes, err := kubeClient.ListNodes()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error listing nodes: %v", err), http.StatusInternalServerError)
			return
		}

		pods, err := kubeClient.ListPods("default")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error listing pods: %v", err), http.StatusInternalServerError)
			return
		}

		nodeMetrics := make([]map[string]interface{}, len(nodes))
		for i, node := range nodes {
			nodeMetrics[i] = map[string]interface{}{
				"nodeName":    node.Name,
				"cpuUsage":    "100m", // Placeholder
				"memoryUsage": "1Gi",  // Placeholder
				"diskUsage":   "20Gi", // Placeholder
			}
		}

		podMetrics := make([]map[string]interface{}, len(pods))
		for i, pod := range pods {
			podMetrics[i] = map[string]interface{}{
				"podName":     pod.Name,
				"namespace":   pod.Namespace,
				"cpuUsage":    "50m",   // Placeholder
				"memoryUsage": "100Mi", // Placeholder
			}
		}

		metricsData = map[string]interface{}{
			"clusterName": "real-cluster", // Placeholder
			"totalNodes":  len(nodes),
			"totalPods":   len(pods),
			"nodeMetrics": nodeMetrics,
			"podMetrics":  podMetrics,
		}
	}

	writeJSON(w, metricsData)
}

// writeJSON is used to send the response as JSON
func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
