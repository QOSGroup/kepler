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
	Run:   show,
}

func show(cmd *cobra.Command, args []string) {
	if csrFile == "" || crtFile == "" {
		fmt.Println("no files to show")
		return
	}

	if verbose {
		if csrFile != "" {
			fmt.Println("csr file:", csrFile)
		}

		if crtFile != "" {
			fmt.Println("crt file:", crtFile)
		}
	}

	// READ CSR
	if csrFile != "" {
		csrBytes := common.MustReadFile(csrFile)

		err := cdc.UnmarshalBinaryBare(csrBytes, &csr)
		if err != nil {
			common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
		}

		fmt.Println("csrFile:", string(MustMarshalJson(csr)))
	}

	// READ CRT
	if crtFile != "" {
		crtBytes := common.MustReadFile(crtFile)

		err := cdc.UnmarshalBinaryBare(crtBytes, &crt)
		if err != nil {
			common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
		}

		fmt.Println("crtFile:", string(MustMarshalJson(csr)))
	}

}

func init() {
	RootCmd.AddCommand(ShowCmd)

	ShowCmd.PersistentFlags().StringVar(&csrFile, "in-csr-file", "", "certificate signing request")
	ShowCmd.PersistentFlags().StringVar(&crtFile, "in-crt-file", "", "certificate signed")

}
