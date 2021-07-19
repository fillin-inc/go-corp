package request

import "net/url"

const HOST = "api.houjin-bangou.nta.go.jp"

const API_VER = 4

const RESPONSE_TYPE = "12"

type URLBuilder interface {
	Validate() error
	URL() (url.URL, error)
}
