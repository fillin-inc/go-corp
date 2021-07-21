package corp

import (
	"encoding/xml"
	"testing"
)

func TestUnmarshalToXML(t *testing.T) {
	str := `
	<?xml version="1.0" encoding="UTF-8"?>
		<corporations>
			<lastUpdateDate>2021-07-16</lastUpdateDate>
			<count>3</count>
			<divideNumber>1</divideNumber>
			<divideSize>1</divideSize>
			<corporation>
				<sequenceNumber>1</sequenceNumber>
				<corporateNumber>5070001032626</corporateNumber>
				<process>01</process>
				<correct>1</correct>
				<updateDate>2018-05-08</updateDate>
				<changeDate>2016-09-05</changeDate>
				<name>株式会社フィルイン</name>
				<nameImageId/>
				<kind>301</kind>
				<prefectureName>群馬県</prefectureName>
				<cityName>高崎市</cityName>
				<streetNumber>八島町５８番地１ウエストワンビル１０Ｆ１０１１号</streetNumber>
				<addressImageId/>
				<prefectureCode>10</prefectureCode>
				<cityCode>202</cityCode>
				<postCode>3700849</postCode>
				<addressOutside/>
				<addressOutsideImageId/>
				<closeDate/>
				<closeCause/>
				<successorCorporateNumber/>
				<changeCause/>
				<assignmentDate>2016-09-05</assignmentDate>
				<latest>0</latest>
				<enName/>
				<enPrefectureName/>
				<enCityName/>
				<enAddressOutside/>
				<furigana>フィルイン</furigana>
				<hihyoji>0</hihyoji>
			</corporation>
			<corporation>
				<sequenceNumber>2</sequenceNumber>
				<corporateNumber>5070001032626</corporateNumber>
				<process>12</process>
				<correct>1</correct>
				<updateDate>2018-05-08</updateDate>
				<changeDate>2018-05-02</changeDate>
				<name>株式会社フィルイン</name>
				<nameImageId/>
				<kind>301</kind>
				<prefectureName>群馬県</prefectureName>
				<cityName>高崎市</cityName>
				<streetNumber>本町４８番地</streetNumber>
				<addressImageId/>
				<prefectureCode>10</prefectureCode>
				<cityCode>202</cityCode>
				<postCode>3700813</postCode>
				<addressOutside/>
				<addressOutsideImageId/>
				<closeDate/>
				<closeCause/>
				<successorCorporateNumber/>
				<changeCause/>
				<assignmentDate>2016-09-05</assignmentDate>
				<latest>0</latest>
				<enName/>
				<enPrefectureName/>
				<enCityName/>
				<enAddressOutside/>
				<furigana>フィルイン</furigana>
				<hihyoji>0</hihyoji>
			</corporation>
			<corporation>
				<sequenceNumber>3</sequenceNumber>
				<corporateNumber>5070001032626</corporateNumber>
				<process>12</process>
				<correct>0</correct>
				<updateDate>2021-06-09</updateDate>
				<changeDate>2021-06-02</changeDate>
				<name>株式会社フィルイン</name>
				<nameImageId/>
				<kind>301</kind>
				<prefectureName>群馬県</prefectureName>
				<cityName>高崎市</cityName>
				<streetNumber>飯塚町１４７番地４</streetNumber>
				<addressImageId/>
				<prefectureCode>10</prefectureCode>
				<cityCode>202</cityCode>
				<postCode>3700069</postCode>
				<addressOutside/>
				<addressOutsideImageId/>
				<closeDate/>
				<closeCause/>
				<successorCorporateNumber/>
				<changeCause/>
				<assignmentDate>2016-09-05</assignmentDate>
				<latest>1</latest>
				<enName/>
				<enPrefectureName/>
				<enCityName/>
				<enAddressOutside/>
				<furigana>フィルイン</furigana>
				<hihyoji>0</hihyoji>
			</corporation>
		</corporations>`

	var res Response
	err := xml.Unmarshal([]byte(str), &res)

	if err != nil {
		t.Errorf("failed to parse XML: %v", err)
	}

	// 0/1 を bool で扱うフィールドのみ検証
	corp := res.Corporations[0]
	if corp.Correct != true {
		t.Errorf("Correct is wrong result:%t expected:%t", corp.Correct, true)
	}

	if corp.Latest != false {
		t.Errorf("Latest is wrong result:%t expected:%t", corp.Latest, false)
	}

	if corp.Hihyoji != false {
		t.Errorf("Hihyoji is wrong result:%t expected:%t", corp.Hihyoji, false)
	}

	corp = res.Corporations[2]
	if corp.Correct != false {
		t.Errorf("Correct is wrong result:%t expected:%t", corp.Correct, false)
	}

	if corp.Latest != true {
		t.Errorf("Latest is wrong result:%t expected:%t", corp.Latest, true)
	}

	if corp.Hihyoji != false {
		t.Errorf("Hihyoji is wrong result:%t expected:%t", corp.Hihyoji, false)
	}
}

func TestUnmarshalToXMLWhenProcessIs99(t *testing.T) {
	str := `
	<?xml version="1.0" encoding="UTF-8"?>
		<corporations>
			<lastUpdateDate>2021-07-16</lastUpdateDate>
			<count>1</count>
			<divideNumber>1</divideNumber>
			<divideSize>1</divideSize>
			<corporation>
				<sequenceNumber>1</sequenceNumber>
				<corporateNumber>5070001032626</corporateNumber>
				<process>99</process>
				<correct/>
				<updateDate>2021-06-09</updateDate>
				<changeDate/>
				<name/>
				<nameImageId/>
				<kind/>
				<prefectureName/>
				<cityName/>
				<streetNumber/>
				<addressImageId/>
				<prefectureCode/>
				<cityCode/>
				<postCode/>
				<addressOutside/>
				<addressOutsideImageId/>
				<closeDate/>
				<closeCause/>
				<successorCorporateNumber/>
				<changeCause/>
				<assignmentDate/>
				<latest/>
				<enName/>
				<enPrefectureName/>
				<enCityName/>
				<enAddressOutside/>
				<furigana/>
				<hihyoji/>
			</corporation>
		</corporations>`

	var res Response
	err := xml.Unmarshal([]byte(str), &res)

	if err != nil {
		t.Errorf("failed to parse XML: %v", err)
	}
}

func TestProcessText(t *testing.T) {
	c := Corporation{}
	for key, val := range processes {
		c.Process = key
		if c.ProcessText() != val {
			t.Errorf("ProcessText return wrong value result:%s expected:%s", c.ProcessText(), val)
		}
	}
}

func TestKindText(t *testing.T) {
	c := Corporation{}
	for key, val := range kinds {
		c.Kind = key
		if c.KindText() != val {
			t.Errorf("KindText return wrong value result:%s expected:%s", c.KindText(), val)
		}
	}
}

func TestKindTextEmpty(t *testing.T) {
	c := Corporation{}
	if c.KindText() != "" {
		t.Errorf("KindText return wrong value result:%s expected:%s", c.KindText(), "")
	}
}

func TestCloseCauseText(t *testing.T) {
	c := Corporation{}
	for key, val := range closeCauses {
		c.CloseCause = key
		if c.CloseCauseText() != val {
			t.Errorf("CloseCauseText return wrong value result:%s expected:%s", c.CloseCauseText(), val)
		}
	}
}

func TestCloseCauseTextEmpty(t *testing.T) {
	c := Corporation{}
	if c.CloseCauseText() != "" {
		t.Errorf("CloseCauseText return wrong value result:%s expected:%s", c.CloseCauseText(), "")
	}
}
