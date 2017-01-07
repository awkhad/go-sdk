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

// DomainDataClaimNoticeAddress Address for a claim notice holder
type DomainDataClaimNoticeAddress struct {

	// City City
	City string `json:"city,omitempty"`

	// CountryCode Country code
	CountryCode string `json:"countryCode,omitempty"`

	// Fax Fax number
	Fax string `json:"fax,omitempty"`

	// FaxExtension Fax number extension
	FaxExtension string `json:"faxExtension,omitempty"`

	// PostalCode Postal zip code
	PostalCode string `json:"postalCode,omitempty"`

	// StateOrProvince State of province
	StateOrProvince string `json:"stateOrProvince,omitempty"`

	// Streets Array of street name
	Streets []string `json:"streets,omitempty"`

	// Voice Phone number
	Voice string `json:"voice,omitempty"`

	// VoiceExtension Phone number extension
	VoiceExtension string `json:"voiceExtension,omitempty"`
}