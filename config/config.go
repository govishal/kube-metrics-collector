package config

import "os"

type Config struct {
	UseFake    bool
	ClusterURL string
}

func LoadConfig() Config {
	useFake := os.Getenv("USE_FAKE") == "true"
	clusterURL := os.Getenv("CLUSTER_URL")

	return Config{
		UseFake:    useFake,
		ClusterURL: clusterURL,
	}
}
