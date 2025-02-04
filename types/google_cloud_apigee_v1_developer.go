// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1Developer struct {
	AccessType       nullable.Nullable[string]                         `json:"accessType,omitempty"`
	AppFamily        nullable.Nullable[string]                         `json:"appFamily,omitempty"`
	Apps             nullable.Nullable[[]string]                       `json:"apps,omitempty"`
	Attributes       nullable.Nullable[[]GoogleCloudApigeeV1Attribute] `json:"attributes,omitempty"`
	Companies        nullable.Nullable[[]string]                       `json:"companies,omitempty"`
	CreatedAt        nullable.Nullable[string]                         `json:"createdAt,omitempty"`
	DeveloperId      nullable.Nullable[string]                         `json:"developerId,omitempty"`
	Email            nullable.Nullable[string]                         `json:"email,omitempty"`
	FirstName        nullable.Nullable[string]                         `json:"firstName,omitempty"`
	LastModifiedAt   nullable.Nullable[string]                         `json:"lastModifiedAt,omitempty"`
	LastName         nullable.Nullable[string]                         `json:"lastName,omitempty"`
	OrganizationName nullable.Nullable[string]                         `json:"organizationName,omitempty"`
	Status           nullable.Nullable[string]                         `json:"status,omitempty"`
	UserName         nullable.Nullable[string]                         `json:"userName,omitempty"`
}
