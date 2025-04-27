package metrics

import (
    "fmt"
    "kube-metrics-collector/kubernetes"
)

type MetricsService struct {
    client kubernetes.KubernetesClient
}

// NewMetricsService creates a new MetricsService with given Kubernetes client
func NewMetricsService(client kubernetes.KubernetesClient) *MetricsService {
    return &MetricsService{client: client}
}

// PrintMetrics fetches and prints pod and node details
func (m *MetricsService) PrintMetrics() error {
    fmt.Println("Collecting Metrics...")

    // List Pods
    pods, err := m.client.ListPods("")
    if err != nil {
        return fmt.Errorf("failed to list pods: %w", err)
    }
    fmt.Println("Pods:")
    for _, pod := range pods {
        fmt.Printf(" - %s (Namespace: %s)\n", pod.Name, pod.Namespace)
    }

    // List Nodes
    nodes, err := m.client.ListNodes()
    if err != nil {
        return fmt.Errorf("failed to list nodes: %w", err)
    }
    fmt.Println("\nNodes:")
    for _, node := range nodes {
        fmt.Printf(" - %s\n", node.Name)
    }

    fmt.Println("\nMetrics collection completed ✅")
    return nil
}