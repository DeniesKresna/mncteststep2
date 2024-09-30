package user

import (
	"github.com/DeniesKresna/mncteststep2/config"
	"github.com/DeniesKresna/mncteststep2/service/middlewares"
	"github.com/DeniesKresna/mncteststep2/service/modules/transaction/handler"
	"github.com/DeniesKresna/mncteststep2/service/modules/transaction/usecase"
	userUsecase "github.com/DeniesKresna/mncteststep2/service/modules/user/usecase"
	"github.com/gin-gonic/gin"
)

func InitRoutes(v1 *gin.RouterGroup, transactionCase usecase.IUsecase, userCase userUsecase.IUsecase, cfg *config.Config) {
	transactionHandler := handler.TransactionCreateHandler(transactionCase)

	authRoute := v1.Use(authCheck(userCase))
	{
		authRoute.POST("/topup", transactionHandler.WalletTopup)
		authRoute.POST("/pay", transactionHandler.PaymentCreate)
		authRoute.GET("/transactions", transactionHandler.TransactionList)
	}

}

func authCheck(userCase userUsecase.IUsecase) gin.HandlerFunc {
	return middlewares.CheckAuth(userCase)
}
