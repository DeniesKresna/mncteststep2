package usecase

import (
	"fmt"

	"github.com/DeniesKresna/gohelper/utstruct"
	"github.com/DeniesKresna/mncteststep2/service/extensions/helper"
	"github.com/DeniesKresna/mncteststep2/service/extensions/terror"
	"github.com/DeniesKresna/mncteststep2/types/constants"
	"github.com/DeniesKresna/mncteststep2/types/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (u UserUsecase) UserGetByPhoneNumber(ctx *gin.Context, phone string) (user models.User, terr terror.ErrInterface) {
	user.PhoneNumber = phone
	terr = u.userRepo.UserGet(ctx, &user)
	return
}

func (u UserUsecase) UserGetByUserID(ctx *gin.Context, id string) (user models.User, terr terror.ErrInterface) {
	user.UserID = id
	terr = u.userRepo.UserGet(ctx, &user)
	return
}

func (u UserUsecase) UserRegister(ctx *gin.Context, payload models.UserRegisterPayload) (user models.User, terr terror.ErrInterface) {
	helper.TxCreate(ctx, u.userRepo.GetDB)
	defer func() {
		helper.TxSubmitTerr(ctx, terr)
	}()

	// check existing user
	{
		_, terr = u.UserGetByPhoneNumber(ctx, payload.PhoneNumber)
		if terr == nil {
			terr = terror.ErrParameter("Phone Number already registered")
			return
		}
		if terr.GetType() != terror.ERROR_TYPE_DATA_NOT_FOUND {
			return
		}
		terr = nil
	}

	hashedPin, err := bcrypt.GenerateFromPassword([]byte(payload.Pin), bcrypt.DefaultCost)
	if err != nil {
		terr = terror.New(err)
		return
	}
	utstruct.InjectStructValue(payload, &user)
	user.Pin = string(hashedPin)
	user.UserID = helper.CreateUUID()

	terr = u.userRepo.UserCreate(ctx, &user)
	if terr != nil {
		return
	}

	newWallet := models.Wallet{
		UserID:  user.UserID,
		Balance: 0,
	}
	fmt.Println(newWallet)
	_, terr = u.transactionCase.WalletCreate(ctx, newWallet)
	if terr != nil {
		return
	}

	return
}

func (u UserUsecase) UserSearch(ctx *gin.Context, filter models.DbSearchObject) (res models.DbSearchObject, terr terror.ErrInterface) {
	filter.Mode = constants.DB_MODE_PAGE

	usersRes, totalData, terr := u.userRepo.UserSearch(ctx, filter)
	if terr != nil {
		return
	}

	filter.ResponseData = usersRes
	filter.TotalData = totalData
	res = filter

	return
}

func (u UserUsecase) UserUpdate(ctx *gin.Context, payload models.User) (res models.User, terr terror.ErrInterface) {
	userID := ctx.GetString(constants.CONTEXT_USER_ID)

	res, terr = u.UserGetByUserID(ctx, userID)
	if terr != nil {
		return
	}

	res.Address = payload.Address
	res.FirstName = payload.FirstName
	res.LastName = payload.LastName

	terr = u.userRepo.UserUpdate(ctx, &res)
	if terr != nil {
		return
	}

	return
}
