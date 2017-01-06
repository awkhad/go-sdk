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

// PackXdslDomainServicesPost ...
type PackXdslDomainServicesPost struct {
	Action string `json:"action,omitempty"`

	AuthInfo string `json:"authInfo,omitempty"`

	Domain string `json:"domain,omitempty"`

	Tld string `json:"tld,omitempty"`
}
