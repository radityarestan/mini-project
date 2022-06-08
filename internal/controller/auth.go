package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/radityarestan/mini-project/internal/shared"
	"github.com/radityarestan/mini-project/internal/shared/dto"
)

const (
	Error = "error"
	Data  = "data"
)

func (ctr *Controller) register(ctx *gin.Context) {
	var (
		req = dto.RegisterRequest{}
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			Error: dto.ErrorResponse{
				Message: shared.ErrBadRequest.Error(),
			},
		})
		return
	}

	response, err := ctr.service.Register(&req)

	if err != nil {
		if errors.Is(err, shared.ErrUserAlreadyExist) {
			ctx.JSON(http.StatusConflict, gin.H{
				Error: dto.ErrorResponse{
					Message: err.Error(),
				},
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			Error: dto.ErrorResponse{
				Message: shared.ErrInternalServerError.Error(),
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		Data: response,
	})
}

func (ctr *Controller) login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
