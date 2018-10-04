package cmd

import (
	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var privateKeyFile string
var publicKeyFile string
var csrFile string

type Serialization interface {
	ToJson() []byte
}

type Csr struct {
	Version   int8                  `json:"version"`
	CA        bool                  `json:"ca"`
	CN        string                `json:"cn"`
	PublicKey ed25519.PubKeyEd25519 `json:"public_key"`
}

func (csr Csr) ToJson() []byte {
	bz, err := cdc.MarshalJSON(csr)
	if err != nil {
		panic(err)
	}
	return bz
}

type Crt struct {
	CSR       Csr                         `json:"csr"`
	Signature [ed25519.SignatureSize]byte `json:"signature"`
}

func (crt Crt) ToJson() []byte {
	bz, err := cdc.MarshalJSON(crt)
	if err != nil {
		panic(err)
	}
	return bz
}

const (
	CsrAminoRoute = "ed25519/csr"
	CrtAminoRoute = "ed25519/crt"
)

var cdc = amino.NewCodec()

func init() {
	cdc.RegisterInterface((*Serialization)(nil), nil)

	cdc.RegisterConcrete(Csr{},
		CsrAminoRoute, nil)
	cdc.RegisterConcrete(Crt{},
		CrtAminoRoute, nil)

	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(ed25519.PubKeyEd25519{},
		ed25519.PubKeyAminoRoute, nil)

	cdc.RegisterInterface((*crypto.PrivKey)(nil), nil)
	cdc.RegisterConcrete(ed25519.PrivKeyEd25519{},
		ed25519.PrivKeyAminoRoute, nil)
}
