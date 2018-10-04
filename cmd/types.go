package cmd

import "github.com/tendermint/tendermint/crypto/ed25519"

type Csr struct {
	version   int8                  `json:"version"`
	ca        bool                  `json:"ca"`
	cn        string                `json:"cn"`
	publicKey ed25519.PubKeyEd25519 `json:"public_key"`
}

type Crt struct {
	csr       Csr                         `json:"csr"`
	signature [ed25519.SignatureSize]byte `json:"signature"`
}
