/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rum_sdk
// ApiGrpUserResult struct for ApiGrpUserResult
type ApiGrpUserResult struct {
	Action string `json:"action,omitempty"`
	EncryptPubkey string `json:"encrypt_pubkey,omitempty"`
	GroupId string `json:"group_id,omitempty"`
	Memo string `json:"memo,omitempty"`
	OwnerPubkey string `json:"owner_pubkey,omitempty"`
	Sign string `json:"sign,omitempty"`
	TrxId string `json:"trx_id,omitempty"`
	UserPubkey string `json:"user_pubkey,omitempty"`
}
