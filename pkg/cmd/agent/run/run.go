package run

import (
	"fmt"
	"os"

	kubernetesHandler "github.com/mateoops/k8s-roa/handlers/kubernetes"
	prometheusHandler "github.com/mateoops/k8s-roa/handlers/prometheus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCmdRun() *cobra.Command {

	cmd :=
		&cobra.Command{
			Use:    "run",
			Hidden: true,
			Run: func(cmd *cobra.Command, args []string) {
				if authMode := viper.Get("authMode"); authMode == "kube-config" {
					fmt.Println("Starting agent...")

					kubernetes := kubernetesHandler.NewKubernetesHandler()
					nodes := kubernetes.ListNodes()
					for _, node := range nodes {
						fmt.Println("Node: ", node.Name)
						metrics := kubernetes.GetNodeUsageMetrics(node.Name)
						fmt.Println("CPU usage: ", metrics.CpuUsage, "m")
						fmt.Println("Memory usage: ", metrics.MemoryUsage, "MB")
					}

					prometheus := prometheusHandler.NewPrometheusHandler()
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
