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

// SmsAlertThreshold A structure describing all information about alert threshold informations
type SmsAlertThreshold struct {
	AlertEmail string `json:"alertEmail,omitempty"`

	AlertNumber string `json:"alertNumber,omitempty"`

	AlertThreshold int64 `json:"alertThreshold,omitempty"`

	Support string `json:"support,omitempty"`
}
