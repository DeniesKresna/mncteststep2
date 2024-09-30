package repository

import (
	"github.com/DeniesKresna/mncteststep2/service/extensions/terror"
	"github.com/DeniesKresna/mncteststep2/types/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func UserCreateRepository(db *gorm.DB) IRepository {
	userRepository := UserRepository{
		db: db,
	}
	return userRepository
}

type IRepository interface {
	GetDB(ctx *gin.Context) (tx interface{})
	UserGet(ctx *gin.Context, user *models.User) (terr terror.ErrInterface)
	UserCreate(ctx *gin.Context, user *models.User) (terr terror.ErrInterface)
	UserSearch(ctx *gin.Context, searchPayload models.DbSearchObject) (users []models.User, totalData int64, terr terror.ErrInterface)
	UserUpdate(ctx *gin.Context, user *models.User) (terr terror.ErrInterface)
}
