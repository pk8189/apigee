// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1DateRange struct {
	End   nullable.Nullable[string] `json:"end,omitempty"`
	Start nullable.Nullable[string] `json:"start,omitempty"`
}
