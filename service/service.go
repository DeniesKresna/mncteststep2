package service

import (
	"fmt"

	"github.com/DeniesKresna/mncteststep2/config"

	userModule "github.com/DeniesKresna/mncteststep2/service/modules/user"
	userrepo "github.com/DeniesKresna/mncteststep2/service/modules/user/repository"
	usercase "github.com/DeniesKresna/mncteststep2/service/modules/user/usecase"

	transactionModule "github.com/DeniesKresna/mncteststep2/service/modules/transaction"
	transactionrepo "github.com/DeniesKresna/mncteststep2/service/modules/transaction/repository"
	transactioncase "github.com/DeniesKresna/mncteststep2/service/modules/transaction/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setRoutes(cfg *config.Config) (r *gin.Engine, err error) {
	transactionRepo := transactionrepo.TransactionCreateRepository(cfg.DB)
	transactionCase := transactioncase.TransactionCreateUsecase(transactionRepo)

	userRepo := userrepo.UserCreateRepository(cfg.DB)
	userCase := usercase.UserCreateUsecase(userRepo, transactionCase)

	r = gin.New()
	r.Use(corsConfig(), recoverHandle())

	api := r.Group("/")
	{
		userModule.InitRoutes(api, userCase, cfg)
		transactionModule.InitRoutes(api, transactionCase, userCase, cfg)
	}

	return
}

func recoverHandle() gin.HandlerFunc {
	return func(ctxOri *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		ctxOri.Next()
	}
}

func corsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Include "Content-Type" in the list of allowed headers
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
}

func Start(cfg *config.Config) (err error) {
	eng, err := setRoutes(cfg)

	eng.Run(cfg.App.Host + ":" + cfg.App.Port)
	return
}
