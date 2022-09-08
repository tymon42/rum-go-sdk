package api

import (
	"github.com/rumsystem/quorum/cmd/cli/api"
	"github.com/rumsystem/quorum/pkg/chainapi/handlers"
)

// CreateGroup create group
func (c ReqClient) CreateGroup(name, consensusType, encryptiionType, appKey string) (res *handlers.CreateGroupResult, err error) {
	createGroup := &api.CreateGroupReqStruct{
		Name:           name,
		ConsensusType:  consensusType,
		EncryptionType: encryptiionType,
		AppKey:         appKey,
	}
	res, err = c.createGroup(createGroup)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c ReqClient) createGroup(data *api.CreateGroupReqStruct) (res *handlers.CreateGroupResult, err error) {
	url := c.Client.BaseURL + "/api/v1/group"
	c.Client.R().SetBody(data).SetResult(res).SetError(err).Post(url)
	if err != nil {
		return &handlers.CreateGroupResult{}, err
	}
	return res, nil
}

// ClearGroup clear group data
func (c ReqClient) ClearGroupData(groupID string) (res *handlers.ClearGroupDataResult, err error) {
	clearGroup := &api.DeleteGroupReqStruct{
		GroupId: groupID,
	}
	res, err = c.clearGroupData(clearGroup)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c ReqClient) clearGroupData(data *api.DeleteGroupReqStruct) (res *handlers.ClearGroupDataResult, err error) {
	url := c.Client.BaseURL + "/api/v1/group/clear"
	c.Client.R().SetBody(data).SetResult(res).SetError(err).Post(url)
	if err != nil {
		return &handlers.ClearGroupDataResult{}, err
	}
	return res, nil
}
