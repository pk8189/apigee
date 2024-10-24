// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1ResourceStatus struct {
	Resource      nullable.Nullable[string]                              `json:"resource,omitempty"`
	Revisions     nullable.Nullable[[]GoogleCloudApigeeV1RevisionStatus] `json:"revisions,omitempty"`
	TotalReplicas nullable.Nullable[int]                                 `json:"totalReplicas,omitempty"`
	Uid           nullable.Nullable[string]                              `json:"uid,omitempty"`
}
