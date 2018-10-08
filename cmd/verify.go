package cmd

import (
	"fmt"
	"time"

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
	ok := false
	for _, value := range trustCrts.PublicKeys {
		if value.Equals(crt.CA.PublicKey) {
			ok = crt.CA.PublicKey.VerifyBytes(crt.CSR.Bytes(cdc), crt.Signature)
			break
		}
	}

	// Check timestamp
	now := time.Now().Unix()
	if now <= crt.CSR.NotBefore.Unix() || now >= crt.CSR.NotAfter.Unix() {
		ok = false
	}

	// TODO: add publicKey to trustCrts if crt.CSR.IsCa is true

	fmt.Println(crtFile, "verify result:", ok)
}

func init() {
	RootCmd.AddCommand(VerifyCmd)

	VerifyCmd.PersistentFlags().StringVar(&crtFile, "in-signed-ca", "root.crt", "certificate signed")
	VerifyCmd.PersistentFlags().StringVar(&trustCrtsFile, "in-trust-crts", "trust.crts", "trust certificate list")
}
