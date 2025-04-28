package kubernetes

import (
	"context"
	"os"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type RealKubeClient struct {
	clientset *kubernetes.Clientset
}

func NewRealKubeClient() (KubernetesClient, error) {
	// First, try in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		// If failed (probably running locally), try kubeconfig
		kubeConfigPath := os.Getenv("KUBECONFIG")
		if kubeConfigPath == "" {
			// If not set, fallback to project relative path
			kubeConfigPath = "./.kube/config"
		}

		config, err = clientcmd.BuildConfigFromFlags("", kubeConfigPath)
		if err != nil {
			return nil, err
		}
	}

	// Create the Kubernetes client using the config
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &RealKubeClient{clientset: clientset}, nil
}

func (r *RealKubeClient) ListPods(namespace string) ([]corev1.Pod, error) {
	pods, err := r.clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods.Items, nil
}

func (r *RealKubeClient) ListNodes() ([]corev1.Node, error) {
	nodes, err := r.clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return nodes.Items, nil
}
