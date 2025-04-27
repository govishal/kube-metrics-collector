package metrics

type NodeMetric struct {
	NodeName    string `json:"nodeName"`
	CpuUsage    string `json:"cpuUsage"`
	MemoryUsage string `json:"memoryUsage"`
	DiskUsage   string `json:"diskUsage"`
}

type PodMetric struct {
	PodName      string `json:"podName"`
	Namespace    string `json:"namespace"`
	CpuUsage     string `json:"cpuUsage"`
	MemoryUsage  string `json:"memoryUsage"`
}

type Metrics struct {
	ClusterName string       `json:"clusterName"`
	TotalNodes  int          `json:"totalNodes"`
	TotalPods   int          `json:"totalPods"`
	NodeMetrics []NodeMetric `json:"nodeMetrics"`
	PodMetrics  []PodMetric  `json:"podMetrics"`
}

func MockMetrics() Metrics {
	metrics := Metrics{
		ClusterName: "fake-cluster",
		TotalNodes:  3,
		TotalPods:   9,
		NodeMetrics: []NodeMetric{
			{"fake-node-1", "220m", "512Mi", "10Gi"},
			{"fake-node-2", "180m", "430Mi", "8Gi"},
			{"fake-node-3", "300m", "700Mi", "12Gi"},
		},
		PodMetrics: []PodMetric{
			{"fake-pod-1", "default", "50m", "100Mi"},
			{"fake-pod-2", "default", "30m", "80Mi"},
			{"fake-pod-3", "kube-system", "70m", "200Mi"},
		},
	}
	return metrics
}
