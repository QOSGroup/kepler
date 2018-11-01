package cert

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func TestGetQscName(t *testing.T) {
	csr := CertificateSigningRequest{}
	csr.Subj.CN = "QOSC1"

	crt := Certificate{}
	crt.CSR = csr

	assert.Equal(t, "QOSC1", crt.QscName())
}

func TestIsBanker(t *testing.T) {
	csr := CertificateSigningRequest{}
	csr.IsBanker = true

	crt := Certificate{}
	crt.CSR = csr

	assert.Equal(t, true, crt.IsBanker())
}

func TestGetPublicKey(t *testing.T) {
	privKey := ed25519.GenPrivKey()
	pubKey := privKey.PubKey()

	csr := CertificateSigningRequest{}
	csr.PublicKey = pubKey.(ed25519.PubKeyEd25519)

	crt := Certificate{}
	crt.CSR = csr

	assert.Equal(t, true, pubKey.Equals(crt.PublicKey()))
}
