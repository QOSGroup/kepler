package cmd

import (
	"fmt"
	"github.com/tendermint/go-amino"

	"github.com/QOSGroup/kepler/cert"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/common"
)

func TrustCmd(cdc *amino.Codec) *cobra.Command {

	var publicKeyFile string
	var trustCrtsFile string

	cmd := &cobra.Command{
		Use:   "trust",
		Short: "add trust certificate",
		Long:  `add trust certificate`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if verbose {
				if publicKeyFile != "" {
					fmt.Println("public key:", publicKeyFile)
				}
				if trustCrtsFile != "" {
					fmt.Println("trust crts File:", trustCrtsFile)
				}
			}

			var publicKey ed25519.PubKeyEd25519
			publicBytes := common.MustReadFile(publicKeyFile)
			err := cdc.UnmarshalJSON(publicBytes, &publicKey)
			if err != nil {
				return err
			}

			var trustCrts cert.TrustCrts
			var trustCrtsBytes []byte
			if common.FileExists(trustCrtsFile) {
				trustCrtsBytes = common.MustReadFile(trustCrtsFile)
				if len(trustCrtsBytes) > 0 {
					err := cdc.UnmarshalJSON(trustCrtsBytes, &trustCrts)
					if err != nil {
						return err
					}
				}
			}

			trustCrts.PublicKeys = append(trustCrts.PublicKeys, publicKey)

			common.MustWriteFile(trustCrtsFile, cert.MustMarshalJson(trustCrts), 0644)

			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&publicKeyFile, "in-public-key", "key.pub", "public key file")
	cmd.PersistentFlags().StringVar(&trustCrtsFile, "out-trust-crts", "trust.crts", "persisted trust certificate")

	return cmd
}
