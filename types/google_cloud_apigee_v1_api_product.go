// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1ApiProduct struct {
	ApiResources          nullable.Nullable[[]string]                                           `json:"apiResources,omitempty"`
	ApprovalType          nullable.Nullable[string]                                             `json:"approvalType,omitempty"`
	Attributes            nullable.Nullable[[]GoogleCloudApigeeV1Attribute]                     `json:"attributes,omitempty"`
	CreatedAt             nullable.Nullable[string]                                             `json:"createdAt,omitempty"`
	Description           nullable.Nullable[string]                                             `json:"description,omitempty"`
	DisplayName           nullable.Nullable[string]                                             `json:"displayName,omitempty"`
	Environments          nullable.Nullable[[]string]                                           `json:"environments,omitempty"`
	GraphqlOperationGroup nullable.Nullable[GoogleCloudApigeeV1GraphQlOperationGroup]           `json:"graphqlOperationGroup,omitempty"`
	LastModifiedAt        nullable.Nullable[string]                                             `json:"lastModifiedAt,omitempty"`
	Name                  nullable.Nullable[string]                                             `json:"name,omitempty"`
	OperationGroup        nullable.Nullable[GoogleCloudApigeeV1OperationGroup]                  `json:"operationGroup,omitempty"`
	Proxies               nullable.Nullable[[]string]                                           `json:"proxies,omitempty"`
	Quota                 nullable.Nullable[string]                                             `json:"quota,omitempty"`
	QuotaCounterScope     nullable.Nullable[GoogleCloudApigeeV1ApiProductQuotaCounterScopeEnum] `json:"quotaCounterScope,omitempty"`
	QuotaInterval         nullable.Nullable[string]                                             `json:"quotaInterval,omitempty"`
	QuotaTimeUnit         nullable.Nullable[string]                                             `json:"quotaTimeUnit,omitempty"`
	Scopes                nullable.Nullable[[]string]                                           `json:"scopes,omitempty"`
}
