// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1EndpointAttachment struct {
	ConnectionState   nullable.Nullable[GoogleCloudApigeeV1EndpointAttachmentConnectionStateEnum] `json:"connectionState,omitempty"`
	Host              nullable.Nullable[string]                                                   `json:"host,omitempty"`
	Location          nullable.Nullable[string]                                                   `json:"location,omitempty"`
	Name              nullable.Nullable[string]                                                   `json:"name,omitempty"`
	ServiceAttachment nullable.Nullable[string]                                                   `json:"serviceAttachment,omitempty"`
	State             nullable.Nullable[GoogleCloudApigeeV1EndpointAttachmentStateEnum]           `json:"state,omitempty"`
}
