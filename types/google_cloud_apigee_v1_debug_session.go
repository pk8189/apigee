// Code generated by Sideko (sideko.dev)
package types

import (
	nullable "apigee_api/nullable"
)

type GoogleCloudApigeeV1DebugSession struct {
	Count      nullable.Nullable[int]    `json:"count,omitempty"`
	CreateTime nullable.Nullable[string] `json:"createTime,omitempty"`
	Filter     nullable.Nullable[string] `json:"filter,omitempty"`
	Name       nullable.Nullable[string] `json:"name,omitempty"`
	Timeout    nullable.Nullable[string] `json:"timeout,omitempty"`
	Tracesize  nullable.Nullable[int]    `json:"tracesize,omitempty"`
	Validity   nullable.Nullable[int]    `json:"validity,omitempty"`
}
