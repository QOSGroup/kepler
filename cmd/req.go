package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/common"
)

var ReqCmd *cobra.Command = &cobra.Command{
	Use:   "req",
	Short: "Certificate Signing Request",
	Long:  `Certificate Signing Request`,
	Run:   req,
}

func req(cmd *cobra.Command, args []string) {
	if verbose {
		if publicKeyFile != "" {
			fmt.Println("public key:", publicKeyFile)
		}
		if csrFile != "" {
			fmt.Println("csr File:", csrFile)
		}
	}

	publicBytes := common.MustReadFile(publicKeyFile)

	err := cdc.UnmarshalBinaryBare(publicBytes, &csr.PublicKey)
	if err != nil {
		common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
	}

	common.MustWriteFile(csrFile, csr.Bytes(), 0644)
}

func init() {
	RootCmd.AddCommand(ReqCmd)

	ReqCmd.PersistentFlags().Int8Var(&csr.Version, "version", 1, "Certificate version")
	ReqCmd.Flags().BoolVar(&csr.CA, "ca", false, "Is it root certificate")
	ReqCmd.PersistentFlags().StringVar(&csr.CN, "cn", "QSC", "Common name")
	ReqCmd.PersistentFlags().BoolVar(&csr.Banker, "banker", false, "Additional attributes")
	ReqCmd.PersistentFlags().StringVar(&publicKeyFile, "in-public-key", "key.pub", "public key filename")
	ReqCmd.PersistentFlags().StringVar(&csrFile, "out-sign-req", "my.csr", "certificate signing request filename")
}
