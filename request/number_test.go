package request

import (
	"reflect"
	"testing"
)

type numberValidationErrTest struct {
	Number Number
	ErrMsg string
}

type numberURLTest struct {
	Number Number
	URL    string
}

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

func TestNumberValidate(t *testing.T) {
	tests := []Number{
		{
			ID:           "ABCDEFG",
			Numbers:      []uint64{1234},
			ResponseType: RESPONSE_TYPE,
			History:      true,
		},
		// Number has max slicer count(10)
		{
			ID:           "ABCDEFG",
			Numbers:      []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			ResponseType: RESPONSE_TYPE,
			History:      true,
		},
		// History is false
		{
			ID:           "ABCDEFG",
			Numbers:      []uint64{1234},
			ResponseType: RESPONSE_TYPE,
			History:      false,
		},
	}

	for i, number := range tests {
		if err := number.Validate(); err != nil {
			t.Errorf("%d: Validation Error %v", i, err)
		}
	}
}

func TestNumberValidateError(t *testing.T) {
	tests := []numberValidationErrTest{
		{
			Number{
				ID:           "",
				Numbers:      []uint64{1234},
				ResponseType: RESPONSE_TYPE,
				History:      false,
			},
			"Key: 'Number.ID' Error:Field validation for 'ID' failed on the 'required' tag",
		},
		{
			Number{
				ID:           "ABCDEFG",
				Numbers:      []uint64{},
				ResponseType: RESPONSE_TYPE,
				History:      false,
			},
			"Key: 'Number.Numbers' Error:Field validation for 'Numbers' failed on the 'min' tag",
		},
		{
			Number{
				ID:           "ABCDEFG",
				Numbers:      []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
				ResponseType: RESPONSE_TYPE,
				History:      false,
			},
			"Key: 'Number.Numbers' Error:Field validation for 'Numbers' failed on the 'max' tag",
		},
		{
			Number{
				ID:           "ABCDEFG",
				Numbers:      []uint64{1234},
				ResponseType: "",
				History:      false,
			},
			"Key: 'Number.ResponseType' Error:Field validation for 'ResponseType' failed on the 'required' tag",
		},
		{
			Number{
				ID:           "ABCDEFG",
				Numbers:      []uint64{1234},
				ResponseType: "01",
				History:      false,
			},
			"Key: 'Number.ResponseType' Error:Field validation for 'ResponseType' failed on the 'eq' tag",
		},
	}

	for i, test := range tests {
		err := test.Number.Validate()
		if err == nil {
			t.Errorf("%d: Validation Error not returns", i)
		}

		if err.Error() != test.ErrMsg {
			t.Errorf("%d: Validation Error Message not matched result:%s expected:%s", i, err.Error(), test.ErrMsg)
		}
	}
}

func TestURL(t *testing.T) {
	tests := []numberURLTest{
		{
			Number{
				ID:           "ABCDEFG",
				Numbers:      []uint64{1234},
				ResponseType: RESPONSE_TYPE,
				History:      true,
			},
			"https://api.houjin-bangou.nta.go.jp/4/num?history=1&id=ABCDEFG&number=1234&type=12",
		},
		{
			// Numbers is multiple value
			Number{
				ID:           "ABCDEFG",
				Numbers:      []uint64{1234, 5678},
				ResponseType: RESPONSE_TYPE,
				History:      true,
			},
			"https://api.houjin-bangou.nta.go.jp/4/num?history=1&id=ABCDEFG&number=1234%2C5678&type=12",
		},
		{
			// History is false
			Number{
				ID:           "ABCDEFG",
				Numbers:      []uint64{1234},
				ResponseType: RESPONSE_TYPE,
				History:      false,
			},
			"https://api.houjin-bangou.nta.go.jp/4/num?history=0&id=ABCDEFG&number=1234&type=12",
		},
	}

	for i, test := range tests {
		url, _ := test.Number.URL()

		if url.String() != test.URL {
			t.Errorf("%d: URL String not match result:%s expected:%s", i, url.String(), test.URL)
		}
	}
}
