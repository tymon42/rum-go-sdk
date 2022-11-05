package rumgosdk

import (
	"github.com/tymon42/rum-go-sdk/model"
)

// SignTx signs a transaction, /api/v1/keystore/signtx
// Param: apiSignTxParam
// Return: apiSignTxResult
func (q *Quorum) SignTx(param *model.ApiSignTxParam) (*model.ApiSignTxResult, error) {
	url := q.ApiServer + "/api/v1/keystore/signtx"
	res := &model.ApiSignTxResult{}
	_, err := q.HttpClient.R().SetBody(param).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, err
}
