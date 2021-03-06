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

// DomainDataClaimNotice Definition of claim notices applying to a domain name
type DomainDataClaimNotice struct {

	// Claims Array of claim notice for the domain
	Claims []*DomainDataClaimNoticeDecision `json:"claims,omitempty"`

	// EndingDate Ending date of claim notice
	EndingDate string `json:"endingDate,omitempty"`

	// ID Claim notice ID
	ID string `json:"id,omitempty"`

	// Label Label referring to claim notice
	Label string `json:"label,omitempty"`

	// StartingDate Beginning date of claim notice
	StartingDate string `json:"startingDate,omitempty"`

	// TType Type of claim notice
	TType string `json:"type,omitempty"`
}
