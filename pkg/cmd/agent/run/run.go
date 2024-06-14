package run

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func NewCmdRun() *cobra.Command {
	cmd :=
		&cobra.Command{
			Use:    "run",
			Hidden: true,
			Run: func(cmd *cobra.Command, args []string) {
				if authMode := viper.Get("authMode"); authMode == "kube-config" {
					fmt.Println("Starting agent...", args)
					connectToCluster()
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

func connectToCluster() {
	kubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range pods.Items {
		fmt.Println("Pod: ", pod.Name)
	}
}
