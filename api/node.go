package api

import (
	"github.com/rumsystem/quorum/cmd/cli/api"
)

// Network get network info
func (c ReqClient) Network() (res *api.NetworkInfoStruct, err error) {
	url := c.Client.BaseURL + "/api/v1/network"
	c.Client.R().SetResult(res).SetError(err).Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Node get node info
func (c ReqClient) Node() (res *api.NodeInfoStruct, err error) {
	url := c.Client.BaseURL + "/api/v1/node"
	c.Client.R().SetResult(res).SetError(err).Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//Ping ping node
func (c ReqClient) Ping() (res *map[string]api.PingInfoItemStruct, err error) {
	url := c.Client.BaseURL + "/api/v1/network/peers/ping"
	ret := make(map[string]api.PingInfoItemStruct)
	c.Client.R().SetResult(&ret).SetError(err).Get(url)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// Stats get network stats summary
func (c ReqClient) Stats() ()
