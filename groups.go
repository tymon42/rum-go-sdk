package rumgosdk

import (
	"github.com/tymon42/rum-go-sdk/model"
)

// CreateGroupUrl create group url, /api/v1/group
// param: model.handlerCreateGroupParam
func (q *Quorum) CreateGroupUrl(param model.HandlersCreateGroupParam) (*model.HandlersCreateGroupResult, error) {
	url := q.ApiServer + "/api/v1/group"
	res := &model.HandlersCreateGroupResult{}
	_, err := q.HttpClient.R().SetBody(param).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ClearGroupData clears group data, /api/v1/clear
func (q *Quorum) ClearGroupData() error {
	url := q.ApiServer + "/api/v1/clear"
	_, err := q.HttpClient.R().Post(url)
	return err
}

// PostToGroup posts object to group, /api/v1/group/content
// param: pbActivity
// return: HandlersTrxResult
func (q *Quorum) PostToGroup(param *model.PbActivity) (*model.HandlersTrxResult, error) {
	url := q.ApiServer + "/api/v1/group/content"
	res := &model.HandlersTrxResult{}
	_, err := q.HttpClient.R().SetBody(param).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// JoinGroup joins group, /api/v1/group/join
// param: handlersGroupSeed
// return: apiJoinGroupResult
func (q *Quorum) JoinGroup(param model.HandlersGroupSeed) (*model.ApiJoinGroupResult, error) {
	url := q.ApiServer + "/api/v1/group/join"
	res := &model.ApiJoinGroupResult{}
	_, err := q.HttpClient.R().SetBody(param).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// JoinGroupByNewSeed joins group by new seed, /api/v1/group/join
// param: seed
// return: apiJoinGroupResult
func (q *Quorum) JoinGroupV2(seed string) (*model.ApiJoinGroupResult, error) {
	url := q.ApiServer + "/api/v1/group/join"
	res := &model.ApiJoinGroupResult{}
	_, err := q.HttpClient.R().SetBody(map[string]string{"seed": seed}).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// LeaveGroup leaves group, /api/v1/group/leave
// param: handlersGroupSeed
// return: HandlersLeaveGroupResult
func (q *Quorum) LeaveGroup(param model.HandlersGroupSeed) (*model.HandlersLeaveGroupResult, error) {
	url := q.ApiServer + "/api/v1/group/leave"
	res := &model.HandlersLeaveGroupResult{}
	_, err := q.HttpClient.R().SetBody(param).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetPubQueue gets pub queue, /api/v1/group/{group_id}/pubqueue
// return handlersPubQueueInfo
func (q *Quorum) GetPubQueue(group_id string) (*model.HandlersPubQueueInfo, error) {
	url := q.ApiServer + "/api/v1/group/" + group_id + "/pubqueue"
	res := &model.HandlersPubQueueInfo{}
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetGroupSeed gets group seed, /api/v1/group/{group_id}/seed
// return handlersGroupSeedResult
func (q *Quorum) GetGroupSeed(group_id string) (*model.HandlersGetGroupSeedResult, error) {
	url := q.ApiServer + "/api/v1/group/" + group_id + "/seed"
	res := &model.HandlersGetGroupSeedResult{}
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetGroups gets groups, /api/v1/groups
// return apiGroupsList
func (q *Quorum) GetGroups() (*model.ApiGroupInfoList, error) {
	url := q.ApiServer + "/api/v1/groups"
	res := &model.ApiGroupInfoList{}
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}
