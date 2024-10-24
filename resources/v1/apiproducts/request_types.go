// Code generated by Sideko (sideko.dev)
package apiproducts

import (
	nullable "apigee_api/nullable"
	types "apigee_api/types"
)

type ListRequest struct {
	Parent         string                                                   `json:"parent"`
	QueryXgafv     nullable.Nullable[types.GetV1ParentApiproductsXgafvEnum] `json:"queryXgafv,omitempty"`
	AccessToken    nullable.Nullable[string]                                `json:"accessToken,omitempty"`
	Alt            nullable.Nullable[types.GetV1ParentApiproductsAltEnum]   `json:"alt,omitempty"`
	Attributename  nullable.Nullable[string]                                `json:"attributename,omitempty"`
	Attributevalue nullable.Nullable[string]                                `json:"attributevalue,omitempty"`
	Callback       nullable.Nullable[string]                                `json:"callback,omitempty"`
	Count          nullable.Nullable[string]                                `json:"count,omitempty"`
	Expand         nullable.Nullable[bool]                                  `json:"expand,omitempty"`
	Fields         nullable.Nullable[string]                                `json:"fields,omitempty"`
	Key            nullable.Nullable[string]                                `json:"key,omitempty"`
	OauthToken     nullable.Nullable[string]                                `json:"oauthToken,omitempty"`
	PrettyPrint    nullable.Nullable[bool]                                  `json:"prettyPrint,omitempty"`
	QuotaUser      nullable.Nullable[string]                                `json:"quotaUser,omitempty"`
	StartKey       nullable.Nullable[string]                                `json:"startKey,omitempty"`
	UploadType     nullable.Nullable[string]                                `json:"uploadType,omitempty"`
	UploadProtocol nullable.Nullable[string]                                `json:"uploadProtocol,omitempty"`
}

type CreateRequest struct {
	Parent         string                                                    `json:"parent"`
	QueryXgafv     nullable.Nullable[types.PostV1ParentApiproductsXgafvEnum] `json:"queryXgafv,omitempty"`
	AccessToken    nullable.Nullable[string]                                 `json:"accessToken,omitempty"`
	Alt            nullable.Nullable[types.PostV1ParentApiproductsAltEnum]   `json:"alt,omitempty"`
	Callback       nullable.Nullable[string]                                 `json:"callback,omitempty"`
	Fields         nullable.Nullable[string]                                 `json:"fields,omitempty"`
	Key            nullable.Nullable[string]                                 `json:"key,omitempty"`
	OauthToken     nullable.Nullable[string]                                 `json:"oauthToken,omitempty"`
	PrettyPrint    nullable.Nullable[bool]                                   `json:"prettyPrint,omitempty"`
	QuotaUser      nullable.Nullable[string]                                 `json:"quotaUser,omitempty"`
	UploadType     nullable.Nullable[string]                                 `json:"uploadType,omitempty"`
	UploadProtocol nullable.Nullable[string]                                 `json:"uploadProtocol,omitempty"`
	Data           types.GoogleCloudApigeeV1ApiProduct                       `json:"data"`
}
