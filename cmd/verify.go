package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/common"
)

// byeCmd represents the bye command
var VerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify certificate signature",
	Long:  `verify certificate signature`,
	Run:   Verify,
}

func Verify(cmd *cobra.Command, args []string) {
	crtBytes := common.MustReadFile(crtFile)
	fmt.Println("csr bytes:", crtBytes)

	err := cdc.UnmarshalJSON(crtBytes, &crt)
	if err != nil {
		common.Exit(fmt.Sprintf("cdc.UnmarshalJSON failed: %v", err))
	}
	fmt.Println("crt:", crt)

	priKeyBytes := common.MustReadFile(privateKeyFile)
	fmt.Println("key.pri bytes:", priKeyBytes)

	pubKeyBytes := common.MustReadFile(publicKeyFile)
	fmt.Println("key.pub bytes:", pubKeyBytes)

	var pubKey ed25519.PubKeyEd25519
	err = cdc.UnmarshalBinaryBare(pubKeyBytes, &pubKey)
	if err != nil {
		common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
	}
	fmt.Println("privKey:", pubKey)

	ok := pubKey.VerifyBytes(crt.CSR.ToJson(), crt.Signature)
	fmt.Println("verify result:", ok)
}

func init() {
	RootCmd.AddCommand(VerifyCmd)

	ReqCmd.PersistentFlags().StringVar(&publicKeyFile, "in-key-pub", "key.pub", "public key")
	ReqCmd.PersistentFlags().StringVar(&crtFile, "in-signed-ca", "my.crt", "certificate signed")

}
