package corp

import (
	"encoding/xml"
	"reflect"
	"testing"
	"time"
)

type dateTest struct {
	Result *Date `xml:"createdAt"`
}

func TestUnmarshalDate(t *testing.T) {
	str := `
	<?xml version="1.0" encoding="UTF-8"?>
	<wrapper>
		<createdAt>2021-07-16</createdAt>
	</wrapper>
	`
	var dt dateTest
	err := xml.Unmarshal([]byte(str), &dt)
	if err != nil {
		t.Error(err)
	}

	loc, _ := time.LoadLocation("Asia/Tokyo")
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
