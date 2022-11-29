/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// ChainPublishQueueItem struct for ChainPublishQueueItem
type ChainPublishQueueItem struct {
	UpdateAt string `json:"UpdateAt,omitempty"`
	// also stored in the key for quick indexing
	GroupId    string `json:"groupId,omitempty"`
	RetryCount int64  `json:"retryCount,omitempty"`
	// in value only
	State       string `json:"state,omitempty"`
	StorageType string `json:"storageType,omitempty"`
	Trx         PbTrx  `json:"trx,omitempty"`
}
