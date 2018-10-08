package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/common"
)

var SignCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign certificate",
	Long:  `Sign certificate`,
	Run:   sign,
}

func sign(cmd *cobra.Command, args []string) {
	if verbose {
		if csrFile != "" {
			fmt.Println("csr file:", csrFile)
		}

		if crtFile != "" {
			fmt.Println("crt file:", crtFile)
		}

		if privateKeyFile != "" {
			fmt.Println("private key:", privateKeyFile)
		}
	}

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

	// Load PubKey
	var pubKey ed25519.PubKeyEd25519
	pubKeyBytes := common.MustReadFile(publicKeyFile)
	err = cdc.UnmarshalBinaryBare(pubKeyBytes, &pubKey)
	if err != nil {
		common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
	}

	// Sign CSR
	csr.Issuer = pubKey
	csr.NotBefore = time.Now()
	csr.NotAfter = time.Now().AddDate(1, 0, 0)
	crt.CSR = csr
	crt.Signature, err = privKey.Sign(csr.Bytes(cdc))
	if err != nil {
		common.Exit(fmt.Sprintf("privKey.Sign failed: %v", err))
	}

	common.MustWriteFile(crtFile, crt.Bytes(cdc), 0644)
}

func init() {
	RootCmd.AddCommand(SignCmd)

	SignCmd.PersistentFlags().StringVar(&csrFile, "in-sign-req", "root.csr", "certificate signing request filename")
	SignCmd.PersistentFlags().StringVar(&crtFile, "out-signed-ca", "root.crt", "certificate signed")
	SignCmd.PersistentFlags().StringVar(&privateKeyFile, "in-key-pri", "key.pri", "private key")
	SignCmd.PersistentFlags().StringVar(&publicKeyFile, "in-key-pub", "key.pub", "public key")

}
