package cmd

import (
	"fmt"
	"os"

	"github.com/QOSGroup/kepler/cert"
	"github.com/spf13/cobra"
)

var (
	privateKeyFile string
	publicKeyFile  string
	csrFile        string
	crtFile        string
	trustCrtsFile  string
	csr            cert.CertificateSigningRequest
	crt            cert.Certificate
	verbose        bool
)

var RootCmd = &cobra.Command{
	Use:   "kepler",
	Short: "An example of cobra",
	Long: `This application shows how to create modern CLI 
applications in go using Cobra CLI library`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
