/*
法人番号システム Web-API のリクエスト URL パッケージです。

法人番号システム Web-API(https://www.houjin-bangou.nta.go.jp/webapi/) 対する
リクエスト URL オブジェクトを生成します。

関連ドキュメント:

・API 仕様書「Web-API(Ver.4.0)のリクエストの設定方法及び提供データの内容について」
https://www.houjin-bangou.nta.go.jp/documents/k-web-api-kinou-ver4.pdf

・API 仕様書「Web-API(Ver.1.1)のリクエストの設定方法及び提供データの内容について」
https://www.houjin-bangou.nta.go.jp/documents/k-web-api-kinou-gaiyo.pdf

・都道府県コード: https://nlftp.mlit.go.jp/ksj/gml/codelist/PrefCd.html

・都道府県コード+市区町村コード: https://www.soumu.go.jp/denshijiti/code.html
*/
package request

import (
	"net/url"

	"github.com/go-playground/validator"
)

// 法人番号システム Web-API バージョン
const API_VER = 4

/*
応答形式: XML形式/Unicode(JIS第一水準から第四水準)に固定
*/
const RESPONSE_TYPE = "12"

var (
	// Web-API Scheme
	Scheme = "https"
	// Web-API Host
	Host     = "api.houjin-bangou.nta.go.jp"
	validate = validator.New()
)

type URLBuilder interface {
	Validate() error
	URL() (url.URL, error)
}

func init() {
	vals := map[string]func(fl validator.FieldLevel) bool{
		"date":        dateValidation,
		"gtedate":     dateEqualOrGreaterValidation,
		"address":     addressValidation,
		"kind":        kindValidation,
		"checkdigits": checkdigitsValidation,
	}

	for name, f := range vals {
		err := validate.RegisterValidation(name, f)
		if err != nil {
			panic(err)
		}
	}
}
