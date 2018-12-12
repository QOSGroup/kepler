package cert

import (
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/types"
	"time"
)

const (
	CommonSubjAminoRoute = "certificate/CommonSubject"
	QSCSubjAminoRoute    = "certificate/QSCSubject"
	QCPSubjAminoRoute    = "certificate/QCPSubject"
)

type Subject interface{}

type CertificateSigningRequest struct {
	ChainId   string        `json:"chain_id"`
	Subj      Subject       `json:"subj"`
	IsCa      bool          `json:"is_ca"`
	NotBefore time.Time     `json:"not_before"`
	NotAfter  time.Time     `json:"not_after"`
	PublicKey crypto.PubKey `json:"public_key"`
}

type Issuer struct {
	Subj      Subject       `json:"subj"`
	PublicKey crypto.PubKey `json:"public_key"`
}

type CommonSubject struct {
	CN string `json:"cn"`
}

type Certificate struct {
	CSR       CertificateSigningRequest `json:"csr"`
	CA        Issuer                    `json:"ca"`
	Signature []byte                    `json:"signature"`
}

func (crt Certificate) ChainId() string {
	return crt.CSR.ChainId
}

func (crt Certificate) PublicKey() crypto.PubKey {
	return crt.CSR.PublicKey
}

type TrustCrts struct {
	PublicKeys []crypto.PubKey `json:"public_keys"`
}

// QSC CA
//-----------------------------------------------------

type QSCSubject struct {
	Name   string        `json:"name"`
	Banker crypto.PubKey `json:"banker"`
}

type QSCCertificate struct {
	*Certificate
}

func (crt QSCCertificate) QSCName() string {
	subj := crt.CSR.Subj.(QSCSubject)
	return subj.Name
}

func (crt QSCCertificate) HasBanker() bool {
	subj := crt.CSR.Subj.(QSCSubject)
	return subj.Banker != nil
}

func (crt QSCCertificate) Banker() types.Address {
	subj := crt.CSR.Subj.(QSCSubject)
	return subj.Banker.Address()
}

// QCP CA
//-----------------------------------------------------

type QCPSubject struct {
	QCPChain string `json:"qcp_chain"`
}

type QCPCertificate struct {
	*Certificate
}

func (crt QCPCertificate) QCPChain() string {
	subj := crt.CSR.Subj.(QCPSubject)
	return subj.QCPChain
}
