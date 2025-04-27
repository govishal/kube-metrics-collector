package kubernetes

import "kube-metrics-collector/config"

func NewKubernetesClient(cfg config.Config) (KubernetesClient, error) {
    if cfg.UseFake {
        return NewFakeKubeClient(), nil
    }
    return NewRealKubeClient()
}