package user

import (
	"github.com/DeniesKresna/mncteststep2/config"
	"github.com/DeniesKresna/mncteststep2/service/modules/user/handler"
	"github.com/DeniesKresna/mncteststep2/service/modules/user/usecase"
	"github.com/gin-gonic/gin"
)

func InitRoutes(v1 *gin.RouterGroup, userCase usecase.IUsecase, cfg *config.Config) {
	userHandler := handler.UserCreateHandler(userCase)

	v1.POST("/login", userHandler.AuthLogin)
	v1.POST("/register", userHandler.UserRegister)
	v1.PUT("/profile", userHandler.UserUpdate)
}
