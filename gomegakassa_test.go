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
	// TODO
}
