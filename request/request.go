package request

import "fmt"

const URL = "https://api.houjin-bangou.nta.go.jp"

const API_VER = 4

const RESPONSE_TYPE = "12"

const NUMBER_END_POINT = "num"

func requestURL(endPoint string) string {
	return fmt.Sprintf("%s/%d/%s", URL, API_VER, endPoint)
}
