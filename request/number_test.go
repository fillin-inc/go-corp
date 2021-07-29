package request

import (
	"fmt"
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
	appID := "your-token"
	nums := []uint64{1234, 5678}
	history := true

	number := NewNumber(appID, nums, history)
	if number.ID != appID {
		t.Error("ID field is not match.")
	}

	if !reflect.DeepEqual(number.Numbers, nums) {
		t.Error("Numbers field is not match.")
	}

	if number.History != history {
		t.Error("History field is not match.")
	}
}

func TestNumberValidate(t *testing.T) {
	tests := []Number{
		{
			ID:           "your-token",
			Numbers:      []uint64{5070001032626},
			ResponseType: RESPONSE_TYPE,
			History:      true,
		},
		// Number has max slice count(10)
		{
			ID: "your-token",
			Numbers: []uint64{
				5070001032626,
				7000020100005,
				9000020102024,
				1010401089234,
				6011001106696,
				5070001032626,
				7000020100005,
				9000020102024,
				1010401089234,
				6011001106696,
			},
			ResponseType: RESPONSE_TYPE,
			History:      true,
		},
		// History is false
		{
			ID:           "your-token",
			Numbers:      []uint64{5070001032626},
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
				Numbers:      []uint64{5070001032626},
				ResponseType: RESPONSE_TYPE,
				History:      false,
			},
			"Key: 'Number.ID' Error:Field validation for 'ID' failed on the 'required' tag",
		},
		{
			Number{
				ID:           "your-token",
				Numbers:      []uint64{},
				ResponseType: RESPONSE_TYPE,
				History:      false,
			},
			"Key: 'Number.Numbers' Error:Field validation for 'Numbers' failed on the 'min' tag",
		},
		{
			Number{
				ID: "your-token",
				Numbers: []uint64{
					5070001032626,
					7000020100005,
					9000020102024,
					1010401089234,
					6011001106696,
					5070001032626,
					7000020100005,
					9000020102024,
					1010401089234,
					6011001106696,
					5070001032626,
				},
				ResponseType: RESPONSE_TYPE,
				History:      false,
			},
			"Key: 'Number.Numbers' Error:Field validation for 'Numbers' failed on the 'max' tag",
		},
		{
			Number{
				ID:           "your-token",
				Numbers:      []uint64{4070001032626},
				ResponseType: RESPONSE_TYPE,
				History:      false,
			},
			"Key: 'Number.Numbers' Error:Field validation for 'Numbers' failed on the 'checkdigits' tag",
		},
		{
			Number{
				ID: "your-token",
				Numbers: []uint64{
					5070001032626,
					6000020100005, // invalid checkdigit
				},
				ResponseType: RESPONSE_TYPE,
				History:      false,
			},
			"Key: 'Number.Numbers' Error:Field validation for 'Numbers' failed on the 'checkdigits' tag",
		},
		{
			Number{
				ID:           "your-token",
				Numbers:      []uint64{5070001032626},
				ResponseType: "",
				History:      false,
			},
			"Key: 'Number.ResponseType' Error:Field validation for 'ResponseType' failed on the 'required' tag",
		},
		{
			Number{
				ID:           "your-token",
				Numbers:      []uint64{5070001032626},
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
				ID:           "your-token",
				Numbers:      []uint64{5070001032626},
				ResponseType: RESPONSE_TYPE,
				History:      true,
			},
			"https://api.houjin-bangou.nta.go.jp/4/num?history=1&id=your-token&number=5070001032626&type=12",
		},
		{
			// Numbers is multiple value
			Number{
				ID: "your-token",
				Numbers: []uint64{
					5070001032626,
					7000020100005,
				},
				ResponseType: RESPONSE_TYPE,
				History:      true,
			},
			"https://api.houjin-bangou.nta.go.jp/4/num?history=1&id=your-token&number=5070001032626%2C7000020100005&type=12",
		},
		{
			// History is false
			Number{
				ID:           "your-token",
				Numbers:      []uint64{5070001032626},
				ResponseType: RESPONSE_TYPE,
				History:      false,
			},
			"https://api.houjin-bangou.nta.go.jp/4/num?history=0&id=your-token&number=5070001032626&type=12",
		},
	}

	for i, test := range tests {
		url, _ := test.Number.URL()

		if url.String() != test.URL {
			t.Errorf("%d: URL String not match result:%s expected:%s", i, url.String(), test.URL)
		}
	}
}

func BenchmarkNewNumber(b *testing.B) {
	appID := "your-token"
	numbers := []uint64{5070001032626}
	history := false

	for i := 0; i < b.N; i++ {
		NewNumber(appID, numbers, history)
	}
}

func BenchmarkNumberValidate(b *testing.B) {
	number := Number{
		ID:           "your-token",
		Numbers:      []uint64{5070001032626},
		ResponseType: RESPONSE_TYPE,
		History:      false,
	}

	for i := 0; i < b.N; i++ {
		number.Validate()
	}
}

func BenchmarkNumberURL(b *testing.B) {
	number := Number{
		ID:           "your-token",
		Numbers:      []uint64{5070001032626},
		ResponseType: RESPONSE_TYPE,
		History:      false,
	}

	for i := 0; i < b.N; i++ {
		number.URL()
	}
}

func ExampleNumber_URL() {
	// 株式会社フィルインの法人番号
	var corpNum uint64 = 5070001032626
	number := NewNumber("your-token", []uint64{corpNum}, false)

	if err := number.Validate(); err != nil {
		fmt.Println(err)
		return
	}

	url, _ := number.URL()
	fmt.Println(url.String())
	// Output: https://api.houjin-bangou.nta.go.jp/4/num?history=0&id=your-token&number=5070001032626&type=12
}
