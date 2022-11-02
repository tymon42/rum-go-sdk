package rumgosdk

import (
	"strings"

	"github.com/go-resty/resty/v2"
)

type Quorum struct {
	ApiServer  string
	HttpClient *resty.Client
}

// Connect a new Quorum instance
func Connect(apiServer string) *Quorum {
	ApiServer := apiServer
	if len(apiServer) > 0 {
		if strings.HasPrefix(apiServer, "https") || strings.HasPrefix(apiServer, "http") {
			ApiServer = apiServer
		} else {
			ApiServer = "https://" + apiServer
		}
	}

	return &Quorum{
		ApiServer:  ApiServer,
		HttpClient: resty.New(),
	}
}
