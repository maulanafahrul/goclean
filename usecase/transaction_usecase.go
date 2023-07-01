package usecase

import (
	"goclean/model"
	"goclean/repo"
)

type TransactionUsecase interface {
	List() (*[]model.TransactionHeader, error)
}

type transactionUsecaseImpl struct {
	trxRepo repo.TransactionRepo
}

func (trxUsecase *transactionUsecaseImpl) List() (*[]model.TransactionHeader, error) {
	return trxUsecase.trxRepo.List()
}

func NewTransactionUsecase(trxRepo repo.TransactionRepo) TransactionUsecase {
	return &transactionUsecaseImpl{
		trxRepo: trxRepo,
	}
}
