// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1ListDataCollectorsResponse struct {
	DataCollectors nullable.Nullable[[]GoogleCloudApigeeV1DataCollector] `json:"dataCollectors,omitempty"`
	NextPageToken  nullable.Nullable[string]                             `json:"nextPageToken,omitempty"`
}
