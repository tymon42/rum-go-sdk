package rumgosdk

import (
	"github.com/tymon42/rum-go-sdk/model"
)

// GetBlock gets block, /api/v1/block/{group_id}/{block_id}
// return pbBlock
func (q *Quorum) GetBlock(group_id string, block_id string) (*model.PbBlock, error) {
	url := q.ApiServer + "/api/v1/block/" + group_id + "/" + block_id
	res := &model.PbBlock{}
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PubQueueAck acknowledges pub queue, /api/v1/trx/ack
// param: apiPubQueueAckPayload
// return: []string
func (q *Quorum) PubQueueAck(param model.ApiPubQueueAckPayload) ([]string, error) {
	url := q.ApiServer + "/api/v1/trx/ack"
	res := []string{}
	_, err := q.HttpClient.R().SetBody(param).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetTrx gets trx of group, /api/v1/trx/{group_id}/{trx_id}
// return pbTrx
func (q *Quorum) GetTrx(group_id string, trx_id string) (*model.PbTrx, error) {
	url := q.ApiServer + "/api/v1/trx/" + group_id + "/" + trx_id
	res := &model.PbTrx{}
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}
