package rumgosdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/rumsystem/quorum/cmd/cli/api"
	qApi "github.com/rumsystem/quorum/pkg/chainapi/api"
	"github.com/rumsystem/quorum/pkg/chainapi/handlers"
	"github.com/tymon42/rum-go-sdk/http"
)

type Quorum struct {
	ApiServer string
}

// Connect a new Quorum instance
func Connect(apiServer string) *Quorum {
	ApiServer := apiServer
	if len(apiServer) > 0 {
		if strings.HasPrefix(apiServer, "https") {
			ApiServer = apiServer
		} else {
			ApiServer = "https://" + apiServer
		}
	}

	return &Quorum{
		ApiServer: ApiServer,
	}
}

// Get Node info
func (q *Quorum) Node() (*api.NodeInfoStruct, error) {
	url := q.ApiServer + "/api/v1/node"
	ret := api.NodeInfoStruct{}
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Get Network info
func (q *Quorum) Network() (*api.NetworkInfoStruct, error) {
	url := q.ApiServer + "/api/v1/network"
	ret := api.NetworkInfoStruct{}
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Ping nodes
func (q *Quorum) Ping() (res *map[string]api.PingInfoItemStruct, err error) {
	url := q.ApiServer + "/api/v1/network/peers/ping"
	ret := make(map[string]api.PingInfoItemStruct)
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Get groups
func (q *Quorum) Groups() (groupsInfo *qApi.GroupInfoList, err error) {
	url := q.ApiServer + "/api/v1/groups"
	ret := qApi.GroupInfoList{}
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Get Content
func (q *Quorum) Content(groupId string, num int, opt api.PagerOpt) (contents *[]api.ContentStruct, err error) {
	url := fmt.Sprintf(
		"%s/app/api/v1/group/%s/content?num=%d&reverse=%v", q.ApiServer, groupId, num, opt.Reverse)
	if opt.StartTrxId != "" {
		url = fmt.Sprintf("%s&starttrx=%s", url, opt.StartTrxId)
	}
	ret := []api.ContentStruct{}
	body, err := http.HttpPost(url, []byte("{}"))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Force sync group
func (q *Quorum) ForceSyncGroup(groupId string) (syncRes *api.GroupForceSyncRetStruct, err error) {
	url := fmt.Sprintf(
		"%s/api/v1/group/%s/startsync", q.ApiServer, groupId)
	ret := api.GroupForceSyncRetStruct{}
	body, err := http.HttpPost(url, []byte(""))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil || ret.GroupId == "" {
		return nil, errors.New(string(body))
	}
	if ret.Error != "" {
		return nil, errors.New(ret.Error)
	}
	return &ret, nil
}

// Get public queue
func (q *Quorum) GetPubQueue(groupId string, trxId string, status string) (*handlers.PubQueueInfo, error) {
	url := fmt.Sprintf(
		"%s/api/v1/group/%s/pubqueue?trx=%s&status=%s", q.ApiServer, groupId, trxId, status)
	ret := handlers.PubQueueInfo{}
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil || ret.GroupId == "" {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Public queue acknowledge
func (q *Quorum) PubQueueAck(trxIds []string) ([]string, error) {
	url := fmt.Sprintf("%s/api/v1/trx/ack", q.ApiServer)
	param := qApi.PubQueueAckPayload{trxIds}
	payload, err := json.Marshal(&param)
	if err != nil {
		return nil, err
	}
	ret := []string{}
	body, err := http.HttpPost(url, payload)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return ret, nil
}

// Is quorum content message?
func (q *Quorum) IsQuorumContentMessage(content api.ContentStruct) bool {
	// only support Note
	if content.TypeUrl == "quorum.pb.Object" {
		innerType, hasKey := content.Content["type"]
		if hasKey && innerType == "Note" {
			return true
		}
	}
	return false
}

// Is Quorum content user infomation?
func (q *Quorum) IsQuorumContentUserInfo(content api.ContentStruct) bool {
	// only support Note
	if content.TypeUrl == "quorum.pb.Person" {
		_, hasKey := content.Content["name"]
		if hasKey {
			return true
		}
	}
	return false
}

// Add group config
func (q *Quorum) AddGroupConfig(groupId, key, tp, value, memo string) (*handlers.AppConfigResult, error) {
	return q.ModifyGroupConfig("add", groupId, key, tp, value, memo)
}

// Delete group config
func (q *Quorum) DelGroupConfig(groupId, key, tp, value, memo string) (*handlers.AppConfigResult, error) {
	return q.ModifyGroupConfig("del", groupId, key, tp, value, memo)
}

// Modify group config
func (q *Quorum) ModifyGroupConfig(action, groupId, key, tp, value, memo string) (*handlers.AppConfigResult, error) {
	data := handlers.AppConfigParam{
		Action:  action,
		GroupId: groupId,
		Name:    key,
		Type:    tp,
		Value:   value,
		Memo:    memo,
	}
	json_data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	url := q.ApiServer + "/api/v1/group/appconfig"
	ret := handlers.AppConfigResult{}
	body, err := http.HttpPost(url, json_data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Get group config list
func (q *Quorum) GetGroupConfigList(groupId string) ([]*handlers.AppConfigKeyListItem, error) {
	url := fmt.Sprintf("%s/api/v1/group/%s/config/keylist", q.ApiServer, groupId)
	ret := []*handlers.AppConfigKeyListItem{}
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return ret, nil
}

// Get Group Config
func (q *Quorum) GetGroupConfig(groupId, key string) (*handlers.AppConfigKeyItem, error) {
	url := fmt.Sprintf("%s/api/v1/group/%s/config/%s", q.ApiServer, groupId, key)
	ret := handlers.AppConfigKeyItem{}
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Update chain config
func (q *Quorum) UpdateChainConfig(groupId, tp, config, memo string) (*handlers.ChainConfigResult, error) {
	data := handlers.ChainConfigParams{
		GroupId: groupId,
		Type:    tp,
		Config:  config,
		Memo:    memo,
	}
	json_data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	url := q.ApiServer + "/api/v1/group/chainconfig"
	ret := handlers.ChainConfigResult{}
	body, err := http.HttpPost(url, json_data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Get chain auth mode
func (q *Quorum) GetChainAuthMode(groupId, trxType string) (*handlers.TrxAuthItem, error) {
	url := fmt.Sprintf("%s/api/v1/group/%s/trx/auth/%s", q.ApiServer, groupId, trxType)
	ret := handlers.TrxAuthItem{}
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Get chain allow list
func (q *Quorum) GetChainAllowList(groupId string) ([]*handlers.ChainSendTrxRuleListItem, error) {
	url := fmt.Sprintf("%s/api/v1/group/%s/trx/allowlist", q.ApiServer, groupId)
	ret := []*handlers.ChainSendTrxRuleListItem{}
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return ret, nil
}

// Get chain deny list
func (q *Quorum) GetChainDenyList(groupId string) ([]*handlers.ChainSendTrxRuleListItem, error) {
	url := fmt.Sprintf("%s/api/v1/group/%s/trx/denylist", q.ApiServer, groupId)
	ret := []*handlers.ChainSendTrxRuleListItem{}
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return ret, nil
}

// Nick
func (q *Quorum) Nick(groupId string, nick string) (*api.NickRespStruct, error) {
	data := api.NickReqStruct{
		Person: api.QuorumPersonStruct{
			Name: nick,
		},
		Target: api.QuorumTargetStruct{
			Id:   groupId,
			Type: "Group",
		},
		Type: "Update",
	}
	json_data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	url := q.ApiServer + "/api/v1/group/profile"
	ret := api.NickRespStruct{}
	body, err := http.HttpPost(url, json_data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Token apply
func (q *Quorum) TokenApply() (*api.TokenRespStruct, error) {
	url := q.ApiServer + "/app/api/v1/token/apply"
	ret := api.TokenRespStruct{}
	body, err := http.HttpPost(url, []byte(""))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Create content
func (q *Quorum) CreateContent(groupId string, name string, content string) (*api.ContentRespStruct, error) {
	data := api.ContentReqStruct{
		Object: api.ContentReqObjectStruct{
			Content: content,
			Name:    name,
			Type:    "Note",
		},
		Target: api.ContentReqTargetStruct{
			Id:   groupId,
			Type: "Group",
		},
		Type: "Add",
	}
	url := q.ApiServer + "/api/v1/group/content"
	ret := api.ContentRespStruct{}
	json_data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body, err := http.HttpPost(url, json_data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Create group
func (q *Quorum) CreateGroup(data api.CreateGroupReqStruct) ([]byte, error) {
	json_data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	url := q.ApiServer + "/api/v1/group"
	return http.HttpPost(url, json_data)
}

// Get group seed
func (q *Quorum) GetGroupSeed(gid string) (*handlers.GroupSeed, error) {
	url := fmt.Sprintf("%s/api/v1/group/%s/seed", q.ApiServer, gid)
	ret := handlers.GroupSeed{}
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Do backup
func (q *Quorum) DoBackup() (*api.BackupResult, error) {
	url := fmt.Sprintf("%s/api/v1/backup", q.ApiServer)
	ret := api.BackupResult{}
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Leave group
func (q *Quorum) LeaveGroup(gid string) (*api.GroupLeaveRetStruct, error) {
	data := api.LeaveGroupReqStruct{gid}
	url := q.ApiServer + "/api/v1/group/leave"
	ret := api.GroupLeaveRetStruct{}

	json_data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body, err := http.HttpPost(url, json_data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}

	return &ret, nil
}

// Delete group
func (q *Quorum) DeleteGroup(groupId string) (*api.GroupDelRetStruct, error) {
	data := api.DeleteGroupReqStruct{groupId}
	url := q.ApiServer + "/api/v1/group"
	ret := api.GroupDelRetStruct{}
	json_data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	body, err := http.HttpDelete(url, json_data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Get Trx infomation
func (q *Quorum) TrxInfo(groupId string, trxId string) (trx *api.TrxStruct, err error) {
	url := fmt.Sprintf("%s/api/v1/trx/%s/%s", q.ApiServer, groupId, trxId)
	ret := api.TrxStruct{}
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Join group
func (q *Quorum) JoinGroup(seed string) (*api.JoinRespStruct, error) {
	url := q.ApiServer + "/api/v1/group/join"
	ret := api.JoinRespStruct{}
	body, err := http.HttpPost(url, []byte(seed))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Get block by Id
func (q *Quorum) GetBlockById(groupId string, id string) (*api.BlockStruct, error) {
	url := fmt.Sprintf("%s/api/v1/block/%s/%s", q.ApiServer, groupId, id)
	ret := api.BlockStruct{}
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return &ret, nil
}

// Announced User
func (q *Quorum) AnnouncedUsers(groupId string) ([]*handlers.AnnouncedUserListItem, error) {
	url := fmt.Sprintf("%s/api/v1/group/%s/announced/users", q.ApiServer, groupId)
	ret := []*handlers.AnnouncedUserListItem{}
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return ret, nil
}

// Announced producers
func (q *Quorum) AnnouncedProducers(groupId string) ([]*handlers.AnnouncedProducerListItem, error) {
	url := fmt.Sprintf("%s/api/v1/group/%s/announced/producers", q.ApiServer, groupId)
	ret := []*handlers.AnnouncedProducerListItem{}
	body, err := http.HttpGet(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}
	return ret, nil
}

// Approve announced user
func (q *Quorum) ApproveAnnouncedUser(groupId string, user *handlers.AnnouncedUserListItem, removal bool) (*api.ApproveGrpUserResult, error) {
	ret := &api.ApproveGrpUserResult{}
	url := q.ApiServer + "/api/v1/group/user"

	action := "add"
	if removal {
		action = "remove"
	}

	data := api.ApproveGrpUserParam{
		Action:     action,
		UserPubkey: user.AnnouncedSignPubkey,
		GroupId:    groupId,
		Memo:       "by cli",
	}
	json_data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body, err := http.HttpPost(url, json_data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}

	return ret, nil
}

// Approve announced producer
func (q *Quorum) ApproveAnnouncedProducer(groupId string, user *handlers.AnnouncedProducerListItem, removal bool) (*handlers.GrpProducerResult, error) {
	ret := &handlers.GrpProducerResult{}
	url := q.ApiServer + "/api/v1/group/producer"

	action := "add"
	if removal {
		action = "remove"
	}

	data := handlers.GrpProducerParam{
		Action:         action,
		ProducerPubkey: user.AnnouncedPubkey,
		GroupId:        groupId,
		Memo:           "by cli",
	}
	json_data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body, err := http.HttpPost(url, json_data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, errors.New(string(body))
	}

	return ret, nil
}
