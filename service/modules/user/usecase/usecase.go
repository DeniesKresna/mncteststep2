package usecase

import (
	"github.com/DeniesKresna/mncteststep2/service/extensions/terror"
	transactionUsecase "github.com/DeniesKresna/mncteststep2/service/modules/transaction/usecase"
	"github.com/DeniesKresna/mncteststep2/service/modules/user/repository"
	"github.com/DeniesKresna/mncteststep2/types/models"
	"github.com/gin-gonic/gin"
)

type UserUsecase struct {
	userRepo        repository.IRepository
	transactionCase transactionUsecase.IUsecase
}

func UserCreateUsecase(userRepo repository.IRepository, transactionCase transactionUsecase.IUsecase) IUsecase {
	userUsecase := UserUsecase{
		userRepo:        userRepo,
		transactionCase: transactionCase,
	}
	return userUsecase
}

type IUsecase interface {
	AuthGetFromContext(ctx *gin.Context) (res models.User, terr terror.ErrInterface)
	AuthLogin(ctx *gin.Context, payload models.LoginRequest) (resp models.LoginResponse, terr terror.ErrInterface)
	UserGetByPhoneNumber(ctx *gin.Context, phone string) (user models.User, terr terror.ErrInterface)
	UserGetByUserID(ctx *gin.Context, id string) (user models.User, terr terror.ErrInterface)
	UserSearch(ctx *gin.Context, filter models.DbSearchObject) (res models.DbSearchObject, terr terror.ErrInterface)
	UserRegister(ctx *gin.Context, payload models.UserRegisterPayload) (user models.User, terr terror.ErrInterface)
	UserUpdate(ctx *gin.Context, payload models.User) (res models.User, terr terror.ErrInterface)
}
