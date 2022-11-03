package rumgosdk

import (
	"github.com/tymon42/rum-go-sdk/model"
)

// AnnounceUserPubkey announces user pubkey, /api/v1/group/announce
// param: handlersAnnounceParam
// return: HandlersAnnounceResult
func (q *Quorum) AnnounceUserPubkey(param model.HandlersAnnounceParam) (*model.HandlersAnnounceResult, error) {
	url := q.ApiServer + "/api/v1/group/announce"
	res := &model.HandlersAnnounceResult{}
	_, err := q.HttpClient.R().SetBody(param).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateProfile updates profile, /api/v1/group/profile
// param: pbActivity
// return: HandlersUpdateProfileResult
func (q *Quorum) UpdateProfile(param model.PbActivity) (*model.HandlersUpdateProfileResult, error) {
	url := q.ApiServer + "/api/v1/group/profile"
	res := &model.HandlersUpdateProfileResult{}
	_, err := q.HttpClient.R().SetBody(param).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetAnnouncedGroupProducer gets announced group producer, /api/v1/group/{group_id}/producers
// return []HandlersAnnouncedProdcerListItempbGroupProducer
func (q *Quorum) GetAnnouncedGroupProducer(group_id string) ([]model.HandlersAnnouncedProducerListItem, error) {
	url := q.ApiServer + "/api/v1/group/" + group_id + "/producers"
	res := []model.HandlersAnnouncedProducerListItem{}
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetAnnouncedGroupUser gets announced group user, /api/v1/group/{group_id}/announced/user/{sign_pubkey}
// return: HandlersAnnouncedUserListResult
func (q *Quorum) GetAnnouncedGroupUser(group_id string, sign_pubkey string) (*model.HandlersAnnouncedUserListItem, error) {
	url := q.ApiServer + "/api/v1/group/" + group_id + "/announced/user/" + sign_pubkey
	res := &model.HandlersAnnouncedUserListItem{}
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetAnnouncedGroupUsers gets announced group users, /api/v1/group/{group_id}/announced/users
// return: HandlersAnnouncedUserListItem
func (q *Quorum) GetAnnouncedGroupUsers(group_id string) (*model.HandlersAnnouncedUserListItem, error) {
	url := q.ApiServer + "/api/v1/group/" + group_id + "/announced/users"
	res := &model.HandlersAnnouncedUserListItem{}
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}
