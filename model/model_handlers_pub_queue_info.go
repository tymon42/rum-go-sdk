/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// HandlersPubQueueInfo struct for HandlersPubQueueInfo
type HandlersPubQueueInfo struct {
	Data    []ChainPublishQueueItem `json:"data,omitempty"`
	GroupId string                  `json:"groupId,omitempty"`
}