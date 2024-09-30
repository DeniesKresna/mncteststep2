package repository

import (
	"github.com/DeniesKresna/mncteststep2/service/extensions/terror"
	"github.com/DeniesKresna/mncteststep2/types/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func TransactionCreateRepository(db *gorm.DB) IRepository {
	transactionRepository := TransactionRepository{
		db: db,
	}
	return &transactionRepository
}

func (r *TransactionRepository) GetDB(ctx *gin.Context) (tx interface{}) {
	return r.db
}

type IRepository interface {
	GetDB(ctx *gin.Context) (tx interface{})
	WalletGet(ctx *gin.Context, payload *models.Wallet, lock bool) (terr terror.ErrInterface)
	WalletUpdate(ctx *gin.Context, payload *models.Wallet) (terr terror.ErrInterface)
	WalletCreate(ctx *gin.Context, payload *models.Wallet) (terr terror.ErrInterface)
	TransactionGet(ctx *gin.Context, payload *models.Transaction, lock bool) (terr terror.ErrInterface)
	TransactionCreate(ctx *gin.Context, payload *models.Transaction) (terr terror.ErrInterface)
	TransactionUpdate(ctx *gin.Context, payload *models.Transaction) (terr terror.ErrInterface)
	TransactionList(ctx *gin.Context, payload models.Transaction) (res []models.Transaction, terr terror.ErrInterface)

	TopupCreate(ctx *gin.Context, payload *models.Topup) (terr terror.ErrInterface)

	PaymentCreate(ctx *gin.Context, payload *models.Payment) (terr terror.ErrInterface)
}
