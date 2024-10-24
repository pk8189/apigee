// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1RevisionStatus struct {
	Errors     nullable.Nullable[[]GoogleCloudApigeeV1UpdateError] `json:"errors,omitempty"`
	JsonSpec   nullable.Nullable[string]                           `json:"jsonSpec,omitempty"`
	Replicas   nullable.Nullable[int]                              `json:"replicas,omitempty"`
	RevisionId nullable.Nullable[string]                           `json:"revisionId,omitempty"`
}
