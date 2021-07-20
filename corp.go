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

func EasyNumberSearch(numbers []uint64) (Response, error) {
	builder := request.NewNumber(appID, numbers, false)
	return responseByURLBuilder(builder)
}

func NumberSearch(numbers []uint64, history bool) (Response, error) {
	builder := request.NewNumber(appID, numbers, history)
	return responseByURLBuilder(builder)
}

func EasyDiffSearch(from string, to string, address string) (Response, error) {
	builder := request.NewDiff(appID, from, to, address, []string{}, 1)
	return responseByURLBuilder(builder)
}

func DiffSearch(from string, to string, address string, kind []string, divide int) (Response, error) {
	builder := request.NewDiff(appID, from, to, address, kind, divide)
	return responseByURLBuilder(builder)
}

func EasyNameSearch(name string, address string) (Response, error) {
	builder := request.NewName(appID, name, 2, 1, address, []string{}, false, true, "", "", 1)
	return responseByURLBuilder(builder)
}

func NameSearch(name string, mode int, target int, address string, kind []string, change bool, close bool, from string, to string, divide int) (Response, error) {
	builder := request.NewName(appID, name, mode, target, address, kind, change, close, from, to, divide)
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
