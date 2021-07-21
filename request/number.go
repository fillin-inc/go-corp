package request

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

/*
法人番号指定検索

Web-API 仕様書の「法人番号を指定して情報を取得する機能について」記載の
リクエスト URL 生成を行います。
*/
type Number struct {
	// 法人番号システム Web-API アプリケーションID
	ID string `validate:"required" url:"id"`
	// 法人番号
	// 1 〜 10 個の法人番号
	Numbers []uint64 `validate:"min=1,max=10" url:"number" del:","`
	// 応答形式
	ResponseType string `validate:"required,eq=12" url:"type"`
	// 変更履歴要否
	History bool `url:"history,int"`
}

// Number 生成
func NewNumber(appID string, numbers []uint64, history bool) *Number {
	return &Number{
		appID,
		numbers,
		RESPONSE_TYPE,
		history,
	}
}

// バリデーション
func (n Number) Validate() error {
	return validate.Struct(n)
}

/*
URL 生成

バリデーション処理が必要な場合は別途 Validate メソッドを実行してください。
*/
func (n Number) URL() (url.URL, error) {
	var u url.URL
	q, err := query.Values(n)
	if err != nil {
		return u, fmt.Errorf("failed to convert query string: %w", err)
	}

	return url.URL{
		Scheme:   Scheme,
		Host:     Host,
		Path:     fmt.Sprintf("/%d/num", API_VER),
		RawQuery: q.Encode(),
	}, nil
}
