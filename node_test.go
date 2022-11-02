package rumgosdk

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/tymon42/rum-go-sdk/model"
)

func TestQuorum_Node(t *testing.T) {
	type fields struct {
		ApiServer  string
		HttpClient *resty.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *model.HandlersNodeInfo
		wantErr bool
	}{
		{
			name: "Test Quorum Node",
			fields: fields{
				ApiServer:  "http://localhost:8002",
				HttpClient: resty.New(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quorum{
				ApiServer:  tt.fields.ApiServer,
				HttpClient: tt.fields.HttpClient,
			}
			got, err := q.GetNode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Quorum.Node() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quorum.Node() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuorum_GetNetwork(t *testing.T) {
	type fields struct {
		ApiServer  string
		HttpClient *resty.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *model.HandlersNetworkInfo
		wantErr bool
	}{
		{
			name: "Test Quorum Network",
			fields: fields{
				ApiServer:  "http://localhost:8002",
				HttpClient: resty.New(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quorum{
				ApiServer:  tt.fields.ApiServer,
				HttpClient: tt.fields.HttpClient,
			}
			got, err := q.GetNetwork()
			if (err != nil) != tt.wantErr {
				t.Errorf("Quorum.GetNetwork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quorum.GetNetwork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuorum_AddPeers(t *testing.T) {
	peers := []string{"/ip4/94.23.17.189/tcp/10666/p2p/16Uiu2HAmGTcDnhj3KVQUwVx8SGLyKBXQwfAxNayJdEwfsnUYKK4u"}
	type fields struct {
		ApiServer  string
		HttpClient *resty.Client
	}
	type args struct {
		peers []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test Quorum AddPeers",
			fields: fields{
				ApiServer:  "http://localhost:8002",
				HttpClient: resty.New(),
			},
			args: args{
				peers: peers,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quorum{
				ApiServer:  tt.fields.ApiServer,
				HttpClient: tt.fields.HttpClient,
			}
			if err := q.AddPeers(tt.args.peers); (err != nil) != tt.wantErr {
				t.Errorf("Quorum.AddPeers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuorum_PingPeers(t *testing.T) {
	type fields struct {
		ApiServer  string
		HttpClient *resty.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *model.ApiAddrProtoPair
		wantErr bool
	}{
		{
			name: "Test Quorum PingPeers",
			fields: fields{
				ApiServer:  "http://localhost:8002",
				HttpClient: resty.New(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quorum{
				ApiServer:  tt.fields.ApiServer,
				HttpClient: tt.fields.HttpClient,
			}
			got, err := q.PingPeers()
			if (err != nil) != tt.wantErr {
				t.Errorf("Quorum.PingPeers() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else {
				fmt.Println(got)
			}
		})
	}
}
