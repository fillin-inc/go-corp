package corp

var (
	// Web API アクセストークン
	apiTkn string
)

func SetToken(tkn string) {
	apiTkn = tkn
}
