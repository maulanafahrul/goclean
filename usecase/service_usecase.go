package usecase

import (
	"goclean/model"
	"goclean/repo"
)

type ServiceUsecase interface {
	Get(int) (*model.ServiceModel, error)
	List() (*[]model.ServiceModel, error)
}

type serviceUsecaseImpl struct {
	svcRepo repo.ServiceRepo
}

func (svcUsecase *serviceUsecaseImpl) Get(id int) (*model.ServiceModel, error) {
	return svcUsecase.svcRepo.Get(id)
}
func (svcUsecase *serviceUsecaseImpl) List() (*[]model.ServiceModel, error) {
	return svcUsecase.svcRepo.List()
}

func NewServiceUseCase(svcRepo repo.ServiceRepo) ServiceUsecase {
	return &serviceUsecaseImpl{
		svcRepo: svcRepo,
	}
}
