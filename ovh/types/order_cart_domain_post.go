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

// OrderCartDomainPost ...
type OrderCartDomainPost struct {
	Domain string `json:"domain,omitempty"`

	Duration string `json:"duration,omitempty"`

	OfferID string `json:"offerId,omitempty"`

	Quantity int64 `json:"quantity,omitempty"`
}
