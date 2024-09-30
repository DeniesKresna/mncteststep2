package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/DeniesKresna/mncteststep2/service/extensions/helper"
	"github.com/DeniesKresna/mncteststep2/service/extensions/terror"
	"github.com/DeniesKresna/mncteststep2/service/modules/user/usecase"
	"github.com/DeniesKresna/mncteststep2/types/constants"
	"github.com/DeniesKresna/mncteststep2/types/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckAuth(userCase usecase.IUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var terr terror.ErrInterface
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			terr = terror.ErrInvalidRule("Unauthenticated")
			ResponseJson(ctx, terr)
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			terr = terror.ErrInvalidRule("Unauthenticated")
			ResponseJson(ctx, terr)
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(utstring.GetEnv(constants.ENV_APP_SECRET, "")), nil
		})

		if err != nil {
			terr = terror.ErrInvalidRule("Unauthenticated")
			ResponseJson(ctx, terr)
			ctx.Abort()
			return
		}

		if !token.Valid {
			terr = terror.ErrInvalidRule("Unauthenticated")
			ResponseJson(ctx, terr)
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			terr = terror.ErrInvalidRule("Unauthenticated")
			ResponseJson(ctx, terr)
			ctx.Abort()
			return
		}

		ctx.Set(constants.CONTEXT_USER_ID, claims.Subject)

		ctx.Next()
	}
}

func ResponseJson(ctx *gin.Context, data interface{}) {
	httpStatusCode := http.StatusOK
	resp, ok := data.(*terror.ErrorModel)
	if ok {
		responseData := models.Response{
			ResponseDesc: resp.Message,
			ResponseCode: resp.Code,
		}

		appEnv := utstring.GetEnv(constants.ENV_APP_ENV, "local")
		if appEnv != "production" {
			responseData.ResponseTrace = resp.Trace
		}

		ctx.JSON(httpStatusCode, responseData)
		return
	}

	responseData := models.Response{
		ResponseDesc: "Success",
		ResponseCode: "00",
	}

	if helper.IsStruct(data) || helper.IsMap(data) || helper.IsSlice(data) {
		responseData.ResponseData = data
	} else {
		responseData.ResponseDesc = fmt.Sprintf("%v", data)
	}

	ctx.JSON(httpStatusCode, responseData)
}
