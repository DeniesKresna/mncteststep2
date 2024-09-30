package handler

import (
	"github.com/DeniesKresna/mncteststep2/service/extensions/terror"
	"github.com/DeniesKresna/mncteststep2/types/models"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func (h UserHandler) UserGetByID(ctx *gin.Context) {
	var (
		id   string
		terr terror.ErrInterface
	)

	id = ctx.Param("id")

	user, terr := h.userUsecase.UserGetByUserID(ctx, id)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	user.Pin = ""
	ResponseJson(ctx, user)
}

func (h UserHandler) UserRegister(ctx *gin.Context) {
	var (
		payload models.UserRegisterPayload
		terr    terror.ErrInterface
	)

	terr = BindAndValidate(ctx, &payload, govalidator.MapData{
		"first_name":   []string{"required"},
		"last_name":    []string{"required"},
		"address":      []string{"required"},
		"phone_number": []string{"required"},
		"pin":          []string{"required"},
	})
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}

	user, terr := h.userUsecase.UserRegister(ctx, payload)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}

	user.Pin = ""
	ResponseJson(ctx, user)
}

func (h UserHandler) UserSearch(ctx *gin.Context) {
	var (
		search models.DbSearchObject
		terr   terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&search); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	res, terr := h.userUsecase.UserSearch(ctx, search)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, res)
}

func (h UserHandler) UserUpdate(ctx *gin.Context) {
	var (
		payload models.User
		terr    terror.ErrInterface
	)

	terr = BindAndValidate(ctx, &payload, govalidator.MapData{
		"first_name": []string{"required"},
		"last_name":  []string{"required"},
		"address":    []string{"required"},
	})
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}

	user, terr := h.userUsecase.UserUpdate(ctx, payload)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}

	user.Pin = ""
	ResponseJson(ctx, user)
}
