package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/common"
)

var MintCmd = &cobra.Command{
	Use:   "mint",
	Short: "find smaller address",
	Long:  `find smaller address`,
	Run:   mint,
}

func mint(cmd *cobra.Command, args []string) {
	if verbose {
		fmt.Println("private key file:", privateKeyFile)
		fmt.Println("public key file:", publicKeyFile)
	}

	min_address := "00000000d697a7be725d9b262917bc6844cc5211"

	for a := 0; a < 10000; a++ {
		privKey := ed25519.GenPrivKey()
		pubKey := privKey.PubKey()
		address := hex.EncodeToString(pubKey.Address())

		if min_address > address {
			min_address = address

			common.EnsureDir(address, 0755)
			priKeyBytes, err := cdc.MarshalJSON(privKey)
			if err != nil {
				common.Exit(fmt.Sprintf("cdc.MarshalJSON failed: %v", err))
			}
			pubKeyBytes, err := cdc.MarshalJSON(pubKey.(ed25519.PubKeyEd25519))
			if err != nil {
				common.Exit(fmt.Sprintf("cdc.MarshalJSON failed: %v", err))
			}
			common.MustWriteFile(address+"/"+privateKeyFile, priKeyBytes, 0644)
			common.MustWriteFile(address+"/"+publicKeyFile, pubKeyBytes, 0644)
			fmt.Println(address, a)
		}
	}
}

func init() {
	RootCmd.AddCommand(MintCmd)

	MintCmd.PersistentFlags().StringVar(&privateKeyFile, "out-private-key", "key.pri", "private key file")
	MintCmd.PersistentFlags().StringVar(&publicKeyFile, "out-public-key", "key.pub", "public key file")
}
