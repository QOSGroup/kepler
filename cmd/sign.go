package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/common"
)

// byeCmd represents the bye command
var SignCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign certificate",
	Long:  `Sign certificate`,
	Run:   Sign,
}

func Sign(cmd *cobra.Command, args []string) {
	csrBytes := common.MustReadFile(csrFile)
	fmt.Println("csr bytes:", csrBytes)

	err := cdc.UnmarshalJSON(csrBytes, &csr)
	if err != nil {
		common.Exit(fmt.Sprintf("cdc.UnmarshalJSON failed: %v", err))
	}
	fmt.Println("csr:", csr)

	priKeyBytes := common.MustReadFile(privateKeyFile)
	fmt.Println("key.pri bytes:", priKeyBytes)

	var privKey ed25519.PrivKeyEd25519
	err = cdc.UnmarshalBinaryBare(priKeyBytes, &privKey)
	if err != nil {
		common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
	}
	fmt.Println("privKey:", privKey)

	crt.Signature, err = privKey.Sign(csrBytes)
	if err != nil {
		common.Exit(fmt.Sprintf("privKey.Sign failed: %v", err))
	}

	common.MustWriteFile(crtFile, crt.ToJson(), 0644)

}
func init() {
	RootCmd.AddCommand(SignCmd)

	ReqCmd.PersistentFlags().StringVar(&csrFile, "in-sign-req", "my.csr", "certificate signing request filename")
	ReqCmd.PersistentFlags().StringVar(&privateKeyFile, "in-key-pri", "key.pri", "private key")
	ReqCmd.PersistentFlags().StringVar(&crtFile, "out-signed-ca", "my.crt", "certificate signed")

}
