package request

import (
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

var (
	prefCodes = []string{
		"01", "02", "03", "04", "05", "06", "07", "08", "09", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30",
		"31", "32", "33", "34", "35", "36", "37", "38", "39", "40",
		"41", "42", "43", "44", "45", "46", "47",
		"99",
	}

	kindCodes = []string{
		"01", "02", "03", "04",
	}

	rCityCode = regexp.MustCompile("^[0-9]{3}$")
)

func dateValidation(fl validator.FieldLevel) bool {
	v := fl.Field().String()
	// 空文字列は別バリデーションで検証
	if v == "" {
		return true
	}
	_, err := time.Parse("2006-01-02", v)

	return err == nil
}

func dateEqualOrGreaterValidation(fl validator.FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()

	topField, topKind, ok := fl.GetStructFieldOK()
	if !ok || topKind != kind {
		return false
	}

	// 比較対象の文字列が空文字列の場合は true
	// 必須か否かは別のバリデーションで検証
	if field.String() == "" || topField.String() == "" {
		return true
	}

	// パースエラーは別のバリデーションで検証するため true
	var fieldTime, topTime time.Time
	var err error
	fieldTime, err = time.Parse("2006-01-02", field.String())
	if err != nil {
		return true
	}
	topTime, err = time.Parse("2006-01-02", topField.String())
	if err != nil {
		return true
	}

	// 比較対象の日時と同一または大きい値の場合 true
	return fieldTime.Equal(topTime) || fieldTime.After(topTime)
}

func addressValidation(fl validator.FieldLevel) bool {
	v := fl.Field().String()
	if v == "" {
		return true
	}

	if len(v) == 2 {
		return containCodes(v, prefCodes)
	}

	if len(v) == 5 {
		prefCode := v[:2]
		cityCode := v[2:]

		if !containCodes(prefCode, prefCodes) {
			return false
		}
		// 市区町村コードは数が多いため正規表現チェックのみで判定
		return rCityCode.MatchString(cityCode)
	}

	return false
}

func kindValidation(fl validator.FieldLevel) bool {
	kinds := fl.Field().Interface().([]string)
	if len(kinds) == 0 {
		return true
	}

	var invalid []string
	for _, kind := range kinds {
		if !containCodes(kind, kindCodes) {
			invalid = append(invalid, kind)
		}
	}

	return len(invalid) == 0
}

func containCodes(target string, codes []string) bool {
	for _, c := range codes {
		if target == c {
			return true
		}
	}
	return false
}
