package cert

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func TestGetQscName(t *testing.T) {
	csr := CertificateSigningRequest{}
	csr.Subj = QSCSubject{Name: "QOSC1"}

	crt := Certificate{}
	crt.CSR = csr

	assert.Equal(t, "QOSC1", crt.CSR.Subj.(QSCSubject).Name)
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
