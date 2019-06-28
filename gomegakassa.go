package gomegakassa

type Params map[string]string

type GoMegakassa struct {
	*Shop
}

func New(shopID uint32, secretKey string) *GoMegakassa {
	return &GoMegakassa{
		Shop: &Shop{
			ID: shopID,

			SecretKey: secretKey,
		},
	}
}
