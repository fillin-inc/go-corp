package request

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

/*
法人名指定検索

Web-API 仕様書の「法人名を指定して情報を取得する機能について」記載の
リクエスト URL 生成を行います。
*/
type Name struct {
	// 法人番号システム Web-API アプリケーションID
	ID string `validate:"required" url:"id"`
	// 商号または名称
	Name string `validate:"required" url:"name"`
	// 検索方式
	// 1:前方一致形式, 2:部分一致形式
	Mode int `validate:"min=1,max=2" url:"mode,omitempty"`
	// 検索対象
	// 1:JIS第一・第二水準(あいまい検索), 2:JIS第一〜第四水準(完全一致検索), 3:英語表記(英語表記登録情報検索)
	Target int `validate:"min=1,max=3" url:"target"`
	// 所在地
	// 空文字, 都道府県コード(2桁)または都道府県コード+市区町村コード(5桁)
	Address string `validate:"address" url:"address,omitempty"`
	// 法人種別
	// 01:国の機関, 02:地方公共団体, 03:設立法人登記, 04:外国会社等・その他
	Kind []string `validate:"max=4,kind" url:"kind,omitempty" del:","`
	// 変更履歴
	Change bool `url:"change,int"`
	// 閉鎖登記取得
	Close bool `url:"close,int"`
	// 法人番号指定年月日開始日
	// 空文字またはYYYY-MM-DD 形式の文字列
	From string `validate:"date" url:"from,omitempty"`
	// 法人番号指定年月日終了日
	// 空文字またはYYYY-MM-DD 形式の文字列
	To string `validate:"date,gtedate=From" url:"to,omitempty"`
	// 分割番号
	// 1〜99999
	Divide int `validate:"min=1,max=99999" url:"divide"`
	// 応答形式
	ResponseType string `validate:"required,eq=12" url:"type"`
}

// Name 生成
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

// バリデーション
func (n Name) Validate() error {
	return validate.Struct(n)
}

/*
URL 生成

バリデーション処理が必要な場合は別途 Validate メソッドを実行してください。
*/
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
