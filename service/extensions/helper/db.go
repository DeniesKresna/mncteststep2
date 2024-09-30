package helper

import (
	"errors"
	"reflect"

	"github.com/DeniesKresna/mncteststep2/service/extensions/terror"
	"github.com/DeniesKresna/mncteststep2/types/constants"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WrapPercentOnStructString(data interface{}) (err error) {
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		err = errors.New("input must be a pointer to a struct")
		return
	}

	structVal := val.Elem()

	for i := 0; i < structVal.NumField(); i++ {
		field := structVal.Field(i)

		if field.Kind() == reflect.String {
			currentValue := field.String()

			if currentValue != "" {
				uppercaseValue := WrapString(currentValue, "%")
				field.SetString(uppercaseValue)
			}
		}
	}

	return
}

type TransactionSetFunc = func(ctx *gin.Context) interface{}

func TxCreate(ctx *gin.Context, fn TransactionSetFunc) {
	if TxGet(ctx) == nil {
		txCtx := fn(ctx)
		var ok bool
		var tx *gorm.DB
		tx, ok = txCtx.(*gorm.DB)
		if !ok {
			return
		}
		ctx.Set(constants.TX_CTX_KEY, tx.Begin())
	}
}

func TxFlush(ctx *gin.Context) {
	if TxGet(ctx) != nil {
		ctx.Set(constants.TX_CTX_KEY, nil)
	}
}

func TxGet(ctx *gin.Context) (tx *gorm.DB) {
	txCtx, exist := ctx.Get(constants.TX_CTX_KEY)
	if !exist {
		return
	}
	var ok bool
	tx, ok = txCtx.(*gorm.DB)
	if !ok {
		return nil
	}
	return
}

func TxGetDefault(ctx *gin.Context, db *gorm.DB) (tx *gorm.DB) {
	tx = TxGet(ctx)
	if tx == nil {
		tx = db
	}
	return
}

func TxSubmitTerr(ctx *gin.Context, terr terror.ErrInterface) {
	if terr != nil {
		TxRollBack(ctx)
	} else {
		TxCommit(ctx)
	}
}

func TxCommit(ctx *gin.Context) {
	defer TxFlush(ctx)
	txCtx, exist := ctx.Get(constants.TX_CTX_KEY)
	if !exist {
		return
	}
	tx, ok := txCtx.(*gorm.DB)
	if !ok {
		return
	}
	tx.Commit()
}

func TxRollBack(ctx *gin.Context) {
	defer TxFlush(ctx)
	txCtx, exist := ctx.Get(constants.TX_CTX_KEY)
	if !exist {
		return
	}
	tx, ok := txCtx.(*gorm.DB)
	if !ok {
		return
	}
	tx.Rollback()
}
