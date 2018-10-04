package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/common"
)

var csr Csr

// byeCmd represents the bye command
var ReqCmd *cobra.Command = &cobra.Command{
	Use:   "req",
	Short: "Certificate Signing Request",
	Long:  `Certificate Signing Request`,
	Run:   Req,
}

func Req(cmd *cobra.Command, args []string) {
	fmt.Println("public key:", publicKeyFile)

	publicBytes := common.MustReadFile(publicKeyFile)
	fmt.Println("public bytes:", publicBytes)

	err := cdc.UnmarshalBinaryBare(publicBytes, &csr.PublicKey)
	if err != nil {
		common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
	}

	common.MustWriteFile(csrFile, csr.ToJson(), 0644)
}

func init() {
	RootCmd.AddCommand(ReqCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// byeCmd.PersistentFlags().String("foo", "", "A help for foo")
	ReqCmd.PersistentFlags().Int8Var(&csr.Version, "version", 1, "Certificate version")
	ReqCmd.Flags().BoolVar(&csr.CA, "ca", false, "Is it root certificate")
	ReqCmd.PersistentFlags().StringVar(&csr.CN, "cn", "QSC", "Common name")
	ReqCmd.PersistentFlags().StringVar(&publicKeyFile, "in-public-key", "key.pub", "public key filename")
	ReqCmd.PersistentFlags().StringVar(&csrFile, "out-sign-req", "root.csr", "certificate signing request filename")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// byeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
