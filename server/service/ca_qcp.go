package service

import (
	"database/sql"
	"fmt"

	"github.com/QOSGroup/kepler/server/module"
)

type CaQcpService struct{}

func (service *CaQcpService) Add(ca module.CaQcp) (cnt int64, err error) {
	cnt, err = module.KEngine.InsertOne(&ca)
	return
}

func (service *CaQcpService) Get(ca module.CaQcp) (*module.CaQcp, error) {
	_, err := module.KEngine.Get(&ca)
	if ca.Id > 0 {
		return &ca, err
	}
	return nil, err
}

// CheckAndUpdateDownload checks if the column download is zero and then set 1 to it
func (service *CaQcpService) CheckAndUpdateDownload(
	ca module.CaQcp) (err error) {
	// cnt, err = module.KEngine.Update(&ca, &module.CaQcp{Id: ca.Id})
	sess := module.KEngine.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		err = fmt.Errorf("fail to session begin: %v", err)
		return
	}
	// if _, err = sess.Where("apply_id = ?", ca.ApplyId)
	// .And("download=?", 0).Update(&ca); err != nil {
	sqlStr := "UPDATE ca_qcp SET download = download + 1 WHERE apply_id = ? AND download = 0"
	var res sql.Result
	if res, err = sess.Exec(sqlStr, ca.ApplyId); err != nil {
		sess.Rollback()
		err = fmt.Errorf("fail to update: %v", err)
		return
	}

	err = sess.Commit()
	var affected int64
	affected, err = res.RowsAffected()
	if err != nil {
		err = fmt.Errorf("update error: %v", err)
		return
	}
	if affected == 0 {
		err = fmt.Errorf("updated nothing")
	}
	return
}

func (service *CaQcpService) FindAll() (cas []*module.CaQcp, err error) {
	err = module.KEngine.OrderBy("id desc").Find(&cas)
	return
}

func (service *CaQcpService) UpdateById(ca module.CaQcp) (cnt int64, err error) {
	cnt, err = module.KEngine.Update(&ca, &module.CaQcp{Id: ca.Id})
	return
}

func (service *CaQcpService) Delete(ca module.CaQcp) (cnt int64, err error) {
	cnt, err = module.KEngine.Delete(&ca)
	return
}

func (service *CaQcpService) Exists(qosChainId string, qcpChainId string) (has bool, err error) {
	ca, err := service.Get(module.CaQcp{
		QosChainId: qosChainId,
		QcpChainId: qcpChainId,
	})
	return ca != nil, err
}
