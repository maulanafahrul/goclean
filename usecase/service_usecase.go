package usecase

import (
	"database/sql"
	"fmt"
	"goclean/apperror"
	"goclean/model"
	"goclean/repo"
)

type ServiceUsecase interface {
	Get(int) (*model.ServiceModel, error)
	List() (*[]model.ServiceModel, error)
	Create(*model.ReqService) error
	Update(*model.ReqService) error
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

func (svcUsecase *serviceUsecaseImpl) Create(payload *model.ReqService) error {
	payloadDB, err := svcUsecase.svcRepo.FindByName(payload.Name)
	if err != nil {
		return fmt.Errorf("serviceUsecaseImpl.Create() : %w", err)
	}
	if payloadDB != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data service dengan nama %v sudah ada", payload.Name),
		}
	}
	// konvert
	svc := &model.ServiceModel{}
	svc.Name = sql.NullString{String: payload.Name, Valid: true}
	svc.Uom = sql.NullString{String: payload.Uom, Valid: true}
	svc.Price = sql.NullFloat64{Float64: payload.Price, Valid: true}

	return svcUsecase.svcRepo.Create(svc)
}

func (svcUsecase *serviceUsecaseImpl) Update(payload *model.ReqService) error {
	payloadDB, err := svcUsecase.svcRepo.FindByName(payload.Name)
	if err != nil {
		return fmt.Errorf("serviceUsecaseImpl.Update() : %w", err)
	}
	if payloadDB == nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data service dengan nama %v belum ada", payload.Name),
		}
	}
	// konvert
	svc := &model.ServiceModel{}
	svc.Id = payloadDB.Id
	svc.Name = sql.NullString{String: payload.Name, Valid: true}
	svc.Uom = sql.NullString{String: payload.Uom, Valid: true}
	svc.Price = sql.NullFloat64{Float64: payload.Price, Valid: true}

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
