/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rum_sdk

// PbBlock struct for PbBlock
type PbBlock struct {
	BlockHash      []int64 `json:"BlockHash,omitempty"`
	BlockId        int64   `json:"BlockId,omitempty"`
	Epoch          int64   `json:"Epoch,omitempty"`
	GroupId        string  `json:"GroupId,omitempty"`
	PrevHash       []int64 `json:"PrevHash,omitempty"`
	ProducerPubkey string  `json:"ProducerPubkey,omitempty"`
	ProducerSign   []int64 `json:"ProducerSign,omitempty"`
	Sudo           bool    `json:"Sudo,omitempty"`
	TimeStamp      string  `json:"TimeStamp,omitempty"`
	Trxs           []PbTrx `json:"Trxs,omitempty"`
}
