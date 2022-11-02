package rumgosdk

import (
	"github.com/tymon42/rum-go-sdk/model"
)

// AnnounceUserPubkey announces user pubkey, /api/v1/group/announce
// param: handlersAnnounceParam
// return: HandlersAnnounceResult
func (q *Quorum) AnnounceUserPubkey(param model.HandlersAnnounceParam) (res *model.HandlersAnnounceResult, err error) {
	url := q.ApiServer + "/api/v1/group/announce"
	q.HttpClient.R().SetBody(param).SetResult(res).SetError(err).Post(url)
	return res, err
}

// UpdateProfile updates profile, /api/v1/group/profile
// param: pbActivity
// return: HandlersUpdateProfileResult
func (q *Quorum) UpdateProfile(param model.PbActivity) (res *model.HandlersUpdateProfileResult, err error) {
	url := q.ApiServer + "/api/v1/group/profile"
	q.HttpClient.R().SetBody(param).SetResult(res).SetError(err).Post(url)
	return res, err
}

// GetAnnouncedGroupProducer gets announced group producer, /api/v1/group/{group_id}/producers
// return []HandlersAnnouncedProdcerListItempbGroupProducer
func (q *Quorum) GetAnnouncedGroupProducer(group_id string) (res []model.HandlersAnnouncedProducerListItem, err error) {
	url := q.ApiServer + "/api/v1/group/" + group_id + "/producers"
	q.HttpClient.R().SetResult(res).SetError(err).Get(url)
	return res, err
}

// GetAnnouncedGroupUser gets announced group user, /api/v1/group/{group_id}/announced/user/{sign_pubkey}
// return: HandlersAnnouncedUserListResult
func (q *Quorum) GetAnnouncedGroupUser(group_id string, sign_pubkey string) (res *model.HandlersAnnouncedUserListItem, err error) {
	url := q.ApiServer + "/api/v1/group/" + group_id + "/announced/user/" + sign_pubkey
	q.HttpClient.R().SetResult(res).SetError(err).Get(url)
	return res, err
}

// GetAnnouncedGroupUsers gets announced group users, /api/v1/group/{group_id}/announced/users
// return: HandlersAnnouncedUserListItem
func (q *Quorum) GetAnnouncedGroupUsers(group_id string) (res *model.HandlersAnnouncedUserListItem, err error) {
	url := q.ApiServer + "/api/v1/group/" + group_id + "/announced/users"
	q.HttpClient.R().SetResult(res).SetError(err).Get(url)
	return res, err
}
