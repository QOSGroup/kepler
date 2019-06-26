package qcp

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

var applyService = service.ApplyQcpService{}
var caService = service.CaQcpService{}
var rootService = service.CaRootService{}

// QCP证书管理API
func Register(r *gin.Engine) {
	r.POST("/qcp/apply", addApply())
	r.GET("/qcp/apply", queryApply())
	r.GET("/qcp/apply/:id", getApply())
	r.PUT("/qcp/apply/:id", updateApply())
	r.GET("/qcp/ca", findCa())
	r.GET("/qcp/ca/:id", getCa())
}

func addApply() gin.HandlerFunc {
	return func(c *gin.Context) {
		var apply module.ApplyQcp
		if err := c.ShouldBind(&apply); err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}

		if ca, err := rootService.Get(module.RootCa{ChainId: apply.QosChainId, Type: module.ROOT}); ca.Id == 0 || err != nil {
			c.JSON(http.StatusOK, types.Error(fmt.Sprintf("no %s public chain", apply.QosChainId)))
			return
		}

		if exists, err := applyService.Exists(apply.QosChainId, apply.QcpChainId, apply.Email); exists || err != nil {
			c.JSON(http.StatusOK, types.Error("repeat apply"))
			return
		}

		if exists, err := caService.Exists(apply.QosChainId, apply.QcpChainId); exists || err != nil {
			c.JSON(http.StatusOK, types.Error(fmt.Sprintf("%s in %s has been registered", apply.QcpChainId, apply.QosChainId)))
			return
		}

		// valid qcp_pub
		var pubKey crypto.PubKey
		err := cert.Codec.UnmarshalJSON([]byte(apply.QcpPub), &pubKey)
		if err != nil {
			c.JSON(http.StatusOK, types.Error("qcp_pub incorrect"))
			return
		}

		apply.CreateTime = time.Now()
		apply.UpdateTime = time.Now()
		res, err := applyService.Add(apply)
		if res != 1 && err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}
		c.JSON(http.StatusOK, types.Ok(res))
	}
}

func queryApply() gin.HandlerFunc {
	return func(c *gin.Context) {
		apply := module.ApplyQcp{Email: c.DefaultQuery("email", ""), Phone: c.DefaultQuery("phone", "")}
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
		apply := module.ApplyQcp{Id: id}
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
		var query module.ApplyQcp
		if err := c.ShouldBind(&query); err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}
		res, err := applyService.Update(query)
		if res != 1 && err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}
		apply, err := applyService.Get(query)
		if err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}

		res, err = addCa(*apply)
		if err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}
		c.JSON(http.StatusOK, types.Ok(res))
	}
}

func addCa(apply module.ApplyQcp) (int64, error) {
	csr := cert.CertificateSigningRequest{}
	qcpSubject := cert.QCPSubject{}
	qcpSubject.ChainId = apply.QosChainId
	qcpSubject.QCPChain = apply.QcpChainId
	cert.Codec.MustUnmarshalJSON([]byte(apply.QcpPub), &csr.PublicKey)
	csr.IsCa = false
	csr.Subj = qcpSubject
	csr.NotBefore = time.Now()
	csr.NotAfter = time.Now().AddDate(1, 0, 0)

	crt := cert.Certificate{}
	crt.CSR = csr

	rootCa, err := rootService.Get(module.RootCa{
		ChainId: apply.QosChainId,
		Type:    module.ROOT_QCP,
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

	ca := module.CaQcp{
		QosChainId: apply.QosChainId,
		QcpChainId: apply.QcpChainId,
		Csr:        string(cert.MustMarshalJson(csr)),
		Crt:        string(cert.MustMarshalJson(crt)),
		ApplyId:    apply.Id,
		CreateTime: time.Now(),
		ExpireTime: csr.NotAfter,
	}
	cnt, err := caService.Add(ca)

	go mail.Send(apply.Email, fmt.Sprintf("qcp crt for %s in %s", apply.QcpChainId, apply.QosChainId), ca.Crt)

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
		ca := module.CaQcp{Id: id}
		res, err := caService.Get(ca)
		if err != nil {
			c.JSON(http.StatusOK, types.Error(err))
			return
		}
		c.IndentedJSON(http.StatusOK, types.Ok(res))
	}
}
