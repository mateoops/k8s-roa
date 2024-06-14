package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdVersion() *cobra.Command {
	cmd :=
		&cobra.Command{
			Use:    "version",
			Hidden: true,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Version: TODO")
			},
		}
	return cmd
}
