package request

import (
	"fmt"
	"reflect"
	"testing"
)

type nameValidationErrTest struct {
	Name   Name
	ErrMsg string
}

type nameURLTest struct {
	Name Name
	URL  string
}

func TestNewName(t *testing.T) {
	appID := "you-token"
	name := "フィルイン"
	mode := 1
	target := 1
	address := "10202"
	kind := []string{"03"}
	change := false
	close := false
	from := "2016-09-01"
	to := "2016-09-10"
	divide := 2

	n := NewName(appID, name, mode, target, address, kind, change, close, from, to, divide)

	if n.ID != appID {
		t.Error("ID is not match")
	}

	if n.Name != name {
		t.Error("Name is not match")
	}

	if n.Mode != mode {
		t.Error("Mode is not match")
	}

	if n.Target != target {
		t.Error("Target is not match")
	}

	if n.Address != address {
		t.Error("Address is not match")
	}

	if !reflect.DeepEqual(n.Kind, kind) {
		t.Error("Kind is not match")
	}

	if n.Change != change {
		t.Error("Change is not match")
	}

	if n.Close != close {
		t.Error("Close is not match")
	}

	if n.From != from {
		t.Error("From is not match")
	}

	if n.To != to {
		t.Error("To is not match")
	}

	if n.Divide != divide {
		t.Error("Devide is not match")
	}

	if n.ResponseType != RESPONSE_TYPE {
		t.Error("ResponseType is not match")
	}
}

func TestNameValidate(t *testing.T) {
	tests := []Name{
		{
			ID:           "you-token",
			Name:         "フィルイン",
			Mode:         1,
			Target:       1,
			Address:      "",
			Kind:         []string{},
			Change:       false,
			Close:        false,
			From:         "",
			To:           "",
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// Mode is 2
		{
			ID:           "you-token",
			Name:         "フィルイン",
			Mode:         2,
			Target:       1,
			Address:      "",
			Kind:         []string{},
			Change:       false,
			Close:        false,
			From:         "",
			To:           "",
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// Target is 2
		{
			ID:           "you-token",
			Name:         "フィルイン",
			Mode:         1,
			Target:       2,
			Address:      "",
			Kind:         []string{},
			Change:       false,
			Close:        false,
			From:         "",
			To:           "",
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// Target is 3
		{
			ID:           "you-token",
			Name:         "フィルイン",
			Mode:         1,
			Target:       3,
			Address:      "",
			Kind:         []string{},
			Change:       false,
			Close:        false,
			From:         "",
			To:           "",
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// Address is PrefCode
		{
			ID:           "you-token",
			Name:         "フィルイン",
			Mode:         1,
			Target:       1,
			Address:      "10",
			Kind:         []string{},
			Change:       false,
			Close:        false,
			From:         "",
			To:           "",
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// Address has PrefCode + CityCode
		{
			ID:           "you-token",
			Name:         "フィルイン",
			Mode:         1,
			Target:       1,
			Address:      "10202",
			Kind:         []string{},
			Change:       false,
			Close:        false,
			From:         "",
			To:           "",
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// Address has 1 KindCode
		{
			ID:           "you-token",
			Name:         "フィルイン",
			Mode:         1,
			Target:       1,
			Address:      "",
			Kind:         []string{"03"},
			Change:       false,
			Close:        false,
			From:         "",
			To:           "",
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// Kind has 2 KindCode
		{
			ID:           "you-token",
			Name:         "フィルイン",
			Mode:         1,
			Target:       1,
			Address:      "",
			Kind:         []string{"01", "03"},
			Change:       false,
			Close:        false,
			From:         "",
			To:           "",
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// Change is true
		{
			ID:           "you-token",
			Name:         "フィルイン",
			Mode:         1,
			Target:       1,
			Address:      "",
			Kind:         []string{},
			Change:       true,
			Close:        false,
			From:         "",
			To:           "",
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// Close is true
		{
			ID:           "you-token",
			Name:         "フィルイン",
			Mode:         1,
			Target:       1,
			Address:      "",
			Kind:         []string{},
			Change:       false,
			Close:        true,
			From:         "",
			To:           "",
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// From has value
		{
			ID:           "you-token",
			Name:         "フィルイン",
			Mode:         1,
			Target:       1,
			Address:      "",
			Kind:         []string{},
			Change:       false,
			Close:        true,
			From:         "2016-09-01",
			To:           "",
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// To have value
		{
			ID:           "you-token",
			Name:         "フィルイン",
			Mode:         1,
			Target:       1,
			Address:      "",
			Kind:         []string{},
			Change:       false,
			Close:        true,
			From:         "",
			To:           "2016-09-10",
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// From and To have same value
		{
			ID:           "you-token",
			Name:         "フィルイン",
			Mode:         1,
			Target:       1,
			Address:      "",
			Kind:         []string{},
			Change:       false,
			Close:        true,
			From:         "2016-09-01",
			To:           "2016-09-01",
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// From and To have different value
		{
			ID:           "you-token",
			Name:         "フィルイン",
			Mode:         1,
			Target:       1,
			Address:      "",
			Kind:         []string{},
			Change:       false,
			Close:        true,
			From:         "2016-09-01",
			To:           "2016-09-10",
			Divide:       1,
			ResponseType: RESPONSE_TYPE,
		},
		// Divide has max value(99999)
		{
			ID:           "you-token",
			Name:         "フィルイン",
			Mode:         1,
			Target:       1,
			Address:      "",
			Kind:         []string{},
			Change:       false,
			Close:        true,
			From:         "",
			To:           "",
			Divide:       99999,
			ResponseType: RESPONSE_TYPE,
		},
	}

	for i, name := range tests {
		if err := name.Validate(); err != nil {
			t.Errorf("%d: Validation Error %v", i, err)
		}
	}
}

func TestNameValidateError(t *testing.T) {
	tests := []nameValidationErrTest{
		{
			// ID is empty
			Name{
				ID:           "",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.ID' Error:Field validation for 'ID' failed on the 'required' tag",
		},
		{
			// Name is empty
			Name{
				ID:           "you-token",
				Name:         "",
				Mode:         1,
				Target:       1,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.Name' Error:Field validation for 'Name' failed on the 'required' tag",
		},
		{
			// Mode is zero
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         0,
				Target:       1,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.Mode' Error:Field validation for 'Mode' failed on the 'min' tag",
		},
		{
			// Mode is 3
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         3,
				Target:       1,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.Mode' Error:Field validation for 'Mode' failed on the 'max' tag",
		},
		{
			// Target is zero
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       0,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.Target' Error:Field validation for 'Target' failed on the 'min' tag",
		},
		{
			// Target is 4
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       4,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.Target' Error:Field validation for 'Target' failed on the 'max' tag",
		},
		{
			// Address is invalid PrefCode
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "48",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.Address' Error:Field validation for 'Address' failed on the 'address' tag",
		},
		{
			// Address is invalid PrefCode + CityCode
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "00202",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.Address' Error:Field validation for 'Address' failed on the 'address' tag",
		},
		{
			// Address is invalid format
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "TEST1",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.Address' Error:Field validation for 'Address' failed on the 'address' tag",
		},
		{
			// Address is invalid CityCode
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "99TES",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.Address' Error:Field validation for 'Address' failed on the 'address' tag",
		},
		{
			// Kind contains invalid KindCode
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "",
				Kind:         []string{"05", "03"},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.Kind' Error:Field validation for 'Kind' failed on the 'kind' tag",
		},
		{
			// From is invalid format
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "2021-07",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.From' Error:Field validation for 'From' failed on the 'date' tag",
		},
		{
			// To is invalid format
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "2021-07",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.To' Error:Field validation for 'To' failed on the 'date' tag",
		},
		{
			// To is past than From
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "2021-07-19",
				To:           "2021-07-18",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.To' Error:Field validation for 'To' failed on the 'gtedate' tag",
		},
		{
			// Divide is less than min value(1)
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       0,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.Divide' Error:Field validation for 'Divide' failed on the 'min' tag",
		},
		{
			// Divide is greater than max value(99999)
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       100000,
				ResponseType: RESPONSE_TYPE,
			},
			"Key: 'Name.Divide' Error:Field validation for 'Divide' failed on the 'max' tag",
		},
		{
			// ResponseType is empty
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: "",
			},
			"Key: 'Name.ResponseType' Error:Field validation for 'ResponseType' failed on the 'required' tag",
		},
		{
			// ResponseType is invalid
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: "01",
			},
			"Key: 'Name.ResponseType' Error:Field validation for 'ResponseType' failed on the 'eq' tag",
		},
	}

	for i, test := range tests {
		err := test.Name.Validate()
		if err == nil {
			t.Errorf("%d: Validation Error not returns", i)
		}

		if err.Error() != test.ErrMsg {
			t.Errorf("%d: Validation Error Message not matched result:%s expected:%s", i, err.Error(), test.ErrMsg)
		}
	}
}

func TestNameURL(t *testing.T) {
	tests := []nameURLTest{
		{
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"https://api.houjin-bangou.nta.go.jp/4/name?change=0&close=0&divide=1&id=you-token&mode=1&name=%E3%83%95%E3%82%A3%E3%83%AB%E3%82%A4%E3%83%B3&target=1&type=12",
		},
		{
			// Address is specified
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "10202",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"https://api.houjin-bangou.nta.go.jp/4/name?address=10202&change=0&close=0&divide=1&id=you-token&mode=1&name=%E3%83%95%E3%82%A3%E3%83%AB%E3%82%A4%E3%83%B3&target=1&type=12",
		},
		{
			// Kind is provided 1 KindCode
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "",
				Kind:         []string{"03"},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"https://api.houjin-bangou.nta.go.jp/4/name?change=0&close=0&divide=1&id=you-token&kind=03&mode=1&name=%E3%83%95%E3%82%A3%E3%83%AB%E3%82%A4%E3%83%B3&target=1&type=12",
		},
		{
			// Kind is provided 3 KindCode
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "",
				Kind:         []string{"01", "02", "03"},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"https://api.houjin-bangou.nta.go.jp/4/name?change=0&close=0&divide=1&id=you-token&kind=01%2C02%2C03&mode=1&name=%E3%83%95%E3%82%A3%E3%83%AB%E3%82%A4%E3%83%B3&target=1&type=12",
		},
		{
			// From is specified
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "2021-07-19",
				To:           "",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"https://api.houjin-bangou.nta.go.jp/4/name?change=0&close=0&divide=1&from=2021-07-19&id=you-token&mode=1&name=%E3%83%95%E3%82%A3%E3%83%AB%E3%82%A4%E3%83%B3&target=1&type=12",
		},
		{
			// To is specified
			Name{
				ID:           "you-token",
				Name:         "フィルイン",
				Mode:         1,
				Target:       1,
				Address:      "",
				Kind:         []string{},
				Change:       false,
				Close:        false,
				From:         "",
				To:           "2021-07-19",
				Divide:       1,
				ResponseType: RESPONSE_TYPE,
			},
			"https://api.houjin-bangou.nta.go.jp/4/name?change=0&close=0&divide=1&id=you-token&mode=1&name=%E3%83%95%E3%82%A3%E3%83%AB%E3%82%A4%E3%83%B3&target=1&to=2021-07-19&type=12",
		},
	}

	for i, test := range tests {
		url, _ := test.Name.URL()

		if url.String() != test.URL {
			t.Errorf("%d: URL String not match result:%s expected:%s", i, url.String(), test.URL)
		}
	}
}

func ExampleName_URL() {
	keyword := "フィルイン"
	mode := 1
	target := 1
	address := "10202"
	kind := []string{"03"}
	name := NewName("your-token", keyword, mode, target, address, kind, false, false, "", "", 1)

	if err := name.Validate(); err != nil {
		fmt.Println(err)
		return
	}

	url, _ := name.URL()
	fmt.Println(url.String())
	// Output: https://api.houjin-bangou.nta.go.jp/4/name?address=10202&change=0&close=0&divide=1&id=your-token&kind=03&mode=1&name=%E3%83%95%E3%82%A3%E3%83%AB%E3%82%A4%E3%83%B3&target=1&type=12
}
