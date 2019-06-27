package service

import (
	"github.com/QOSGroup/kepler/server/module"
	"github.com/QOSGroup/kepler/server/types"
)

type ApplyQcpService struct{}

func (service *ApplyQcpService) Add(apply module.ApplyQcp) (cnt int64, err error) {
	cnt, err = module.KEngine.InsertOne(&apply)
	return
}

func (service *ApplyQcpService) Get(apply module.ApplyQcp) (*module.ApplyQcp, error) {
	_, err := module.KEngine.Get(&apply)
	if apply.Id > 0 {
		return &apply, err
	}
	return nil, err
}

func (service *ApplyQcpService) Find(apply module.ApplyQcp, page types.Page) (cas []*module.ApplyQcp, err error) {
	err = module.KEngine.OrderBy("id desc").Limit(page.Limit(), page.Start()).Find(&cas, &apply)
	return
}

func (service *ApplyQcpService) FindAll() (cas []*module.ApplyQcp, err error) {
	err = module.KEngine.OrderBy("id desc").Find(&cas)
	return
}

func (service *ApplyQcpService) Update(apply module.ApplyQcp) (cnt int64, err error) {
	cnt, err = module.KEngine.Update(&apply)
	return
}

func (service *ApplyQcpService) Delete(apply module.ApplyQcp) (cnt int64, err error) {
	cnt, err = module.KEngine.Delete(&apply)
	return
}

func (service *ApplyQcpService) Exists(qosChainId string, qcpChainId string, email string) (has bool, err error) {
	apply, err := service.Get(module.ApplyQcp{
		QosChainId: qosChainId,
		QcpChainId: qcpChainId,
		Email:      email,
	})
	return apply != nil, err
}
