package gomegakassa

import (
	"fmt"
	"testing"
)

var mk = New(
	1,
	"0123456789abcdef",
)

func TestPayment(t *testing.T) {
	p := mk.NewPayment(&PaymentConfig{
		Description: "iPhone 8 plus 32 Gb",

		Currency: "RUB",
		Amount:   100.50,

		OrderID: "123456",

		Language: "ru",
	})

	sign := p.Sign()
	signMustBe := "ac1cbfe5be0a124e20316ea5165b6e15"

	if sign != signMustBe {
		fmt.Println(p)

		t.Errorf("sign isn't valid (%s != %s)", sign, signMustBe)
	}

	p.Amount = 322.22
	p.Params["user_id"] = "qwerty"

	sign = p.Sign()
	signMustBe = "8c651a3d63949e65b7fa3cd1f4bca77d"

	if sign != signMustBe {
		fmt.Println(p)

		t.Errorf("sign isn't valid (%s != %s)", sign, signMustBe)
	}
}

func TestNotification(t *testing.T) {
	formParams := map[string]string{
		"uid":                  "1",
		"amount":               "122",
		"amount_shop":          "120",
		"amount_client":        "125",
		"currency":             "RUB",
		"order_id":             "1",
		"payment_method_id":    "1",
		"payment_method_title": "QIWI",
		"creation_time":        "2005-08-09T00:00:00Z",
		"payment_time":         "2006-08-09T00:00:00Z",
		"client_email":         "test@test.ru",
		"status":               "success",
		"debug":                "1",
		"p_user_id":            "23",
		// sign must be
		"signature": "79a34b4b5cc604f79e74fe2a88682240",
	}

	notification, err := mk.Verify(formParams)
	if err != nil {
		t.Errorf(err.Error())

		return
	}

	userID := notification.GetParam("user_id")
	userIDMustBe := "23"

	if userID != userIDMustBe {
		t.Errorf("can't get user_id from notification (%s != %s)", userID, userIDMustBe)

		return
	}
}
