package service

import (
	"github.com/QOSGroup/kepler/server/module"
)

type CaQscService struct{}

func (service *CaQscService) Add(ca module.CaQsc) (cnt int64, err error) {
	cnt, err = module.KEngine.InsertOne(&ca)
	return
}

func (service *CaQscService) Get(ca module.CaQsc) (*module.CaQsc, error) {
	_, err := module.KEngine.Get(&ca)
	if ca.Id > 0 {
		return &ca, err
	}
	return nil, err
}

func (service *CaQscService) FindAll() (cas []*module.CaQsc, err error) {
	err = module.KEngine.OrderBy("id desc").Find(&cas)
	return
}

func (service *CaQscService) UpdateById(ca module.CaQsc) (cnt int64, err error) {
	cnt, err = module.KEngine.Update(&ca, &module.CaQsc{Id: ca.Id})
	return
}

func (service *CaQscService) Delete(ca module.CaQsc) (cnt int64, err error) {
	cnt, err = module.KEngine.Delete(&ca)
	return
}

func (service *CaQscService) Exists(qosChainId string, ascName string) (has bool, err error) {
	ca, err := service.Get(module.CaQsc{
		QosChainId: qosChainId,
		Name:       ascName,
	})
	return ca != nil, err
}
