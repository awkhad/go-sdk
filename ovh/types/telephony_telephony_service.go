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

// TelephonyService Telephony service
type TelephonyService struct {

	// Country The country of the number
	Country string `json:"country,omitempty"`

	// CountryCode The country code of the number
	CountryCode int64 `json:"countryCode,omitempty"`

	CurrentOutplan *OrderPrice `json:"currentOutplan,omitempty"`

	Description string `json:"description,omitempty"`

	FeatureType string `json:"featureType,omitempty"`

	Offers []string `json:"offers,omitempty"`

	Properties []string `json:"properties,omitempty"`

	// Rio The identifier to use to port the number
	Rio string `json:"rio,omitempty"`

	ServiceName string `json:"serviceName,omitempty"`

	ServiceType string `json:"serviceType,omitempty"`

	SimultaneousLines int64 `json:"simultaneousLines,omitempty"`
}
