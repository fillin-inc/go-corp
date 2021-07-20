package request

import (
	"reflect"
	"testing"
)

type diffValidationErrTest struct {
	Diff   Diff
	ErrMsg string
}

type diffURLTest struct {
	Diff Diff
	URL  string
}

func TestNewDiff(t *testing.T) {
	appID := "ABCDEFG"
	from := "2021-07-19"
	to := "2021-07-19"
	address := "10202"
	kind := []string{"03"}
	divide := 2

	diff := NewDiff(appID, from, to, address, kind, divide)

	if diff.ID != appID {
		t.Error("ID フィールドの値が一致しません。")
	}

	if diff.From != from {
		t.Error("From フィールドの値が一致しません。")
	}

	if diff.To != to {
		t.Error("To フィールドの値が一致しません。")
	}

	if diff.Address != address {
		t.Error("Address フィールドの値が一致しません。")
	}

	if !reflect.DeepEqual(diff.Kind, kind) {
		t.Error("Kind フィールドの値が一致しません。")
	}

	if diff.Divide != divide {
		t.Error("Divide フィールドの値が一致しません。")
	}

	if diff.ResponseType != RESPONSE_TYPE {
		t.Error("ResponseType フィールドの値が一致しません。")
	}
}

func TestDiffValidate(t *testing.T) {
	tests := []Diff{
		{
			ID:           "ABCDEFG",
			From:         "2021-07-19",
			To:           "2021-07-19",
			Address:      "",
			Kind:         []string{},
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// Address has PrefCode
		{
			ID:           "ABCDEFG",
			From:         "2021-07-19",
			To:           "2021-07-19",
			Address:      "10",
			Kind:         []string{},
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// Address has PrefCode + CityCode
		{
			ID:           "ABCDEFG",
			From:         "2021-07-19",
			To:           "2021-07-19",
			Address:      "10202",
			Kind:         []string{},
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// Kind has 1 KindCode
		{
			ID:           "ABCDEFG",
			From:         "2021-07-19",
			To:           "2021-07-19",
			Address:      "",
			Kind:         []string{"03"},
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// Kind has 2 KindCode
		{
			ID:           "ABCDEFG",
			From:         "2021-07-19",
			To:           "2021-07-19",
			Address:      "",
			Kind:         []string{"01", "03"},
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// Divide has max value(99999)
		{
			ID:           "ABCDEFG",
			From:         "2021-07-19",
			To:           "2021-07-19",
			Address:      "",
			Kind:         []string{},
			Divide:       99999,
			ResponseType: RESPONSE_TYPE,
		},
	}

	for i, diff := range tests {
		if err := diff.Validate(); err != nil {
			t.Errorf("%d: Validation Error %v", i, err)
		}
	}
}

func TestDiffValidateError(t *testing.T) {
	tests := []diffValidationErrTest{
		{
			// ID is empty
			Diff{
				ID:           "",
				From:         "2021-07-19",
				To:           "2021-07-19",
				Address:      "",
				Kind:         []string{},
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Diff.ID' Error:Field validation for 'ID' failed on the 'required' tag",
		},
		{
			// From is empty
			Diff{
				ID:           "ABCDEFG",
				From:         "",
				To:           "2021-07-19",
				Address:      "",
				Kind:         []string{},
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Diff.From' Error:Field validation for 'From' failed on the 'required' tag",
		},
		{
			// From is invalid format
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07",
				To:           "2021-07-19",
				Address:      "",
				Kind:         []string{},
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Diff.From' Error:Field validation for 'From' failed on the 'date' tag"},
		{
			// TO is empty
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "",
				Address:      "",
				Kind:         []string{},
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Diff.To' Error:Field validation for 'To' failed on the 'required' tag",
		},
		{
			// To is invalid format
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "2021-08",
				Address:      "",
				Kind:         []string{},
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Diff.To' Error:Field validation for 'To' failed on the 'date' tag",
		},
		{
			// To is past than From
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "2021-07-18",
				Address:      "",
				Kind:         []string{},
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Diff.To' Error:Field validation for 'To' failed on the 'gtedate' tag",
		},
		{
			// Address is invalid PrefCode
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "2021-07-19",
				Address:      "48",
				Kind:         []string{},
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Diff.Address' Error:Field validation for 'Address' failed on the 'address' tag",
		},
		{
			// Address is invalid PrefCode + CityCode
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "2021-07-19",
				Address:      "00202",
				Kind:         []string{},
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Diff.Address' Error:Field validation for 'Address' failed on the 'address' tag",
		},
		{
			// Address is invalid format
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "2021-07-19",
				Address:      "TEST1",
				Kind:         []string{},
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Diff.Address' Error:Field validation for 'Address' failed on the 'address' tag",
		},
		{
			// Address is invalid CityCode
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "2021-07-19",
				Address:      "99TES",
				Kind:         []string{},
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Diff.Address' Error:Field validation for 'Address' failed on the 'address' tag",
		},
		{
			// Kind contains invalid KindCode
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "2021-07-19",
				Address:      "",
				Kind:         []string{"05", "03"},
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Diff.Kind' Error:Field validation for 'Kind' failed on the 'kind' tag",
		},
		{
			// Divide is less than min value(1)
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "2021-07-19",
				Address:      "",
				Kind:         []string{},
				Divide:       0,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Diff.Divide' Error:Field validation for 'Divide' failed on the 'min' tag",
		},
		{
			// Divide is greater than max value(99999)
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "2021-07-19",
				Address:      "",
				Kind:         []string{},
				Divide:       100000,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Diff.Divide' Error:Field validation for 'Divide' failed on the 'max' tag",
		},
		{
			// ResponseType is empty
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "2021-07-19",
				Address:      "",
				Kind:         []string{},
				Divide:       1,
				ResponseType: "",
			},
			"Key: 'Diff.ResponseType' Error:Field validation for 'ResponseType' failed on the 'required' tag",
		},
		{
			// ResponseType is invalid
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "2021-07-19",
				Address:      "",
				Kind:         []string{},
				Divide:       1,
				ResponseType: "01",
			},
			"Key: 'Diff.ResponseType' Error:Field validation for 'ResponseType' failed on the 'eq' tag",
		},
	}

	for i, test := range tests {
		err := test.Diff.Validate()
		if err == nil {
			t.Errorf("%d: Validation Error not returns", i)
		}

		if err.Error() != test.ErrMsg {
			t.Errorf("%d: Validation Error Message not matched result:%s expected:%s", i, err.Error(), test.ErrMsg)
		}
	}
}

func TestDiffURL(t *testing.T) {
	tests := []diffURLTest{
		{
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "2021-07-20",
				Address:      "",
				Kind:         []string{},
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"https://api.houjin-bangou.nta.go.jp/4/diff?devide=1&from=2021-07-19&id=ABCDEFG&to=2021-07-20&type=12",
		},
		{
			// Address is specified
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "2021-07-20",
				Address:      "10202",
				Kind:         []string{},
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"https://api.houjin-bangou.nta.go.jp/4/diff?address=10202&devide=1&from=2021-07-19&id=ABCDEFG&to=2021-07-20&type=12",
		},
		{
			// Kind is provided 1 KindCode
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "2021-07-20",
				Address:      "",
				Kind:         []string{"03"},
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"https://api.houjin-bangou.nta.go.jp/4/diff?devide=1&from=2021-07-19&id=ABCDEFG&kind=03&to=2021-07-20&type=12",
		},
		{
			// Kind is provided 3 KindCode
			Diff{
				ID:           "ABCDEFG",
				From:         "2021-07-19",
				To:           "2021-07-20",
				Address:      "",
				Kind:         []string{"01", "02", "03"},
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"https://api.houjin-bangou.nta.go.jp/4/diff?devide=1&from=2021-07-19&id=ABCDEFG&kind=01%2C02%2C03&to=2021-07-20&type=12",
		},
	}

	for i, test := range tests {
		url, _ := test.Diff.URL()

		if url.String() != test.URL {
			t.Errorf("%d: URL String not match result:%s expected:%s", i, url.String(), test.URL)
		}
	}
}
