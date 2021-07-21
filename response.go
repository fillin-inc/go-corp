package corp

var (
	// 処理区分
	processes = map[string]string{
		"01": "新規",
		"11": "商号又は名称の変更",
		"12": "国内所在地の変更",
		"13": "国外所在地の変更",
		"21": "登記記録の閉鎖等",
		"22": "登記記録の復活等",
		"71": "吸収合併",
		"72": "吸収合併無効",
		"81": "商号の登記の抹消",
		"99": "削除",
	}

	// 法人種別
	kinds = map[uint16]string{
		101: "国の機関",
		201: "地方公共団体",
		301: "株式会社",
		302: "有限会社",
		303: "合名会社",
		304: "合資会社",
		305: "合同会社",
		399: "その他の設立登記法人",
		401: "外国会社等",
		499: "その他",
	}

	// 登記記録の閉鎖等の事由
	closeCauses = map[string]string{
		"01": "精算の結了等",
		"11": "合併による解散等",
		"21": "登記官による閉鎖",
		"31": "その他の精算の結了等",
	}
)

/*
Response は法人番号システム Web-API から取得できる XML データを扱います。

詳細については Web-API 仕様書巻末のリソース定義書を参照してください。
*/
type Response struct {
	// 最終更新年月日
	LastUpdateDate *Date `xml:"lastUpdateDate"`
	// 総件数
	// 一致するデータがない場合は 0
	Count uint32 `xml:"count"`
	// 分割番号
	DivideNumber uint32 `xml:"divideNumber"`
	// 分割数
	DevideSize uint32 `xml:"divideSize"`
	// 法人等要素
	Corporations []Corporation `xml:"corporation"`
}

// Corporation は法人番号システム Web-API から取得できるレスポンスのうち, 法人情報部分の構造体です。
type Corporation struct {
	// 一連番号
	SequenceNumber uint32 `xml:"sequenceNumber"`
	// 法人番号
	CorporateNumber uint64 `xml:"corporateNumber"`
	// 処理区分
	// 01:新規, 11:商号又は名称の変更, 12:国内所在地の変更, 13: 国外所在地の変更,
	// 21:登記記録の閉鎖等, 22:登記記録の復活等,
	// 71:吸収合併, 72:吸収合併無効, 81:商号の登記の抹消, 99:削除
	Process string `xml:"process"`
	// 訂正区分
	// false:訂正以外, true:訂正
	Correct bool `xml:"correct"`
	// 更新年月日
	UpdateDate *Date `xml:"updateDate"`
	// 変更年月日
	ChangeDate *Date `xml:"ChangeDate"`
	// 商号または名称
	Name string `xml:"name"`
	// 商号または名称イメージID
	NameImageId string `xml:"nameImageId"`
	// 法人種別
	// 101:国の機関, 201:地方公共団体,
	// 301:株式会社, 302:有限会社, 303:合名会社, 304:合資会社, 305:合同会社, 399:その他の設立登記法人
	// 401:外国会社等, 402:その他
	Kind uint16 `xml:"kind"`
	// 国内所在地(都道府県)
	PrefectureName string `xml:"prefectureName"`
	// 国内所在地(市区町村)
	CityName string `xml:"cityName"`
	// 国内所在地(丁目番地等)
	StreetNumber string `xml:"streetNumber"`
	// 国内所在地イメージID
	AddressImageId string `xml:"addressImageId"`
	// 都道府県コード
	// JIS X 401 に準ずる
	PrefectureCode uint8 `xml:"prefectureCode"`
	// 市区町村コード
	// JIS X 402 に準ずる
	CityCode uint16 `xml:"cityCode"`
	// 郵便番号
	PostCode string `xml:"postCode"`
	// 国外所在地
	AddressOutside string `xml:"addressOutside"`
	// 国外所在地イメージID
	AddressOutsideImageId string `xml:"addressOutsideImageId"`
	// 登記記録の閉鎖等年月日
	CloseDate *Date `xml:"closeDate"`
	// 登記記録の閉鎖等の事由
	// 01:精算の結了等, 11:合併による解散等, 21:登記官による閉鎖, 31:その他の精算の結了等
	CloseCause string `xml:"closeCause"`
	// 承継先法人番号
	SuccessorCorporateNumber uint64 `xml:"successorCorporateNumber"`
	// 変更事由の詳細
	ChangeCause string `xml:"changeCause"`
	// 法人番号指定年月日
	AssignmentDate *Date `xml:"assignmentDate"`
	// 最新履歴
	// false: 過去情報, true:最新情報
	Latest bool `xml:"latest"`
	// 商号または名称(英語表記)
	EnName string `xml:"enName"`
	// 国内所在地(都道府県)(英語表記)
	EnPrefectureName string `xml:"enPrefectureName"`
	// 国内所在地(市区町村)(英語表記)
	EnCityName string `xml:"enCityName"`
	// 国内所在地(丁目番地等)(英語表記)
	EnAddressOutside string `xml:"enAddressOutside"`
	// フリガナ
	Furigana string `xml:"furigana"`
	// 検索対象除外
	// false:検索対象, true:検索対象除外
	Hihyoji bool `xml:"hihyoji"`
}

// ProcessText は Process(処理区分) の表示用テキストを返します。
func (c Corporation) ProcessText() string {
	return processes[c.Process]
}

// KindText は Kind(法人種別) の表示用テキストを返します。
func (c Corporation) KindText() string {
	if c.Kind == 0 {
		return ""
	}
	return kinds[c.Kind]
}

// CloseCauseText は CloseCause(登記事項の閉鎖等の事由)の表示用テキストを返します。
func (c Corporation) CloseCauseText() string {
	if c.CloseCause == "" {
		return ""
	}
	return closeCauses[c.CloseCause]
}

/*
Available は法人情報が有効か判定します。

処理区分(Process) が 99 の場合,
一連番号(SequenceNumber), 法人番号(CorporateNumber), 更新年月日(UpdateDate) を除き
すべてブランクとなります。

このデータは実際には利用できないため無効と判定されます。
*/
func (c Corporation) Available() bool {
	return c.Process != "99"
}
