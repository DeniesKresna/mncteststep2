package handler

import (
	"github.com/DeniesKresna/mncteststep2/service/extensions/terror"
	"github.com/DeniesKresna/mncteststep2/types/models"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func (h UserHandler) AuthLogin(ctx *gin.Context) {
	var (
		payload models.LoginRequest
		terr    terror.ErrInterface
	)

	terr = BindAndValidate(ctx, &payload, govalidator.MapData{
		"phone_number": []string{"required"},
		"pin":          []string{"required"},
	})
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}

	authResp, terr := h.userUsecase.AuthLogin(ctx, payload)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, authResp)
}
