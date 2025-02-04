// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1SharedFlowRevision1 struct {
	ConfigurationVersion       nullable.Nullable[GoogleCloudApigeeV1ConfigVersion]                                `json:"configurationVersion,omitempty"`
	ContextInfo                nullable.Nullable[string]                                                          `json:"contextInfo,omitempty"`
	CreatedAt                  nullable.Nullable[string]                                                          `json:"createdAt,omitempty"`
	Description                nullable.Nullable[string]                                                          `json:"description,omitempty"`
	DisplayName                nullable.Nullable[string]                                                          `json:"displayName,omitempty"`
	EntityMetaDataAsProperties nullable.Nullable[GoogleCloudApigeeV1SharedFlowRevisionEntityMetaDataAsProperties] `json:"entityMetaDataAsProperties,omitempty"`
	LastModifiedAt             nullable.Nullable[string]                                                          `json:"lastModifiedAt,omitempty"`
	Name                       nullable.Nullable[string]                                                          `json:"name,omitempty"`
	Policies                   nullable.Nullable[[]string]                                                        `json:"policies,omitempty"`
	ResourceFiles              nullable.Nullable[GoogleCloudApigeeV1ResourceFiles]                                `json:"resourceFiles,omitempty"`
	Resources                  nullable.Nullable[[]string]                                                        `json:"resources,omitempty"`
	Revision                   nullable.Nullable[string]                                                          `json:"revision,omitempty"`
	SharedFlows                nullable.Nullable[[]string]                                                        `json:"sharedFlows,omitempty"`
	Type                       nullable.Nullable[string]                                                          `json:"type,omitempty"`
}
