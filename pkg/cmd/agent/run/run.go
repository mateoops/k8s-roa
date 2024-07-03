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
	prometheusMemoryGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "roa_nodes",
			Name:      "memory_usage",
		},
		[]string{"node"})

	prometheusCpuGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "roa_nodes",
			Name:      "cpu_usage",
		},
		[]string{"node"})
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
					prometheus.MustRegister(prometheusCpuGauge)
					prometheus.MustRegister(prometheusMemoryGauge)

					kubernetes := kubernetesHandler.NewKubernetesHandler()
					interval := time.Duration(viper.GetInt("scrapInterval"))

					go func() {
						for {
							nodes := kubernetes.ListNodes()
							for _, node := range nodes {
								metrics := kubernetes.GetNodeUsageMetrics(node.Name)
								prometheusCpuGauge.WithLabelValues(node.Name).Add(float64(metrics.CpuUsage))
								prometheusMemoryGauge.WithLabelValues(node.Name).Add(float64(metrics.MemoryUsage))

								time.Sleep(interval * time.Second)
							}
						}
					}()

					prometheus := prometheusHandler.NewPrometheusHandler()
					prometheus.ExposeMetrics()

					fmt.Println("Agent is running!")

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
