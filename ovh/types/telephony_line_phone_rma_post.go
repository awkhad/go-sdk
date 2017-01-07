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

// TelephonyLinePhoneRmaPost ...
type TelephonyLinePhoneRmaPost struct {
	MondialRelayID string `json:"mondialRelayId,omitempty"`

	NewMerchandise string `json:"newMerchandise,omitempty"`

	ShippingContactID int64 `json:"shippingContactId,omitempty"`

	TType string `json:"type,omitempty"`
}