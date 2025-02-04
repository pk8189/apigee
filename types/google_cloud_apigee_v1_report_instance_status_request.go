// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1ReportInstanceStatusRequest struct {
	InstanceUid nullable.Nullable[string]                              `json:"instanceUid,omitempty"`
	ReportTime  nullable.Nullable[string]                              `json:"reportTime,omitempty"`
	Resources   nullable.Nullable[[]GoogleCloudApigeeV1ResourceStatus] `json:"resources,omitempty"`
}
