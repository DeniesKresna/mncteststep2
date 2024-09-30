package usecase

import (
	"time"

	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/DeniesKresna/gohelper/utstruct"
	"github.com/DeniesKresna/mncteststep2/service/extensions/terror"
	"github.com/DeniesKresna/mncteststep2/types/constants"
	"github.com/DeniesKresna/mncteststep2/types/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (u UserUsecase) AuthGetFromContext(ctx *gin.Context) (res models.User, terr terror.ErrInterface) {
	const ERR_UNAUTHENTICATED = "Unauthenticated"

	userID := ctx.GetString(constants.CONTEXT_USER_ID)

	userRes, terr := u.UserGetByUserID(ctx, userID)
	if terr != nil {
		terr = terror.ErrInvalidRule(ERR_UNAUTHENTICATED)
		return
	}

	utstruct.InjectStructValue(userRes, &res)

	return
}

func (u UserUsecase) AuthLogin(ctx *gin.Context, payload models.LoginRequest) (resp models.LoginResponse, terr terror.ErrInterface) {
	const ERR_MESSAGE = "Phone Number and PIN doesn’t match."
	user, terr := u.UserGetByPhoneNumber(ctx, payload.PhoneNumber)
	if terr != nil {
		if terr.GetType() == terror.ERROR_TYPE_DATA_NOT_FOUND {
			terr = terror.ErrParameter(ERR_MESSAGE)
			return
		} else {
			return
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Pin), []byte(payload.Pin))
	if err != nil {
		terr = terror.ErrParameter(ERR_MESSAGE)
		return
	}

	var (
		tokenString string
		expires     time.Time
	)

	{
		expires = time.Now().Add(time.Hour * 2)

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			ExpiresAt: expires.Unix(),
			Issuer:    utstring.GetEnv(constants.ENV_APP_NAME),
			Subject:   user.UserID,
		})

		tokenString, err = token.SignedString([]byte(utstring.GetEnv(constants.ENV_APP_SECRET, "")))
		if err != nil {
			terr = terror.New(err)
			return
		}
	}

	resp = models.LoginResponse{
		AccessToken:  tokenString,
		RefreshToken: tokenString,
	}

	return
}
