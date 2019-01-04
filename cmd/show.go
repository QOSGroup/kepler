package cmd

import (
	"fmt"
	"github.com/QOSGroup/kepler/cert"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/libs/common"
)

func ShowCmd(cdc *amino.Codec) *cobra.Command {

	var csrFile string
	var crtFile string

	cmd := &cobra.Command{
		Use:   "show",
		Short: "display csr or crt contents",
		Long:  `display csr or crt contents`,
		RunE: func(cmd *cobra.Command, args []string) error {

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
				csr := cert.CertificateSigningRequest{}

				csrBytes := common.MustReadFile(csrFile)

				err := cdc.UnmarshalJSON(csrBytes, &csr)
				if err != nil {
					return err
				}

				fmt.Println("csrFile:", string(cert.MustMarshalJson(csr)))
			}

			// READ CRT
			if crtFile != "" {
				crt := cert.Certificate{}

				crtBytes := common.MustReadFile(crtFile)

				err := cdc.UnmarshalJSON(crtBytes, &crt)
				if err != nil {
					return err
				}

				fmt.Println("crtFile:", string(cert.MustMarshalJson(crt)))
			}

			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&csrFile, "in-csr-file", "", "certificate signing request")
	cmd.PersistentFlags().StringVar(&crtFile, "in-crt-file", "", "certificate signed")

	cmd.MarkFlagRequired(csrFile)
	cmd.MarkFlagRequired(crtFile)

	return cmd
}
