/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// HandlersAnnounceParam struct for HandlersAnnounceParam
type HandlersAnnounceParam struct {
	Action  string `json:"action"`
	GroupId string `json:"group_id"`
	Memo    string `json:"memo,omitempty"`
	Type    string `json:"type"`
}
