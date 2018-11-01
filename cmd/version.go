package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "0.4.0"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}
