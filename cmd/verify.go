package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// byeCmd represents the bye command
var VerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify certificate signature",
	Long:  `verify certificate signature`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("verify signature")
	},
}

func init() {
	RootCmd.AddCommand(VerifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// byeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// byeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
