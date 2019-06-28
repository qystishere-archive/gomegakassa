package gomegakassa

import (
	"time"
)

type Notification struct {
	ID uint32

	Amount       float32
	AmountShop   float32
	AmountClient float32

	Currency string

	OrderID string

	PaymentMethodID    int
	PaymentMethodTitle string

	ClientEmail string

	Debug bool

	Params map[string]string

	CreatedAt time.Time
	PaidAt    *time.Time
}
