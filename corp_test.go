package corp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/fillin-inc/go-corp/request"
)

var (
	testFillinCorpNum uint64 = 5070001032626
	testGunmaCorpNum  uint64 = 7000020100005
)

func TestByNumber(t *testing.T) {
	ts := testServer("./testdata/response/by_number.xml")
	defer ts.Close()

	SetAppID("your-token")
	setTestEnvToRequest(ts)

	res, err := ByNumber(testFillinCorpNum)
	if err != nil {
		t.Errorf("error! %v", err)
	}

	if res.Count != 1 {
		t.Errorf("count valuw is wrong. result:%d expected:%d", res.Count, 1)
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
}

func TestByNumberWithHistory(t *testing.T) {
	ts := testServer("./testdata/response/by_number_with_history.xml")
	defer ts.Close()

	SetAppID("your-token")
	setTestEnvToRequest(ts)

	res, err := ByNumberWithHistory(testFillinCorpNum)
	if err != nil {
		t.Errorf("error! %v", err)
	}

	if res.Count != 3 {
		t.Errorf("count valuw is wrong. result:%d expected:%d", res.Count, 3)
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
}

func TestByNumbers(t *testing.T) {
	ts := testServer("./testdata/response/by_numbers.xml")
	defer ts.Close()

	SetAppID("your-token")
	setTestEnvToRequest(ts)

	res, err := ByNumbers([]uint64{testFillinCorpNum, testGunmaCorpNum})
	if err != nil {
		t.Errorf("error! %v", err)
	}

	if res.Count != 2 {
		t.Errorf("count valuw is wrong. result:%d expected:%d", res.Count, 2)
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
		t.Errorf("corporation name is wtong. result:%d expected:%d", res.Corporations[1].CorporateNumber, testFillinCorpNum)
	}

	if res.Corporations[1].Name != "群馬県" {
		t.Errorf("corporation name is wtong. result:%s expected:%s", res.Corporations[1].Name, "群馬県")
	}
}

func TestDiffSearch(t *testing.T) {
	ts := testServer("./testdata/response/diff_search.xml")
	defer ts.Close()

	SetAppID("your-token")
	setTestEnvToRequest(ts)

	res, err := DiffSearch("2021-06-09", "2021-06-09", "10202")
	if err != nil {
		t.Errorf("error! %v", err)
	}

	if res.Count != 1 {
		t.Errorf("count valuw is wrong. result:%d expected:%d", res.Count, 1)
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
}

func TestNameSearch(t *testing.T) {
	ts := testServer("./testdata/response/name_search.xml")
	defer ts.Close()

	SetAppID("your-token")
	setTestEnvToRequest(ts)

	res, err := NameSearch("フィルイン", "10202")
	if err != nil {
		t.Errorf("error! %v", err)
	}

	if res.Count != 1 {
		t.Errorf("count valuw is wrong. result:%d expected:%d", res.Count, 1)
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
}

func TestSetAppID(t *testing.T) {
	tkn := "1234567890"
	SetAppID(tkn)

	if appID != tkn {
		t.Errorf("token is wrong. result:%s expected:%s", appID, tkn)
	}
}

func testServer(xmlPath string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		data, _ := ioutil.ReadFile(xmlPath)
		_, err := w.Write(data)
		if err != nil {
			panic(err)
		}
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
	ts := testServer("./testdata/response/by_number.xml")
	setTestEnvToRequest(ts)
	defer ts.Close()

	// 法人番号 Web-API アプリケーションIDを設定
	SetAppID(os.Getenv("CORP_API_TOKEN"))

	// 株式会社フィルインの法人番号
	var corpNum uint64 = 5070001032626

	res, _ := ByNumber(corpNum)
	fmt.Println(res.Corporations[0].Name)
	// Output: 株式会社フィルイン
}

func ExampleDiffSearch() {
	// テスト環境用の設定
	// 実際に利用する際には不要
	ts := testServer("./testdata/response/diff_search.xml")
	setTestEnvToRequest(ts)
	defer ts.Close()

	// 法人番号 Web-API アプリケーションIDを設定
	SetAppID(os.Getenv("CORP_API_TOKEN"))

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
	ts := testServer("./testdata/response/name_search.xml")
	setTestEnvToRequest(ts)
	defer ts.Close()

	// 法人番号 Web-API アプリケーションIDを設定
	SetAppID(os.Getenv("CORP_API_TOKEN"))

	name := "フィルイン"
	// 群馬県高崎市に限定
	address := "10202"
	res, _ := NameSearch(name, address)
	fmt.Println(res.Corporations[0].Name)
	// Output: 株式会社フィルイン
}
