// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1RevenueShareRange struct {
	End             nullable.Nullable[string]  `json:"end,omitempty"`
	SharePercentage nullable.Nullable[float64] `json:"sharePercentage,omitempty"`
	Start           nullable.Nullable[string]  `json:"start,omitempty"`
}
