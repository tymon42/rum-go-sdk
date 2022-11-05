package rumgosdk

import (
	"github.com/tymon42/rum-go-sdk/model"
)

// MgrAppConfig set app config
// Param: handlersAppConfigParam
// Return: handlersAppConfigResult
func (q *Quorum) MgrAppConfig(param *model.HandlersAppConfigParam) (*model.HandlersAppConfigResult, error) {
	url := q.ApiServer + "/api/v1/group/appconfig"
	res := &model.HandlersAppConfigResult{}
	_, err := q.HttpClient.R().SetBody(param).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, err
}

// ChainConfig set chain config
// Param: handlersChainConfigParams
// Return: handlersChainConfigResult
func (q *Quorum) ChainConfig(param *model.HandlersChainConfigParams) (*model.HandlersChainConfigResult, error) {
	url := q.ApiServer + "/api/v1/group/chainconfig"
	res := &model.HandlersChainConfigResult{}
	_, err := q.HttpClient.R().SetBody(param).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, err
}

// AddProducer add a peer to the group producer list
// Param: handlersGrpProducerParam
// Return: handlersGrpProducerResult
func (q *Quorum) AddProducer(param *model.HandlersGrpProducerParam) (*model.HandlersGrpProducerResult, error) {
	url := q.ApiServer + "/api/v1/group/producer"
	res := &model.HandlersGrpProducerResult{}
	_, err := q.HttpClient.R().SetBody(param).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, err
}

// AddUsers add a user to a private group users list
// Param: apiGrpUserParam
// Return: apiGrpUserResult
func (q *Quorum) AddUsers(param *model.ApiGrpUserParam) (*model.ApiGrpUserResult, error) {
	url := q.ApiServer + "/api/v1/group/user"
	res := &model.ApiGrpUserResult{}
	_, err := q.HttpClient.R().SetBody(param).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, err
}

// GetAppConfigKey get app config key list, /api/v1/group/{group_id}/config/keylist
// Param: 1. group_id
// Return: handlersAppConfigListItemResult
func (q *Quorum) GetAppConfigKey(groupID string) ([]*model.HandlersAppConfigKeyListItem, error) {
	url := q.ApiServer + "/api/v1/group/" + groupID + "/config/keylist"
	var res []*model.HandlersAppConfigKeyListItem
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, err
}

// GetAppConfigItem get app config item, /api/v1/group/{group_id}/config/{key}
// Param: 1. group_id 2. key
// Return: handlersAppConfigKeyIterm
func (q *Quorum) GetAppConfigItem(groupID, key string) (*model.HandlersAppConfigKeyItem, error) {
	url := q.ApiServer + "/api/v1/group/" + groupID + "/config/" + key
	res := &model.HandlersAppConfigKeyItem{}
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, err
}

// GetGroupProducers get the list of group producers, /api/v1/group/{group_id}/producers
// Param: 1. group_id
// Return: []handlersProducerListItem
func (q *Quorum) GetGroupProducers(groupID string) ([]*model.HandlersProducerListItem, error) {
	url := q.ApiServer + "/api/v1/group/" + groupID + "/producers"
	var res []*model.HandlersProducerListItem
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, err
}

// GetChainTrxAllowList gets group allow list, /api/v1/group/{group_id}/trx/allowlist
// Param: 1. group_id
// Return: []handlersChainSendTrxRuleListItem
func (q *Quorum) GetChainTrxAllowList(groupID string) ([]*model.HandlersChainSendTrxRuleListItem, error) {
	url := q.ApiServer + "/api/v1/group/" + groupID + "/trx/allowlist"
	var res []*model.HandlersChainSendTrxRuleListItem
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, err
}

// GetChainTrxAuthMode gets group auth mode, /api/v1/group/{group_id}/trx/auth/{trx_type}
// Param: 1. group_id 2. trx_type
// Return: []handlersTrxAuthItem
func (q *Quorum) GetChainTrxAuthMode(groupID, trxType string) ([]*model.HandlersTrxAuthItem, error) {
	url := q.ApiServer + "/api/v1/group/" + groupID + "/trx/auth/" + trxType
	var res []*model.HandlersTrxAuthItem
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, err
}

// GetDeniedUserList gets the list of denied users, /api/v1/group/{group_id}/trx/denylist
// Param: 1. group_id
// Return: []handlersChainSendTrxRuleListItem
func (q *Quorum) GetDeniedUserList(groupID string) ([]*model.HandlersChainSendTrxRuleListItem, error) {
	url := q.ApiServer + "/api/v1/group/" + groupID + "/trx/denylist"
	var res []*model.HandlersChainSendTrxRuleListItem
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, err
}
