package corp

type Response struct {
	LastUpdateDate string        `xml:"lastUpdateDate"`
	Count          uint32        `xml:"count"`
	DivideNumber   uint16        `xml:"divideNumber"`
	DevideSize     uint16        `xml:"divideSize"`
	Corporations   []Corporation `xml:"corporation"`
}

type Corporation struct {
	SequenceNumber           uint16 `xml:"sequenceNumber"`
	CorporateNumber          uint64 `xml:"corporateNumber"`
	Process                  string `xml:"process"`
	Correct                  uint8  `xml:"correct"`
	UpdateDate               string `xml:"updateDate"`
	ChangeDate               string `xml:"ChangeDate"`
	Name                     string `xml:"name"`
	NameImageId              string `xml:"nameImageId"`
	Kind                     uint16 `xml:"kind"`
	PrefectureName           string `xml:"prefectureName"`
	CityName                 string `xml:"cityName"`
	StreetNumber             string `xml:"streetNumber"`
	AddressImageId           string `xml:"addressImageId"`
	PrefectureCode           uint8  `xml:"prefectureCode"`
	CityCode                 uint16 `xml:"cityCode"`
	PostCode                 string `xml:"postCode"`
	AddressOutside           string `xml:"addressOutside"`
	AddressOutsideImageId    string `xml:"addressOutsideImageId"`
	CloseDate                string `xml:"closeDate"`
	CloseCause               string `xml:"closeCause"`
	SuccessorCorporateNumber uint64 `xml:"successorCorporateNumber"`
	ChangeCause              string `xml:"changeCause"`
	AssignmentDate           string `xml:"assignmentDate"`
	Latest                   bool   `xml:"latest"`
	EnName                   string `xml:"enName"`
	EnPrefectureName         string `xml:"enPrefectureName"`
	EnCityName               string `xml:"enCityName"`
	EnAddressOutside         string `xml:"enAddressOutside"`
	Furigana                 string `xml:"furigana"`
	Hihyoji                  bool   `xml:"hihyoji"`
}
