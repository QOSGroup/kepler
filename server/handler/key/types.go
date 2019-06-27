package key

import (
	"github.com/QOSGroup/kepler/cert"
)

type KeyValue struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type KeyData struct {
	PrivKey KeyValue `json:"priv_key"`
	PubKey  KeyValue `json:"pub_key"`
}

func NewKeyData(privAminoJson []byte, pubAminoJson []byte) *KeyData {
	data := &KeyData{}
	cert.Codec.MustUnmarshalJSON(privAminoJson, &data.PrivKey)
	cert.Codec.MustUnmarshalJSON(pubAminoJson, &data.PubKey)
	return data
}
