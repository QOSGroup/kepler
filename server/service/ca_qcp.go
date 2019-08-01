package service

import (
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
