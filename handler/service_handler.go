package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"goclean/apperror"
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

func (svcHandler serviceHandlerImpl) GetServiceByIdHandler(ctx *gin.Context) {
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
	if svc == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": fmt.Sprintf("data dengan id: %d tidak ada", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    svc,
	})
}

func (svcHandler serviceHandlerImpl) GetAllServiceHandler(ctx *gin.Context) {
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

func (svcHandler serviceHandlerImpl) AddServiceHandler(ctx *gin.Context) {
	payload := &model.ReqService{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	// validate
	if len(payload.Name) > 15 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Panjang Nama tidak boleh lebih dari 15 karakter",
		})
		return
	}
	if payload.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Nama gak boleh kosong",
		})
		return
	}
	// konvert
	svc := &model.ServiceModel{}
	svc.Name = sql.NullString{String: payload.Name, Valid: true}
	svc.Uom = sql.NullString{String: payload.Uom, Valid: true}
	svc.Price = sql.NullFloat64{Float64: payload.Price, Valid: true}

	err := svcHandler.svcUsecase.Create(svc)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("ServiceHandler.Create() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("ServiceHandler.Create() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menyimpan data service",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (svcHandler serviceHandlerImpl) UpdateServiceHandler(ctx *gin.Context) {
	payload := &model.ReqService{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	// validate
	if len(payload.Name) > 15 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Panjang Nama tidak boleh lebih dari 15 karakter",
		})
		return
	}
	if payload.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Nama gak boleh kosong",
		})
		return
	}
	svc := &model.ServiceModel{}
	svc.Name = sql.NullString{String: payload.Name, Valid: true}
	svc.Uom = sql.NullString{String: payload.Uom, Valid: true}
	svc.Price = sql.NullFloat64{Float64: payload.Price, Valid: true}
	err := svcHandler.svcUsecase.Update(svc)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("ServiceHandler.Update() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("ServiceHandler.Update() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika update data service",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (svcHandler serviceHandlerImpl) DeleteServiceHandler(ctx *gin.Context) {
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
	err = svcHandler.svcUsecase.Delete(id)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("ServiceHandler.Delete() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("ServiceHandler.Delete() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika Delete data service",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})

}

func NewServiceHandler(srv *gin.Engine, svcUsecase usecase.ServiceUsecase) ServiceHandler {
	svcHandler := &serviceHandlerImpl{
		svcUsecase: svcUsecase,
	}
	srv.GET("/service", svcHandler.GetAllServiceHandler)
	srv.GET("/service/:id", svcHandler.GetServiceByIdHandler)
	srv.POST("/service", svcHandler.AddServiceHandler)
	srv.PUT("/service", svcHandler.UpdateServiceHandler)
	srv.DELETE("/service/:id", svcHandler.DeleteServiceHandler)

	return svcHandler
}
