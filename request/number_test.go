package request

import (
	"errors"
	"reflect"
	"testing"

	"github.com/go-playground/validator"
)

func TestNewNumber(t *testing.T) {
	appID := "ABCDEFG"
	nums := []uint64{1234, 5678}
	history := true

	number := NewNumber(appID, nums, history)
	if number.ID != appID {
		t.Error("ID フィールドの値が一致しません。")
	}

	if !reflect.DeepEqual(number.Numbers, nums) {
		t.Error("Numbers フィールドの値が一致しません。")
	}

	if number.History != history {
		t.Error("History フィールドの値が一致しません。")
	}
}

func TestValidate(t *testing.T) {
	n := Number{
		ID:           "ABCDEFG",
		Numbers:      []uint64{1234, 5678},
		ResponseType: RESPONSE_TYPE,
		History:      true,
	}

	if err := n.Validate(); err != nil {
		t.Errorf("バリデーションエラーが発生しました。%v", err)
	}
}

func TestValidateReturnErrorWhenAppIDIsEmpty(t *testing.T) {
	n := Number{
		ID:           "",
		Numbers:      []uint64{1234, 5678},
		ResponseType: RESPONSE_TYPE,
		History:      false,
	}

	err := n.Validate()

	if errors.Is(err, &validator.ValidationErrors{}) {
		t.Errorf("エラーの型情報が異なります。result:%T expected:validator.ValidationErrors", err)
	}

	if err.Error() != "Key: 'Number.ID' Error:Field validation for 'ID' failed on the 'required' tag" {
		t.Errorf("エラーメッセージが異なります。%v", err.Error())
	}
}

func TestValidateReturnErrorWhenNumbersIsEmpty(t *testing.T) {
	n := Number{
		ID:           "ABCDEFG",
		Numbers:      []uint64{},
		ResponseType: RESPONSE_TYPE,
		History:      false,
	}

	err := n.Validate()
	if err.Error() != "Key: 'Number.Numbers' Error:Field validation for 'Numbers' failed on the 'min' tag" {
		t.Errorf("エラーメッセージが異なります。%v", err.Error())
	}
}

func TestValidateReturnNoErrorWhenNumbersIs1(t *testing.T) {
	n := Number{
		ID:           "ABCDEFG",
		Numbers:      []uint64{1234},
		ResponseType: RESPONSE_TYPE,
		History:      false,
	}

	if err := n.Validate(); err != nil {
		t.Errorf("バリデーションエラーが発生しました。%v", err)
	}
}

func TestValidateReturnNoErrorWhenNumbersIs10(t *testing.T) {
	n := Number{
		ID:           "ABCDEFG",
		Numbers:      []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		ResponseType: RESPONSE_TYPE,
		History:      false,
	}

	if err := n.Validate(); err != nil {
		t.Errorf("バリデーションエラーが発生しました。%v", err)
	}
}

func TestValidateReturnErrorWhenNumbersIs11(t *testing.T) {
	n := Number{
		ID:           "ABCDEFG",
		Numbers:      []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		ResponseType: RESPONSE_TYPE,
		History:      false,
	}

	err := n.Validate()
	if err.Error() != "Key: 'Number.Numbers' Error:Field validation for 'Numbers' failed on the 'max' tag" {
		t.Errorf("エラーメッセージが異なります。%v", err.Error())
	}
}

func TestValidateReturnErrorWhenResponseTypeIsEmpty(t *testing.T) {
	n := Number{
		ID:           "ABCDEFG",
		Numbers:      []uint64{1234},
		ResponseType: "",
		History:      false,
	}

	err := n.Validate()
	if err.Error() != "Key: 'Number.ResponseType' Error:Field validation for 'ResponseType' failed on the 'required' tag" {
		t.Errorf("エラーメッセージが異なります。%v", err.Error())
	}
}

func TestValidateReturnErrorWhenResponseTypeIsInvalid(t *testing.T) {
	n := Number{
		ID:           "ABCDEFG",
		Numbers:      []uint64{1234},
		ResponseType: "01",
		History:      false,
	}

	err := n.Validate()
	if err.Error() != "Key: 'Number.ResponseType' Error:Field validation for 'ResponseType' failed on the 'eq' tag" {
		t.Errorf("エラーメッセージが異なります。%v", err.Error())
	}
}

func TestURL(t *testing.T) {
	appID := "ABCDEFG"
	nums := []uint64{1234, 5678}
	history := true
	number := &Number{appID, nums, RESPONSE_TYPE, history}

	u, err := number.URL()
	if err != nil {
		t.Errorf("エラーが発生しました。%v", err)
	}

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
