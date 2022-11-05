package rumgosdk

import (
	"strconv"

	"github.com/tymon42/rum-go-sdk/model"
)

// GetGroupContents gets contents  in a group, /app/api/v1/group/{group_id}/contents
// Param: 1. group_id, 2. num, 3. reverse, 4. starttrx, 5. includestarttrx, 6. nonce, 7. appapiSenderList
// Return: appapiGroupContentObjectItem
func (q *Quorum) GetGroupContents(groupID string, num int, reverse bool, starttrx string, includestarttrx string, nonce int, appapiSenderList *model.AppapiSenderList) ([]*model.AppapiGroupContentObjectItem, error) {
	url := q.ApiServer + "/app/api/v1/group/" + groupID + "/contents"
	var res []*model.AppapiGroupContentObjectItem
	_, err := q.HttpClient.R().SetQueryParams(map[string]string{
		"num":             string(num),
		"reverse":         strconv.FormatBool(reverse),
		"starttrx":        starttrx,
		"includestarttrx": string(includestarttrx),
		"nonce":           string(nonce),
	}).SetBody(appapiSenderList).SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, err
}

// CreateToken creates a new auth token, only allow access from localhost, /app/api/v1/token/create
// Param: appapi.CreateJWTParams
// Return: appapiTokenItem
func (q *Quorum) CreateToken(param *model.AppapiCreateJwtParams) (*model.AppapiTokenItem, error) {
	url := q.ApiServer + "/app/api/v1/token/create"
	res := &model.AppapiTokenItem{}
	_, err := q.HttpClient.R().SetBody(param).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, err
}

// RefreshToken gets a new auth token, /app/api/v1/token/refresh
// Headers: Authorization: {token}
// Return: appapiTokenItem
func (q *Quorum) RefreshToken(token string) (*model.AppapiTokenItem, error) {
	url := q.ApiServer + "/app/api/v1/token/refresh"
	res := &model.AppapiTokenItem{}
	_, err := q.HttpClient.R().SetHeader("Authorization", token).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, err
}
