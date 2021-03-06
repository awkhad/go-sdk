/*
 * OVH API - EU
 *
 * Build your own OVH world.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: api-subscribe@ml.ovh.net
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package types

// DBaasLogsOutputGraylogStreamAlertPost ...
type DBaasLogsOutputGraylogStreamAlertPost struct {
	Backlog int64 `json:"backlog,omitempty"`

	ConditionType string `json:"conditionType,omitempty"`

	ConstraintType string `json:"constraintType,omitempty"`

	Field string `json:"field,omitempty"`

	Grace int64 `json:"grace,omitempty"`

	Threshold int64 `json:"threshold,omitempty"`

	ThresholdType string `json:"thresholdType,omitempty"`

	Time int64 `json:"time,omitempty"`

	Title string `json:"title,omitempty"`

	Value string `json:"value,omitempty"`
}
