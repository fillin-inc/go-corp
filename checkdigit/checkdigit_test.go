package checkdigit

import "testing"

type testCalc struct {
	CorpNum    uint64
	CheckDigit int
}

type testErr struct {
	CorpNum uint64
	ErrMsg  string
}

func TestCalculateCheckDigit(t *testing.T) {
	tests := []testCalc{
		// 株式会社フィルイン
		{
			uint64(5070001032626),
			5,
		},
		// 群馬県
		{
			uint64(7000020100005),
			7,
		},
		// 高崎市
		{
			uint64(9000020102024),
			9,
		},
		// グーグル合同会社
		{
			uint64(1010401089234),
			1,
		},
		// アマゾンウェブサービスジャパンジャパン株式会社
		{
			uint64(6011001106696),
			6,
		},
		// Facebook Technoligies Japan合同会社
		{
			uint64(8010403022079),
			8,
		},
		// Apple Japan合同会社
		{
			uint64(3011103003992),
			3,
		},
	}

	for i, test := range tests {
		checkDigit, err := CalculateCheckDigit(test.CorpNum)
		if err != nil {
			t.Error(err)
		}

		if checkDigit != test.CheckDigit {
			t.Errorf("%d:CheckDigit is not match. result:%d expected:%d", i, checkDigit, test.CheckDigit)
		}
	}
}

func TestCalculateCheckDigitError(t *testing.T) {
	tests := []testErr{
		// 桁数誤り
		{
			uint64(1234),
			"Corporate Number is 13-digit number. Not 4 Digit.",
		},
	}

	for i, test := range tests {
		_, err := CalculateCheckDigit(test.CorpNum)

		if err == nil {
			t.Errorf("%d: not returns", i)
		}

		if err.Error() != test.ErrMsg {
			t.Errorf("%d: Error Message not matched result:%s expected:%s", i, err.Error(), test.ErrMsg)
		}
	}
}

func TestIsValid(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		tests := []uint64{
			// 株式会社フィルイン
			5070001032626,
			// 群馬県
			7000020100005,
			// 高崎市
			9000020102024,
			// グーグル合同会社
			1010401089234,
			// アマゾンウェブサービスジャパンジャパン株式会社
			6011001106696,
			// Facebook Technoligies Japan合同会社
			8010403022079,
			// Apple Japan合同会社
			3011103003992,
		}

		for i, corpNum := range tests {
			ret, err := IsValid(corpNum)

			if err != nil {
				t.Error(err)
			}

			if !ret {
				t.Errorf("%d: IsValid return false.", i)
			}
		}
	})

	t.Run("invalid", func(t *testing.T) {
		tests := []uint64{
			4070001032626,
			6000020100005,
			8000020102024,
			2010401089234,
			7011001106696,
			9010403022079,
			1011103003992,
		}

		for i, corpNum := range tests {
			ret, err := IsValid(corpNum)

			if err != nil {
				t.Error(err)
			}

			if ret {
				t.Errorf("%d: IsValid return true.", i)
			}
		}
	})
}

func TestIsValidError(t *testing.T) {
	tests := []testErr{
		// 桁数誤り
		{
			uint64(1234),
			"Corporate Number is 13-digit number. Not 4 Digit.",
		},
	}

	for i, test := range tests {
		_, err := CalculateCheckDigit(test.CorpNum)

		if err == nil {
			t.Errorf("%d: not returns", i)
		}

		if err.Error() != test.ErrMsg {
			t.Errorf("%d: Error Message not matched result:%s expected:%s", i, err.Error(), test.ErrMsg)
		}
	}
}
