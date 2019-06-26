package qsc

import (
	"fmt"
	"github.com/QOSGroup/kepler/cert"
	"github.com/QOSGroup/kepler/server/mail"
	"github.com/QOSGroup/kepler/server/module"
	"github.com/QOSGroup/kepler/server/service"
	"github.com/QOSGroup/kepler/server/types"
	"github.com/gin-gonic/gin"
	"github.com/tendermint/tendermint/crypto"
	"net/http"
	"strconv"
	"time"
)

var applyService = service.ApplyQscService{}
var caService = service.CaQscService{}
var rootService = service.CaRootService{}

// QCP证书管理API
func Register(r *gin.Engine) {
	r.POST("/qsc/apply", addApply())
	r.GET("/qsc/apply", queryApply())
	r.GET("/qsc/apply/:id", getApply())
	r.PUT("/qsc/apply/:id", updateApply())
	r.GET("/qsc/ca", findCa())
	r.GET("/qsc/ca/:id", getCa())
}

func addApply() gin.HandlerFunc {
	return func(c *gin.Context) {
		var apply module.ApplyQsc
		if err := c.ShouldBind(&apply); err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}

		if ca, err := rootService.Get(module.RootCa{ChainId: apply.QosChainId, Type: module.ROOT}); ca.Id == 0 || err != nil {
			c.JSON(http.StatusOK, types.Error(fmt.Sprintf("no %s public chain", apply.QosChainId)))
			return
		}

		if exists, err := applyService.Exists(apply.QosChainId, apply.QscName, apply.Email); exists || err != nil {
			c.JSON(http.StatusOK, types.Error("repeat apply"))
			return
		}

		if exists, err := caService.Exists(apply.QosChainId, apply.QscName); exists || err != nil {
			c.JSON(http.StatusOK, types.Error(fmt.Sprintf("%s in %s has been registered", apply.QscName, apply.QosChainId)))
			return
		}

		// valid qsc_pub
		var pubKey crypto.PubKey
		err := cert.Codec.UnmarshalJSON([]byte(apply.QscPub), &pubKey)
		if err != nil {
			c.JSON(http.StatusOK, types.Error("qsc_pub incorrect"))
			return
		}

		//valid banker_pub
		if len(apply.BankerPub) != 0 {
			err = cert.Codec.UnmarshalJSON([]byte(apply.BankerPub), &pubKey)
			if err != nil {
				c.JSON(http.StatusOK, types.Error("banker_pub incorrect"))
				return
			}
		}

		apply.CreateTime = time.Now()
		apply.UpdateTime = time.Now()
		res, err := applyService.Add(apply)
		if err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}

		a, err := applyService.Get(module.ApplyQsc{QosChainId: apply.QosChainId, QscName: apply.QscName, Email: apply.Email})
		if err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}
		res, err = addCa(*a)

		c.JSON(http.StatusOK, types.Ok(res))
	}
}

func queryApply() gin.HandlerFunc {
	return func(c *gin.Context) {
		apply := module.ApplyQsc{Email: c.DefaultQuery("email", ""), Phone: c.DefaultQuery("phone", "")}
		var page types.Page
		c.ShouldBind(&page)
		res, err := applyService.Find(apply, page)
		if err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}
		c.IndentedJSON(http.StatusOK, types.Ok(res))
	}
}

func getApply() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, types.Error(err))
		}
		apply := module.ApplyQsc{Id: id}
		res, err := applyService.Get(apply)
		if err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}
		c.IndentedJSON(http.StatusOK, types.Ok(res))
	}
}

func updateApply() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query module.ApplyQsc
		if err := c.ShouldBind(&query); err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}
		res, err := applyService.Update(query)
		if res != 1 && err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}
		c.JSON(http.StatusOK, types.Ok(res))
	}
}

func addCa(apply module.ApplyQsc) (int64, error) {
	csr := cert.CertificateSigningRequest{}
	qscSubject := cert.QSCSubject{}
	qscSubject.ChainId = apply.QosChainId
	qscSubject.Name = apply.QscName
	cert.Codec.MustUnmarshalJSON([]byte(apply.BankerPub), &qscSubject.Banker)
	cert.Codec.MustUnmarshalJSON([]byte(apply.QscPub), &csr.PublicKey)
	csr.IsCa = false
	csr.Subj = qscSubject
	csr.NotBefore = time.Now()
	csr.NotAfter = time.Now().AddDate(1, 0, 0)

	crt := cert.Certificate{}
	crt.CSR = csr

	rootCa, err := rootService.Get(module.RootCa{
		ChainId: apply.QosChainId,
		Type:    module.ROOT_QSC,
	})
	if err != nil {
		return 0, err
	}
	cert.Codec.UnmarshalJSON([]byte(rootCa.PubKey), &crt.CA.PublicKey)
	var rootPriv crypto.PrivKey
	cert.Codec.UnmarshalJSON([]byte(rootCa.PrivKey), &rootPriv)
	crt.Signature, err = rootPriv.Sign(cert.MustMarshalJson(csr))
	if err != nil {
		return 0, err
	}

	ca := module.CaQsc{
		QosChainId: apply.QosChainId,
		Name:       apply.QscName,
		Csr:        string(cert.MustMarshalJson(csr)),
		Crt:        string(cert.MustMarshalJson(crt)),
		ApplyId:    apply.Id,
		CreateTime: time.Now(),
		ExpireTime: csr.NotAfter,
	}
	cnt, err := caService.Add(ca)

	go mail.Send(apply.Email, fmt.Sprintf("qsc crt for %s in %s", apply.QscName, apply.QosChainId), ca.Crt)

	return cnt, err
}

func findCa() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := caService.FindAll()
		if err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}
		c.IndentedJSON(http.StatusOK, types.Ok(res))
	}
}

func getCa() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, types.Error(err))
		}
		ca := module.CaQsc{Id: id}
		res, err := caService.Get(ca)
		if err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}
		c.IndentedJSON(http.StatusOK, types.Ok(res))
	}
}
