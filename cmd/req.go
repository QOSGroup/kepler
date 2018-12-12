package cmd

import (
	"fmt"
	"github.com/QOSGroup/kepler/cert"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/common"
)

var commonSubject cert.CommonSubject
var qscSubject cert.QSCSubject
var bankerFile string
var qcpSubject cert.QCPSubject

func ReqCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "req",
		Short: "QSC Certificate Signing Request",
		Long:  `Certificate Signing Request`,
		Run: func(cmd *cobra.Command, args []string) {
			if verbose {
				if publicKeyFile != "" {
					fmt.Println("public key:", publicKeyFile)
				}
				if csrFile != "" {
					fmt.Println("csr File:", csrFile)
				}
			}

			publicBytes := common.MustReadFile(publicKeyFile)

			err := cdc.UnmarshalJSON(publicBytes, &csr.PublicKey)
			if err != nil {
				common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
			}

			csr.Subj = commonSubject

			common.MustWriteFile(csrFile, MustMarshalBinaryBare(csr), 0644)
		},
	}

	cmd.Flags().BoolVar(&csr.IsCa, "is-ca", false, "Is it root certificate")
	cmd.PersistentFlags().StringVar(&commonSubject.CN, "cn", "CA", "Common name")
	cmd.PersistentFlags().StringVar(&publicKeyFile, "in-public-key", "key.pub", "public key filename")
	cmd.PersistentFlags().StringVar(&csrFile, "out-sign-req", "root.csr", "certificate signing request filename")

	return cmd
}

func QSCReqCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "req-qsc",
		Short: "QSC Certificate Signing Request",
		Long:  `Certificate Signing Request`,
		Run: func(cmd *cobra.Command, args []string) {
			if verbose {
				if publicKeyFile != "" {
					fmt.Println("public key:", publicKeyFile)
				}
				if csrFile != "" {
					fmt.Println("csr File:", csrFile)
				}
				if bankerFile != "" {
					fmt.Println("banker File:", bankerFile)
				}
			}

			publicBytes := common.MustReadFile(publicKeyFile)

			err := cdc.UnmarshalJSON(publicBytes, &csr.PublicKey)
			if err != nil {
				common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
			}

			bankerBytes := common.MustReadFile(bankerFile)

			err = cdc.UnmarshalJSON(bankerBytes, &qscSubject.Banker)
			if err != nil {
				common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
			}

			csr.Subj = qscSubject

			common.MustWriteFile(csrFile, MustMarshalBinaryBare(csr), 0644)
		},
	}

	cmd.Flags().BoolVar(&csr.IsCa, "is-ca", false, "Is it root certificate")
	cmd.PersistentFlags().StringVar(&qscSubject.Name, "name", "", "QSC name")
	cmd.PersistentFlags().StringVar(&bankerFile, "banker", "", "banker public key filename")
	cmd.PersistentFlags().StringVar(&csr.ChainId, "chain-id", "", "ChainId where CA can be used")
	cmd.PersistentFlags().StringVar(&publicKeyFile, "in-public-key", "key.pub", "public key filename")
	cmd.PersistentFlags().StringVar(&csrFile, "out-sign-req", "root.csr", "certificate signing request filename")

	return cmd
}

func QCPReqCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "req-qcp",
		Short: "QCP Certificate Signing Request",
		Long:  `Certificate Signing Request`,
		Run: func(cmd *cobra.Command, args []string) {
			if verbose {
				if publicKeyFile != "" {
					fmt.Println("public key:", publicKeyFile)
				}
				if csrFile != "" {
					fmt.Println("csr File:", csrFile)
				}
			}

			publicBytes := common.MustReadFile(publicKeyFile)

			err := cdc.UnmarshalJSON(publicBytes, &csr.PublicKey)
			if err != nil {
				common.Exit(fmt.Sprintf("cdc.UnmarshalBinaryBare failed: %v", err))
			}

			csr.Subj = qcpSubject

			common.MustWriteFile(csrFile, MustMarshalBinaryBare(csr), 0644)
		},
	}

	cmd.Flags().BoolVar(&csr.IsCa, "is-ca", false, "Is it root certificate")
	cmd.PersistentFlags().StringVar(&qcpSubject.QCPChain, "qcp-chain", "", "QCP ChainId")
	cmd.PersistentFlags().StringVar(&csr.ChainId, "chain-id", "", "ChainId where CA can be used")
	cmd.PersistentFlags().StringVar(&publicKeyFile, "in-public-key", "key.pub", "public key filename")
	cmd.PersistentFlags().StringVar(&csrFile, "out-sign-req", "root.csr", "certificate signing request filename")

	return cmd
}

func init() {
	RootCmd.AddCommand(ReqCmd())
	RootCmd.AddCommand(QSCReqCmd())
	RootCmd.AddCommand(QCPReqCmd())
}
