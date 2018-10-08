package cmd

import (
	"fmt"
	"os"

	"github.com/kepler/cert"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var (
	privateKeyFile string
	publicKeyFile  string
	csrFile        string
	crtFile        string
	trustCrtsFile  string
	csr            cert.CertificateSigningRequest
	crt            cert.Certificate
	trustCrts      cert.TrustCrts
	verbose        bool
)

var RootCmd = &cobra.Command{
	Use:   "kepler",
	Short: "An example of cobra",
	Long: `This application shows how to create modern CLI 
applications in go using Cobra CLI library`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var cdc = amino.NewCodec()

func init() {
	cdc.RegisterInterface((*cert.Serialization)(nil), nil)
	cdc.RegisterConcrete(cert.CertificateSigningRequest{},
		cert.CsrAminoRoute, nil)
	cdc.RegisterConcrete(cert.Certificate{},
		cert.CrtAminoRoute, nil)
	cdc.RegisterConcrete(cert.TrustCrts{},
		cert.TrustCrtsAminoRoute, nil)

	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(ed25519.PubKeyEd25519{},
		ed25519.PubKeyAminoRoute, nil)

	cdc.RegisterInterface((*crypto.PrivKey)(nil), nil)
	cdc.RegisterConcrete(ed25519.PrivKeyEd25519{},
		ed25519.PrivKeyAminoRoute, nil)

	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
