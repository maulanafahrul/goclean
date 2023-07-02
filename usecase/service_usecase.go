package usecase

import (
	"fmt"
	"goclean/apperror"
	"goclean/model"
	"goclean/repo"
)

type ServiceUsecase interface {
	Get(int) (*model.ServiceModel, error)
	List() (*[]model.ServiceModel, error)
	Create(*model.ServiceModel) error
	Update(*model.ServiceModel) error
	Delete(int) error
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

func (svcUsecase *serviceUsecaseImpl) Create(svc *model.ServiceModel) error {
	payloadDB, err := svcUsecase.svcRepo.FindByName(svc.Name)
	if err != nil {
		return fmt.Errorf("serviceUsecaseImpl.Create() : %w", err)
	}
	if payloadDB != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data service dengan nama %v sudah ada", svc.Name),
		}
	}

	return svcUsecase.svcRepo.Create(svc)
}

func (svcUsecase *serviceUsecaseImpl) Update(svc *model.ServiceModel) error {
	payloadDB, err := svcUsecase.svcRepo.FindByName(svc.Name)
	if err != nil {
		return fmt.Errorf("serviceUsecaseImpl.Update() : %w", err)
	}
	if payloadDB == nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data service dengan nama %v belum ada", svc.Name),
		}
	}
	// konvert

	svc.Id = payloadDB.Id

	return svcUsecase.svcRepo.Update(svc)

}

func (svcUsecase *serviceUsecaseImpl) Delete(id int) error {
	svc, err := svcUsecase.svcRepo.Get(id)
	if err != nil {
		return fmt.Errorf("serviceUsecaseImpl.Delete() : %w", err)
	}
	if svc == nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data service dengan id %d belum ada", id),
		}
	}

	return svcUsecase.svcRepo.Delete(id)
}

func NewServiceUseCase(svcRepo repo.ServiceRepo) ServiceUsecase {
	return &serviceUsecaseImpl{
		svcRepo: svcRepo,
	}
}
