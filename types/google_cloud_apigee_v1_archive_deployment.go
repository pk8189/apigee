// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1ArchiveDeployment struct {
	CreatedAt nullable.Nullable[string]                                     `json:"createdAt,omitempty"`
	GcsUri    nullable.Nullable[string]                                     `json:"gcsUri,omitempty"`
	Labels    nullable.Nullable[GoogleCloudApigeeV1ArchiveDeploymentLabels] `json:"labels,omitempty"`
	Name      nullable.Nullable[string]                                     `json:"name,omitempty"`
	Operation nullable.Nullable[string]                                     `json:"operation,omitempty"`
	UpdatedAt nullable.Nullable[string]                                     `json:"updatedAt,omitempty"`
}
