package corp

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"testing"
	"time"
)

type testDateMarshal struct {
	date     Date
	expected string
}

type testDateUnmarshal struct {
	str      string
	expected Date
}

func TestDateMarshalXML(t *testing.T) {
	tests := []testDateMarshal{
		{
			date:     Date(time.Date(2021, 7, 16, 0, 0, 0, 0, currentLocation())),
			expected: "<Date>2021-07-16</Date>",
		},
		{
			date:     Date(time.Date(2021, 8, 31, 0, 0, 0, 0, currentLocation())),
			expected: "<Date>2021-08-31</Date>",
		},
	}

	for i, test := range tests {
		b, err := xml.Marshal(test.date)
		if err != nil {
			t.Errorf("%d: MarshalXML return error:%v", i, err)
		}

		str := string(b)
		if str != test.expected {
			t.Errorf("%d: failed to MarshalXML. result:%v, expected:%v", i, str, test.expected)
		}
	}
}

func TestDateUnmarshalXML(t *testing.T) {
	tests := []testDateUnmarshal{
		{
			str:      "<Date>2021-07-16</Date>",
			expected: Date(time.Date(2021, 7, 16, 0, 0, 0, 0, currentLocation())),
		},
		{
			str:      "<Date>2021-08-31</Date>",
			expected: Date(time.Date(2021, 8, 31, 0, 0, 0, 0, currentLocation())),
		},
	}

	for i, test := range tests {
		var d Date
		err := xml.Unmarshal([]byte(test.str), &d)
		if err != nil {
			t.Errorf("%d: UnmarshalXML return error:%v", i, err)
		}

		if !reflect.DeepEqual(d, test.expected) {
			t.Errorf("%d: failed to UnmarshalXML. result:%v, expected:%v", i, d, test.expected)
		}
	}
}

func TestDateMarshalJSON(t *testing.T) {
	tests := []testDateMarshal{
		{
			date:     Date(time.Date(2021, 7, 16, 0, 0, 0, 0, currentLocation())),
			expected: `"2021-07-16"`,
		},
		{
			date:     Date(time.Date(2021, 8, 31, 0, 0, 0, 0, currentLocation())),
			expected: `"2021-08-31"`,
		},
	}

	for i, test := range tests {
		b, err := json.Marshal(test.date)
		if err != nil {
			t.Errorf("%d: MarshalJSON return error:%v", i, err)
		}

		str := string(b)
		if str != test.expected {
			t.Errorf("%d: failed to MarshalJSON. result:%v, expected:%v", i, str, test.expected)
		}
	}
}

func TestDateUnmarshalJSON(t *testing.T) {
	tests := []testDateUnmarshal{
		{
			str:      `"2021-07-16"`,
			expected: Date(time.Date(2021, 7, 16, 0, 0, 0, 0, currentLocation())),
		},
		{
			str:      `"2021-08-31"`,
			expected: Date(time.Date(2021, 8, 31, 0, 0, 0, 0, currentLocation())),
		},
	}

	for i, test := range tests {
		var d Date
		err := json.Unmarshal([]byte(test.str), &d)
		if err != nil {
			t.Errorf("%d: UnmarshalJSON return error:%v", i, err)
		}

		if !reflect.DeepEqual(d, test.expected) {
			t.Errorf("%d: failed to UnmarshalJSON. result:%v, expected:%v", i, d, test.expected)
		}
	}
}

func TestDateToTime(t *testing.T) {
	time := time.Now()
	d := Date(time)

	if reflect.TypeOf(d.Time()) != reflect.TypeOf(time) {
		t.Errorf("Type is wrrong. result:%T expected:%T", d.Time(), time)
	}
}

func TestDateToString(t *testing.T) {
	loc, _ := time.LoadLocation(location)
	time := time.Date(2021, 7, 16, 0, 0, 0, 0, loc)
	d := Date(time)

	expected := "2021-07-16"
	if d.String() != expected {
		t.Errorf("Unexpeted date string. result:%s expected:%s", d.String(), expected)
	}
}
