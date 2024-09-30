package models

import (
	"time"

	"github.com/DeniesKresna/mncteststep2/types/constants"
	"gorm.io/gorm"
)

type Wallet struct {
	ID        int64          `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	UserID    string         `json:"user_id"`
	Balance   int64          `json:"balance"`
}

type Transaction struct {
	ID            int64                       `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time                   `json:"created_at"`
	UpdatedAt     time.Time                   `json:"updated_at"`
	DeletedAt     gorm.DeletedAt              `gorm:"index" json:"deleted_at"`
	UserID        string                      `json:"user_id"`
	Status        constants.TransactionStatus `json:"status"`
	ServiceID     string                      `json:"service_id"`
	ServiceName   constants.ServiceName       `json:"service"`
	Type          constants.TransactionType   `json:"type"`
	Amount        int64                       `json:"amount"`
	Remarks       string                      `json:"remarks"`
	BalanceBefore int64                       `json:"balance_before"`
	BalanceAfter  int64                       `json:"balance_after"`
}

type Payment struct {
	ID        int64          `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	PaymentID string         `gorm:"index" json:"payment_id"`
	Amount    int64          `json:"source"`
	Remarks   string         `json:"remarks"`
}

type Topup struct {
	ID        int64          `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	TopupID   string         `gorm:"index" json:"topup_id"`
	Amount    int64          `json:"amount"`
	UserID    string         `json:"user_id"`
}

type Transfer struct {
	ID                        int64                   `gorm:"primarykey" json:"id"`
	CreatedAt                 time.Time               `json:"created_at"`
	UpdatedAt                 time.Time               `json:"updated_at"`
	DeletedAt                 gorm.DeletedAt          `gorm:"index" json:"deleted_at"`
	TransferID                string                  `json:"transfer_id"`
	Type                      constants.PaymentType   `json:"type"`
	Status                    constants.PaymentStatus `json:"status"`
	TransferUserSourceID      string                  `json:"source"`
	TransferUserDestinationID string                  `json:"destination"`
	Amount                    int64                   `json:"amount"`
	Remarks                   string                  `json:"remarks"`
}

// request

type WalletTopupRequest struct {
	Amount int64 `json:"amount"`
}

type WalletTopupResponse struct {
	TopupID       string    `json:"topup_id"`
	AmountTopup   int64     `json:"amount_top_up"`
	BalanceBefore int64     `json:"balance_before"`
	BalanceAfter  int64     `json:"balance_after"`
	CreatedDate   time.Time `json:"created_date"`
}

type PaymentCreateRequest struct {
	Amount  int64  `json:"amount"`
	Remarks string `json:"remarks"`
}

type PaymentCreateResponse struct {
	PaymentID     string    `json:"payment_id"`
	Amount        int64     `json:"amount"`
	Remarks       string    `json:"remarks"`
	BalanceBefore int64     `json:"balance_before"`
	BalanceAfter  int64     `json:"balance_after"`
	CreatedDate   time.Time `json:"created_date"`
}
