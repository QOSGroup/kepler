package cmd

import (
	"fmt"
	"os"

	"github.com/QOSGroup/kepler/cert"
	"github.com/spf13/cobra"
)

var (
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "kepler",
	Short: "An example of cobra",
	Long: `This application shows how to create modern CLI 
applications in go using Cobra CLI library`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "verbose output")

	rootCmd.AddCommand(GenKeyCmd(cert.Codec))
	rootCmd.AddCommand(ReqCmd(cert.Codec))
	rootCmd.AddCommand(QSCReqCmd(cert.Codec))
	rootCmd.AddCommand(QCPReqCmd(cert.Codec))
	rootCmd.AddCommand(ShowCmd(cert.Codec))
	rootCmd.AddCommand(SignCmd(cert.Codec))
	rootCmd.AddCommand(TrustCmd(cert.Codec))
	rootCmd.AddCommand(VerifyCmd(cert.Codec))
	rootCmd.AddCommand(VersionCmd)

}
