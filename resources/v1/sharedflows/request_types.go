// Code generated by Sideko (sideko.dev)
package sharedflows

import (
	nullable "apigee_api/nullable"
	types "apigee_api/types"
)

type ListRequest struct {
	Parent           string                                                   `json:"parent"`
	QueryXgafv       nullable.Nullable[types.GetV1ParentSharedflowsXgafvEnum] `json:"queryXgafv,omitempty"`
	AccessToken      nullable.Nullable[string]                                `json:"accessToken,omitempty"`
	Alt              nullable.Nullable[types.GetV1ParentSharedflowsAltEnum]   `json:"alt,omitempty"`
	Callback         nullable.Nullable[string]                                `json:"callback,omitempty"`
	Fields           nullable.Nullable[string]                                `json:"fields,omitempty"`
	IncludeMetaData  nullable.Nullable[bool]                                  `json:"includeMetaData,omitempty"`
	IncludeRevisions nullable.Nullable[bool]                                  `json:"includeRevisions,omitempty"`
	Key              nullable.Nullable[string]                                `json:"key,omitempty"`
	OauthToken       nullable.Nullable[string]                                `json:"oauthToken,omitempty"`
	PrettyPrint      nullable.Nullable[bool]                                  `json:"prettyPrint,omitempty"`
	QuotaUser        nullable.Nullable[string]                                `json:"quotaUser,omitempty"`
	UploadType       nullable.Nullable[string]                                `json:"uploadType,omitempty"`
	UploadProtocol   nullable.Nullable[string]                                `json:"uploadProtocol,omitempty"`
}

type CreateRequest struct {
	Parent         string                                                    `json:"parent"`
	QueryXgafv     nullable.Nullable[types.PostV1ParentSharedflowsXgafvEnum] `json:"queryXgafv,omitempty"`
	AccessToken    nullable.Nullable[string]                                 `json:"accessToken,omitempty"`
	Action         nullable.Nullable[string]                                 `json:"action,omitempty"`
	Alt            nullable.Nullable[types.PostV1ParentSharedflowsAltEnum]   `json:"alt,omitempty"`
	Callback       nullable.Nullable[string]                                 `json:"callback,omitempty"`
	Fields         nullable.Nullable[string]                                 `json:"fields,omitempty"`
	Key            nullable.Nullable[string]                                 `json:"key,omitempty"`
	Name           nullable.Nullable[string]                                 `json:"name,omitempty"`
	OauthToken     nullable.Nullable[string]                                 `json:"oauthToken,omitempty"`
	PrettyPrint    nullable.Nullable[bool]                                   `json:"prettyPrint,omitempty"`
	QuotaUser      nullable.Nullable[string]                                 `json:"quotaUser,omitempty"`
	UploadType     nullable.Nullable[string]                                 `json:"uploadType,omitempty"`
	UploadProtocol nullable.Nullable[string]                                 `json:"uploadProtocol,omitempty"`
	Data           types.GoogleApiHttpBody                                   `json:"data"`
}
