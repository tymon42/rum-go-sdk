package rumgosdk

import "github.com/tymon42/rum-go-sdk/model"

// GetNode info, /api/v1/node
func (q *Quorum) GetNode() (res *model.HandlersNodeInfo, err error) {
	url := q.ApiServer + "/api/v1/node"
	q.HttpClient.R().SetResult(res).SetError(err).Get(url)
	return res, err
}

// GetNetwork gets node network info, /api/v1/network
func (q *Quorum) GetNetwork() (res *model.HandlersNetworkInfo, err error) {
	url := q.ApiServer + "/api/v1/network"
	q.HttpClient.R().SetResult(res).SetError(err).Get(url)
	return res, err
}

// AddPeers adds peers to the node, /api/v1/network/peers
func (q *Quorum) AddPeers(peers []string) (err error) {
	url := q.ApiServer + "/api/v1/network/peers"
	q.HttpClient.R().SetBody(peers).SetError(err).Post(url)
	return err
}

// PingPeers pings peers, /api/v1/network/peers/ping
func (q *Quorum) PingPeers() (res *model.ApiAddrProtoPair, err error) {
	url := q.ApiServer + "/api/v1/network/peers/ping"
	q.HttpClient.R().SetResult(res).SetError(err).Get(url)
	return res, err
}

// AddRelayServers adds relay servers to the node, /api/v1/network/relay
func (q *Quorum) AddRelayServers(relayServers []string) (res *model.ApiAddRelayServersResp, err error) {
	url := q.ApiServer + "/api/v1/network/relay"
	q.HttpClient.R().SetBody(relayServers).SetResult(res).Post(url)
	return res, err
}

// PubsubPing pings pubsub, /api/v1/psping
func (q *Quorum) PubsubPing(data model.ApiPsPingParam) (res *model.HandlersPingResp, err error) {
	url := q.ApiServer + "/api/v1/psping"
	q.HttpClient.R().SetBody(data).SetResult(res).Post(url)
	return res, err
}
