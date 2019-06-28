package gomegakassa

import (
	"time"
)

type Notification struct {
	UID string

	Amount       float32
	AmountShop   float32
	AmountClient float32

	Currency string

	OrderID string

	PaymentMethodID    int
	PaymentMethodTitle string

	ClientEmail string

	Debug bool

	Params Params

	CreatedAt time.Time
	PaidAt    *time.Time
}

func (n *Notification) GetParam(name string) string {
	v, ok := n.Params["p_" + name]
	if !ok {
		return ""
	}

	return v
}
