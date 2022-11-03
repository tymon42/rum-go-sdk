package rumgosdk

import "github.com/tymon42/rum-go-sdk/model"

// GetNode info, /api/v1/node
func (q *Quorum) GetNode() (*model.HandlersNodeInfo, error) {
	url := q.ApiServer + "/api/v1/node"
	res := &model.HandlersNodeInfo{}
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetNetwork gets node network info, /api/v1/network
func (q *Quorum) GetNetwork() (*model.HandlersNetworkInfo, error) {
	url := q.ApiServer + "/api/v1/network"
	res := &model.HandlersNetworkInfo{}
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AddPeers adds peers to the node, /api/v1/network/peers
func (q *Quorum) AddPeers(peers []string) error {
	url := q.ApiServer + "/api/v1/network/peers"
	_, err := q.HttpClient.R().SetBody(peers).Post(url)
	return err
}

// PingPeers pings peers, /api/v1/network/peers/ping
func (q *Quorum) PingPeers() (*model.ApiAddrProtoPair, error) {
	url := q.ApiServer + "/api/v1/network/peers/ping"
	res := &model.ApiAddrProtoPair{}
	_, err := q.HttpClient.R().SetResult(res).Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AddRelayServers adds relay servers to the node, /api/v1/network/relay
func (q *Quorum) AddRelayServers(relayServers []string) (*model.ApiAddRelayServersResp, error) {
	url := q.ApiServer + "/api/v1/network/relay"
	res := &model.ApiAddRelayServersResp{}
	_, err := q.HttpClient.R().SetBody(relayServers).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PubsubPing pings pubsub, /api/v1/psping
func (q *Quorum) PubsubPing(data model.ApiPsPingParam) (*model.HandlersPingResp, error) {
	url := q.ApiServer + "/api/v1/psping"
	res := &model.HandlersPingResp{}
	_, err := q.HttpClient.R().SetBody(data).SetResult(res).Post(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}
