package corp

import (
	"encoding/xml"
	"time"
)

const DATE_FORMAT = "2006-01-02"

var location = "Asia/Tokyo"

type Date time.Time

func (date Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(date.String(), start)
}

func (date *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "" {
		*date = Date(time.Time{})
		return nil
	}

	t, err := time.ParseInLocation(DATE_FORMAT, s, currentLocation())
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

func (date Date) String() string {
	return date.Time().In(currentLocation()).Format(DATE_FORMAT)
}

func currentLocation() *time.Location {
	loc, _ := time.LoadLocation(location)
	return loc
}
