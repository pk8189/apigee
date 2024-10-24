// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1DeveloperApp struct {
	ApiProducts    nullable.Nullable[[]string]                        `json:"apiProducts,omitempty"`
	AppFamily      nullable.Nullable[string]                          `json:"appFamily,omitempty"`
	AppId          nullable.Nullable[string]                          `json:"appId,omitempty"`
	Attributes     nullable.Nullable[[]GoogleCloudApigeeV1Attribute]  `json:"attributes,omitempty"`
	CallbackUrl    nullable.Nullable[string]                          `json:"callbackUrl,omitempty"`
	CreatedAt      nullable.Nullable[string]                          `json:"createdAt,omitempty"`
	Credentials    nullable.Nullable[[]GoogleCloudApigeeV1Credential] `json:"credentials,omitempty"`
	DeveloperId    nullable.Nullable[string]                          `json:"developerId,omitempty"`
	KeyExpiresIn   nullable.Nullable[string]                          `json:"keyExpiresIn,omitempty"`
	LastModifiedAt nullable.Nullable[string]                          `json:"lastModifiedAt,omitempty"`
	Name           nullable.Nullable[string]                          `json:"name,omitempty"`
	Scopes         nullable.Nullable[[]string]                        `json:"scopes,omitempty"`
	Status         nullable.Nullable[string]                          `json:"status,omitempty"`
}
