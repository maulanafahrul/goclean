package handler

import (
	"fmt"
	"goclean/model"
	"goclean/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
}

type transactionHandlerImpl struct {
	trxUsecase usecase.TransactionUsecase
}

func (trxHandler transactionHandlerImpl) GetAllTransactionHandler(ctx *gin.Context) {
	trx, err := trxHandler.trxUsecase.List()
	if err != nil {
		fmt.Printf("serviceHandlerImpl.GetAllService() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data transaction",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    trx,
	})
}

func (trxHandler transactionHandlerImpl) addTransactionHandler(ctx *gin.Context) {
	trxheader := model.TransactionHeader{}
	if err := ctx.ShouldBindJSON(&trxheader); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

}

func NewTransactionHandler(srv *gin.Engine, trxUsecase usecase.TransactionUsecase) TransactionHandler {
	trxHandler := &transactionHandlerImpl{
		trxUsecase: trxUsecase,
	}
	srv.GET("/transaction", trxHandler.GetAllTransactionHandler)
	srv.POST("/transaction", trxHandler.addTransactionHandler)
	return trxHandler
}
