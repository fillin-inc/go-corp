package corp

import (
	"encoding/xml"
	"testing"
)

func TestUnmarshalToXML(t *testing.T) {
	str := `<corporations>
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
		t.Errorf("XMLのパースに失敗しました。%v", err)
	}
}
