package request

import (
	"github.com/go-playground/validator"
	"github.com/google/go-querystring/query"
)

// 法人番号検索
type Number struct {
	// アプリケーションID
	Id string `validate:"required" url:"id"`
	// 法人番号
	Numbers []uint64 `validate:"required,min=1" url:"number" del:","`
	// 応答形式
	ResponseType string `validate:"required,eq=12" url:"type"`
	// 変更履歴要否
	History bool `validate:"required" url:"history,int"`
}

func NewNumber(appId string, numbers []uint64, history bool) *Number {
	return &Number{
		appId,
		numbers,
		RESPONSE_TYPE,
		history,
	}
}

func (n Number) queryString() (string, error) {
	v, err := query.Values(n)
	if err != nil {
		return "", err
	}
	return v.Encode(), err
}

func (n Number) BuildURL() (string, error) {
	validator := validator.New()
	err := validator.Struct(n)
	if err != nil {
		return "", err
	}

	q, err := n.queryString()
	if err != nil {
		return "", err
	}
	return requestURL(NUMBER_END_POINT) + "?" + q, nil
}
