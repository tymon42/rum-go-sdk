package rumgosdk

import (
	"github.com/rumsystem/quorum/cmd/cli/api"
	qApi "github.com/rumsystem/quorum/pkg/chainapi/api"
	"github.com/rumsystem/quorum/pkg/chainapi/handlers"
	"strings"
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

	api.ApiServer = apiServer
	return &Quorum{
		ApiServer: ApiServer,
	}
}

// Get Node info
func (q *Quorum) Node() (*api.NodeInfoStruct, error) {
	return api.Node()
}

// Get Network info
func (q *Quorum) Network() (*api.NetworkInfoStruct, error) {
	return api.Network()
}

// Ping nodes
func (q *Quorum) Ping() (res *map[string]api.PingInfoItemStruct, err error) {
	return api.Ping()
}

// Get groups
func (q *Quorum) Groups() (groupsInfo *qApi.GroupInfoList, err error) {
	return api.Groups()
}

// Get Content
func (q *Quorum) Content(groupId string, opt api.PagerOpt) (contents *[]api.ContentStruct, err error) {
	return api.Content(groupId, opt)
}

// Force sync group
func (q *Quorum) ForceSyncGroup(groupId string) (syncRes *api.GroupForceSyncRetStruct, err error) {
	return api.ForceSyncGroup(groupId)
}

// Get public queue
func (q *Quorum) GetPubQueue(groupId string, trxId string, status string) (*handlers.PubQueueInfo, error) {
	return api.GetPubQueue(groupId, trxId, status)
}

// Public queue acknowledge
func (q *Quorum) PubQueueAck(trxIds []string) ([]string, error) {
	return api.PubQueueAck(trxIds)
}

// Is quorum content message?
func (q *Quorum) IsQuorumContentMessage(content api.ContentStruct) bool {
	return api.IsQuorumContentMessage(content)
}

// Is Quorum content user infomation?
func (q *Quorum) IsQuorumContentUserInfo(content api.ContentStruct) bool {
	return api.IsQuorumContentUserInfo(content)
}

// Add group config
func (q *Quorum) AddGroupConfig(groupId, key, tp, value, memo string) (*handlers.AppConfigResult, error) {
	return api.AddGroupConfig(groupId, key, tp, value, memo)
}

// Delete group config
func (q *Quorum) DelGroupConfig(groupId, key, tp, value, memo string) (*handlers.AppConfigResult, error) {
	return api.DelGroupConfig(groupId, key, tp, value, memo)
}

// Modify group config
func (q *Quorum) ModifyGroupConfig(action, groupId, key, tp, value, memo string) (*handlers.AppConfigResult, error) {
	return api.ModifyGroupConfig(action, groupId, key, tp, value, memo)
}

// Get group config list
func (q *Quorum) GetGroupConfigList(groupId string) ([]*handlers.AppConfigKeyListItem, error) {
	return api.GetGroupConfigList(groupId)
}

// Update chain config
func (q *Quorum) UpdateChainConfig(groupId, tp, config, memo string) (*handlers.ChainConfigResult, error) {
	return api.UpdateChainConfig(groupId, tp, config, memo)
}

// Get chain auth mode
func (q *Quorum) GetChainAuthMode(groupId, trxType string) (*handlers.TrxAuthItem, error) {
	return api.GetChainAuthMode(groupId, trxType)
}

// Get chain allow list
func (q *Quorum) GetChainAllowList(groupId string) ([]*handlers.ChainSendTrxRuleListItem, error) {
	return api.GetChainAllowList(groupId)
}

// Get chain deny list
func (q *Quorum) GetChainDenyList(groupId string) ([]*handlers.ChainSendTrxRuleListItem, error) {
	return api.GetChainDenyList(groupId)
}

// Nick
func (q *Quorum) Nick(groupId string, nick string) (*api.NickRespStruct, error) {
	return api.Nick(groupId, nick)
}

// Token apply
func (q *Quorum) TokenApply() (*api.TokenRespStruct, error) {
	return api.TokenApply()
}

// Create content
func (q *Quorum) CreateContent(groupId string, content string) (*api.ContentRespStruct, error) {
	return api.CreateContent(groupId, content)
}

// Create group
func (q *Quorum) CreateGroup(data api.CreateGroupReqStruct) ([]byte, error) {
	return api.CreateGroup(data)
}

// Get group seed
func (q *Quorum) GetGroupSeed(gid string) (*handlers.GroupSeed, error) {
	return api.GetGroupSeed(gid)
}

// Do backup
func (q *Quorum) DoBackup() (*api.BackupResult, error) {
	return api.DoBackup()
}

// Leave group
func (q *Quorum) LeaveGroup(groupId string) (*api.GroupLeaveRetStruct, error) {
	return api.LeaveGroup(groupId)
}

// Delete group
func (q *Quorum) DeleteGroup(groupId string) (*api.GroupDelRetStruct, error) {
	return api.DelGroup(groupId)
}

// Get Trx infomation
func (q *Quorum) TrxInfo(groupId string, trxId string) (trx *api.TrxStruct, err error) {
	return api.TrxInfo(groupId, trxId)
}

// Join group
func (q *Quorum) JoinGroup(seed string) (*api.JoinRespStruct, error) {
	return api.JoinGroup(seed)
}

// Get block by Id
func (q *Quorum) GetBlockById(groupId string, blockId string) (*api.BlockStruct, error) {
	return api.GetBlockById(groupId, blockId)
}

// Announced User
func (q *Quorum) AnnouncedUsers(groupId string) ([]*handlers.AnnouncedUserListItem, error) {
	return api.AnnouncedUsers(groupId)
}

// Announced producers
func (q *Quorum) AnnouncedProducers(groupId string) ([]*handlers.AnnouncedProducerListItem, error) {
	return api.AnnouncedProducers(groupId)
}

// Approve announced user
func (q *Quorum) ApproveAnnouncedUser(groupId string, user *handlers.AnnouncedUserListItem, removal bool) (*api.ApproveGrpUserResult, error) {
	return api.ApproveAnnouncedUser(groupId, user, removal)
}

// Approve announced producer
func (q *Quorum) ApproveAnnouncedProducer(groupId string, user *handlers.AnnouncedProducerListItem, removal bool) (*handlers.GrpProducerResult, error) {
	return api.ApproveAnnouncedProducer(groupId, user, removal)
}
