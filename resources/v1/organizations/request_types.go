// Code generated by Sideko (sideko.dev)
package organizations

import (
	nullable "apigee_api/nullable"
	types "apigee_api/types"
)

type CreateRequest struct {
	QueryXgafv     nullable.Nullable[types.PostV1OrganizationsXgafvEnum] `json:"queryXgafv,omitempty"`
	AccessToken    nullable.Nullable[string]                             `json:"accessToken,omitempty"`
	Alt            nullable.Nullable[types.PostV1OrganizationsAltEnum]   `json:"alt,omitempty"`
	Callback       nullable.Nullable[string]                             `json:"callback,omitempty"`
	Fields         nullable.Nullable[string]                             `json:"fields,omitempty"`
	Key            nullable.Nullable[string]                             `json:"key,omitempty"`
	OauthToken     nullable.Nullable[string]                             `json:"oauthToken,omitempty"`
	Parent         nullable.Nullable[string]                             `json:"parent,omitempty"`
	PrettyPrint    nullable.Nullable[bool]                               `json:"prettyPrint,omitempty"`
	QuotaUser      nullable.Nullable[string]                             `json:"quotaUser,omitempty"`
	UploadType     nullable.Nullable[string]                             `json:"uploadType,omitempty"`
	UploadProtocol nullable.Nullable[string]                             `json:"uploadProtocol,omitempty"`
	Data           types.GoogleCloudApigeeV1Organization                 `json:"data"`
}
