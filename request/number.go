package request

import (
	"fmt"
	"net/url"

	"github.com/go-playground/validator"
	"github.com/google/go-querystring/query"
)

// 法人番号検索
type Number struct {
	// アプリケーションID
	ID string `validate:"required" url:"id"`
	// 法人番号
	Numbers []uint64 `validate:"min=1,max=10" url:"number" del:","`
	// 応答形式
	ResponseType string `validate:"required,eq=12" url:"type"`
	// 変更履歴要否
	History bool `url:"history,int"`
}

func NewNumber(appID string, numbers []uint64, history bool) *Number {
	return &Number{
		appID,
		numbers,
		RESPONSE_TYPE,
		history,
	}
}

func (n Number) Validate() error {
	validator := validator.New()
	return validator.Struct(n)
}

func (n Number) URL() (url.URL, error) {
	var u url.URL
	q, err := query.Values(n)
	if err != nil {
		return u, fmt.Errorf("failed to convert query string: %w", err)
	}

	return url.URL{
		Scheme:   "https",
		Host:     HOST,
		Path:     fmt.Sprintf("/%d/num", API_VER),
		RawQuery: q.Encode(),
	}, nil
}
