package corp

import (
	"encoding/xml"
	"reflect"
	"testing"
	"time"
)

func TestDateMarshalXML(t *testing.T) {
	loc, _ := time.LoadLocation(location)
	d := Date(time.Date(2021, 7, 16, 0, 0, 0, 0, loc))

	b, err := xml.Marshal(d)
	if err != nil {
		t.Error(err)
	}

	str := string(b)
	expected := "<Date>2021-07-16</Date>"
	if str != expected {
		t.Errorf("failed to MarshalXML. result:%v, expected:%v", str, expected)
	}
}

func TestDateUnmarshalXML(t *testing.T) {
	str := `<Date>2021-07-16</Date>`

	var d Date
	err := xml.Unmarshal([]byte(str), &d)
	if err != nil {
		t.Error(err)
	}

	loc, _ := time.LoadLocation(location)
	expected := time.Date(2021, 7, 16, 0, 0, 0, 0, loc).Unix()
	if d.Time().Unix() != expected {
		t.Errorf("failed to UnmarshalXML. result:%v, expected:%v", d.Time().Unix(), expected)
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
