package kubernetes

import (
    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FakeKubeClient struct{}

func NewFakeKubeClient() KubernetesClient {
    return &FakeKubeClient{}
}

func (f *FakeKubeClient) ListPods(namespace string) ([]corev1.Pod, error) {
    return []corev1.Pod{
        {ObjectMeta: metav1.ObjectMeta{Name: "fake-pod-1"}},
        {ObjectMeta: metav1.ObjectMeta{Name: "fake-pod-2"}},
    }, nil
}

func (f *FakeKubeClient) ListNodes() ([]corev1.Node, error) {
    return []corev1.Node{
        {ObjectMeta: metav1.ObjectMeta{Name: "fake-node-1"}},
    }, nil
}