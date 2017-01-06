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

// DBaaSLogsRole Graylog role
type DBaaSLogsRole struct {

	// Description Role description
	Description string `json:"description,omitempty"`

	// Name Role name
	Name string `json:"name,omitempty"`

	// OptionID Associated DBaaS Logs option
	OptionID string `json:"optionId,omitempty"`

	// RoleID Role UUID
	RoleID string `json:"roleId,omitempty"`
}
