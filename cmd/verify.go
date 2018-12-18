package cmd

import (
	"fmt"
	"github.com/QOSGroup/kepler/cert"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/common"
)

var VerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify certificate signature",
	Long:  `verify certificate signature`,
	Run:   verify,
}

func verify(cmd *cobra.Command, args []string) {
	if verbose {
		if crtFile != "" {
			fmt.Println("crt File:", crtFile)
		}
		if trustCrtsFile != "" {
			fmt.Println("trust crts File:", trustCrtsFile)
		}
	}

	var trustCrts cert.TrustCrts
	trustCrtsBytes := common.MustReadFile(trustCrtsFile)
	err := cdc.UnmarshalJSON(trustCrtsBytes, &trustCrts)
	if err != nil {
		common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
	}

	crtBytes := common.MustReadFile(crtFile)
	err = cdc.UnmarshalBinaryBare(crtBytes, &crt)
	if err != nil {
		common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
	}

	// Check issuer
	ok := cert.VerityCrt(trustCrts.PublicKeys, crt)

	fmt.Println(crtFile, "verify result:", ok)
}

func init() {
	RootCmd.AddCommand(VerifyCmd)

	VerifyCmd.PersistentFlags().StringVar(&crtFile, "in-signed-ca", "root.crt", "certificate signed")
	VerifyCmd.PersistentFlags().StringVar(&trustCrtsFile, "in-trust-crts", "trust.crts", "trust certificate list")
}
