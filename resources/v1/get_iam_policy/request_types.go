// Code generated by Sideko (sideko.dev)
package get_iam_policy

import (
	nullable "apigee_api/nullable"
	types "apigee_api/types"
)

type ListRequest struct {
	Resource                      string                                                      `json:"resource"`
	QueryXgafv                    nullable.Nullable[types.GetV1ResourceGetIamPolicyXgafvEnum] `json:"queryXgafv,omitempty"`
	AccessToken                   nullable.Nullable[string]                                   `json:"accessToken,omitempty"`
	Alt                           nullable.Nullable[types.GetV1ResourceGetIamPolicyAltEnum]   `json:"alt,omitempty"`
	Callback                      nullable.Nullable[string]                                   `json:"callback,omitempty"`
	Fields                        nullable.Nullable[string]                                   `json:"fields,omitempty"`
	Key                           nullable.Nullable[string]                                   `json:"key,omitempty"`
	OauthToken                    nullable.Nullable[string]                                   `json:"oauthToken,omitempty"`
	OptionsRequestedPolicyVersion nullable.Nullable[int]                                      `json:"optionsRequestedPolicyVersion,omitempty"`
	PrettyPrint                   nullable.Nullable[bool]                                     `json:"prettyPrint,omitempty"`
	QuotaUser                     nullable.Nullable[string]                                   `json:"quotaUser,omitempty"`
	UploadType                    nullable.Nullable[string]                                   `json:"uploadType,omitempty"`
	UploadProtocol                nullable.Nullable[string]                                   `json:"uploadProtocol,omitempty"`
}
