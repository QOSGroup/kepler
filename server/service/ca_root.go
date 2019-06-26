package service

import (
	"github.com/QOSGroup/kepler/server/module"
)

type CaRootService struct{}

func (service *CaRootService) Get(ca module.RootCa) (*module.RootCa, error) {
	_, err := module.KEngine.Get(&ca)
	if ca.Id == 0 || ca.Type != module.ROOT {
		ca.Type = module.ROOT
		_, err = module.KEngine.Get(&ca)
	}
	return &ca, err
}
