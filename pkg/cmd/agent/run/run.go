package run

import (
	"fmt"
	"os"
	"time"

	kubernetesHandler "github.com/mateoops/k8s-roa/handlers/kubernetes"
	prometheusHandler "github.com/mateoops/k8s-roa/handlers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	prometheusNodeMemoryGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "roa_nodes",
			Name:      "memory_usage",
		},
		[]string{"node"})

	prometheusPodMemoryGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "roa_pods",
			Name:      "memory_usage",
		},
		[]string{"namespace", "pod", "container"})

	prometheusNodeCpuGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "roa_nodes",
			Name:      "cpu_usage",
		},
		[]string{"node"})

	prometheusPodCpuGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "roa_pods",
			Name:      "cpu_usage",
		},
		[]string{"namespace", "pod", "container"})
)

func NewCmdRun() *cobra.Command {

	cmd :=
		&cobra.Command{
			Use:    "run",
			Hidden: true,
			Run: func(cmd *cobra.Command, args []string) {
				if authMode := viper.Get("authMode"); authMode == "kube-config" {
					fmt.Println("Starting agent...")

					// register Prometheus gauges
					registerPrometheusGauges()

					kubernetes := kubernetesHandler.NewKubernetesHandler()
					interval := time.Duration(viper.GetInt("scrapInterval"))

					// scrap metrics
					go scrapNodesMetrics(kubernetes, interval)
					go scrapPodsMetrics(kubernetes, interval)

					// prepare prometheus handler and expose metrics
					prometheus := prometheusHandler.NewPrometheusHandler()
					fmt.Println("Agent is running!")
					prometheus.ExposeMetrics()

				} else if authMode == "native" {
					fmt.Println("Native auth option not supported yet.")
					os.Exit(1)
				} else {
					fmt.Println("Not supported auth option", authMode)
					os.Exit(1)
				}
			},
		}
	return cmd
}

func registerPrometheusGauges() {
	prometheus.MustRegister(prometheusNodeCpuGauge)
	prometheus.MustRegister(prometheusNodeMemoryGauge)
	prometheus.MustRegister(prometheusPodMemoryGauge)
	prometheus.MustRegister(prometheusPodCpuGauge)
}

func scrapNodesMetrics(kubernetes *kubernetesHandler.KubernetesHandler, interval time.Duration) {
	for {
		nodes := kubernetes.ListNodes()
		for _, node := range nodes {
			metrics := kubernetes.GetNodeUsageMetrics(node.Name)
			prometheusNodeCpuGauge.WithLabelValues(node.Name).Add(float64(metrics.CpuUsage))
			prometheusNodeMemoryGauge.WithLabelValues(node.Name).Add(float64(metrics.MemoryUsage))
		}
		time.Sleep(interval * time.Second)
	}
}

func scrapPodsMetrics(kubernetes *kubernetesHandler.KubernetesHandler, interval time.Duration) {
	for {
		pods := kubernetes.ListPods()
		for _, pod := range pods {
			metricsPod := kubernetes.GetPodUsageMetrics(pod.Name, pod.Namespace)
			for _, metricsContainer := range metricsPod {
				prometheusPodCpuGauge.WithLabelValues(pod.Namespace, pod.Name, metricsContainer.Container.Name).Add(float64(metricsContainer.CpuUsage))
				prometheusPodMemoryGauge.WithLabelValues(pod.Namespace, pod.Name, metricsContainer.Container.Name).Add(float64(metricsContainer.MemoryUsage))
			}
		}
		time.Sleep(interval * time.Second)
	}
}
