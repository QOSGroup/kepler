package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tendermint/tendermint/version"
)

var (
	Version = "0.1.0"
)

// VersionCmd ...
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}