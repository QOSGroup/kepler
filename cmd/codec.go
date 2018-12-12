package cmd

import (
	"github.com/QOSGroup/kepler/cert"
)

var cdc = cert.MakeCodec()

func MustMarshalJson(obj interface{}) []byte {
	bz, err := cdc.MarshalJSON(obj)
	if err != nil {
		panic(err)
	}
	return bz
}

func MustMarshalBinaryBare(obj interface{}) []byte {
	bz, err := cdc.MarshalBinaryBare(obj)
	if err != nil {
		panic(err)
	}
	return bz
}
