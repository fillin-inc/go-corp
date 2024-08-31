package corp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/fillin-inc/go-corp/request"
)

var (
	testFillinCorpNum uint64 = 5070001032626
	testGunmaCorpNum  uint64 = 7000020100005
)

func TestByNumber(t *testing.T) {
	t.Run("Basic Usage", func(t *testing.T) {
		ts := testServer(http.StatusOK, "./testdata/response/by_number.xml")
		defer ts.Close()

		SetAppID("your-token")
		setTestEnvToRequest(ts)

		res, err := ByNumber(testFillinCorpNum)
		if err != nil {
			t.Errorf("error! %v", err)
		}

		if res.Count != 1 {
			t.Errorf("count value is wrong. result:%d expected:%d", res.Count, 1)
		}

		if len(res.Corporations) != 1 {
			t.Errorf("corporations length is wtong. result:%d expected:%d", len(res.Corporations), 1)
		}

		if res.Corporations[0].CorporateNumber != testFillinCorpNum {
			t.Errorf("corporation name is wtong. result:%d expected:%d", res.Corporations[0].CorporateNumber, testFillinCorpNum)
		}

		if res.Corporations[0].Name != "株式会社フィルイン" {
			t.Errorf("corporation name is wtong. result:%s expected:%s", res.Corporations[0].Name, "株式会社フィルイン")
		}
	})

	t.Run("Multiple Corporate Numbers", func(t *testing.T) {
		ts := testServer(http.StatusOK, "./testdata/response/by_numbers.xml")
		defer ts.Close()

		SetAppID("your-token")
		setTestEnvToRequest(ts)

		res, err := ByNumber(testFillinCorpNum, testGunmaCorpNum)
		if err != nil {
			t.Errorf("error! %v", err)
		}

		if res.Count != 2 {
			t.Errorf("count value is wrong. result:%d expected:%d", res.Count, 2)
		}

		if len(res.Corporations) != 2 {
			t.Errorf("corporations length is wtong. result:%d expected:%d", len(res.Corporations), 2)
		}

		if res.Corporations[0].CorporateNumber != testFillinCorpNum {
			t.Errorf("corporation name is wtong. result:%d expected:%d", res.Corporations[0].CorporateNumber, testFillinCorpNum)
		}

		if res.Corporations[0].Name != "株式会社フィルイン" {
			t.Errorf("corporation name is wtong. result:%s expected:%s", res.Corporations[0].Name, "株式会社フィルイン")
		}

		if res.Corporations[1].CorporateNumber != testGunmaCorpNum {
			t.Errorf("corporation name is wtong. result:%d expected:%d", res.Corporations[1].CorporateNumber, testGunmaCorpNum)
		}

		if res.Corporations[1].Name != "群馬県" {
			t.Errorf("corporation name is wtong. result:%s expected:%s", res.Corporations[1].Name, "群馬県")
		}
	})

	t.Run("Errors", func(t *testing.T) {
		tests := []struct {
			name          string
			statusCode    int
			contentType   string
			content       string
			expectedError string
		}{
			{
				"HTTP_Status_400",
				http.StatusBadRequest,
				"application/csv",
				"042,法人番号は10件以内で指定してください。",
				"042:法人番号は10件以内で指定してください。",
			},
			{
				"HTTP_Status_403",
				http.StatusForbidden,
				"text/html",
				"",
				"同一アプリケーションIDで一定期間内に多数のアクセスが実行されたため制限されています。",
			},
			{
				"HTTP_Status_404",
				http.StatusNotFound,
				"text/html",
				"",
				"アプリケーションIDが登録されていないまたは無効です。",
			},
			{
				"HTTP_Status_500",
				http.StatusInternalServerError,
				"text/html",
				"",
				"法人番号システム Web-API に問題が発生しています。",
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				ts := testErrorServer(test.statusCode, test.contentType, test.content)
				defer ts.Close()

				SetAppID("your-token")
				setTestEnvToRequest(ts)

				_, err := ByNumber(testFillinCorpNum)
				if err == nil {
					t.Errorf("No error occurred.")
				}

				if strings.TrimSpace(err.Error()) != test.expectedError {
					t.Log(err.Error())
					t.Log(test.expectedError)
				}
			})
		}
	})
}

func TestByNumberWithHistory(t *testing.T) {
	t.Run("Basic Usage", func(t *testing.T) {
		ts := testServer(http.StatusOK, "./testdata/response/by_number_with_history.xml")
		defer ts.Close()

		SetAppID("your-token")
		setTestEnvToRequest(ts)

		res, err := ByNumberWithHistory(testFillinCorpNum)
		if err != nil {
			t.Errorf("error! %v", err)
		}

		if res.Count != 3 {
			t.Errorf("count value is wrong. result:%d expected:%d", res.Count, 3)
		}

		if len(res.Corporations) != 3 {
			t.Errorf("corporations length is wtong. result:%d expected:%d", len(res.Corporations), 3)
		}

		postCodes := []string{"3700849", "3700813", "3700069"}
		for i, postCode := range postCodes {
			corp := res.Corporations[i]
			if corp.CorporateNumber != testFillinCorpNum {
				t.Errorf("corporation name is wtong. result:%d expected:%d", corp.CorporateNumber, testFillinCorpNum)
			}

			if res.Corporations[i].PostCode != postCode {
				t.Errorf("corporation name is wtong. result:%s expected:%s", corp.PostCode, postCode)
			}
		}
	})

	t.Run("Errors", func(t *testing.T) {
		tests := []struct {
			name          string
			statusCode    int
			contentType   string
			content       string
			expectedError string
		}{
			{
				"HTTP_Status_400",
				http.StatusBadRequest,
				"application/csv",
				"042,法人番号は10件以内で指定してください。",
				"042:法人番号は10件以内で指定してください。",
			},
			{
				"HTTP_Status_403",
				http.StatusForbidden,
				"text/html",
				"",
				"同一アプリケーションIDで一定期間内に多数のアクセスが実行されたため制限されています。",
			},
			{
				"HTTP_Status_404",
				http.StatusNotFound,
				"text/html",
				"",
				"アプリケーションIDが登録されていないまたは無効です。",
			},
			{
				"HTTP_Status_500",
				http.StatusInternalServerError,
				"text/html",
				"",
				"法人番号システム Web-API に問題が発生しています。",
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				ts := testErrorServer(test.statusCode, test.contentType, test.content)
				defer ts.Close()

				SetAppID("your-token")
				setTestEnvToRequest(ts)

				_, err := ByNumber(testFillinCorpNum)
				if err == nil {
					t.Errorf("No error occurred.")
				}

				if strings.TrimSpace(err.Error()) != test.expectedError {
					t.Errorf("Unexpected error received: %s, expected: %s", err.Error(), test.expectedError)
				}
			})
		}
	})
}

func TestDiffSearch(t *testing.T) {
	t.Run("Basic Usage", func(t *testing.T) {
		ts := testServer(http.StatusOK, "./testdata/response/diff_search.xml")
		defer ts.Close()

		SetAppID("your-token")
		setTestEnvToRequest(ts)

		res, err := DiffSearch("2021-06-09", "2021-06-09", "10202")
		if err != nil {
			t.Errorf("error! %v", err)
		}

		if res.Count != 1 {
			t.Errorf("count value is wrong. result:%d expected:%d", res.Count, 1)
		}

		if len(res.Corporations) != 1 {
			t.Errorf("corporations length is wtong. result:%d expected:%d", len(res.Corporations), 1)
		}

		if res.Corporations[0].CorporateNumber != testFillinCorpNum {
			t.Errorf("corporation name is wtong. result:%d expected:%d", res.Corporations[0].CorporateNumber, testFillinCorpNum)
		}

		if res.Corporations[0].Name != "株式会社フィルイン" {
			t.Errorf("corporation name is wtong. result:%s expected:%s", res.Corporations[0].Name, "株式会社フィルイン")
		}
	})

	t.Run("Errors", func(t *testing.T) {
		tests := []struct {
			name          string
			from          string
			to            string
			address       string
			statusCode    int
			contentType   string
			content       string
			expectedError string
		}{
			{
				"HTTP_Status_400",
				"2015-06-09",
				"2015-06-09",
				"10202",
				http.StatusBadRequest,
				"application/csv",
				"013,取得期間開始日は2015-12-01以降を指定してください。",
				"013:取得期間開始日は2015-12-01以降を指定してください。",
			},
			{
				"HTTP_Status_403",
				"2021-06-09",
				"2021-06-09",
				"10202",
				http.StatusForbidden,
				"text/html",
				"",
				"同一アプリケーションIDで一定期間内に多数のアクセスが実行されたため制限されています。",
			},
			{
				"HTTP_Status_404",
				"2021-06-09",
				"2021-06-09",
				"10202",
				http.StatusNotFound,
				"text/html",
				"",
				"アプリケーションIDが登録されていないまたは無効です。",
			},
			{
				"HTTP_Status_500",
				"2021-06-09",
				"2021-06-09",
				"10202",
				http.StatusInternalServerError,
				"text/html",
				"",
				"法人番号システム Web-API に問題が発生しています。",
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				ts := testErrorServer(test.statusCode, test.contentType, test.content)
				defer ts.Close()

				SetAppID("your-token")
				setTestEnvToRequest(ts)

				_, err := DiffSearch(test.from, test.to, test.address)
				if err == nil {
					t.Errorf("No error occurred.")
				}

				if strings.TrimSpace(err.Error()) != test.expectedError {
					t.Errorf("Unexpected error received: %s, expected: %s", err.Error(), test.expectedError)
				}
			})
		}
	})
}

func TestNameSearch(t *testing.T) {
	t.Run("Basic Usage", func(t *testing.T) {
		ts := testServer(http.StatusOK, "./testdata/response/name_search.xml")
		defer ts.Close()

		SetAppID("your-token")
		setTestEnvToRequest(ts)

		res, err := NameSearch("フィルイン", "10202")
		if err != nil {
			t.Errorf("error! %v", err)
		}

		if res.Count != 1 {
			t.Errorf("count value is wrong. result:%d expected:%d", res.Count, 1)
		}

		if len(res.Corporations) != 1 {
			t.Errorf("corporations length is wtong. result:%d expected:%d", len(res.Corporations), 1)
		}

		if res.Corporations[0].CorporateNumber != testFillinCorpNum {
			t.Errorf("corporation name is wtong. result:%d expected:%d", res.Corporations[0].CorporateNumber, testFillinCorpNum)
		}

		if res.Corporations[0].Name != "株式会社フィルイン" {
			t.Errorf("corporation name is wtong. result:%s expected:%s", res.Corporations[0].Name, "株式会社フィルイン")
		}
	})

	t.Run("Errors", func(t *testing.T) {
		tests := []struct {
			name          string
			searchName    string
			address       string
			statusCode    int
			contentType   string
			content       string
			expectedError string
		}{
			{
				"HTTP_Status_400",
				"フィルイン",
				"10202",
				http.StatusBadRequest,
				"application/csv",
				"101,商号又は名称には（全角文字|半角英数字記号）をutf-8でエンコードして設定してください。。",
				"101:商号又は名称には（全角文字|半角英数字記号）をutf-8でエンコードして設定してください。。",
			},
			{
				"HTTP_Status_403",
				"フィルイン",
				"10202",
				http.StatusForbidden,
				"text/html",
				"",
				"同一アプリケーションIDで一定期間内に多数のアクセスが実行されたため制限されています。",
			},
			{
				"HTTP_Status_404",
				"フィルイン",
				"10202",
				http.StatusNotFound,
				"text/html",
				"",
				"アプリケーションIDが登録されていないまたは無効です。",
			},
			{
				"HTTP_Status_500",
				"フィルイン",
				"10202",
				http.StatusInternalServerError,
				"text/html",
				"",
				"法人番号システム Web-API に問題が発生しています。",
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				ts := testErrorServer(test.statusCode, test.contentType, test.content)
				defer ts.Close()

				SetAppID("your-token")
				setTestEnvToRequest(ts)

				_, err := NameSearch(test.searchName, test.address)
				if err == nil {
					t.Errorf("No error occurred.")
				}

				if strings.TrimSpace(err.Error()) != test.expectedError {
					t.Errorf("Unexpected error received: %s, expected: %s", err.Error(), test.expectedError)
				}
			})
		}
	})
}

func TestSetAppID(t *testing.T) {
	tkn := "1234567890"
	SetAppID(tkn)

	if appID != tkn {
		t.Errorf("token is wrong. result:%s expected:%s", appID, tkn)
	}
}

func testServer(statusCode int, xmlPath string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(statusCode)
		data, _ := os.ReadFile(xmlPath)
		_, err := w.Write(data)
		if err != nil {
			panic(err)
		}
	}))
}

func testErrorServer(statusCode int, contentType string, content string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, content)
	}))
}

func setTestEnvToRequest(ts *httptest.Server) {
	u, _ := url.ParseRequestURI(ts.URL)
	request.Scheme = u.Scheme
	request.Host = u.Host
}

func ExampleByNumber() {
	// テスト環境用の設定
	// 実際に利用する際には不要
	ts := testServer(http.StatusOK, "./testdata/response/by_number.xml")
	setTestEnvToRequest(ts)
	defer ts.Close()

	// 法人番号 Web-API アプリケーションIDを設定
	SetAppID("your-token")

	// 株式会社フィルインの法人番号
	var corpNum uint64 = 5070001032626

	res, _ := ByNumber(corpNum)
	fmt.Println(res.Corporations[0].Name)
	// Output: 株式会社フィルイン
}

func ExampleDiffSearch() {
	// テスト環境用の設定
	// 実際に利用する際には不要
	ts := testServer(http.StatusOK, "./testdata/response/diff_search.xml")
	setTestEnvToRequest(ts)
	defer ts.Close()

	// 法人番号 Web-API アプリケーションIDを設定
	SetAppID("your-token")

	from := "2021-06-09"
	to := "2021-06-09"
	// 群馬県高崎市に限定
	address := "10202"
	res, _ := DiffSearch(from, to, address)
	fmt.Println(res.Corporations[0].Name)
	// Output: 株式会社フィルイン
}

func ExampleNameSearch() {
	// テスト環境用の設定
	// 実際に利用する際には不要
	ts := testServer(http.StatusOK, "./testdata/response/name_search.xml")
	setTestEnvToRequest(ts)
	defer ts.Close()

	// 法人番号 Web-API アプリケーションIDを設定
	SetAppID("your-token")

	name := "フィルイン"
	// 群馬県高崎市に限定
	address := "10202"
	res, _ := NameSearch(name, address)
	fmt.Println(res.Corporations[0].Name)
	// Output: 株式会社フィルイン
}
