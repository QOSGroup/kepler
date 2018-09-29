package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/common"
)

var privateKeyFile string
var publicKeyFile string

// helloCmd represents the hello command
var GenkeyCmd = &cobra.Command{
	Use:   "genkey",
	Short: "Generating public/private ed25519 key pair",
	Long:  `Generating public/private ed25519 key pair`,
	Run: Genkey,
}

func Genkey(cmd *cobra.Command, args []string) {
	privKey := ed25519.GenPrivKey()
	pubKey := privKey.PubKey()

	common.MustWriteFile(privateKeyFile, privKey.Bytes(), 0644)
	common.MustWriteFile(publicKeyFile, pubKey.Bytes(), 0644)

	fmt.Println("priv key:", privateKeyFile)
	fmt.Println("pub key:", publicKeyFile)
}

func init() {
	RootCmd.AddCommand(GenkeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	GenkeyCmd.PersistentFlags().StringVar(&privateKeyFile,"out-private-key", "private.key", "private key filename")
	GenkeyCmd.PersistentFlags().StringVar(&publicKeyFile,"out-public-key", "public.key", "public key filename")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
