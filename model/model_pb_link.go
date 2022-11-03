/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// PbLink struct for PbLink
type PbLink struct {
	Height    int32    `json:"height,omitempty"`
	Href      string   `json:"href,omitempty"`
	Hreflang  string   `json:"hreflang,omitempty"`
	MediaType string   `json:"mediaType,omitempty"`
	Name      string   `json:"name,omitempty"`
	Preview   PbObject `json:"preview,omitempty"`
	Rel       []string `json:"rel,omitempty"`
	Width     int32    `json:"width,omitempty"`
}