// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1ApiProxyRevision1 struct {
	Archive                    nullable.Nullable[string]                                                        `json:"archive,omitempty"`
	Basepaths                  nullable.Nullable[[]string]                                                      `json:"basepaths,omitempty"`
	ConfigurationVersion       nullable.Nullable[GoogleCloudApigeeV1ConfigVersion]                              `json:"configurationVersion,omitempty"`
	ContextInfo                nullable.Nullable[string]                                                        `json:"contextInfo,omitempty"`
	CreatedAt                  nullable.Nullable[string]                                                        `json:"createdAt,omitempty"`
	Description                nullable.Nullable[string]                                                        `json:"description,omitempty"`
	DisplayName                nullable.Nullable[string]                                                        `json:"displayName,omitempty"`
	EntityMetaDataAsProperties nullable.Nullable[GoogleCloudApigeeV1ApiProxyRevisionEntityMetaDataAsProperties] `json:"entityMetaDataAsProperties,omitempty"`
	IntegrationEndpoints       nullable.Nullable[[]string]                                                      `json:"integrationEndpoints,omitempty"`
	LastModifiedAt             nullable.Nullable[string]                                                        `json:"lastModifiedAt,omitempty"`
	Name                       nullable.Nullable[string]                                                        `json:"name,omitempty"`
	Policies                   nullable.Nullable[[]string]                                                      `json:"policies,omitempty"`
	Proxies                    nullable.Nullable[[]string]                                                      `json:"proxies,omitempty"`
	ProxyEndpoints             nullable.Nullable[[]string]                                                      `json:"proxyEndpoints,omitempty"`
	ResourceFiles              nullable.Nullable[GoogleCloudApigeeV1ResourceFiles]                              `json:"resourceFiles,omitempty"`
	Resources                  nullable.Nullable[[]string]                                                      `json:"resources,omitempty"`
	Revision                   nullable.Nullable[string]                                                        `json:"revision,omitempty"`
	SharedFlows                nullable.Nullable[[]string]                                                      `json:"sharedFlows,omitempty"`
	Spec                       nullable.Nullable[string]                                                        `json:"spec,omitempty"`
	TargetEndpoints            nullable.Nullable[[]string]                                                      `json:"targetEndpoints,omitempty"`
	TargetServers              nullable.Nullable[[]string]                                                      `json:"targetServers,omitempty"`
	Targets                    nullable.Nullable[[]string]                                                      `json:"targets,omitempty"`
	Teams                      nullable.Nullable[[]string]                                                      `json:"teams,omitempty"`
	Type                       nullable.Nullable[string]                                                        `json:"type,omitempty"`
}
