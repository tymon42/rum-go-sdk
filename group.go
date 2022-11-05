package rumgosdk

import (
	"github.com/tymon42/rum-go-sdk/model"
)

// StartSync starts sync, /api/v1/group/{group_id}/startsync
// Param: 1. group_id
// Return: apiStartSyncResult
func (q *Quorum) StartSync(groupID string) (*model.ApiStartSyncResult, error) {
	url := q.ApiServer + "/api/v1/group/" + groupID + "/startsync"
	res := &model.ApiStartSyncResult{}
	_, err := q.HttpClient.R().SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, err
}
