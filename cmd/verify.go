package cmd

import (
	"fmt"
	"github.com/QOSGroup/kepler/cert"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/libs/common"
)

func VerifyCmd(cdc *amino.Codec) *cobra.Command {

	var flagCrtFile string
	var flagTrustCrtsFile string

	cmd := &cobra.Command{Use: "verify",
		Short: "verify certificate signature",
		Long:  `verify certificate signature`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if verbose {
				if flagCrtFile != "" {
					fmt.Println("crt File:", flagCrtFile)
				}
				if flagTrustCrtsFile != "" {
					fmt.Println("trust crts File:", flagTrustCrtsFile)
				}
			}

			var trustCrts cert.TrustCrts
			trustCrtsBytes := common.MustReadFile(flagTrustCrtsFile)
			err := cdc.UnmarshalJSON(trustCrtsBytes, &trustCrts)
			if err != nil {
				return err
			}

			crtBytes := common.MustReadFile(flagCrtFile)
			crt := cert.Certificate{}
			err = cdc.UnmarshalJSON(crtBytes, &crt)
			if err != nil {
				return err
			}

			// Check issuer
			ok := cert.VerityCrt(trustCrts.PublicKeys, crt)

			fmt.Println(flagCrtFile, "verify result:", ok)

			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&flagCrtFile, "in-signed-ca", "root.crt", "certificate signed")
	cmd.PersistentFlags().StringVar(&flagTrustCrtsFile, "in-trust-crts", "trust.crts", "trust certificate list")

	return cmd
}
