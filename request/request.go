package request

import (
	"net/url"

	"github.com/go-playground/validator"
)

const API_VER = 4

const RESPONSE_TYPE = "12"

var (
	Scheme   = "https"
	Host     = "api.houjin-bangou.nta.go.jp"
	validate = validator.New()
)

type URLBuilder interface {
	Validate() error
	URL() (url.URL, error)
}

func init() {
	vals := map[string]func(fl validator.FieldLevel) bool{
		"date":    dateValidation,
		"gtedate": dateEqualOrGreaterValidation,
		"address": addressValidation,
		"kind":    kindValidation,
	}

	for name, f := range vals {
		err := validate.RegisterValidation(name, f)
		if err != nil {
			panic(err)
		}
	}
}
