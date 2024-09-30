package usecase

import (
	"github.com/DeniesKresna/mncteststep2/service/extensions/terror"
	"github.com/DeniesKresna/mncteststep2/service/modules/transaction/repository"
	"github.com/DeniesKresna/mncteststep2/types/models"
	"github.com/gin-gonic/gin"
)

type TransactionUsecase struct {
	transactionRepo repository.IRepository
}

func TransactionCreateUsecase(transactionRepo repository.IRepository) IUsecase {
	transactionUsecase := TransactionUsecase{
		transactionRepo: transactionRepo,
	}
	return &transactionUsecase
}

type IUsecase interface {
	WalletGetByUserID(ctx *gin.Context, userID string, lock bool) (res models.Wallet, terr terror.ErrInterface)
	WalletCreate(ctx *gin.Context, payload models.Wallet) (res models.Wallet, terr terror.ErrInterface)
	WalletUpdate(ctx *gin.Context, payload models.Wallet) (res models.Wallet, terr terror.ErrInterface)
	WalletTopup(ctx *gin.Context, amount int64) (res models.Transaction, terr terror.ErrInterface)
	TransactionCreate(ctx *gin.Context, payload models.Transaction) (res models.Transaction, terr terror.ErrInterface)
	PaymentCreate(ctx *gin.Context, payload models.PaymentCreateRequest) (res models.Transaction, terr terror.ErrInterface)
	TransactionList(ctx *gin.Context) (res []models.Transaction, terr terror.ErrInterface)
}
