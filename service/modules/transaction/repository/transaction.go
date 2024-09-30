package repository

import (
	"errors"

	"github.com/DeniesKresna/mncteststep2/service/extensions/helper"
	"github.com/DeniesKresna/mncteststep2/service/extensions/terror"
	"github.com/DeniesKresna/mncteststep2/types/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *TransactionRepository) WalletGet(ctx *gin.Context, payload *models.Wallet, lock bool) (terr terror.ErrInterface) {
	db := helper.TxGetDefault(ctx, r.db).Where(payload)
	if lock {
		db = db.Clauses(clause.Locking{Strength: "UPDATE"})
	}

	err := db.First(payload).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData("Wallet not found")
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r *TransactionRepository) WalletCreate(ctx *gin.Context, payload *models.Wallet) (terr terror.ErrInterface) {
	db := helper.TxGetDefault(ctx, r.db)

	err := db.Create(payload).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}

func (r *TransactionRepository) WalletUpdate(ctx *gin.Context, payload *models.Wallet) (terr terror.ErrInterface) {
	db := helper.TxGetDefault(ctx, r.db)

	err := db.Updates(payload).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}

func (r *TransactionRepository) TransactionGet(ctx *gin.Context, payload *models.Transaction, lock bool) (terr terror.ErrInterface) {
	db := helper.TxGetDefault(ctx, r.db).Where(payload)
	if lock {
		db = db.Clauses(clause.Locking{Strength: "UPDATE"})
	}

	err := db.First(payload).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData("Transaction not found")
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r *TransactionRepository) TransactionList(ctx *gin.Context, payload models.Transaction) (res []models.Transaction, terr terror.ErrInterface) {
	db := helper.TxGetDefault(ctx, r.db).Where(payload)

	err := db.Order("created_at desc").Find(&res).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}

func (r *TransactionRepository) TransactionCreate(ctx *gin.Context, payload *models.Transaction) (terr terror.ErrInterface) {
	db := helper.TxGetDefault(ctx, r.db)

	err := db.Create(payload).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}

func (r *TransactionRepository) TransactionUpdate(ctx *gin.Context, payload *models.Transaction) (terr terror.ErrInterface) {
	db := helper.TxGetDefault(ctx, r.db)

	err := db.Updates(payload).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}

func (r *TransactionRepository) TopupCreate(ctx *gin.Context, payload *models.Topup) (terr terror.ErrInterface) {
	db := helper.TxGetDefault(ctx, r.db)

	err := db.Create(payload).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}

func (r *TransactionRepository) PaymentCreate(ctx *gin.Context, payload *models.Payment) (terr terror.ErrInterface) {
	db := helper.TxGetDefault(ctx, r.db)

	err := db.Create(payload).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}
