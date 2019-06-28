package gomegakassa

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

var ErrNoRequiredParam = errors.New("required params not provided")
var ErrBadSignature = errors.New("bad signature")

var wrapParseError = func(paramName string, err error) error {
	return errors.New(paramName + " parse error: " + err.Error())
}

var requiredFormParams = []string{
	"uid", "amount", "amount_shop", "amount_client", "currency", "order_id",
	"payment_method_title", "creation_time", "client_email", "status", "signature",
}

type Shop struct {
	ID uint32

	SecretKey string
}

func (s *Shop) NewPayment(cfg *PaymentConfig) *Payment {
	p := &Payment{
		ShopID: s.ID,

		Description: cfg.Description,

		Currency: cfg.Currency,
		Amount:   cfg.Amount,

		OrderID: cfg.OrderID,

		Language: cfg.Language,

		Params: cfg.Params,

		shop: s,
	}

	if cfg.MethodID != nil {
		p.MethodID = *cfg.MethodID
	}

	if cfg.ClientEmail != nil {
		p.ClientEmail = *cfg.ClientEmail
	}

	if cfg.ClientPhone != nil {
		p.ClientPhone = *cfg.ClientPhone
	}

	if cfg.Debug {
		p.Debug = "1"
	}

	if p.Params == nil {
		p.Params = map[string]string{}
	}

	return p
}

func (s *Shop) Verify(formParams map[string]string) (*Notification, error) {
	for _, v := range requiredFormParams {
		_, ok := formParams[v]
		if !ok {
			return nil, ErrNoRequiredParam
		}
	}

	sign := md5(fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s:%s:%s:%s:%s:%s:%s:%s",
		formParams["uid"], formParams["amount"], formParams["amount_shop"], formParams["amount_client"],
		formParams["currency"], formParams["order_id"], formParams["payment_method_id"], formParams["payment_method_title"],
		formParams["creation_time"], formParams["payment_time"], formParams["client_email"], formParams["status"],
		formParams["debug"], s.SecretKey,
	))

	if sign != formParams["signature"] {
		return nil, ErrBadSignature
	}

	amount, err := strconv.ParseFloat(formParams["amount"], 32)
	if err != nil {
		return nil, wrapParseError("amount", err)
	}

	amountShop, err := strconv.ParseFloat(formParams["amount_shop"], 32)
	if err != nil {
		return nil, wrapParseError("amount_shop", err)
	}

	amountClient, err := strconv.ParseFloat(formParams["amount_client"], 32)
	if err != nil {
		return nil, wrapParseError("amount_client", err)
	}

	paymentMethodID, err := strconv.ParseInt(formParams["payment_method_id"], 32, 8)
	if err != nil {
		return nil, wrapParseError("payment_method_id", err)
	}

	createdAt, err := time.Parse(time.RFC3339, formParams["creation_time"])
	if err != nil {
		return nil, wrapParseError("creation_time", err)
	}

	var paidAt *time.Time
	if formParams["payment_time"] != "" {
		cPaidAt, err := time.Parse(time.RFC3339, formParams["payment_time"])
		if err != nil {
			return nil, wrapParseError("payment_time", err)
		}

		paidAt = &cPaidAt
	}

	n := &Notification{
		UID: formParams["uid"],

		Amount:       float32(amount),
		AmountShop:   float32(amountShop),
		AmountClient: float32(amountClient),

		Currency: formParams["currency"],

		OrderID: formParams["order_id"],

		PaymentMethodID:    int(paymentMethodID),
		PaymentMethodTitle: formParams["payment_method_title"],

		ClientEmail: formParams["client_email"],

		Debug: formParams["debug"] == "1",

		Params: map[string]string{},

		CreatedAt: createdAt,
		PaidAt:    paidAt,
	}

	for k, v := range formParams {
		found := false

		for _, vF := range requiredFormParams {
			if v == vF {
				found = true
				break
			}
		}

		if !found {
			n.Params[k] = v
		}
	}

	return n, nil
}
