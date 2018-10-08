package cmd

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/QOSGroup/kepler/cert"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/common"
)

var TrustCmd *cobra.Command = &cobra.Command{
	Use:   "trust",
	Short: "add trust certificate",
	Long:  `add trust certificate`,
	Run:   trust,
}

func trust(cmd *cobra.Command, args []string) {
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
	err := cdc.UnmarshalBinaryBare(publicBytes, &publicKey)
	if err != nil {
		common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
	}

	var trustCrts cert.TrustCrts
	var trustCrtsBytes []byte
	if common.FileExists(trustCrtsFile) {
		trustCrtsBytes = common.MustReadFile(trustCrtsFile)
		if len(trustCrtsBytes) > 0 {
			err := cdc.UnmarshalJSON(trustCrtsBytes, &trustCrts)
			if err != nil {
				common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
			}
		}
	}
	trustCrts.PublicKeys = append(trustCrts.PublicKeys, publicKey)

	common.MustWriteFile(trustCrtsFile, trustCrts.Json(cdc), 0644)
}

func init() {
	RootCmd.AddCommand(TrustCmd)

	TrustCmd.PersistentFlags().StringVar(&publicKeyFile, "in-public-key", "key.pub", "public key file")
	TrustCmd.PersistentFlags().StringVar(&trustCrtsFile, "out-trust-crts", "trust.crts", "persisted trust certificate")
}
