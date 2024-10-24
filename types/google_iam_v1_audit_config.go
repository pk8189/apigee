// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleIamV1AuditConfig struct {
	AuditLogConfigs nullable.Nullable[[]GoogleIamV1AuditLogConfig] `json:"auditLogConfigs,omitempty"`
	Service         nullable.Nullable[string]                      `json:"service,omitempty"`
}
