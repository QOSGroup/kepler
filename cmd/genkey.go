package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/common"
)

var GenkeyCmd = &cobra.Command{
	Use:   "genkey",
	Short: "Generating public/private ed25519 key pair",
	Long:  `Generating public/private ed25519 key pair`,
	Run:   genkey,
}

func genkey(cmd *cobra.Command, args []string) {
	if verbose {
		fmt.Println("private key file:", privateKeyFile)
		fmt.Println("public key file:", publicKeyFile)
	}

	privKey := ed25519.GenPrivKey()
	priKeyBytes, err := cdc.MarshalJSON(privKey)
	if err != nil {
		common.Exit(fmt.Sprintf("cdc.MarshalJSON failed: %v", err))
	}
	common.MustWriteFile(privateKeyFile, priKeyBytes, 0644)

	pubKey := privKey.PubKey()
	pubKeyBytes, err := cdc.MarshalJSON(pubKey.(ed25519.PubKeyEd25519))
	if err != nil {
		common.Exit(fmt.Sprintf("cdc.MarshalJSON failed: %v", err))
	}
	common.MustWriteFile(publicKeyFile, pubKeyBytes, 0644)
}

func init() {
	RootCmd.AddCommand(GenkeyCmd)

	GenkeyCmd.PersistentFlags().StringVar(&privateKeyFile, "out-private-key", "key.pri", "private key file")
	GenkeyCmd.PersistentFlags().StringVar(&publicKeyFile, "out-public-key", "key.pub", "public key file")
}
