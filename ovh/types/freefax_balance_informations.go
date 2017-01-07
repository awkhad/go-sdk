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

// FreefaxBalanceInformations Return credit balance informations structure
type FreefaxBalanceInformations struct {

	// Faxs The number of equivalement remaining french faxs
	Faxs int64 `json:"faxs,omitempty"`

	// Points Total balance available in points
	Points int64 `json:"points,omitempty"`
}