package cmd

import (
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var (
	privateKeyFile string
	publicKeyFile  string
	csrFile        string
	crtFile        string
	csr            CertificateSigningRequest
	crt            Certificate
	verbose        bool
)

type Serialization interface {
	Json() []byte
	Bytes() []byte
}

type CertificateSigningRequest struct {
	Version   int8                  `json:"version"`
	CA        bool                  `json:"ca"`
	CN        string                `json:"cn"`
	Banker    bool                  `json:"banker"`
	PublicKey ed25519.PubKeyEd25519 `json:"public_key"`
}

func (csr CertificateSigningRequest) Json() []byte {
	bz, err := cdc.MarshalJSON(csr)
	if err != nil {
		panic(err)
	}
	return bz
}

func (csr CertificateSigningRequest) Bytes() []byte {
	bz, err := cdc.MarshalBinaryBare(csr)
	if err != nil {
		panic(err)
	}
	return bz
}

type Certificate struct {
	CSR       CertificateSigningRequest `json:"csr"`
	Signature []byte                    `json:"signature"`
}

func (crt Certificate) Json() []byte {
	bz, err := cdc.MarshalJSON(crt)
	if err != nil {
		panic(err)
	}
	return bz
}

func (crt Certificate) Bytes() []byte {
	bz, err := cdc.MarshalBinaryBare(crt)
	if err != nil {
		panic(err)
	}
	return bz
}

const (
	CsrAminoRoute = "certificate/csr"
	CrtAminoRoute = "certificate/crt"
)

var cdc = amino.NewCodec()

func init() {
	cdc.RegisterInterface((*Serialization)(nil), nil)
	cdc.RegisterConcrete(CertificateSigningRequest{},
		CsrAminoRoute, nil)
	cdc.RegisterConcrete(Certificate{},
		CrtAminoRoute, nil)

	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(ed25519.PubKeyEd25519{},
		ed25519.PubKeyAminoRoute, nil)

	cdc.RegisterInterface((*crypto.PrivKey)(nil), nil)
	cdc.RegisterConcrete(ed25519.PrivKeyEd25519{},
		ed25519.PrivKeyAminoRoute, nil)
}
