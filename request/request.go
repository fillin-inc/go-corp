package request

import (
	"net/url"

	"github.com/go-playground/validator"
)

const HOST = "api.houjin-bangou.nta.go.jp"

const API_VER = 4

const RESPONSE_TYPE = "12"

var (
	validate *validator.Validate
)

type URLBuilder interface {
	Validate() error
	URL() (url.URL, error)
}

func init() {
	validate = validator.New()
	validate.RegisterValidation("date", dateValidation)
	validate.RegisterValidation("gtedate", dateEqualOrGreaterValidation)
	validate.RegisterValidation("address", addressValidation)
	validate.RegisterValidation("kind", kindValidation)
}
