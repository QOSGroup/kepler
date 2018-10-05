package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/common"
)

var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "display csr or crt contents",
	Long:  `display csr or crt contents`,
	Run:   Show,
}

func Show(cmd *cobra.Command, args []string) {
	if csrFile == "" || crtFile == "" {
		fmt.Println("please set csr file or crt file")
		return
	}

	// READ CSR
	if csrFile != "" {
		csrBytes := common.MustReadFile(csrFile)

		err := cdc.UnmarshalBinaryBare(csrBytes, &csr)
		if err != nil {
			common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
		}

		fmt.Println("csrFile:", string(csr.Json()))
	}

	// READ CRT
	if crtFile != "" {
		crtBytes := common.MustReadFile(crtFile)

		err := cdc.UnmarshalBinaryBare(crtBytes, &crt)
		if err != nil {
			common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
		}

		fmt.Println("crtFile:", string(crt.Json()))
	}

}

func init() {
	RootCmd.AddCommand(ShowCmd)

	ReqCmd.PersistentFlags().StringVar(&csrFile, "in-csr", "", "certificate signing request filename")
	ReqCmd.PersistentFlags().StringVar(&crtFile, "in-cst", "", "certificate signed")

}
