package cert

import (
	"time"

	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

const (
	CsrAminoRoute       = "certificate/csr"
	CrtAminoRoute       = "certificate/crt"
	TrustCrtsAminoRoute = "certificate/trustCrts"
)

type Subject struct {
	CN string `json:"cn"`
}

type Serialization interface {
	Json(cdc *amino.Codec) []byte
	Bytes(cdc *amino.Codec) []byte
}

var _ Serialization = CertificateSigningRequest{}

type CertificateSigningRequest struct {
	Version   int8                  `json:"version"`
	IsCa      bool                  `json:"is_ca"`
	Subj      Subject               `json:"subj"`
	IsBanker  bool                  `json:"is_banker"`
	NotBefore time.Time             `json:"not_before"`
	NotAfter  time.Time             `json:"not_after"`
	PublicKey ed25519.PubKeyEd25519 `json:"public_key"`
}

func (csr CertificateSigningRequest) Json(cdc *amino.Codec) []byte {
	bz, err := cdc.MarshalJSON(csr)
	if err != nil {
		panic(err)
	}
	return bz
}

func (csr CertificateSigningRequest) Bytes(cdc *amino.Codec) []byte {
	bz, err := cdc.MarshalBinaryBare(csr)
	if err != nil {
		panic(err)
	}
	return bz
}

var _ Serialization = Certificate{}

type Issuer struct {
	PublicKey ed25519.PubKeyEd25519 `json:"public_key"`
	Subj      Subject               `json:"subj"`
}

type Certificate struct {
	CSR       CertificateSigningRequest `json:"csr"`
	CA        Issuer                    `json:"ca"`
	Signature []byte                    `json:"signature"`
}

func (crt Certificate) Json(cdc *amino.Codec) []byte {
	bz, err := cdc.MarshalJSON(crt)
	if err != nil {
		panic(err)
	}
	return bz
}

func (crt Certificate) Bytes(cdc *amino.Codec) []byte {
	bz, err := cdc.MarshalBinaryBare(crt)
	if err != nil {
		panic(err)
	}
	return bz
}

var _ Serialization = TrustCrts{}

type TrustCrts struct {
	PublicKeys []ed25519.PubKeyEd25519 `json:"public_keys"`
}

func (certs TrustCrts) Json(cdc *amino.Codec) []byte {
	bz, err := cdc.MarshalJSON(certs)
	if err != nil {
		panic(err)
	}
	return bz
}

func (certs TrustCrts) Bytes(cdc *amino.Codec) []byte {
	bz, err := cdc.MarshalBinaryBare(certs)
	if err != nil {
		panic(err)
	}
	return bz
}
