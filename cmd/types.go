package cmd

import "github.com/tendermint/tendermint/crypto/ed25519"

type Csr struct {
	Version   int8                  `json:"version"`
	CA        bool                  `json:"ca"`
	CN        string                `json:"cn"`
	PublicKey ed25519.PubKeyEd25519 `json:"public_key"`
}
