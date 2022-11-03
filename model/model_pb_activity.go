/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// PbActivity struct for PbActivity
type PbActivity struct {
	Actor        PbObject             `json:"actor,omitempty"`
	Attachments  []PbObject           `json:"attachments,omitempty"`
	AttributedTo []PbObject           `json:"attributedTo,omitempty"`
	Audience     PbObject             `json:"audience,omitempty"`
	Bcc          []PbObject           `json:"bcc,omitempty"`
	Bto          []PbObject           `json:"bto,omitempty"`
	Cc           []PbObject           `json:"cc,omitempty"`
	Content      string               `json:"content,omitempty"`
	Context      PbObject             `json:"context,omitempty"`
	Duration     string               `json:"duration,omitempty"`
	Endtime      TimestamppbTimestamp `json:"endtime,omitempty"`
	Generator    PbObject             `json:"generator,omitempty"`
	Icon         []PbObject           `json:"icon,omitempty"`
	Id           string               `json:"id,omitempty"`
	Image        []PbObject           `json:"image,omitempty"`
	InReplyTo    PbObject             `json:"inReplyTo,omitempty"`
	Instrument   PbObject             `json:"instrument,omitempty"`
	Location     PbObject             `json:"location,omitempty"`
	MediaType    string               `json:"mediaType,omitempty"`
	Name         string               `json:"name,omitempty"`
	Object       PbObject             `json:"object,omitempty"`
	Origin       PbObject             `json:"origin,omitempty"`
	Person       PbPerson             `json:"person,omitempty"`
	Preview      PbObject             `json:"preview,omitempty"`
	Published    TimestamppbTimestamp `json:"published,omitempty"`
	Replies      PbObject             `json:"replies,omitempty"`
	Result       PbObject             `json:"result,omitempty"`
	StartTime    TimestamppbTimestamp `json:"startTime,omitempty"`
	Summary      string               `json:"summary,omitempty"`
	Tag          []PbObject           `json:"tag,omitempty"`
	Target       PbObject             `json:"target,omitempty"`
	To           []PbObject           `json:"to,omitempty"`
	Type         string               `json:"type,omitempty"`
	Updated      TimestamppbTimestamp `json:"updated,omitempty"`
	Url          []PbLink             `json:"url,omitempty"`
}