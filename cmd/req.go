package cmd

import (
	"fmt"
	"github.com/QOSGroup/kepler/cert"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/libs/common"
)

func ReqCmd(cdc *amino.Codec) *cobra.Command {
	var publicKeyFile string
	var csrFile string
	var isCa bool
	var cn string

	cmd := &cobra.Command{
		Use:   "req",
		Short: "Certificate Signing Request",
		Long:  `Certificate Signing Request`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if verbose {
				if publicKeyFile != "" {
					fmt.Println("public key:", publicKeyFile)
				}
				if csrFile != "" {
					fmt.Println("csr File:", csrFile)
				}
			}

			publicBytes := common.MustReadFile(publicKeyFile)

			csr := cert.CertificateSigningRequest{}

			err := cdc.UnmarshalJSON(publicBytes, &csr.PublicKey)
			if err != nil {
				return err
			}

			commonSubject := cert.CommonSubject{}
			commonSubject.CN = cn

			csr.IsCa = isCa
			csr.Subj = commonSubject

			common.MustWriteFile(csrFile, cert.MustMarshalJson(csr), 0644)

			return nil
		},
	}

	cmd.PersistentFlags().BoolVar(&isCa, "is-ca", false, "Is it root certificate")
	cmd.PersistentFlags().StringVar(&cn, "cn", "CA", "Common name")
	cmd.PersistentFlags().StringVar(&publicKeyFile, "in-public-key", "key.pub", "public key filename")
	cmd.PersistentFlags().StringVar(&csrFile, "out-sign-req", "root.csr", "certificate signing request filename")

	return cmd
}

func QSCReqCmd(cdc *amino.Codec) *cobra.Command {
	var publicKeyFile string
	var csrFile string
	var isCa bool
	var chainId string
	var name string
	var banker string

	cmd := &cobra.Command{
		Use:   "req-qsc",
		Short: "QSC Certificate Signing Request",
		Long:  `QSC Certificate Signing Request`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if verbose {
				if publicKeyFile != "" {
					fmt.Println("public key:", publicKeyFile)
				}
				if csrFile != "" {
					fmt.Println("csr File:", csrFile)
				}
				if banker != "" {
					fmt.Println("banker File:", banker)
				}
			}

			publicBytes := common.MustReadFile(publicKeyFile)

			csr := cert.CertificateSigningRequest{}

			err := cdc.UnmarshalJSON(publicBytes, &csr.PublicKey)
			if err != nil {
				return err
			}

			bankerBytes := common.MustReadFile(banker)

			qscSubject := cert.QSCSubject{}
			qscSubject.Name = name
			qscSubject.ChainId = chainId

			err = cdc.UnmarshalJSON(bankerBytes, &qscSubject.Banker)
			if err != nil {
				return err
			}

			csr.IsCa = isCa
			csr.Subj = qscSubject

			common.MustWriteFile(csrFile, cert.MustMarshalJson(csr), 0644)

			return nil
		},
	}

	cmd.PersistentFlags().BoolVar(&isCa, "is-ca", false, "Is it root certificate")
	cmd.PersistentFlags().StringVar(&name, "name", "", "QSC name")
	cmd.PersistentFlags().StringVar(&banker, "banker", "banker.pub", "banker public key filename")
	cmd.PersistentFlags().StringVar(&chainId, "chain-id", "", "ChainId where CA can be used")
	cmd.PersistentFlags().StringVar(&publicKeyFile, "in-public-key", "key.pub", "public key filename")
	cmd.PersistentFlags().StringVar(&csrFile, "out-sign-req", "root.csr", "certificate signing request filename")

	cmd.MarkFlagRequired(name)

	return cmd
}

func QCPReqCmd(cdc *amino.Codec) *cobra.Command {
	var publicKeyFile string
	var csrFile string
	var isCa bool
	var chainId string
	var qcpChain string

	cmd := &cobra.Command{
		Use:   "req-qcp",
		Short: "QCP Certificate Signing Request",
		Long:  `QCP Certificate Signing Request`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if verbose {
				if publicKeyFile != "" {
					fmt.Println("public key:", publicKeyFile)
				}
				if csrFile != "" {
					fmt.Println("csr File:", csrFile)
				}
			}

			publicBytes := common.MustReadFile(publicKeyFile)

			csr := cert.CertificateSigningRequest{}

			err := cdc.UnmarshalJSON(publicBytes, &csr.PublicKey)
			if err != nil {
				return err
			}

			qcpSubject := cert.QCPSubject{}
			qcpSubject.ChainId = chainId
			qcpSubject.QCPChain = qcpChain

			csr.IsCa = isCa
			csr.Subj = qcpSubject

			common.MustWriteFile(csrFile, cert.MustMarshalJson(csr), 0644)

			return nil
		},
	}

	cmd.PersistentFlags().BoolVar(&isCa, "is-ca", false, "Is it root certificate")
	cmd.PersistentFlags().StringVar(&qcpChain, "qcp-chain", "", "QCP ChainId")
	cmd.PersistentFlags().StringVar(&chainId, "chain-id", "", "ChainId where CA can be used")
	cmd.PersistentFlags().StringVar(&publicKeyFile, "in-public-key", "key.pub", "public key filename")
	cmd.PersistentFlags().StringVar(&csrFile, "out-sign-req", "root.csr", "certificate signing request filename")

	cmd.MarkFlagRequired(qcpChain)
	cmd.MarkFlagRequired(chainId)

	return cmd
}
