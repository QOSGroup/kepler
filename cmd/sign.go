package cmd

import (
	"fmt"
	"github.com/QOSGroup/kepler/cert"
	"github.com/tendermint/go-amino"
	"time"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/common"
)

func SignCmd(cdc *amino.Codec) *cobra.Command {

	var csrFile string
	var crtFile string
	var privateKeyFile string
	var publicKeyFile string

	cmd := &cobra.Command{
		Use:   "sign",
		Short: "Sign certificate",
		Long:  `Sign certificate`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if verbose {
				if csrFile != "" {
					fmt.Println("csr file:", csrFile)
				}

				if crtFile != "" {
					fmt.Println("crt file:", crtFile)
				}

				if privateKeyFile != "" {
					fmt.Println("private key:", privateKeyFile)
				}
			}

			csr := cert.CertificateSigningRequest{}

			// Load CSR
			csrBytes := common.MustReadFile(csrFile)

			err := cdc.UnmarshalJSON(csrBytes, &csr)
			if err != nil {
				return err
			}

			// Load PrivKey
			var privKey ed25519.PrivKeyEd25519
			priKeyBytes := common.MustReadFile(privateKeyFile)
			err = cdc.UnmarshalJSON(priKeyBytes, &privKey)
			if err != nil {
				return err
			}

			// Load PubKey
			var pubKey ed25519.PubKeyEd25519
			pubKeyBytes := common.MustReadFile(publicKeyFile)
			err = cdc.UnmarshalJSON(pubKeyBytes, &pubKey)
			if err != nil {
				return err
			}

			// Sign CSR
			crt := cert.Certificate{}
			crt.CA.PublicKey = pubKey
			csr.NotBefore = time.Now()
			csr.NotAfter = time.Now().AddDate(1, 0, 0)
			crt.CSR = csr
			crt.Signature, err = privKey.Sign(cert.MustMarshalJson(csr))
			if err != nil {
				return err
			}

			common.MustWriteFile(crtFile, cert.MustMarshalJson(crt), 0644)

			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&csrFile, "in-sign-req", "root.csr", "certificate signing request filename")
	cmd.PersistentFlags().StringVar(&crtFile, "out-signed-ca", "root.crt", "certificate signed")
	cmd.PersistentFlags().StringVar(&privateKeyFile, "in-key-pri", "key.pri", "private key")
	cmd.PersistentFlags().StringVar(&publicKeyFile, "in-key-pub", "key.pub", "public key")

	return cmd
}
