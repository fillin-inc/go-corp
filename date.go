package corp

import (
	"encoding/xml"
	"time"
)

const DATE_FORMAT = "2006-01-02"

var location = "Asia/Tokyo"

type Date time.Time

func (date *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "" {
		*date = Date(time.Time{})
		return nil
	}

	loc, _ := time.LoadLocation(location)
	t, err := time.ParseInLocation(DATE_FORMAT, s, loc)
	if err != nil {
		return err
	}

	*date = Date(t)
	return nil
}

// convert to time.Time
func (date Date) Time() time.Time {
	return time.Time(date)
}
