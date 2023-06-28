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
		return fmt.Errorf("serviceUsecaseImpl.InsertService() : %w", err)
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

func NewServiceUseCase(svcRepo repo.ServiceRepo) ServiceUsecase {
	return &serviceUsecaseImpl{
		svcRepo: svcRepo,
	}
}
