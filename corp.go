package corp

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/fillin-inc/go-corp/request"
)

var (
	// Web API アクセストークン
	appID string

	fetch = func(URL string, options interface{}) ([]byte, error) {
		var body []byte

		res, err := http.Get(URL)
		if err != nil {
			return body, err
		}
		defer res.Body.Close()

		return ioutil.ReadAll(res.Body)
	}
)

func ByNumber(number uint64) (Response, error) {
	builder := request.NewNumber(appID, []uint64{number}, false)
	return responseByURLBuilder(builder)
}

func ByNumberWithHistory(number uint64) (Response, error) {
	builder := request.NewNumber(appID, []uint64{number}, true)
	return responseByURLBuilder(builder)
}

func ByNumbers(numbers []uint64) (Response, error) {
	builder := request.NewNumber(appID, numbers, false)
	return responseByURLBuilder(builder)
}

func DiffSearch(from string, to string, address string) (Response, error) {
	builder := request.NewDiff(appID, from, to, address, []string{}, 1)
	return responseByURLBuilder(builder)
}

func NameSearch(name string, address string) (Response, error) {
	builder := request.NewName(appID, name, 2, 1, address, []string{}, false, true, "", "", 1)
	return responseByURLBuilder(builder)
}

func SetAppID(tkn string) {
	appID = tkn
}

func SetFetch(f func(URL string, options interface{}) ([]byte, error)) {
	fetch = f
}

func responseByURLBuilder(builder request.URLBuilder) (Response, error) {
	if err := builder.Validate(); err != nil {
		return Response{}, err
	}

	u, err := builder.URL()
	if err != nil {
		return Response{}, err
	}

	var body []byte
	var res Response
	body, err = fetch(u.String(), nil)
	if err != nil {
		return Response{}, err
	}

	err = xml.Unmarshal(body, &res)
	return res, err
}
