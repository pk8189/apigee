// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1ListInstanceAttachmentsResponse struct {
	Attachments   nullable.Nullable[[]GoogleCloudApigeeV1InstanceAttachment] `json:"attachments,omitempty"`
	NextPageToken nullable.Nullable[string]                                  `json:"nextPageToken,omitempty"`
}
