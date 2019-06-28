package gomegakassa

import (
	"fmt"
)

type Payment struct {
	ShopID uint32

	Description string

	Currency string
	Amount   float32

	OrderID  string
	MethodID string

	ClientEmail string
	ClientPhone string

	Debug string

	Language string

	Params Params

	shop *Shop
}

type PaymentConfig struct {
	Description string

	Currency string
	Amount   float32

	OrderID  string
	MethodID *string

	ClientEmail *string
	ClientPhone *string

	Debug bool

	Language string

	Params Params
}

func (p *Payment) Sign() string {
	if p.shop == nil {
		return ""
	}

	return md5(fmt.Sprintf(
		"%s%s",
		p.shop.SecretKey,
		md5(fmt.Sprintf("%d:%.2f:%s:%s:%s:%s:%s:%s:%s",
			p.shop.ID,
			p.Amount,
			p.Currency,
			p.Description,
			p.OrderID,
			p.MethodID,
			p.ClientEmail,
			p.Debug,
			p.shop.SecretKey,
		)),
	))
}

func (p *Payment) Raw() map[string]string {
	if p.shop == nil {
		return map[string]string{}
	}

	m := map[string]string{
		"shop_id": fmt.Sprintf("%d", p.shop.ID),

		"description": p.Description,

		"currency": p.Currency,
		"amount":   fmt.Sprintf("%.2f", p.Amount),

		"order_id":  p.OrderID,
		"method_id": p.MethodID,

		"client_email": p.ClientEmail,
		"client_phone": p.ClientPhone,

		"debug": p.Debug,

		"signature": p.Sign(),

		"language": p.Language,
	}

	for k, v := range p.Params {
		m["p_"+k] = v
	}

	return m
}
