package kubernetes

import (
	"context"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	metricsClient "k8s.io/metrics/pkg/client/clientset/versioned"
)

type KubernetesHandler struct {
	clientSet        *kubernetes.Clientset
	metricsClientSet *metricsClient.Clientset
}

func NewKubernetesHandler() *KubernetesHandler {
	kubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		panic(err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	metricsClientSet, err := metricsClient.NewForConfig(config)

	return &KubernetesHandler{clientSet: clientSet, metricsClientSet: metricsClientSet}
}

func (kubernetesHandler *KubernetesHandler) ListNodes() {
	nodes, err := kubernetesHandler.clientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, node := range nodes.Items {
		fmt.Println("Node: ", node.Name)
	}
}

func (kubernetesHandler *KubernetesHandler) ListNodesMetrics() {
	nodes, err := kubernetesHandler.clientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, node := range nodes.Items {
		fmt.Println("Node: ", node.Name)
		metrics, err := kubernetesHandler.metricsClientSet.MetricsV1beta1().NodeMetricses().Get(context.TODO(), node.Name, metav1.GetOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(metrics)
	}

}
