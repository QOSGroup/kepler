package cert

import (
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
	RegisterCodec(cdc)
}

func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterInterface((*Subject)(nil), nil)
	cdc.RegisterConcrete(CommonSubject{},
		CommonSubjAminoRoute, nil)
	cdc.RegisterConcrete(QSCSubject{},
		QSCSubjAminoRoute, nil)
	cdc.RegisterConcrete(QCPSubject{},
		QCPSubjAminoRoute, nil)
}

func MakeCodec() *amino.Codec {
	cdc := amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	RegisterCodec(cdc)
	return cdc
}
