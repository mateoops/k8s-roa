package root

import (
	"fmt"
	"os"

	runCmd "github.com/mateoops/k8s-roa/pkg/cmd/agent/run"
	versionCmd "github.com/mateoops/k8s-roa/pkg/cmd/agent/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCmdRoot() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "roa-agent",
		Short: "ROA Agent ",
		Long:  "Kubernetes Resource Optimization Advisor agent",
	}

	initCmdRoot(cmd)
	cobra.OnInitialize(initConfig)

	return cmd, nil
}

func initCmdRoot(cmd *cobra.Command) {

	cmd.PersistentFlags().Bool("help", false, "Show help for command")

	cmd.AddCommand(versionCmd.NewCmdVersion())
	cmd.AddCommand(runCmd.NewCmdRun())
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.SetConfigType("yaml")
	viper.SetConfigName("roa-agent")
	viper.AddConfigPath(home)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		viper.AddConfigPath("./configs/")
		viper.ReadInConfig()
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
