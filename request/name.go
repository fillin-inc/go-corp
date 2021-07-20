package request

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// 法人名検索
type Name struct {
	// アプリケーションID
	ID string `validate:"required" url:"id"`
	// 商号または名称
	Name string `validate:"required" url:"name"`
	// 検索方式
	Mode int `validate:"min=1,max=2" url:"mode,omitempty"`
	// 検索対象
	Target int `validate:"min=1,max=3" url:"target"`
	// 所在地
	Address string `validate:"address" url:"address,omitempty"`
	// 法人種別
	Kind []string `validate:"max=4,kind" url:"kind,omitempty" del:","`
	// 変更履歴
	Change bool `url:"change,int"`
	// 閉鎖登記取得
	Close bool `url:"close,int"`
	// 法人番号指定年月日開始日
	From string `validate:"date" url:"from,omitempty"`
	// 法人番号指定年月日終了日
	To string `validate:"date,gtedate=From" url:"to,omitempty"`
	// 分割番号
	Divide int `validate:"min=1,max=99999" url:"divide"`
	// 応答形式
	ResponseType string `validate:"required,eq=12" url:"type"`
}

func NewName(appID string, name string, mode int, target int, address string, kind []string, change bool, close bool, from string, to string, divide int) *Name {
	return &Name{
		appID,
		name,
		mode,
		target,
		address,
		kind,
		change,
		close,
		from,
		to,
		divide,
		RESPONSE_TYPE,
	}
}

func (n Name) Validate() error {
	return validate.Struct(n)
}

func (n Name) URL() (url.URL, error) {
	var u url.URL
	q, err := query.Values(n)
	if err != nil {
		return u, fmt.Errorf("failed to convert query string: %w", err)
	}

	return url.URL{
		Scheme:   Scheme,
		Host:     Host,
		Path:     fmt.Sprintf("/%d/name", API_VER),
		RawQuery: q.Encode(),
	}, nil
}
