package request

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

/*
期間指定検索

Web-API 仕様書の「取得期間を指定して情報を取得する機能について」記載の
リクエスト URL 生成を行います。
*/
type Diff struct {
	// 法人番号システム Web-API アプリケーションID
	ID string `validate:"required" url:"id"`
	// 取得期間開始日
	// YYYY-MM-DD 形式の文字列
	From string `validate:"required,date" url:"from"`
	// 取得期間終了日
	// YYYY-MM-DD 形式の文字列
	To string `validate:"required,date,gtedate=From" url:"to"`
	// 所在地
	// 空文字, 都道府県コード(2桁)または都道府県コード+市区町村コード(5桁)
	Address string `validate:"address" url:"address,omitempty"`
	// 法人種別
	// 01:国の機関, 02:地方公共団体, 03:設立法人登記, 04:外国会社等・その他
	Kind []string `validate:"max=4,kind" url:"kind,omitempty" del:","`
	// 分割番号
	// 1〜99999
	Divide int `validate:"min=1,max=99999" url:"devide"`
	// 応答形式
	ResponseType string `validate:"required,eq=12" url:"type"`
}

// Diff 生成
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

// バリデーション
func (d Diff) Validate() error {
	return validate.Struct(d)
}

/*
URL 生成

バリデーション処理が必要な場合は別途 Validate メソッドを実行してください。
*/
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
