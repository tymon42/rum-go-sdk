package rumgosdk

import "github.com/tymon42/rum-go-sdk/model"

// GetBlock gets block, /api/v1/block/{group_id}/{block_id}
// return pbBlock
func (q *Quorum) GetBlock(group_id string, block_id string) (res *model.PbBlock, err error) {
	url := q.ApiServer + "/api/v1/block/" + group_id + "/" + block_id
	q.HttpClient.R().SetResult(res).SetError(err).Get(url)
	return res, err
}

// PubQueueAck acknowledges pub queue, /api/v1/trx/ack
// param: apiPubQueueAckPayload
// return: []string
func (q *Quorum) PubQueueAck(param model.ApiPubQueueAckPayload) (res []string, err error) {
	url := q.ApiServer + "/api/v1/trx/ack"
	q.HttpClient.R().SetBody(param).SetResult(res).SetError(err).Post(url)
	return res, err
}

// GetTrx gets trx of group, /api/v1/trx/{group_id}/{trx_id}
// return pbTrx
func (q *Quorum) GetTrx(group_id string, trx_id string) (res *model.PbTrx, err error) {
	url := q.ApiServer + "/api/v1/trx/" + group_id + "/" + trx_id
	q.HttpClient.R().SetResult(res).SetError(err).Get(url)
	return res, err
}
