package kubernetes

import (
    corev1 "k8s.io/api/core/v1"
)

type KubernetesClient interface {
    ListPods(namespace string) ([]corev1.Pod, error)
    ListNodes() ([]corev1.Node, error)
}