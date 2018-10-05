package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/common"
)

var VerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify certificate signature",
	Long:  `verify certificate signature`,
	Run:   verify,
}

func verify(cmd *cobra.Command, args []string) {
	if verbose {
		if crtFile != "" {
			fmt.Println("crt File:", crtFile)
		}
	}
	crtBytes := common.MustReadFile(crtFile)

	err := cdc.UnmarshalBinaryBare(crtBytes, &crt)
	if err != nil {
		common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
	}

	ok := crt.Issuer.VerifyBytes(crt.CSR.Bytes(cdc), crt.Signature)
	fmt.Println("verify result:", ok)
}

func init() {
	RootCmd.AddCommand(VerifyCmd)

	ReqCmd.PersistentFlags().StringVar(&crtFile, "in-signed-ca", "my.crt", "certificate signed")
}
