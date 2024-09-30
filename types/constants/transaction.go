package constants

type PaymentStatus string

func (p *PaymentStatus) ToString() string {
	return string(*p)
}

type PaymentType string

func (p *PaymentType) ToString() string {
	return string(*p)
}

type ServiceName string

func (p *ServiceName) ToString() string {
	return string(*p)
}

type TransactionType string

func (p *TransactionType) ToString() string {
	return string(*p)
}

type TransactionStatus string

func (p *TransactionStatus) ToString() string {
	return string(*p)
}

const (
	PAYMENT_STATUS_SUCCESS PaymentStatus = "SUCCESS"
	PAYMENT_STATUS_FAILED  PaymentStatus = "FAILED"
	PAYMENT_STATUS_PENDING PaymentStatus = "PENDING"

	PAYMENT_TYPE_DEBIT  PaymentType = "DEBIT"
	PAYMENT_TYPE_CREDIT PaymentType = "CREDIT"

	SERVICE_NAME_TOPUP    ServiceName = "TOPUP"
	SERVICE_NAME_PAYMENT  ServiceName = "PAYMENT"
	SERVICE_NAME_TRANSFER ServiceName = "TRANSFER"

	TRANSACTION_TYPE_DEBIT  TransactionType = "DEBIT"
	TRANSACTION_TYPE_CREDIT TransactionType = "CREDIT"

	TRANSACTION_STATUS_SUCCESS TransactionStatus = "SUCCESS"
	TRANSACTION_STATUS_FAILED  TransactionStatus = "FAILED"
	TRANSACTION_STATUS_PENDING TransactionStatus = "PENDING"
)
