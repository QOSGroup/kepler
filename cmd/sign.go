package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/common"
)

var SignCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign certificate",
	Long:  `Sign certificate`,
	Run:   Sign,
}

func Sign(cmd *cobra.Command, args []string) {
	// Load CSR
	csrBytes := common.MustReadFile(csrFile)

	err := cdc.UnmarshalBinaryBare(csrBytes, &csr)
	if err != nil {
		common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
	}

	// Load PrivKey
	var privKey ed25519.PrivKeyEd25519
	priKeyBytes := common.MustReadFile(privateKeyFile)
	err = cdc.UnmarshalBinaryBare(priKeyBytes, &privKey)
	if err != nil {
		common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
	}

	// Sign CSR
	crt.CSR = csr
	crt.Signature, err = privKey.Sign(csr.Bytes())
	if err != nil {
		common.Exit(fmt.Sprintf("privKey.Sign failed: %v", err))
	}

	common.MustWriteFile(crtFile, crt.Bytes(), 0644)
}

func init() {
	RootCmd.AddCommand(SignCmd)

	ReqCmd.PersistentFlags().StringVar(&csrFile, "in-sign-req", "my.csr", "certificate signing request filename")
	ReqCmd.PersistentFlags().StringVar(&privateKeyFile, "in-key-pri", "key.pri", "private key")
	ReqCmd.PersistentFlags().StringVar(&crtFile, "out-signed-ca", "my.crt", "certificate signed")

}
