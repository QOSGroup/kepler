package cmd

import (
	"fmt"
	"github.com/tendermint/go-amino"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/common"
)

func GenKeyCmd(cdc *amino.Codec) *cobra.Command {
	var privateKeyFile string
	var publicKeyFile string

	cmd := &cobra.Command{
		Use:   "genkey",
		Short: "Generating public/private ed25519 key pair",
		Long:  `Generating public/private ed25519 key pair`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if verbose {
				fmt.Println("private key file:", privateKeyFile)
				fmt.Println("public key file:", publicKeyFile)
			}

			privKey := ed25519.GenPrivKey()
			priKeyBytes, err := cdc.MarshalJSON(privKey)
			if err != nil {
				return err
			}
			common.MustWriteFile(privateKeyFile, priKeyBytes, 0644)

			pubKey := privKey.PubKey()
			pubKeyBytes, err := cdc.MarshalJSON(pubKey)
			if err != nil {
				return err
			}
			common.MustWriteFile(publicKeyFile, pubKeyBytes, 0644)

			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&privateKeyFile, "out-private-key", "key.pri", "private key file")
	cmd.PersistentFlags().StringVar(&publicKeyFile, "out-public-key", "key.pub", "public key file")

	return cmd
}
