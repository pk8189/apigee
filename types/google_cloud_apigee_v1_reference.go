// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1Reference struct {
	Description  nullable.Nullable[string] `json:"description,omitempty"`
	Name         nullable.Nullable[string] `json:"name,omitempty"`
	Refers       nullable.Nullable[string] `json:"refers,omitempty"`
	ResourceType nullable.Nullable[string] `json:"resourceType,omitempty"`
}
