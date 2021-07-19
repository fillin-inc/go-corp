package request

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
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

func TestBuildURL(t *testing.T) {
	appId := "ABCDEFG"
	nums := []uint64{1234, 5678}
	history := true
	number := &Number{appId, nums, RESPONSE_TYPE, history}

	ret, _ := number.BuildURL()
	u, _ := url.Parse(ret)

	if !strings.HasPrefix(ret, URL) {
		t.Errorf("URL が誤っています。result:%v expected:%v", ret, URL)
	}

	path := fmt.Sprintf("/%d/%s", API_VER, NUMBER_END_POINT)
	if u.Path != path {
		t.Errorf("path が誤っています。result:%v expected:%v", u.Path, path)
	}

	query := u.Query()
	if query["id"][0] != appId {
		t.Errorf("id の値が異なります。result:%v expected:%v", query["id"][0], appId)
	}

	if query["number"][0] != "1234,5678" {
		t.Errorf("number の値が異なります。result:%v expected:%v", query["number"][0], "1234,5678")
	}

	if query["type"][0] != RESPONSE_TYPE {
		t.Errorf("type の値が異なります。result:%v expected:%v", query["type"], RESPONSE_TYPE)
	}

	if query["history"][0] != "1" {
		t.Errorf("history の値が異なります。result:%v expected:%v", query["history"], "1")
	}
}
