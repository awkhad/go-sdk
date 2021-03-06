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

// DedicatedServerEmailAlert Service monitoring Email alert
type DedicatedServerEmailAlert struct {

	// AlertID This monitoring id
	AlertID int64 `json:"alertId,omitempty"`

	// Email Alert destination
	Email string `json:"email,omitempty"`

	// Enabled Is this monitor enabled
	Enabled bool `json:"enabled,omitempty"`

	// Language Alert language
	Language string `json:"language,omitempty"`
}
