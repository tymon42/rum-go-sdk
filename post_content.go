package rumgosdk

import (
	"errors"
	"fmt"
)

func (q *Quorum) postFileSplit(cont, splitHash, groupID string) (string, error) {
	res, err := q.CreateContent(groupID, splitHash, cont)
	// fmt.Printf("res.TrxId: %v\n", res.TrxId)
	if err != nil {
		return "", err
	}
	return res.TrxId, nil
}

func (q *Quorum) checkSplitTrxId(groupId, trxId string) (bool, error) {
	res, err := q.TrxInfo(groupId, trxId)
	if err != nil {
		return false, err
	}
	if res.TrxId == "" {
		return false, errors.New("trxId is empty")
	}
	acked, err := q.PubQueueAck([]string{trxId})
	if err != nil {
		return false, err
	}
	if trxId == acked[0] {
		return true, nil
	}
	fmt.Printf("res.TrxId: %v\n", res.TrxId)
	return true, nil
}
