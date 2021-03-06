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

// IPMitigationIP Your IP on mitigation
type IPMitigationIP struct {

	// Auto Set on true if your ip is on auto-mitigation
	Auto bool `json:"auto,omitempty"`

	IPOnMitigation string `json:"ipOnMitigation,omitempty"`

	// Permanent Set on true if your ip is on permanent mitigation
	Permanent bool `json:"permanent,omitempty"`

	// State Current state of your ip on mitigation
	State string `json:"state,omitempty"`
}
