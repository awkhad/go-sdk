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

// TelephonySetDefaultSipDomainPost ...
type TelephonySetDefaultSipDomainPost struct {
	Country string `json:"country,omitempty"`

	Domain string `json:"domain,omitempty"`

	TType string `json:"type,omitempty"`
}
