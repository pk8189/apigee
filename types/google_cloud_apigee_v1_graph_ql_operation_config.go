// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1GraphQlOperationConfig struct {
	ApiSource  nullable.Nullable[string]                                `json:"apiSource,omitempty"`
	Attributes nullable.Nullable[[]GoogleCloudApigeeV1Attribute]        `json:"attributes,omitempty"`
	Operations nullable.Nullable[[]GoogleCloudApigeeV1GraphQlOperation] `json:"operations,omitempty"`
	Quota      nullable.Nullable[GoogleCloudApigeeV1Quota]              `json:"quota,omitempty"`
}
