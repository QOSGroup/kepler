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
	// TODO: Compatible with the openssl
	CN string `json:"cn"`
}

type Serialization interface {
	Json(cdc *amino.Codec) []byte
	Bytes(cdc *amino.Codec) []byte
}

var _ Serialization = CertificateSigningRequest{}

type CertificateSigningRequest struct {
	Subj      Subject               `json:"subj"`
	IsCa      bool                  `json:"is_ca"`
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
	Subj      Subject               `json:"subj"`
	PublicKey ed25519.PubKeyEd25519 `json:"public_key"`
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

func (crt Certificate) QscName() string {
	return crt.CSR.Subj.CN
}

func (crt Certificate) IsBanker() bool {
	return crt.CSR.IsBanker
}

func (crt Certificate) PublicKey() ed25519.PubKeyEd25519 {
	return crt.CSR.PublicKey
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
