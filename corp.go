/*
法人番号システム Web-API 用パッケージです。

法人番号システム Web-API(https://www.houjin-bangou.nta.go.jp/webapi/) 対し
リクエストを行い法人情報を取得します。

利用には次のリンクからアプリケーション ID の申請と取得が必要です。
https://www.houjin-bangou.nta.go.jp/webapi/riyo-todokede/

関連ドキュメント:

・API 仕様書「Web-API(Ver.4.0)のリクエストの設定方法及び提供データの内容について」
https://www.houjin-bangou.nta.go.jp/documents/k-web-api-kinou-ver4.pdf

・API 仕様書「Web-API(Ver.1.1)のリクエストの設定方法及び提供データの内容について」
https://www.houjin-bangou.nta.go.jp/documents/k-web-api-kinou-gaiyo.pdf

・都道府県コード: https://nlftp.mlit.go.jp/ksj/gml/codelist/PrefCd.html

・都道府県コード+市区町村コード: https://www.soumu.go.jp/denshijiti/code.html
*/
package corp

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/fillin-inc/go-corp/request"
)

var (
	// 法人番号 Web-API アプリケーション ID
	appID string

	fetch = func(URL string, options interface{}) (int, []byte, error) {
		var body []byte

		res, err := http.Get(URL)
		if err != nil {
			return http.StatusInternalServerError, body, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		return res.StatusCode, body, err
	}
)

// ByNumber は法人番号を引数に指定することで最新の法人情報を取得できます。
func ByNumber(numbers ...uint64) (Response, error) {
	builder := request.NewNumber(appID, numbers, false)
	return responseByURLBuilder(builder)
}

/*
ByNumberWithHistory は法人番号を引数に指定することで変更履歴を含む法人情報を取得できます。

「変更履歴」とは, 例えば本店所在地を 1 度変更している法人の場合には
変更前と変更後の 2 つの法人情報が取得できます。
*/
func ByNumberWithHistory(numbers ...uint64) (Response, error) {
	builder := request.NewNumber(appID, numbers, true)
	return responseByURLBuilder(builder)
}

/*
DiffSearch は対象期間と地域で変更があった法人情報を検索します。

from, to は YYYY-MM-DD 形式の日付文字列を指定してください。

address は空文字, 「都道府県コード」(2文字)または「都道府県コード+市区町村コード」(5文字)を
指定できます。空文字の場合, from, to に指定した期間のみで検索を行います。

各コードについては次のリンクを参照してください。

・都道府県コード: https://nlftp.mlit.go.jp/ksj/gml/codelist/PrefCd.html

・都道府県コード+市区町村コード: https://www.soumu.go.jp/denshijiti/code.html
*/
func DiffSearch(from string, to string, address string) (Response, error) {
	builder := request.NewDiff(appID, from, to, address, []string{}, 1)
	return responseByURLBuilder(builder)
}

/*
NameSearch は法人名と地域で法人情報を検索します。

このメソッドでは法人名を部分一致のあいまい検索で探します。

address は空文字, 「都道府県コード」(2文字)または「都道府県コード+市区町村コード」(5文字)を
指定できます。空文字の場合, address に指定した法人名のみで検索を行います。

各コードについては次のリンクを参照してください。

・都道府県コード: https://nlftp.mlit.go.jp/ksj/gml/codelist/PrefCd.html

・都道府県コード+市区町村コード: https://www.soumu.go.jp/denshijiti/code.html
*/
func NameSearch(name string, address string) (Response, error) {
	builder := request.NewName(appID, name, 2, 1, address, []string{}, false, true, "", "", 1)
	return responseByURLBuilder(builder)
}

// SetAppID は法人番号 Web-API のアクセスに必要なアプリケーション ID を設定します。
func SetAppID(tkn string) {
	appID = tkn
}

/*
SetFetch は法人番号 Web-API からデータ取得処理を設定します。

標準では単純な fetch 処理が利用可能です。ログ処理など特別な事情がある場合に利用してください。
*/
func SetFetch(f func(URL string, options interface{}) (int, []byte, error)) {
	fetch = f
}

func responseByURLBuilder(builder request.URLBuilder) (Response, error) {
	if err := builder.Validate(); err != nil {
		return Response{}, err
	}

	u, err := builder.URL()
	if err != nil {
		return Response{}, err
	}

	var statusCode int
	var body []byte
	var res Response
	statusCode, body, err = fetch(u.String(), nil)
	if err != nil {
		return res, err
	}

	// エラー情報を取得した場合
	// Web-API 仕様書「HTTPステータスコード、エラーコード及びエラーメッセージ一覧」参照
	if statusCode == http.StatusBadRequest {
		str := string(body)
		strs := strings.Split(str, ",")
		if len(strs) == 2 {
			return res, fmt.Errorf("%s:%s", strs[0], strs[1])
		}
		return res, fmt.Errorf(str)
	}
	if statusCode == http.StatusForbidden {
		return res, fmt.Errorf(
			"同一アプリケーションIDで一定期間内に多数のアクセスが実行されたため制限されています。",
		)
	}
	if statusCode == http.StatusNotFound {
		return res, fmt.Errorf("アプリケーションIDが登録されていないまたは無効です。")
	}
	if statusCode == http.StatusInternalServerError {
		return res, fmt.Errorf("法人番号システム Web-API に問題が発生しています。")
	}

	err = xml.Unmarshal(body, &res)
	return res, err
}
