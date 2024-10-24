// Code generated by Sideko (sideko.dev)
package set_sync_authorization

import (
	nullable "apigee_api/nullable"
	types "apigee_api/types"
)

type CreateRequest struct {
	Name           string                                                           `json:"name"`
	QueryXgafv     nullable.Nullable[types.PostV1NameSetSyncAuthorizationXgafvEnum] `json:"queryXgafv,omitempty"`
	AccessToken    nullable.Nullable[string]                                        `json:"accessToken,omitempty"`
	Alt            nullable.Nullable[types.PostV1NameSetSyncAuthorizationAltEnum]   `json:"alt,omitempty"`
	Callback       nullable.Nullable[string]                                        `json:"callback,omitempty"`
	Fields         nullable.Nullable[string]                                        `json:"fields,omitempty"`
	Key            nullable.Nullable[string]                                        `json:"key,omitempty"`
	OauthToken     nullable.Nullable[string]                                        `json:"oauthToken,omitempty"`
	PrettyPrint    nullable.Nullable[bool]                                          `json:"prettyPrint,omitempty"`
	QuotaUser      nullable.Nullable[string]                                        `json:"quotaUser,omitempty"`
	UploadType     nullable.Nullable[string]                                        `json:"uploadType,omitempty"`
	UploadProtocol nullable.Nullable[string]                                        `json:"uploadProtocol,omitempty"`
	Data           types.GoogleCloudApigeeV1SyncAuthorization                       `json:"data"`
}
