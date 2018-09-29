
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// byeCmd represents the bye command
var SignCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign certificate",
	Long:  `Sign certificate`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sign signature")
	},
}

func init() {
	RootCmd.AddCommand(SignCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// byeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// byeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
