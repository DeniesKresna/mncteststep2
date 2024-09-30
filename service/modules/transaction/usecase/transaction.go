package usecase

import (
	"github.com/DeniesKresna/mncteststep2/service/extensions/helper"
	"github.com/DeniesKresna/mncteststep2/service/extensions/terror"
	"github.com/DeniesKresna/mncteststep2/types/constants"
	"github.com/DeniesKresna/mncteststep2/types/models"
	"github.com/gin-gonic/gin"
)

func (u *TransactionUsecase) WalletGetByUserID(ctx *gin.Context, userID string, lock bool) (res models.Wallet, terr terror.ErrInterface) {
	res.UserID = userID
	terr = u.transactionRepo.WalletGet(ctx, &res, lock)
	return
}

func (u *TransactionUsecase) WalletCreate(ctx *gin.Context, payload models.Wallet) (res models.Wallet, terr terror.ErrInterface) {
	helper.TxCreate(ctx, u.transactionRepo.GetDB)
	defer func() {
		helper.TxSubmitTerr(ctx, terr)
	}()

	payload.Balance = 0
	terr = u.transactionRepo.WalletCreate(ctx, &payload)
	res = payload
	return
}

func (u *TransactionUsecase) WalletUpdate(ctx *gin.Context, payload models.Wallet) (res models.Wallet, terr terror.ErrInterface) {
	helper.TxCreate(ctx, u.transactionRepo.GetDB)
	defer func() {
		helper.TxSubmitTerr(ctx, terr)
	}()

	terr = u.transactionRepo.WalletUpdate(ctx, &res)
	return
}

func (u *TransactionUsecase) WalletTopup(ctx *gin.Context, amount int64) (res models.Transaction, terr terror.ErrInterface) {
	helper.TxCreate(ctx, u.transactionRepo.GetDB)
	defer func() {
		helper.TxSubmitTerr(ctx, terr)
	}()

	userID := ctx.GetString(constants.CONTEXT_USER_ID)

	topupPayload := models.Topup{
		TopupID: helper.CreateUUID(),
		Amount:  amount,
		UserID:  userID,
	}

	terr = u.transactionRepo.TopupCreate(ctx, &topupPayload)
	if terr != nil {
		return
	}

	res, terr = u.TransactionCreate(ctx, models.Transaction{
		UserID:      userID,
		Status:      constants.TRANSACTION_STATUS_SUCCESS,
		ServiceName: constants.SERVICE_NAME_TOPUP,
		ServiceID:   topupPayload.TopupID,
		Type:        constants.TRANSACTION_TYPE_CREDIT,
		Amount:      amount,
		Remarks:     "",
	})
	return
}

func (u *TransactionUsecase) PaymentCreate(ctx *gin.Context, payload models.PaymentCreateRequest) (res models.Transaction, terr terror.ErrInterface) {
	helper.TxCreate(ctx, u.transactionRepo.GetDB)
	defer func() {
		helper.TxSubmitTerr(ctx, terr)
	}()

	userID := ctx.GetString(constants.CONTEXT_USER_ID)

	paymentPayload := models.Payment{
		PaymentID: helper.CreateUUID(),
		Amount:    payload.Amount,
		Remarks:   payload.Remarks,
	}

	terr = u.transactionRepo.PaymentCreate(ctx, &paymentPayload)
	if terr != nil {
		return
	}

	res, terr = u.TransactionCreate(ctx, models.Transaction{
		UserID:      userID,
		Status:      constants.TRANSACTION_STATUS_SUCCESS,
		ServiceName: constants.SERVICE_NAME_PAYMENT,
		ServiceID:   paymentPayload.PaymentID,
		Type:        constants.TRANSACTION_TYPE_DEBIT,
		Amount:      payload.Amount,
		Remarks:     payload.Remarks,
	})
	return
}

func (u *TransactionUsecase) TransactionCreate(ctx *gin.Context, payload models.Transaction) (res models.Transaction, terr terror.ErrInterface) {
	helper.TxCreate(ctx, u.transactionRepo.GetDB)
	defer func() {
		helper.TxSubmitTerr(ctx, terr)
	}()

	res.Status = constants.TRANSACTION_STATUS_PENDING

	wallet, terr := u.WalletGetByUserID(ctx, payload.UserID, true)
	if terr != nil {
		return
	}

	payload.BalanceBefore = wallet.Balance

	if payload.Type == constants.TRANSACTION_TYPE_CREDIT {
		wallet.Balance += payload.Amount
	} else if payload.Type == constants.TRANSACTION_TYPE_DEBIT {
		if payload.BalanceBefore < payload.Amount {
			terr = terror.ErrInvalidRule("Balance is not enough")
			return
		}
		wallet.Balance -= payload.Amount
	} else {
		terr = terror.ErrParameter("transaction type not found")
		return
	}
	payload.BalanceAfter = wallet.Balance

	terr = u.transactionRepo.WalletUpdate(ctx, &wallet)
	if terr != nil {
		return
	}

	payload.Status = constants.TRANSACTION_STATUS_SUCCESS
	terr = u.transactionRepo.TransactionCreate(ctx, &payload)
	if terr != nil {
		return
	}
	res = payload

	return
}

func (u *TransactionUsecase) TransactionList(ctx *gin.Context) (res []models.Transaction, terr terror.ErrInterface) {
	userID := ctx.GetString(constants.CONTEXT_USER_ID)

	return u.transactionRepo.TransactionList(ctx, models.Transaction{
		UserID: userID,
	})
}
