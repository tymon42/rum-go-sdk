/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rum_sdk

// HandlersChainSendTrxRuleListItem struct for HandlersChainSendTrxRuleListItem
type HandlersChainSendTrxRuleListItem struct {
	GroupOwnerPubkey string   `json:"groupOwnerPubkey"`
	GroupOwnerSign   string   `json:"groupOwnerSign"`
	Memo             string   `json:"memo,omitempty"`
	Pubkey           string   `json:"pubkey"`
	TimeStamp        int64    `json:"timeStamp"`
	TrxType          []string `json:"trxType"`
}
