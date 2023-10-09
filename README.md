# rum-go-sdk  
> ‼️ Because of the high speed of quorum evolution, this SDK may not be 100% suitable with latest quorum version.  

# What is it?  
This is a third party quorum SDK in golang. 

## What is rum?  
Rum System is flexable communicating protocol. You can treat it like well-backuped DB for the app that you want to build. [More about Rum](https://github.com/rumsystem/quorum)  

# Usage 
install: `go get github.com/tymon42/rum-go-sdk`

<!-- ```
import (
    "github.com/tymon42/rum-go-sdk"
)

rum_client := rumgosdk.Connect('http://127.0.0.1:8002')

rum_client.Node()
``` -->

<!-- ## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
``` -->


<!-- ## Documentation for API Endpoints

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AppsApi* | [**AppApiV1GroupGroupIdContentGet**](docs/AppsApi.md#appapiv1groupgroupidcontentget) | **Get** /app/api/v1/group/{group_id}/content | GetGroupContents
*AppsApi* | [**AppApiV1TokenDelete**](docs/AppsApi.md#appapiv1tokendelete) | **Delete** /app/api/v1/token | RemoveToken
*AppsApi* | [**AppApiV1TokenListGet**](docs/AppsApi.md#appapiv1tokenlistget) | **Get** /app/api/v1/token/list | ListToken
*AppsApi* | [**AppApiV1TokenPost**](docs/AppsApi.md#appapiv1tokenpost) | **Post** /app/api/v1/token | CreateToken
*AppsApi* | [**AppApiV1TokenRefreshPost**](docs/AppsApi.md#appapiv1tokenrefreshpost) | **Post** /app/api/v1/token/refresh | RefreshToken
*AppsApi* | [**AppApiV1TokenRevokePost**](docs/AppsApi.md#appapiv1tokenrevokepost) | **Post** /app/api/v1/token/revoke | RevokeToken
*ChainApi* | [**ApiV1BlockGroupIdEpochGet**](docs/ChainApi.md#apiv1blockgroupidepochget) | **Get** /api/v1/block/{group_id}/{epoch} | GetBlock
*ChainApi* | [**ApiV1TrxGroupIdTrxIdGet**](docs/ChainApi.md#apiv1trxgroupidtrxidget) | **Get** /api/v1/trx/{group_id}/{trx_id} | GetTrx
*GroupApi* | [**ApiV1GroupGroupIdStartsyncPost**](docs/GroupApi.md#apiv1groupgroupidstartsyncpost) | **Post** /api/v1/group/{group_id}/startsync | StartSync
*GroupsApi* | [**ApiV1GroupClearPost**](docs/GroupsApi.md#apiv1groupclearpost) | **Post** /api/v1/group/clear | ClearGroupData
*GroupsApi* | [**ApiV1GroupGroupIdContentPost**](docs/GroupsApi.md#apiv1groupgroupidcontentpost) | **Post** /api/v1/group/{group_id}/content | PostToGroup
*GroupsApi* | [**ApiV1GroupGroupIdGet**](docs/GroupsApi.md#apiv1groupgroupidget) | **Get** /api/v1/group/{group_id} | GetGroupById
*GroupsApi* | [**ApiV1GroupGroupIdSeedGet**](docs/GroupsApi.md#apiv1groupgroupidseedget) | **Get** /api/v1/group/{group_id}/seed | Get group seed
*GroupsApi* | [**ApiV1GroupLeavePost**](docs/GroupsApi.md#apiv1groupleavepost) | **Post** /api/v1/group/leave | LeaveGroup
*GroupsApi* | [**ApiV1GroupPost**](docs/GroupsApi.md#apiv1grouppost) | **Post** /api/v1/group | CreateGroupUrl
*GroupsApi* | [**ApiV1GroupsGet**](docs/GroupsApi.md#apiv1groupsget) | **Get** /api/v1/groups | GetGroups
*GroupsApi* | [**ApiV2GroupJoinPost**](docs/GroupsApi.md#apiv2groupjoinpost) | **Post** /api/v2/group/join | JoinGroup
*KeystoreApi* | [**ApiV1KeystoreSigntxPost**](docs/KeystoreApi.md#apiv1keystoresigntxpost) | **Post** /api/v1/keystore/signtx | SignTx
*LightNodeApi* | [**ApiV1NodeGroupIdAnnouncePost**](docs/LightNodeApi.md#apiv1nodegroupidannouncepost) | **Post** /api/v1/node/{group_id}/announce | NSdkAnnounce
*LightNodeApi* | [**ApiV1NodeGroupIdAnnouncedProducerGet**](docs/LightNodeApi.md#apiv1nodegroupidannouncedproducerget) | **Get** /api/v1/node/{group_id}/announced/producer | GetNSdkAnnouncedProducer
*LightNodeApi* | [**ApiV1NodeGroupIdAnnouncedUserGet**](docs/LightNodeApi.md#apiv1nodegroupidannounceduserget) | **Get** /api/v1/node/{group_id}/announced/user | GetNSdkAnnouncedUser
*LightNodeApi* | [**ApiV1NodeGroupIdAppconfigByKeyGet**](docs/LightNodeApi.md#apiv1nodegroupidappconfigbykeyget) | **Get** /api/v1/node/{group_id}/appconfig/by/{key} | GetNSdkAppconfigByKey
*LightNodeApi* | [**ApiV1NodeGroupIdAppconfigKeylistGet**](docs/LightNodeApi.md#apiv1nodegroupidappconfigkeylistget) | **Get** /api/v1/node/{group_id}/appconfig/keylist | GetNSdkAppconfigKeylist
*LightNodeApi* | [**ApiV1NodeGroupIdAuthAlwlistGet**](docs/LightNodeApi.md#apiv1nodegroupidauthalwlistget) | **Get** /api/v1/node/{group_id}/auth/alwlist | GetNSdkAllowList
*LightNodeApi* | [**ApiV1NodeGroupIdAuthByTrxTypeGet**](docs/LightNodeApi.md#apiv1nodegroupidauthbytrxtypeget) | **Get** /api/v1/node/{group_id}/auth/by/{trx_type} | GetNSdkAuthType
*LightNodeApi* | [**ApiV1NodeGroupIdAuthDenylistGet**](docs/LightNodeApi.md#apiv1nodegroupidauthdenylistget) | **Get** /api/v1/node/{group_id}/auth/denylist | GetNSdkDenyList
*LightNodeApi* | [**ApiV1NodeGroupIdGroupctnGet**](docs/LightNodeApi.md#apiv1nodegroupidgroupctnget) | **Get** /api/v1/node/{group_id}/groupctn | GetNSdkContent
*LightNodeApi* | [**ApiV1NodeGroupIdInfoGet**](docs/LightNodeApi.md#apiv1nodegroupidinfoget) | **Get** /api/v1/node/{group_id}/info | GetNSdkGroupInfo
*LightNodeApi* | [**ApiV1NodeGroupIdProducersGet**](docs/LightNodeApi.md#apiv1nodegroupidproducersget) | **Get** /api/v1/node/{group_id}/producers | GetNSdkGroupProducers
*LightNodeApi* | [**ApiV1NodeGroupIdTrxPost**](docs/LightNodeApi.md#apiv1nodegroupidtrxpost) | **Post** /api/v1/node/{group_id}/trx | NSdkSendTrx
*LightNodeApi* | [**ApiV1NodeGroupIdUserencryptpubkeysGet**](docs/LightNodeApi.md#apiv1nodegroupiduserencryptpubkeysget) | **Get** /api/v1/node/{group_id}/userencryptpubkeys | GetNSdkUserEncryptPubKeys
*ManagementApi* | [**ApiV1GroupAppconfigPost**](docs/ManagementApi.md#apiv1groupappconfigpost) | **Post** /api/v1/group/appconfig | MgrAppConfig
*ManagementApi* | [**ApiV1GroupChainconfigPost**](docs/ManagementApi.md#apiv1groupchainconfigpost) | **Post** /api/v1/group/chainconfig | chainconfig
*ManagementApi* | [**ApiV1GroupGroupIdAppconfigKeyGet**](docs/ManagementApi.md#apiv1groupgroupidappconfigkeyget) | **Get** /api/v1/group/{group_id}/appconfig/{key} | GetAppConfigItem
*ManagementApi* | [**ApiV1GroupGroupIdAppconfigKeylistGet**](docs/ManagementApi.md#apiv1groupgroupidappconfigkeylistget) | **Get** /api/v1/group/{group_id}/appconfig/keylist | GetAppConfigKey
*ManagementApi* | [**ApiV1GroupGroupIdProducersGet**](docs/ManagementApi.md#apiv1groupgroupidproducersget) | **Get** /api/v1/group/{group_id}/producers | GetGroupProducers
*ManagementApi* | [**ApiV1GroupGroupIdTrxAllowlistGet**](docs/ManagementApi.md#apiv1groupgroupidtrxallowlistget) | **Get** /api/v1/group/{group_id}/trx/allowlist | GetChainTrxAllowList
*ManagementApi* | [**ApiV1GroupGroupIdTrxAuthTrxTypeGet**](docs/ManagementApi.md#apiv1groupgroupidtrxauthtrxtypeget) | **Get** /api/v1/group/{group_id}/trx/auth/{trx_type} | GetChainTrxAuthMode
*ManagementApi* | [**ApiV1GroupGroupIdTrxDenylistGet**](docs/ManagementApi.md#apiv1groupgroupidtrxdenylistget) | **Get** /api/v1/group/{group_id}/trx/denylist | GetDeniedUserList
*ManagementApi* | [**ApiV1GroupProducerPost**](docs/ManagementApi.md#apiv1groupproducerpost) | **Post** /api/v1/group/producer | AddProducer
*ManagementApi* | [**ApiV1GroupUserPost**](docs/ManagementApi.md#apiv1groupuserpost) | **Post** /api/v1/group/user | AddUsers
*NodeApi* | [**ApiV1NetworkGet**](docs/NodeApi.md#apiv1networkget) | **Get** /api/v1/network | GetNetwork
*NodeApi* | [**ApiV1NetworkPeersPost**](docs/NodeApi.md#apiv1networkpeerspost) | **Post** /api/v1/network/peers | AddPeers
*NodeApi* | [**ApiV1NetworkRelayPost**](docs/NodeApi.md#apiv1networkrelaypost) | **Post** /api/v1/network/relay | AddRelayServers
*NodeApi* | [**ApiV1NodeGet**](docs/NodeApi.md#apiv1nodeget) | **Get** /api/v1/node | GetNodeInfo
*ToolsApi* | [**ApiV1ToolsPubkeytoaddrPost**](docs/ToolsApi.md#apiv1toolspubkeytoaddrpost) | **Post** /api/v1/tools/pubkeytoaddr | PubkeyToEthaddr
*UserApi* | [**ApiV1GroupAnnouncePost**](docs/UserApi.md#apiv1groupannouncepost) | **Post** /api/v1/group/announce | AnnounceUserPubkey
*UserApi* | [**ApiV1GroupGroupIdAnnouncedProducersGet**](docs/UserApi.md#apiv1groupgroupidannouncedproducersget) | **Get** /api/v1/group/{group_id}/announced/producers | GetAnnouncedGroupProducer
*UserApi* | [**ApiV1GroupGroupIdAnnouncedUserSignPubkeyGet**](docs/UserApi.md#apiv1groupgroupidannouncedusersignpubkeyget) | **Get** /api/v1/group/{group_id}/announced/user/{sign_pubkey} | GetAnnouncedGroupUser
*UserApi* | [**ApiV1GroupGroupIdAnnouncedUsersGet**](docs/UserApi.md#apiv1groupgroupidannouncedusersget) | **Get** /api/v1/group/{group_id}/announced/users | GetAnnouncedGroupUsers -->
