// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1ListOrganizationsResponse struct {
	Organizations nullable.Nullable[[]GoogleCloudApigeeV1OrganizationProjectMapping] `json:"organizations,omitempty"`
}
