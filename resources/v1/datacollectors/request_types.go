// Code generated by Sideko (sideko.dev)
package datacollectors

import (
	nullable "apigee_api/nullable"
	types "apigee_api/types"
)

type ListRequest struct {
	Parent         string                                                      `json:"parent"`
	QueryXgafv     nullable.Nullable[types.GetV1ParentDatacollectorsXgafvEnum] `json:"queryXgafv,omitempty"`
	AccessToken    nullable.Nullable[string]                                   `json:"accessToken,omitempty"`
	Alt            nullable.Nullable[types.GetV1ParentDatacollectorsAltEnum]   `json:"alt,omitempty"`
	Callback       nullable.Nullable[string]                                   `json:"callback,omitempty"`
	Fields         nullable.Nullable[string]                                   `json:"fields,omitempty"`
	Key            nullable.Nullable[string]                                   `json:"key,omitempty"`
	OauthToken     nullable.Nullable[string]                                   `json:"oauthToken,omitempty"`
	PageSize       nullable.Nullable[int]                                      `json:"pageSize,omitempty"`
	PageToken      nullable.Nullable[string]                                   `json:"pageToken,omitempty"`
	PrettyPrint    nullable.Nullable[bool]                                     `json:"prettyPrint,omitempty"`
	QuotaUser      nullable.Nullable[string]                                   `json:"quotaUser,omitempty"`
	UploadType     nullable.Nullable[string]                                   `json:"uploadType,omitempty"`
	UploadProtocol nullable.Nullable[string]                                   `json:"uploadProtocol,omitempty"`
}

type CreateRequest struct {
	Parent          string                                                       `json:"parent"`
	QueryXgafv      nullable.Nullable[types.PostV1ParentDatacollectorsXgafvEnum] `json:"queryXgafv,omitempty"`
	AccessToken     nullable.Nullable[string]                                    `json:"accessToken,omitempty"`
	Alt             nullable.Nullable[types.PostV1ParentDatacollectorsAltEnum]   `json:"alt,omitempty"`
	Callback        nullable.Nullable[string]                                    `json:"callback,omitempty"`
	DataCollectorId nullable.Nullable[string]                                    `json:"dataCollectorId,omitempty"`
	Fields          nullable.Nullable[string]                                    `json:"fields,omitempty"`
	Key             nullable.Nullable[string]                                    `json:"key,omitempty"`
	OauthToken      nullable.Nullable[string]                                    `json:"oauthToken,omitempty"`
	PrettyPrint     nullable.Nullable[bool]                                      `json:"prettyPrint,omitempty"`
	QuotaUser       nullable.Nullable[string]                                    `json:"quotaUser,omitempty"`
	UploadType      nullable.Nullable[string]                                    `json:"uploadType,omitempty"`
	UploadProtocol  nullable.Nullable[string]                                    `json:"uploadProtocol,omitempty"`
	Data            types.GoogleCloudApigeeV1DataCollector                       `json:"data"`
}
