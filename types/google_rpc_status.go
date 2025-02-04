// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleRpcStatus struct {
	Code    nullable.Nullable[int]                      `json:"code,omitempty"`
	Details nullable.Nullable[[]map[string]interface{}] `json:"details,omitempty"`
	Message nullable.Nullable[string]                   `json:"message,omitempty"`
}
