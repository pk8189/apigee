// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1ListSharedFlowsResponse struct {
	SharedFlows nullable.Nullable[[]GoogleCloudApigeeV1SharedFlow] `json:"sharedFlows,omitempty"`
}
