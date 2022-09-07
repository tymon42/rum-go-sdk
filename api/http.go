package api

import (
	"strings"

	"github.com/imroc/req/v3"
)

type ReqClient struct {
	Client *req.Client
}

func NewReqClient(baseURL string) *ReqClient {
	ApiServer := baseURL
	if len(baseURL) > 0 {
		if strings.HasPrefix(baseURL, "https") || strings.HasPrefix(baseURL, "http") {
			ApiServer = baseURL
		} else {
			ApiServer = "https://" + baseURL
		}
	}
	return &ReqClient{
		Client: req.C().SetBaseURL(ApiServer),
	}
}
