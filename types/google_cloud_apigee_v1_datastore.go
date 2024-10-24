// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1Datastore struct {
	CreateTime      nullable.Nullable[string]                             `json:"createTime,omitempty"`
	DatastoreConfig nullable.Nullable[GoogleCloudApigeeV1DatastoreConfig] `json:"datastoreConfig,omitempty"`
	DisplayName     nullable.Nullable[string]                             `json:"displayName,omitempty"`
	LastUpdateTime  nullable.Nullable[string]                             `json:"lastUpdateTime,omitempty"`
	Org             nullable.Nullable[string]                             `json:"org,omitempty"`
	Self            nullable.Nullable[string]                             `json:"self,omitempty"`
	TargetType      nullable.Nullable[string]                             `json:"targetType,omitempty"`
}
