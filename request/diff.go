package request

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// 期間検索
type Diff struct {
	// アプリケーションID
	ID string `validate:"required" url:"id"`
	// 取得期間開始日
	From string `validate:"required,date" url:"from"`
	// 取得期間終了日
	To string `validate:"required,date,gtedate=From" url:"to"`
	// 所在地
	Address string `validate:"address" url:"address,omitempty"`
	// 法人種別
	Kind []string `validate:"max=4,kind" url:"kind,omitempty" del:","`
	// 分割番号
	Divide int `validate:"min=1,max=99999" url:"devide"`
	// 応答形式
	ResponseType string `validate:"required,eq=12" url:"type"`
}

func NewDiff(appID string, from string, to string, address string, kind []string, divide int) *Diff {
	return &Diff{
		appID,
		from,
		to,
		address,
		kind,
		divide,
		RESPONSE_TYPE,
	}
}

func (d Diff) Validate() error {
	return validate.Struct(d)
}

func (d Diff) URL() (url.URL, error) {
	var u url.URL
	q, err := query.Values(d)
	if err != nil {
		return u, fmt.Errorf("failed to convert query string: %w", err)
	}

	return url.URL{
		Scheme:   Scheme,
		Host:     Host,
		Path:     fmt.Sprintf("/%d/diff", API_VER),
		RawQuery: q.Encode(),
	}, nil
}
