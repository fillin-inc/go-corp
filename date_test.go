package corp

import (
	"encoding/xml"
	"reflect"
	"testing"
	"time"
)

type dateXMLTest struct {
	Result *Date `xml:"createdAt"`
}

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
	str := `
	<?xml version="1.0" encoding="UTF-8"?>
	<wrapper>
		<createdAt>2021-07-16</createdAt>
	</wrapper>
	`
	var dt dateXMLTest
	err := xml.Unmarshal([]byte(str), &dt)
	if err != nil {
		t.Error(err)
	}

	loc, _ := time.LoadLocation(location)
	expected := time.Date(2021, 7, 16, 0, 0, 0, 0, loc).Unix()
	if dt.Result.Time().Unix() != expected {
		t.Errorf("failed to parse. result:%v, expected:%v", dt.Result.Time().Unix(), expected)
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
