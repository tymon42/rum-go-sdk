package rumgosdk

import (
	"github.com/tymon42/rum-go-sdk/model"
)

// CreateGroupUrl create group url, /api/v1/group
// param: model.handlerCreateGroupParam
func (q *Quorum) CreateGroupUrl(param model.HandlersCreateGroupParam) (res *model.HandlersCreateGroupResult, err error) {
	url := q.ApiServer + "/api/v1/group"
	q.HttpClient.R().SetBody(param).SetResult(res).SetError(err).Post(url)
	return res, err
}

// ClearGroupData clears group data, /api/v1/clear
func (q *Quorum) ClearGroupData() (err error) {
	url := q.ApiServer + "/api/v1/clear"
	q.HttpClient.R().SetError(err).Post(url)
	return err
}

// PostToGroup posts object to group, /api/v1/group/content
// param: pbActivity
// return: HandlersTrxResult
func (q *Quorum) PostToGroup(param *model.PbActivity) (res *model.HandlersTrxResult, err error) {
	url := q.ApiServer + "/api/v1/group/content"
	q.HttpClient.R().SetBody(param).SetResult(res).SetError(err).Post(url)
	return res, err
}

// JoinGroup joins group, /api/v1/group/join
// param: handlersGroupSeed
// return: apiJoinGroupResult
func (q *Quorum) JoinGroup(param model.HandlersGroupSeed) (res *model.ApiJoinGroupResult, err error) {
	url := q.ApiServer + "/api/v1/group/join"
	q.HttpClient.R().SetBody(param).SetResult(res).SetError(err).Post(url)
	return res, err
}

// LeaveGroup leaves group, /api/v1/group/leave
// param: handlersGroupSeed
// return: HandlersLeaveGroupResult
func (q *Quorum) LeaveGroup(param model.HandlersGroupSeed) (res *model.HandlersLeaveGroupResult, err error) {
	url := q.ApiServer + "/api/v1/group/leave"
	q.HttpClient.R().SetBody(param).SetResult(res).SetError(err).Post(url)
	return res, err
}

// GetPubQueue gets pub queue, /api/v1/group/{group_id}/pubqueue
// return handlersPubQueueInfo
func (q *Quorum) GetPubQueue(group_id string) (res *model.HandlersPubQueueInfo, err error) {
	url := q.ApiServer + "/api/v1/group/" + group_id + "/pubqueue"
	q.HttpClient.R().SetResult(res).SetError(err).Get(url)
	return res, err
}

// GetGroupSeed gets group seed, /api/v1/group/{group_id}/seed
// return handlersGroupSeedResult
func (q *Quorum) GetGroupSeed(group_id string) (res *model.HandlersGetGroupSeedResult, err error) {
	url := q.ApiServer + "/api/v1/group/" + group_id + "/seed"
	q.HttpClient.R().SetResult(res).SetError(err).Get(url)
	return res, err
}

// GetGroups gets groups, /api/v1/groups
// return apiGroupsList
func (q *Quorum) GetGroups() (res *model.ApiGroupInfoList, err error) {
	url := q.ApiServer + "/api/v1/groups"
	q.HttpClient.R().SetResult(res).SetError(err).Get(url)
	return res, err
}
