/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rum_sdk

// PbAnnounceItem struct for PbAnnounceItem
type PbAnnounceItem struct {
	Action             int64  `json:"Action,omitempty"`
	AnnouncerSignature string `json:"AnnouncerSignature,omitempty"`
	EncryptPubkey      string `json:"EncryptPubkey,omitempty"`
	GroupId            string `json:"GroupId,omitempty"`
	Memo               string `json:"Memo,omitempty"`
	OwnerPubkey        string `json:"OwnerPubkey,omitempty"`
	OwnerSignature     string `json:"OwnerSignature,omitempty"`
	Result             int64  `json:"Result,omitempty"`
	SignPubkey         string `json:"SignPubkey,omitempty"`
	TimeStamp          string `json:"TimeStamp,omitempty"`
	Type               int64  `json:"Type,omitempty"`
}
