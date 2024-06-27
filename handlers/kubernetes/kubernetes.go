package kubernetes

import (
	"context"
	"path/filepath"

	"github.com/mateoops/k8s-roa/handlers/prometheus"
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

func (kubernetesHandler *KubernetesHandler) ListNodes() []prometheus.NodeMetrics {
	nodes, err := kubernetesHandler.clientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	var nodesList []prometheus.NodeMetrics
	for _, node := range nodes.Items {
		nodesList = append(nodesList, prometheus.NodeMetrics{Name: node.Name})
	}
	return nodesList
}

func (kubernetesHandler *KubernetesHandler) GetNodeUsageMetrics(nodeName string) prometheus.NodeUsageMetrics {
	metrics, err := kubernetesHandler.metricsClientSet.MetricsV1beta1().NodeMetricses().Get(context.TODO(), nodeName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	// CPU usage measured in millicpu (1000 millicpu = 1cpu)
	cpuUsage := metrics.Usage.Cpu().AsDec().UnscaledBig().Int64() / 1024 / 1024
	// Memory usage measured in MB
	memoryUsage := metrics.Usage.Memory().MilliValue() / 1024 / 1024 / 1024
	return prometheus.NodeUsageMetrics{CpuUsage: cpuUsage, MemoryUsage: memoryUsage}
}
