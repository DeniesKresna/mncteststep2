package repository

import (
	"errors"

	"github.com/DeniesKresna/mncteststep2/service/extensions/helper"
	"github.com/DeniesKresna/mncteststep2/service/extensions/terror"
	"github.com/DeniesKresna/mncteststep2/types/constants"
	"github.com/DeniesKresna/mncteststep2/types/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (r UserRepository) GetDB(ctx *gin.Context) (tx interface{}) {
	return r.db
}

func (r UserRepository) UserGet(ctx *gin.Context, user *models.User) (terr terror.ErrInterface) {
	err := r.db.Where(user).First(user).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r UserRepository) UserCreate(ctx *gin.Context, user *models.User) (terr terror.ErrInterface) {
	err := r.db.Create(user).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}

func (r UserRepository) UserSearch(ctx *gin.Context, searchPayload models.DbSearchObject) (users []models.User, totalData int64, terr terror.ErrInterface) {
	queryDB := r.db.Session(&gorm.Session{})

	{
		if searchPayload.PayloadData != nil {
			if val, ok := searchPayload.PayloadData["email"]; ok {
				if valStr, ok := val.(string); ok {
					queryDB = queryDB.Where("email like ?", helper.WrapString(valStr, "%"))
				}
			}
			if val, ok := searchPayload.PayloadData["name"]; ok {
				if valStr, ok := val.(string); ok {
					queryDB = queryDB.Where("name like ?", helper.WrapString(valStr, "%"))
				}
			}
		}
	}

	if searchPayload.Mode == constants.DB_MODE_DATA || searchPayload.Mode == constants.DB_MODE_PAGE {
		offset := (searchPayload.Page - 1) * searchPayload.Limit
		for _, v := range searchPayload.Order {
			queryDB = queryDB.Order(v)
		}
		err := queryDB.Limit(int(searchPayload.Limit)).Offset(int(offset)).Find(&users).Error
		if err != nil {
			terr = terror.New(err)
			return
		}
	}

	if searchPayload.Mode == constants.DB_MODE_COUNT || searchPayload.Mode == constants.DB_MODE_PAGE {
		err := queryDB.Model(&models.User{}).Limit(-1).Offset(-1).Count(&totalData).Error
		if err != nil {
			terr = terror.New(err)
			return
		}
	}

	return
}

func (r UserRepository) UserUpdate(ctx *gin.Context, user *models.User) (terr terror.ErrInterface) {
	err := r.db.Save(user).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}
