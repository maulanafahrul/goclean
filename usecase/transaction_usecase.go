package usecase

import (
	"fmt"
	"goclean/apperror"
	"goclean/model"
	"goclean/repo"
)

type TransactionUsecase interface {
	List() (*[]model.TransactionHeader, error)
	Create(*model.TransactionHeader) error
}

type transactionUsecaseImpl struct {
	trxRepo repo.TransactionRepo
	svcRepo repo.ServiceRepo
}

func (trxUsecase *transactionUsecaseImpl) List() (*[]model.TransactionHeader, error) {
	return trxUsecase.trxRepo.List()
}

func (trxUsecase *transactionUsecaseImpl) Create(trxHeader *model.TransactionHeader) error {
	details := []model.TransactionDetail{}
	for _, det := range trxHeader.ArrDetail {
		isNameExist, err := trxUsecase.svcRepo.FindByName(det.ServiceName)
		if err != nil {
			return fmt.Errorf("transactionUsecaseImpl.FindByName() : %w", err)
		}
		if isNameExist == nil {
			return apperror.AppError{
				ErrorCode:    1,
				ErrorMassage: fmt.Sprintf("data service dengan nama %v belum ada", isNameExist.Name),
			}
		}
		detail := model.TransactionDetail{
			ServiceId:   isNameExist.Id,
			ServiceName: isNameExist.Name,
			Uom:         isNameExist.Uom,
			Price:       isNameExist.Price,
		}
		details = append(details, detail)
	}
	transaction := model.TransactionHeader{
		No:        trxHeader.No,
		StartDate: trxHeader.StartDate,
		EndDate:   trxHeader.EndDate,
		CustName:  trxHeader.CustName,
		Phone:     trxHeader.Phone,
		ArrDetail: details,
	}
	return trxUsecase.trxRepo.Create(&transaction)
}

func NewTransactionUsecase(trxRepo repo.TransactionRepo) TransactionUsecase {
	return &transactionUsecaseImpl{
		trxRepo: trxRepo,
	}
}
