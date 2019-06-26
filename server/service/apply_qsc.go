package service

import (
	"github.com/QOSGroup/kepler/server/module"
	"github.com/QOSGroup/kepler/server/types"
)

type ApplyQscService struct{}

func (service *ApplyQscService) Add(apply module.ApplyQsc) (cnt int64, err error) {
	cnt, err = module.KEngine.InsertOne(&apply)
	return
}

func (service *ApplyQscService) Get(apply module.ApplyQsc) (*module.ApplyQsc, error) {
	_, err := module.KEngine.Get(&apply)
	if apply.Id > 0 {
		return &apply, err
	}
	return nil, err
}

func (service *ApplyQscService) Find(apply module.ApplyQsc, page types.Page) (cas []*module.ApplyQsc, err error) {
	err = module.KEngine.OrderBy("id desc").Limit(page.Limit(), page.Start()).Find(&cas, &apply)
	return
}

func (service *ApplyQscService) FindAll() (cas []*module.ApplyQsc, err error) {
	err = module.KEngine.OrderBy("id desc").Find(&cas)
	return
}

func (service *ApplyQscService) Update(apply module.ApplyQsc) (cnt int64, err error) {
	cnt, err = module.KEngine.Update(&apply)
	return
}

func (service *ApplyQscService) Delete(apply module.ApplyQsc) (cnt int64, err error) {
	cnt, err = module.KEngine.Delete(&apply)
	return
}

func (service *ApplyQscService) Exists(qosChainId string, qscName string, email string) (has bool, err error) {
	apply, err := service.Get(module.ApplyQsc{
		QosChainId: qosChainId,
		QscName:    qscName,
		Email:      email,
	})
	return apply != nil, err
}
