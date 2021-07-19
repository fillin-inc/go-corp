package request

import (
	"reflect"
	"testing"
)

func TestNewNumber(t *testing.T) {
	appId := "ABCDEFG"
	nums := []uint64{1234, 5678}
	history := true

	number := NewNumber(appId, nums, history)
	if !reflect.DeepEqual(number.Numbers, nums) {
		t.Error("Numbers フィールドの値が一致しません。")
	}

	if number.History != history {
		t.Error("History フィールドの値が一致しません。")
	}
}

func TestValidateWithoutError(t *testing.T) {
	appId := "ABCDEFG"
	nums := []uint64{1234, 5678}
	history := true

	number := NewNumber(appId, nums, history)
	err := number.Validate()
	if err != nil {
		t.Errorf("バリデーションでエラーが発生しました。%v", err)
	}
}

func TestValidateWithError(t *testing.T) {
	appId := ""
	nums := []uint64{1234, 5678}
	history := true

	number := NewNumber(appId, nums, history)
	err := number.Validate()
	if err == nil {
		t.Errorf("バリデーションでエラーが発生しませんでした。%v", err)
	}
}

func TestURL(t *testing.T) {
	appId := "ABCDEFG"
	nums := []uint64{1234, 5678}
	history := true
	number := &Number{appId, nums, RESPONSE_TYPE, history}

	u, _ := number.URL()

	if u.Scheme != "https" {
		t.Errorf("Scheme は https 限定です。%v", u.Scheme)
	}

	if u.Host != HOST {
		t.Errorf("Host が異なります。%v", u.Host)
	}

	if u.Path != "/4/num" {
		t.Errorf("Path が異なります。%v", u.Path)
	}

	query := "history=1&id=ABCDEFG&number=1234%2C5678&type=12"
	if u.RawQuery != query {
		t.Errorf("Query が異なります。%v", u.RawQuery)
	}

	url := "https://" + HOST + "/4/num?" + query
	if u.String() != url {
		t.Errorf("URL文字列が異なります。%v", u.String())
	}
}
