package main

import (
	"fmt"
	"os"

	"github.com/mateoops/k8s-roa/pkg/cmd/agent/root"
)

func main() {

	rootCmd, err := root.NewCmdRoot()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	rootCmd.Execute()
}
