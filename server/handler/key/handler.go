package key

import (
	"github.com/QOSGroup/kepler/cert"
	"github.com/QOSGroup/kepler/server/types"
	"github.com/gin-gonic/gin"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"net/http"
)

// 公私钥生成API
func Register(r *gin.Engine) {
	r.GET("/key/gen", addKey())
}

func addKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		privKey := ed25519.GenPrivKey()
		res := NewKeyData(cert.Codec.MustMarshalJSON(privKey), cert.Codec.MustMarshalJSON(privKey.PubKey()))
		c.IndentedJSON(http.StatusOK, types.OkWithMsg(res, "请将priv_key和pub_key部分内容分别存放为.pri和.pub文件，JSON格式，注意要去除空格"))
	}
}
