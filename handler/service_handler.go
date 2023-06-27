package handler

import (
	"fmt"
	"goclean/model"
	"goclean/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceHandler interface {
}

type serviceHandlerImpl struct {
	svcUsecase usecase.ServiceUsecase
}

func (svcHandler serviceHandlerImpl) GetServiceById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id harus angka",
		})
		return
	}

	svc, err := svcHandler.svcUsecase.Get(id)
	if err != nil {
		fmt.Printf("serviceHandlerImpl.GetServiceById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data service",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    svc,
	})
}

func (svcHandler serviceHandlerImpl) GetAllService(ctx *gin.Context) {
	svc, err := svcHandler.svcUsecase.List()
	if err != nil {
		fmt.Printf("serviceHandlerImpl.GetAllService() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data service",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    svc,
	})
}

func (svcHandler serviceHandlerImpl) AddService(ctx *gin.Context) {
	payload := &model.ServiceModel{}
	if err := ctx.ShouldBindJSON(&payload); err 
}

func NewServiceHandler(srv *gin.Engine, svcUsecase usecase.ServiceUsecase) ServiceHandler {
	svcHandler := &serviceHandlerImpl{
		svcUsecase: svcUsecase,
	}
	srv.GET("/service", svcHandler.GetAllService)
	srv.GET("/service/:id", svcHandler.GetServiceById)
	srv.POST("/service", svcHandler.AddService)

	return svcHandler
}
