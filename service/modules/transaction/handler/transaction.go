package handler

import (
	"github.com/DeniesKresna/mncteststep2/service/extensions/terror"
	"github.com/DeniesKresna/mncteststep2/types/constants"
	"github.com/DeniesKresna/mncteststep2/types/models"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func (h TransactionHandler) WalletTopup(ctx *gin.Context) {
	var (
		payload models.WalletTopupRequest
		terr    terror.ErrInterface
	)

	terr = BindAndValidate(ctx, &payload, govalidator.MapData{
		"amount": []string{"required", "numeric", "min:1"},
	})
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}

	res, terr := h.transactionUsecase.WalletTopup(ctx, payload.Amount)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}

	resp := models.WalletTopupResponse{
		TopupID:       res.ServiceID,
		AmountTopup:   res.Amount,
		BalanceBefore: res.BalanceBefore,
		BalanceAfter:  res.BalanceAfter,
		CreatedDate:   res.CreatedAt,
	}
	ResponseJson(ctx, resp)
}

func (h TransactionHandler) TransactionList(ctx *gin.Context) {
	var (
		terr terror.ErrInterface
	)

	txs, terr := h.transactionUsecase.TransactionList(ctx)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}

	var dataStack []map[string]interface{}
	for _, tx := range txs {
		var data = make(map[string]interface{})
		serviceKey := "transfer_id"
		if tx.ServiceName == constants.SERVICE_NAME_PAYMENT {
			serviceKey = "payment_id"
		} else if tx.ServiceName == constants.SERVICE_NAME_TOPUP {
			serviceKey = "top_up_id"
		}
		data[serviceKey] = tx.ServiceID
		data["status"] = tx.Status
		data["user_id"] = tx.UserID
		data["transaction_type"] = tx.Type
		data["amount"] = tx.Amount
		data["remarks"] = tx.Remarks
		data["balance_before"] = tx.BalanceBefore
		data["balance_after"] = tx.BalanceAfter
		data["created_at"] = tx.CreatedAt

		dataStack = append(dataStack, data)
	}

	ResponseJson(ctx, dataStack)
}

func (h TransactionHandler) PaymentCreate(ctx *gin.Context) {
	var (
		payload models.PaymentCreateRequest
		terr    terror.ErrInterface
	)

	terr = BindAndValidate(ctx, &payload, govalidator.MapData{
		"amount":  []string{"required", "numeric", "min:1"},
		"remarks": []string{"required"},
	})
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}

	res, terr := h.transactionUsecase.PaymentCreate(ctx, payload)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}

	resp := models.PaymentCreateResponse{
		PaymentID:     res.ServiceID,
		Amount:        res.Amount,
		Remarks:       res.Remarks,
		BalanceBefore: res.BalanceBefore,
		BalanceAfter:  res.BalanceAfter,
		CreatedDate:   res.CreatedAt,
	}
	ResponseJson(ctx, resp)
}
