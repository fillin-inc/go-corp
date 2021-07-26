/*
法人番号のチェックデジットパッケージです。

法人番号は 1 桁目をチェックデジットとして 2 〜 13 桁の数値の整合性チェックを行うことができます。
こんパッケージではチェックデジットの算出と検証をサポートします。

チェックデジットの詳細については次の URL の PDF をご参照ください。
https://www.houjin-bangou.nta.go.jp/documents/checkdigit.pd://www.houjin-bangou.nta.go.jp/documents/checkdigit.pdf
*/
package checkdigit

import (
	"fmt"
	"strconv"
)

// CalculateCheckDigit は法人番号の 2 〜 13 桁の数値からチェックデジットを算出します。
func CalculateCheckDigit(corpNum uint64) (int, error) {
	str, err := corpNumStr(corpNum)
	if err != nil {
		return 0, err
	}

	var even, odd int
	even, err = totalOfEvenRows(str)
	if err != nil {
		return 0, err
	}

	odd, err = totalOfOddRows(str)
	if err != nil {
		return 0, err
	}

	return calculateDigit(even, odd), nil
}

/*
IsValid は法人番号が正しい形式か判定します。

法人番号の 1 桁目のチェックデジットと 2 〜 13 桁の算出値を用いて正しい形式か判定することができます。
詳細については次の URL(PDF) を確認してください。

https://www.houjin-bangou.nta.go.jp/documents/checkdigit.pdf
*/
func IsValid(corpNum uint64) (bool, error) {
	calcDigit, err := CalculateCheckDigit(corpNum)
	if err != nil {
		return false, err
	}

	checkDigit, err := checkDigit(corpNum)
	if err != nil {
		return false, err
	}

	return calcDigit == checkDigit, nil
}

// corpNumStr は法人番号を文字列に変換し桁数を確認します。
func corpNumStr(corpNum uint64) (string, error) {
	str := strconv.FormatUint(corpNum, 10)
	if len(str) != 13 {
		return "", fmt.Errorf("Corporate Number is 13-digit number. Not %d Digit.", len(str))
	}

	return str, nil
}

// totalOfEvenRows は法人番号文字列の最下位から偶数行の和を取得します。
func totalOfEvenRows(corpNumStr string) (int, error) {
	total := 0
	evenStrs := []string{
		corpNumStr[1:2],
		corpNumStr[3:4],
		corpNumStr[5:6],
		corpNumStr[7:8],
		corpNumStr[9:10],
		corpNumStr[11:12],
	}

	for _, str := range evenStrs {
		num, err := strconv.Atoi(str)
		if err != nil {
			return 0, err
		}
		total += num
	}
	return total, nil
}

// totalOfEvenRows は法人番号文字列の最下位から奇数行の和を取得します。
func totalOfOddRows(corpNumStr string) (int, error) {
	total := 0
	oddStrs := []string{
		corpNumStr[2:3],
		corpNumStr[4:5],
		corpNumStr[6:7],
		corpNumStr[8:9],
		corpNumStr[10:11],
		corpNumStr[12:13],
	}

	for _, str := range oddStrs {
		num, err := strconv.Atoi(str)
		if err != nil {
			return 0, err
		}
		total += num
	}
	return total, nil
}

// calculateDigit は偶数列, 奇数列それぞれの合計値からチェックデジットを算出します。
// https://www.houjin-bangou.nta.go.jp/documents/checkdigit.pdf の計算式を参照
func calculateDigit(even, odd int) int {
	return 9 - (((even * 2) + odd) % 9)
}

// checkDigit は処理対象の法人番号の 1 桁目のチェックデジットを取得します。
func checkDigit(corpNum uint64) (int, error) {
	str, err := corpNumStr(corpNum)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(str[:1])
}
