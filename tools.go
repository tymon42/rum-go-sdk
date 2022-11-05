package rumgosdk

import (
	"github.com/tymon42/rum-go-sdk/model"
)

// PubkeyToEthaddr converts a based64 encoded libp2p pubkey to an ethereum address, /api/v1/tools/pubkeytoaddr
// Param: apiPubkeyParam
// Return: map[string]string
func (q *Quorum) PubkeyToEthaddr(param *model.ApiPubkeyParam) (map[string]string, error) {
	url := q.ApiServer + "/api/v1/tools/pubkeytoaddr"
	res := map[string]string{}
	_, err := q.HttpClient.R().SetBody(param).SetResult(&res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, err
}
